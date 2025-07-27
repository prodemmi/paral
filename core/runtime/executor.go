package runtime

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"os/exec"
	"paral/core/variable"
	"strings"
	"time"
)

type Executor struct {
	Runtime  *Runtime
	Reporter *Reporter
}

type ExecutionConfig struct {
	SummaryOnly bool
	Verbose     bool
	Silent      bool
	Short       bool
	DryRun      bool
	NoColor     bool
	Profile     bool
	LogLevel    string
	MaxWorkers  int
	Timeout     int
	OutputFile  string
	Writer      io.Writer // Optional custom writer
}

type ExecutionStats struct {
	TotalJobs          int
	TotalCommands      int
	SuccessfulCommands int
	FailedCommands     int
	StartTime          time.Time
	EndTime            time.Time
	JobStats           map[string]JobStats
}

type ExecutionContext struct {
	Config *ExecutionConfig
	Stats  *ExecutionStats
	Colors *ColorScheme
	Writer io.Writer
}

type JobStats struct {
	Total      int
	Successful int
	Failed     int
	Duration   time.Duration
}

type ColorScheme struct {
	Blue    *color.Color
	Green   *color.Color
	Red     *color.Color
	Yellow  *color.Color
	Cyan    *color.Color
	Magenta *color.Color
	Gray    *color.Color
	White   *color.Color
}

func NewExecutor(runtime *Runtime) *Executor {
	return &Executor{
		Runtime: runtime,
	}
}

func (c *Executor) ExecWithConfig(config *ExecutionConfig) {
	ctx := c.createExecutionContext(config)

	if config.Profile {
		ctx.Stats.StartTime = time.Now()
	}

	if !config.Silent {
		c.printHeader(ctx)
		c.printVariables(ctx)
	}

	if config.DryRun {
		c.executeDryRun(ctx)
	} else {
		c.executeSequential(ctx)
	}

	if config.Profile {
		ctx.Stats.EndTime = time.Now()
	}

	if !config.Silent && !config.SummaryOnly {
		_, _ = fmt.Fprintln(ctx.Writer)
	}

	if !config.Silent {
		c.printSummary(ctx)
	}
}

func (c *Executor) createExecutionContext(config *ExecutionConfig) *ExecutionContext {
	writer := config.Writer
	if writer == nil {
		if config.OutputFile != "" {
			file, err := os.Create(config.OutputFile)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Failed to create output file: %v\n", err)
				writer = os.Stdout
			} else {
				writer = file
			}
		} else {
			writer = os.Stdout
		}
	}

	colors := c.createColorScheme(config.NoColor)

	// Count only non-manual jobs for the initial stats
	nonManualJobs := 0
	for _, task := range c.Runtime.Jobs {
		if !task.IsManual {
			nonManualJobs++
		}
	}

	stats := &ExecutionStats{
		TotalJobs: nonManualJobs,
		JobStats:  make(map[string]JobStats),
		StartTime: time.Now(),
	}

	return &ExecutionContext{
		Config: config,
		Stats:  stats,
		Colors: colors,
		Writer: writer,
	}
}

func (c *Executor) createColorScheme(noColor bool) *ColorScheme {
	if noColor {
		// Return colorless scheme
		return &ColorScheme{
			Blue:    color.New(),
			Green:   color.New(),
			Red:     color.New(),
			Yellow:  color.New(),
			Cyan:    color.New(),
			Magenta: color.New(),
			Gray:    color.New(),
			White:   color.New(),
		}
	}

	return &ColorScheme{
		Blue:    color.New(color.FgBlue, color.Bold),
		Green:   color.New(color.FgGreen),
		Red:     color.New(color.FgRed),
		Yellow:  color.New(color.FgYellow, color.Bold),
		Cyan:    color.New(color.FgCyan, color.Bold),
		Magenta: color.New(color.FgMagenta),
		Gray:    color.New(color.FgHiBlack),
		White:   color.New(color.FgWhite),
	}
}

func (c *Executor) printHeader(ctx *ExecutionContext) {
	if ctx.Config.Short {
		_, _ = fmt.Fprintf(ctx.Writer, "Running %s\n", c.Runtime.Metadata.Filename)
		return
	}

	_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "\n🚀 Running %s\n\n", c.Runtime.Metadata.Filename)
}

func (c *Executor) printVariables(ctx *ExecutionContext) {
	if len(c.Runtime.Vars) == 0 {
		return
	}

	if ctx.Config.Short {
		_, _ = fmt.Fprintf(ctx.Writer, "Variables: ")
		for i, v := range c.Runtime.Vars {
			if i > 0 {
				_, _ = fmt.Fprint(ctx.Writer, ", ")
			}
			_, valueStr := v.Format()
			_, _ = fmt.Fprintf(ctx.Writer, "%s=%s", v.Name, valueStr)
		}
		_, _ = fmt.Fprintln(ctx.Writer)
		return
	}

	_, _ = ctx.Colors.Gray.Fprint(ctx.Writer, "Variables: ")
	for i, v := range c.Runtime.Vars {
		if i > 0 {
			_, _ = ctx.Colors.Gray.Fprint(ctx.Writer, ", ")
		}
		_, valueStr := v.Format()
		_, _ = ctx.Colors.Gray.Fprintf(ctx.Writer, "%s=%s", v.Name, valueStr)
	}
	_, _ = fmt.Fprintln(ctx.Writer)
	_, _ = fmt.Fprintln(ctx.Writer)
}

func (c *Executor) executeDryRun(ctx *ExecutionContext) {
	_, _ = fmt.Fprintln(ctx.Writer, "DRY RUN - Commands that would be executed:")
	_, _ = fmt.Fprintln(ctx.Writer)

	for _, task := range c.Runtime.Jobs {
		if ctx.Config.Verbose && len(task.Description) > 0 {
			_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "📦 Task: %s (%s)\n", task.Name, task.Description)
		} else {
			_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "📦 Task: %s\n", task.Name)
		}

		var forValues []interface{}
		var forEnabled bool

		// Check for 'for' directive
		for _, directive := range task.Directives {
			if directive.Type == "for" && len(directive.Params) > 0 {
				varName := fmt.Sprint(directive.Params[0])
				for _, v := range c.Runtime.Vars {
					if v.Name == varName {
						if listVal, ok := v.Value.(variable.ListValue); ok {
							forValues = listVal.Value
							forEnabled = true
							break
						}
					}
				}
			}
		}

		if forEnabled {
			_, _ = ctx.Colors.Cyan.Fprintf(ctx.Writer, "   🔄 Would loop over: %v\n", forValues)
			for _, val := range forValues {
				_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "   → Item: %v\n", val)
				c.printJobCommandsDryRun(task, val, ctx)
			}
		} else {
			c.printJobCommandsDryRun(task, nil, ctx)
		}
		_, _ = fmt.Fprintln(ctx.Writer)
	}
}

func (c *Executor) printJobCommandsDryRun(task *Task, forValue interface{}, ctx *ExecutionContext) {
	for _, commandBlock := range task.Commands {
		cmd := commandBlock.GetResult(forValue, c.getForIndex(task, forValue))
		displayCmd := fmt.Sprint(cmd)

		for _, function := range commandBlock.Functions {
			function.SetLoopData(forValue, c.getForIndex(task, forValue))
			_, err := function.Call()
			if err != nil {
				if !ctx.Config.Silent {
					throwMsg := fmt.Sprintf("Error in function '%s': %v\n", function.Type, err)
					c.Reporter.ThrowSyntaxError(throwMsg, &function.Metadata)
				}
				continue
			}

			if strings.HasPrefix(strings.TrimSpace(commandBlock.RawText), "@") {
				switch function.Type {
				case "trigger":
					if len(function.GetCalculatedArgs()) > 0 {
						jobName := fmt.Sprint(function.GetCalculatedArgsByIndex(0))
						displayCmd = fmt.Sprintf("would trigger task '%s'", jobName)
					}
				case "print", "println", "printf", "sprintf":
					if result := function.GetResult(); result != nil {
						displayCmd = strings.TrimRight(fmt.Sprint(result), "\n")
					}
				default:
					displayCmd = commandBlock.RawText
				}
			}
		}

		if strings.HasPrefix(displayCmd, "would trigger") {
			_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "     🚀 %s\n", displayCmd)
		} else {
			_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "     ▶ %s\n", displayCmd)
		}
	}
}

func (c *Executor) executeSequential(ctx *ExecutionContext) {
	for _, task := range c.Runtime.Jobs {
		if task.IsManual {
			continue
		}
		jobStats := c.executeJob(task, ctx)
		ctx.Stats.TotalCommands += jobStats.Total
		ctx.Stats.SuccessfulCommands += jobStats.Successful
		ctx.Stats.FailedCommands += jobStats.Failed
		ctx.Stats.JobStats[task.Name] = jobStats
	}
}

func (c *Executor) executeJob(task *Task, ctx *ExecutionContext) JobStats {
	startTime := time.Now()
	if ctx.Config.Verbose {
		if len(task.Description) > 0 {
			_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "📦 Task: %s (%s)\n", task.Name, task.Description)
		} else {
			_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "📦 Task: %s\n", task.Name)
		}
	} else if !ctx.Config.Silent && !ctx.Config.Short {
		if ctx.Config.Verbose && len(task.Description) > 0 {
			_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "📦 Task: %s (%s)\n", task.Name, task.Description)
		} else {
			_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "📦 Task: %s\n", task.Name)
		}
	}

	stats := JobStats{}

	var forValues []interface{}
	var forEnabled bool

	// Check for 'for' directive
	for _, directive := range task.Directives {
		if directive.Type == "for" && len(directive.Params) > 0 {
			varName := fmt.Sprint(directive.Params[0])
			for _, v := range c.Runtime.Vars {
				if v.Name == varName {
					if listVal, ok := v.Value.(variable.ListValue); ok {
						forValues = listVal.Value
						forEnabled = true
						break
					}
				}
			}
		}
	}

	if forEnabled {
		if !ctx.Config.Silent && ctx.Config.Verbose {
			_, _ = ctx.Colors.Cyan.Fprintf(ctx.Writer, "   🔄 Looping over: %v\n", forValues)
		}

		for _, val := range forValues {
			if !ctx.Config.Silent && ctx.Config.Verbose {
				_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "   → Item: %v\n", val)
			}
			jobStats := c.RunJobCommands(task, val, ctx)
			stats.Total += jobStats.Total
			stats.Successful += jobStats.Successful
			stats.Failed += jobStats.Failed
		}
	} else {
		stats = c.RunJobCommands(task, nil, ctx)
	}

	stats.Duration = time.Since(startTime)

	if !ctx.Config.Silent {
		_, _ = fmt.Fprintln(ctx.Writer)
	}

	return stats
}

func (c *Executor) RunJobCommands(task *Task, forValue interface{}, ctx *ExecutionContext) JobStats {
	stats := JobStats{}
	forIndex := c.getForIndex(task, forValue)

	for _, commandBlock := range task.Commands {
		stats.Total++

		// Get the command result using GetResult
		cmd := commandBlock.GetResult(forValue, forIndex)
		functionOutput := ""
		hasFunction := false
		isPrintFunction := false
		isTriggerFunction := false
		var funcRaw interface{}

		// Process functions
		for _, function := range commandBlock.Functions {
			var err error
			function.SetLoopData(forValue, c.getForIndex(task, forValue))
			funcRaw, err = function.Call()
			hasFunction = true

			if err != nil {
				stats.Failed++
				if !ctx.Config.Silent {
					c.Reporter.ThrowSyntaxError(fmt.Sprintf("Error in function '@%s': \n\t%v\n", function.Type, err), &function.Metadata)
				}
				continue
			}

			// Handle different function types
			switch function.Type {
			case "trigger":
				isTriggerFunction = true
				if result := function.GetResult(); result != nil {
					functionOutput = strings.TrimRight(fmt.Sprint(result), "\n")
					if !ctx.Config.Silent {
						// Display the actual output from the triggered task
						lines := strings.Split(functionOutput, "\n")
						for _, line := range lines {
							if strings.TrimSpace(line) != "" {
								if ctx.Config.Short {
									_, _ = fmt.Fprintf(ctx.Writer, "► %s\n", line)
								} else {
									// Find the triggered task to get its description
									triggeredJobName := fmt.Sprint(function.GetCalculatedArgsByIndex(0))
									triggeredJob := c.Runtime.GetTaskByID(triggeredJobName)

									if ctx.Config.Verbose && triggeredJob != nil && len(triggeredJob.Description) > 0 {
										_, _ = ctx.Colors.Cyan.Fprintf(ctx.Writer, "     ▶ Trigging the %q task (%s):\n", triggeredJobName, triggeredJob.Description)
									} else {
										_, _ = ctx.Colors.Cyan.Fprintf(ctx.Writer, "     ▶ Trigging the %q task:\n", triggeredJobName)
									}
									_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "       ▶ %s\n", line)
								}
							}
						}
					}
				}
				stats.Successful++
			case "print", "println", "printf", "sprintf":
				isPrintFunction = true
				if result := function.GetResult(); result != nil {
					functionOutput = strings.TrimRight(fmt.Sprint(result), "\n")
					if !ctx.Config.Silent {
						if ctx.Config.Short {
							_, _ = fmt.Fprintf(ctx.Writer, "► %s\n", functionOutput)
						} else {
							_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "     ▶ %s\n", functionOutput)
							if ctx.Config.Verbose {
								_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "     Output: \n")
								_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "             %s\n", functionOutput)
							}
						}
					}
				}
			default:
				// Handle non-print function-only commands
				if strings.HasPrefix(strings.TrimSpace(commandBlock.RawText), "@") {
					if !ctx.Config.Silent {
						if ctx.Config.Short {
							_, _ = fmt.Fprintf(ctx.Writer, "► %s ✓\n", commandBlock.RawText)
						} else {
							_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "     ▶ %s ✓\n", strings.TrimSpace(commandBlock.RawText))
							if ctx.Config.Verbose {
								if result := function.GetResult(); result != nil {
									_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "     Output: \n")
									_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "             %s\n", strings.TrimRight(fmt.Sprint(result), "\n"))
								}
							}
						}
					}
				}
			}

			funcRawStr := fmt.Sprint(funcRaw)
			switch res := function.GetResult().(type) {
			case string:
				cmd = strings.ReplaceAll(fmt.Sprint(cmd), funcRawStr, res)
			case []interface{}:
				strValues := make([]string, len(res))
				for i, item := range res {
					strValues[i] = fmt.Sprint(item)
				}
				cmd = strings.ReplaceAll(fmt.Sprint(cmd), funcRawStr, strings.Join(strValues, " "))
			default:
				cmd = strings.ReplaceAll(fmt.Sprint(cmd), funcRawStr, fmt.Sprintf("%v", res))
			}
		}

		// Increment stats for successful function-only commands
		if hasFunction && !isPrintFunction && !isTriggerFunction && strings.HasPrefix(strings.TrimSpace(commandBlock.RawText), "@") {
			stats.Successful++
		}

		// Use GetResult as fallback if no functions were processed
		if !hasFunction && cmd != "" {
			functionOutput = strings.TrimRight(fmt.Sprint(cmd), "\n")
		}

		// Execute shell command (skip if it was a trigger function since we already handled it)
		cmdStr := strings.TrimSpace(fmt.Sprint(cmd))
		if cmdStr != "" && !strings.HasPrefix(strings.TrimSpace(commandBlock.RawText), "@") && !isTriggerFunction {
			output, success := c.executeShellCommand(cmdStr, forValue, ctx)
			if success {
				if len(output) > 0 && !ctx.Config.Silent {
					lines := strings.Split(strings.TrimRight(string(output), "\n"), "\n")
					for _, line := range lines {
						if ctx.Config.Short {
							_, _ = fmt.Fprintf(ctx.Writer, "► %s\n", line)
						} else {
							_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "     ▶ %s\n", line)
							if ctx.Config.Verbose {
								_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "     Output: \n")
								_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "             %s\n", line)
							}
						}
					}
				}
				stats.Successful++
			} else {
				stats.Failed++
			}
		} else if isPrintFunction {
			// Print functions already handled above
			stats.Successful++
		}
	}

	return stats
}

func (c *Executor) displayCommand(commandBlock Command, functionOutput string, forValue interface{}, forIndex int, ctx *ExecutionContext) {
	// Type assertion to get RawText - you'll need to adjust this based on your CommandBlock type
	rawText := strings.TrimSpace(commandBlock.GetResult(forValue, forIndex))

	if ctx.Config.Short {
		if strings.HasPrefix(strings.TrimSpace(rawText), "@") && functionOutput != "" {
			_, _ = fmt.Fprintf(ctx.Writer, "► %s\n", functionOutput)
		} else {
			_, _ = fmt.Fprintf(ctx.Writer, "► %s\n", rawText)
		}
		return
	}

	if strings.HasPrefix(strings.TrimSpace(rawText), "@") && functionOutput != "" {
		_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "     ▶ %s\n", functionOutput)
		if ctx.Config.Verbose {
			_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "     Output: \n")
			_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "             %s\n", functionOutput)
		}
	} else {
		_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "     ▶ %s\n", rawText)
	}
}

func (c *Executor) getForIndex(task *Task, forValue interface{}) int {
	for _, directive := range task.Directives {
		if directive.Type == "for" && len(directive.Params) > 0 {
			varName := fmt.Sprint(directive.Params[0])
			for _, v := range c.Runtime.Vars {
				if v.Name == varName {
					if listVal, ok := v.Value.(variable.ListValue); ok {
						for idx, val := range listVal.Value {
							if val == forValue {
								return idx
							}
						}
					}
				}
			}
		}
	}
	return -1
}

func (c *Executor) executeShellCommand(command string, forValue interface{}, ctx *ExecutionContext) ([]byte, bool) {
	// Replace variable
	for _, v := range c.Runtime.Vars {
		placeholder := fmt.Sprintf("$%s", v.Name)
		switch val := v.Value.(type) {
		case variable.StringValue:
			command = strings.ReplaceAll(command, placeholder, val.Value)
		case variable.ListValue:
			strValues := make([]string, len(val.Value))
			for i, item := range val.Value {
				strValues[i] = fmt.Sprint(item)
			}
			command = strings.ReplaceAll(command, placeholder, strings.Join(strValues, " "))
		default:
			command = strings.ReplaceAll(command, placeholder, fmt.Sprint(val))
		}
	}

	// Replace @value
	if forValue != nil {
		valueStr := fmt.Sprint(forValue)
		sanitizedValue := strings.ReplaceAll(valueStr, ":", "-")
		command = strings.ReplaceAll(command, "@value", sanitizedValue)
	}

	// Execute with timeout if specified
	var cmd *exec.Cmd
	if ctx.Config.Timeout > 0 {
		// Implementation would require context.WithTimeout
		cmd = exec.Command("timeout", fmt.Sprintf("%d", ctx.Config.Timeout), "sh", "-c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	output, err := cmd.CombinedOutput()

	if ctx.Config.Silent {
		return output, err == nil
	}

	if err != nil {
		if !ctx.Config.Short {
			_, _ = ctx.Colors.Red.Fprintf(ctx.Writer, "       ❌ Error: %v\n", err)
			if len(output) > 0 && ctx.Config.Verbose {
				_, _ = ctx.Colors.Red.Fprintf(ctx.Writer, "       Output: %s", string(output))
			}
		}
		return output, false
	}

	return output, true
}

func (c *Executor) printSummary(ctx *ExecutionContext) {
	if ctx.Config.Short {
		c.printShortSummary(ctx)
		return
	}

	_, _ = fmt.Fprintln(ctx.Writer, "─────────────────────────────────────")

	_, _ = ctx.Colors.Cyan.Fprintln(ctx.Writer, "📊 Execution Summary:")
	_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "   Jobs executed: %d\n", ctx.Stats.TotalJobs)
	_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "   Total commands: %d\n", ctx.Stats.TotalCommands)

	if ctx.Stats.SuccessfulCommands == ctx.Stats.TotalCommands {
		_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "   ✅ All %d commands succeeded\n", ctx.Stats.SuccessfulCommands)
	} else {
		_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "   ✅ Successful: %d\n", ctx.Stats.SuccessfulCommands)
		_, _ = ctx.Colors.Red.Fprintf(ctx.Writer, "   ❌ Failed: %d\n", ctx.Stats.FailedCommands)
	}

	// Show timing information if profiling is enabled
	if ctx.Config.Profile {
		duration := ctx.Stats.EndTime.Sub(ctx.Stats.StartTime)
		_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "   ⏱️  Total execution time: %v\n", duration)

		if ctx.Config.Verbose {
			_, _ = ctx.Colors.White.Fprintln(ctx.Writer, "   📊 Task breakdown:")
			for jobName, jobStats := range ctx.Stats.JobStats {
				_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "     %s: %v (✅%d ❌%d)\n",
					jobName, jobStats.Duration, jobStats.Successful, jobStats.Failed)
			}
		}
	}

	_, _ = fmt.Fprintln(ctx.Writer, "─────────────────────────────────────")

	if ctx.Stats.SuccessfulCommands == ctx.Stats.TotalCommands {
		_, _ = ctx.Colors.Green.Fprintln(ctx.Writer, "Execution completed successfully!")
	} else {
		_, _ = ctx.Colors.Red.Fprintln(ctx.Writer, "Execution completed with errors!")
	}
}

func (c *Executor) printShortSummary(ctx *ExecutionContext) {
	if ctx.Stats.SuccessfulCommands == ctx.Stats.TotalCommands {
		_, _ = fmt.Fprintf(ctx.Writer, "✅ %d/%d commands succeeded", ctx.Stats.SuccessfulCommands, ctx.Stats.TotalCommands)
	} else {
		_, _ = fmt.Fprintf(ctx.Writer, "  %d/%d commands succeeded", ctx.Stats.SuccessfulCommands, ctx.Stats.TotalCommands)
	}

	if ctx.Config.Profile {
		duration := ctx.Stats.EndTime.Sub(ctx.Stats.StartTime)
		_, _ = fmt.Fprintf(ctx.Writer, " (%v)", duration)
	}

	_, _ = fmt.Fprintln(ctx.Writer)
}
