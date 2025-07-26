package core

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Core struct {
	Filename string
	Content  string
	Vars     []Var
	Jobs     []Job
	mu       sync.Mutex
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

type JobStats struct {
	Total      int
	Successful int
	Failed     int
	Duration   time.Duration
}

type ExecutionContext struct {
	Config *ExecutionConfig
	Stats  *ExecutionStats
	Colors *ColorScheme
	Writer io.Writer
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

type Metadata struct {
	Filename string
	Line     int
	Column   int
}

func NewCore(filename string, content string) *Core {
	return &Core{
		Filename: filename,
		Content:  content,
		Vars:     make([]Var, 0),
		Jobs:     make([]Job, 0),
	}
}

func (c *Core) AddVar(v Var) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Vars = append(c.Vars, v)
}

func (c *Core) AddJob(j Job) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Jobs = append(c.Jobs, j)
}

func (c *Core) Exec() {
	config := &ExecutionConfig{
		SummaryOnly: false,
		Verbose:     false,
		Silent:      false,
		Short:       false,
		DryRun:      false,
		NoColor:     false,
		Profile:     false,
		LogLevel:    "info",
		MaxWorkers:  1,
		Timeout:     0,
	}
	c.ExecWithConfig(config)
}

func (c *Core) ExecWithConfig(config *ExecutionConfig) {
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

func (c *Core) createExecutionContext(config *ExecutionConfig) *ExecutionContext {
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

	stats := &ExecutionStats{
		TotalJobs: len(c.Jobs),
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

func (c *Core) createColorScheme(noColor bool) *ColorScheme {
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

func (c *Core) printHeader(ctx *ExecutionContext) {
	if ctx.Config.Short {
		_, _ = fmt.Fprintf(ctx.Writer, "Running %s\n", c.Filename)
		return
	}

	_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "\n🚀 Running %s\n\n", c.Filename)
}

func (c *Core) printVariables(ctx *ExecutionContext) {
	if len(c.Vars) == 0 {
		return
	}

	if ctx.Config.Short {
		_, _ = fmt.Fprintf(ctx.Writer, "Variables: ")
		for i, v := range c.Vars {
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
	for i, v := range c.Vars {
		if i > 0 {
			_, _ = ctx.Colors.Gray.Fprint(ctx.Writer, ", ")
		}
		_, valueStr := v.Format()
		_, _ = ctx.Colors.Gray.Fprintf(ctx.Writer, "%s=%s", v.Name, valueStr)
	}
	_, _ = fmt.Fprintln(ctx.Writer)
	_, _ = fmt.Fprintln(ctx.Writer)
}

func (c *Core) executeDryRun(ctx *ExecutionContext) {
	_, _ = fmt.Fprintln(ctx.Writer, "DRY RUN - Commands that would be executed:")
	_, _ = fmt.Fprintln(ctx.Writer)

	for _, job := range c.Jobs {
		_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "📦 Job: %s\n", job.Name)

		var forValues []interface{}
		var forEnabled bool

		// Check for 'for' directive
		for _, directive := range job.Directives {
			if directive.Type == "for" && len(directive.Params) > 0 {
				varName := fmt.Sprint(directive.Params[0])
				for _, v := range c.Vars {
					if v.Name == varName {
						if listVal, ok := v.Value.(ListValue); ok {
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
				c.printJobCommandsDryRun(&job, val, ctx)
			}
		} else {
			c.printJobCommandsDryRun(&job, nil, ctx)
		}
		_, _ = fmt.Fprintln(ctx.Writer)
	}
}

func (c *Core) printJobCommandsDryRun(job *Job, forValue interface{}, ctx *ExecutionContext) {
	for _, commandBlock := range job.Commands {
		cmd := commandBlock.GetResult(forValue, c.getForIndex(job, forValue))
		displayCmd := fmt.Sprint(cmd)

		for _, function := range commandBlock.Functions {
			function.forValues = forValue
			function.forIndex = c.getForIndex(job, forValue)
			_, err := function.call()
			if err != nil {
				if !ctx.Config.Silent {
					ThrowSyntaxError(fmt.Sprintf("Error in function '%s': %v\n", function.Type, err), function.Metadata.Filename, function.Metadata.Line, function.Metadata.Column)
				}
				continue
			}

			if strings.HasPrefix(strings.TrimSpace(commandBlock.RawText), "@") {
				switch function.Type {
				case "print", "println", "printf", "sprintf":
					if result := function.GetResult(); result != nil {
						displayCmd = strings.TrimRight(fmt.Sprint(result), "\n")
					}
				default:
					displayCmd = commandBlock.RawText
				}
			}
		}

		_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "     ▶ %s\n", displayCmd)
	}
}

func (c *Core) executeSequential(ctx *ExecutionContext) {
	for _, job := range c.Jobs {
		jobStats := c.executeJob(job, ctx)
		ctx.Stats.TotalCommands += jobStats.Total
		ctx.Stats.SuccessfulCommands += jobStats.Successful
		ctx.Stats.FailedCommands += jobStats.Failed
		ctx.Stats.JobStats[job.Name] = jobStats
	}
}

func (c *Core) executeJob(job Job, ctx *ExecutionContext) JobStats {
	startTime := time.Now()

	if !ctx.Config.Silent && !ctx.Config.Short {
		_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "📦 Job: %s\n", job.Name)
	} else if ctx.Config.Verbose {
		_, _ = fmt.Fprintf(ctx.Writer, "Executing job: %s\n", job.Name)
	}

	stats := JobStats{}

	var forValues []interface{}
	var forEnabled bool

	// Check for 'for' directive
	for _, directive := range job.Directives {
		if directive.Type == "for" && len(directive.Params) > 0 {
			varName := fmt.Sprint(directive.Params[0])
			for _, v := range c.Vars {
				if v.Name == varName {
					if listVal, ok := v.Value.(ListValue); ok {
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
			jobStats := c.RunJobCommands(&job, val, ctx)
			stats.Total += jobStats.Total
			stats.Successful += jobStats.Successful
			stats.Failed += jobStats.Failed
		}
	} else {
		stats = c.RunJobCommands(&job, nil, ctx)
	}

	stats.Duration = time.Since(startTime)

	if !ctx.Config.Silent {
		_, _ = fmt.Fprintln(ctx.Writer)
	}

	return stats
}

func (c *Core) RunJobCommands(job *Job, forValue interface{}, ctx *ExecutionContext) JobStats {
	stats := JobStats{}
	forIndex := c.getForIndex(job, forValue)

	for _, commandBlock := range job.Commands {
		stats.Total++

		// Get the command result using GetResult
		cmd := commandBlock.GetResult(forValue, forIndex)
		functionOutput := ""
		hasFunction := false
		isPrintFunction := false
		var funcRaw interface{}

		// Process functions
		for _, function := range commandBlock.Functions {
			function.forValues = forValue
			function.forIndex = c.getForIndex(job, forValue)

			var err error
			funcRaw, err = function.call()
			hasFunction = true

			if err != nil {
				stats.Failed++
				if !ctx.Config.Silent {
					ThrowSyntaxError(fmt.Sprintf("Error in function '@%s': \n\t%v\n", function.Type, err), function.Metadata.Filename, function.Metadata.Line, function.Metadata.Column)
				}
				continue
			}

			// Handle print functions
			switch function.Type {
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
		if hasFunction && !isPrintFunction && strings.HasPrefix(strings.TrimSpace(commandBlock.RawText), "@") {
			stats.Successful++
		}

		// Use GetResult as fallback if no functions were processed
		if !hasFunction && cmd != "" {
			functionOutput = strings.TrimRight(fmt.Sprint(cmd), "\n")
		}

		// Execute shell command
		cmdStr := strings.TrimSpace(fmt.Sprint(cmd))
		if cmdStr != "" && !strings.HasPrefix(strings.TrimSpace(commandBlock.RawText), "@") {
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

func (c *Core) displayCommand(commandBlock Command, functionOutput string, forValue interface{}, forIndex int, ctx *ExecutionContext) {
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

func (c *Core) getForIndex(job *Job, forValue interface{}) int {
	for _, directive := range job.Directives {
		if directive.Type == "for" && len(directive.Params) > 0 {
			varName := fmt.Sprint(directive.Params[0])
			for _, v := range c.Vars {
				if v.Name == varName {
					if listVal, ok := v.Value.(ListValue); ok {
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

func (c *Core) executeShellCommand(command string, forValue interface{}, ctx *ExecutionContext) ([]byte, bool) {
	// Replace variables
	for _, v := range c.Vars {
		placeholder := fmt.Sprintf("$%s", v.Name)
		switch val := v.Value.(type) {
		case StringValue:
			command = strings.ReplaceAll(command, placeholder, val.Value)
		case ListValue:
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

func (c *Core) printSummary(ctx *ExecutionContext) {
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
			_, _ = ctx.Colors.White.Fprintln(ctx.Writer, "   📊 Job breakdown:")
			for jobName, jobStats := range ctx.Stats.JobStats {
				_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "     %s: %v (✅%d ❌%d)\n",
					jobName, jobStats.Duration, jobStats.Successful, jobStats.Failed)
			}
		}
	}

	_, _ = fmt.Fprintln(ctx.Writer, "─────────────────────────────────────")

	if ctx.Stats.SuccessfulCommands == ctx.Stats.TotalCommands {
		_, _ = ctx.Colors.Green.Fprintln(ctx.Writer, "Execution completed successfully!\n")
	} else {
		_, _ = ctx.Colors.Red.Fprintln(ctx.Writer, "Execution completed with errors!\n")
	}
}

func (c *Core) printShortSummary(ctx *ExecutionContext) {
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

func (c *Core) Describe() {
	c.mu.Lock()
	defer c.mu.Unlock()

	blue := color.New(color.FgBlue, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	white := color.New(color.FgWhite)

	_, _ = blue.Printf("\n📋 %s Summary\n\n", c.Filename)

	// Variables
	_, _ = yellow.Println("🔸 Variables:")
	if len(c.Vars) == 0 {
		_, _ = white.Println("  None")
	} else {
		for _, v := range c.Vars {
			typeStr, valueStr := v.Format()
			_, _ = white.Printf("  %s (%s) = %s\n", v.Name, typeStr, valueStr)
		}
	}

	// Jobs
	fmt.Println()
	_, _ = green.Println("🔸 Jobs:")
	if len(c.Jobs) == 0 {
		_, _ = white.Println("  None")
	} else {
		for _, j := range c.Jobs {
			_, _ = white.Printf("  %s:\n", j.Name)

			// Directives
			if len(j.Directives) > 0 {
				_, _ = white.Print("    Directives: ")
				for i, dir := range j.Directives {
					if i > 0 {
						_, _ = white.Print(", ")
					}
					args := formatParams(dir.Params)
					_, _ = white.Printf("@%s%s", dir.Type, args)
				}
				fmt.Println()
			}

			// Commands
			_, _ = white.Printf("    Commands (%d):\n", len(j.Commands))
			for i, command := range j.Commands {
				_, _ = white.Printf("      %d. %s\n", i+1, command.RawText)
			}
			fmt.Println()
		}
	}
}

func formatParams(params []interface{}) string {
	if len(params) == 0 {
		return ""
	}
	args := make([]string, len(params))
	for i, v := range params {
		args[i] = fmt.Sprint(v)
	}

	return fmt.Sprintf("(%s)", strings.Join(args, ", "))
}
