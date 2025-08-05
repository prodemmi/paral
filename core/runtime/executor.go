package runtime

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"paral/core/variable"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type TaskExecutor struct {
	Runtime         *Runtime
	Reporter        *Reporter
	CommandExecutor *CommandExecutor
}

type TaskOutput struct {
	TaskName    string
	Description string
	OutputLines []string
	IsVerbose   bool
}

type ThreadSafeWriter struct {
	writer io.Writer
	mutex  sync.Mutex
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
	Writer      io.Writer
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

type CommandExecutor struct {
	Runtime  *Runtime
	Reporter *Reporter
}

func NewExecutor(runtime *Runtime) *TaskExecutor {
	return &TaskExecutor{
		Runtime:         runtime,
		Reporter:        runtime.Reporter,
		CommandExecutor: NewCommandExecutor(runtime),
	}
}

func NewCommandExecutor(runtime *Runtime) *CommandExecutor {
	return &CommandExecutor{
		Runtime:  runtime,
		Reporter: runtime.Reporter,
	}
}

func NewThreadSafeWriter(w io.Writer) *ThreadSafeWriter {
	return &ThreadSafeWriter{writer: w}
}

func (t *ThreadSafeWriter) Write(p []byte) (n int, err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.writer.Write(p)
}

func (ce *CommandExecutor) ExecuteShellCommand(pipeline string, forValue interface{}, ctx *ExecutionContext) ([]byte, bool) {
	for _, v := range ce.Runtime.Vars {
		placeholder := fmt.Sprintf("$%s", v.Name)
		switch val := v.Value.(type) {
		case variable.StringValue:
			pipeline = strings.ReplaceAll(pipeline, placeholder, val.Value)
		case variable.ListValue:
			strValues := make([]string, len(val.Value))
			for i, item := range val.Value {
				strValues[i] = fmt.Sprint(item)
			}
			pipeline = strings.ReplaceAll(pipeline, placeholder, strings.Join(strValues, " "))
		default:
			pipeline = strings.ReplaceAll(pipeline, placeholder, fmt.Sprint(val))
		}
	}
	valueStr := fmt.Sprint(forValue)
	pipeline = strings.ReplaceAll(pipeline, "@value", valueStr)
	var cmd *exec.Cmd
	if ctx.Config.Timeout > 0 {
		cmd = exec.Command("timeout", fmt.Sprintf("%d", ctx.Config.Timeout), "sh", "-c", pipeline)
	} else {
		cmd = exec.Command("sh", "-c", pipeline)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		ce.Reporter.Warn(fmt.Sprintf("Command failed: %s, error: %v", pipeline, err), nil)
		return output, false
	}
	return output, true
}

func (c *TaskExecutor) ExecWithConfig(config *ExecutionConfig) {
	ctx := c.createExecutionContext(config)
	if config.Profile {
		ctx.Stats.StartTime = time.Now()
	}
	if !config.Silent {
		c.printVariables(ctx)
	}
	if config.DryRun {
		_, _ = fmt.Fprintln(ctx.Writer, "DRY RUN - Commands that would be executed:")
		c.executeDryRun(ctx)
	} else {
		c.executeParallel(ctx)
	}
	if config.Profile {
		ctx.Stats.EndTime = time.Now()
	}
	if !config.Silent {
		c.printSummary(ctx)
	}
}

func (c *TaskExecutor) createExecutionContext(config *ExecutionConfig) *ExecutionContext {
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
	safeWriter := NewThreadSafeWriter(writer)
	colors := c.createColorScheme(config.NoColor)
	nonManualJobs := 0
	for _, task := range c.Runtime.Tasks {
		if !task.IsManual() {
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
		Writer: safeWriter,
	}
}

func (c *TaskExecutor) createColorScheme(noColor bool) *ColorScheme {
	if noColor {
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

func (c *TaskExecutor) executeParallel(ctx *ExecutionContext) {
	independentTasks, deferTasks, err := c.Runtime.GetExecutionOrder()
	if err != nil {
		c.Reporter.ThrowRuntimeError(err.Error(), c.Runtime.Metadata)
		return
	}
	outputChan := make(chan TaskOutput, len(independentTasks)+len(deferTasks))
	var printWg sync.WaitGroup
	if !ctx.Config.Silent {
		printWg.Add(1)
		go func() {
			defer printWg.Done()
			c.processOutput(outputChan, ctx)
		}()
	}
	if onlyIndependentTasksWithoutDeps(independentTasks) {
		c.executeTasksInParallel(independentTasks, ctx, outputChan)
	} else {
		c.executeTasksWithDependencies(independentTasks, ctx, outputChan)
	}
	if len(deferTasks) > 0 {
		c.executeTasksInParallel(deferTasks, ctx, outputChan)
	}
	close(outputChan)
	if !ctx.Config.Silent {
		printWg.Wait()
	}
}

func (c *TaskExecutor) processOutput(outputChan chan TaskOutput, ctx *ExecutionContext) {
	for output := range outputChan {
		for _, line := range output.OutputLines {
			if ctx.Config.Silent {
				continue
			}
			if output.IsVerbose {
				c.printVerboseLine(line, ctx)
			} else if !ctx.Config.Short {
				_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "%s\n", line)
			} else {
				_, _ = fmt.Fprintf(ctx.Writer, "%s\n", line)
			}
		}
	}
	_, _ = fmt.Fprintln(ctx.Writer)
}

func (c *TaskExecutor) printVerboseLine(line string, ctx *ExecutionContext) {
	switch {
	case strings.HasPrefix(line, "📦 Task"):
		_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "%s\n", line)
	case strings.HasPrefix(line, "  🔄 Looping"):
		_, _ = ctx.Colors.Cyan.Fprintf(ctx.Writer, "%s\n", line)
	case strings.HasPrefix(line, "    → Item"):
		_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "%s\n", line)
	case strings.Contains(line, "▶"):
		_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "%s\n", line)
	case strings.Contains(line, "🚀"):
		_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "%s\n", line)
	case strings.Contains(line, "❌"):
		_, _ = ctx.Colors.Red.Fprintf(ctx.Writer, "%s\n", line)
	case strings.Contains(line, "stash"):
		_, _ = ctx.Colors.Cyan.Fprintf(ctx.Writer, "%s\n", line)
	default:
		_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "%s\n", line)
	}
}

func (c *TaskExecutor) printVariables(ctx *ExecutionContext) {
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
	_, _ = ctx.Colors.Magenta.Fprint(ctx.Writer, "Variables: ")
	for i, v := range c.Runtime.Vars {
		if i > 0 {
			_, _ = ctx.Colors.Magenta.Fprint(ctx.Writer, ", ")
		}
		_, valueStr := v.Format()
		_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "%s=%s", v.Name, valueStr)
	}
	_, _ = fmt.Fprintln(ctx.Writer)
}

func (c *TaskExecutor) executeDryRun(ctx *ExecutionContext) {
	independentTasks, deferTasks, err := c.Runtime.GetExecutionOrder()
	if err != nil {
		_, _ = ctx.Colors.Red.Fprintf(ctx.Writer, "❌ Dependency error: %v\n", err)
		return
	}
	_, _ = ctx.Colors.Cyan.Fprintln(ctx.Writer, "🔄 Tasks with dependencies (execution order):")
	for _, task := range independentTasks {
		c.printTaskDryRun(task, ctx)
	}
	if len(deferTasks) > 0 {
		_, _ = ctx.Colors.Cyan.Fprintln(ctx.Writer, "⏳ Tasks waiting for all others to complete:")
		for _, task := range deferTasks {
			c.printTaskDryRun(task, ctx)
		}
	}
}

func (c *TaskExecutor) printTaskDryRun(task *Task, ctx *ExecutionContext) {
	taskId := task.GetTaskId()
	dependencies := c.Runtime.GetTaskDependencies(task)
	if ctx.Config.Verbose && len(task.Description) > 0 {
		_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "📦 Task: %s (%s)\n", taskId, task.Description)
	} else {
		_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "📦 Task: %s\n", taskId)
	}
	if len(dependencies) > 0 {
		_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "  📋 Depends on: %v\n", dependencies)
	}
	if c.Runtime.HasDeferDirective(task) {
		_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "  ⏳ Waits for all other tasks to complete\n")
	}
	var forValues interface{}
	var forEnabled bool
	var isMatrix bool
	for _, directive := range task.Directives {
		if directive.Type == "for" && len(directive.Params) > 0 {
			varName := fmt.Sprint(directive.Params[0])
			for _, v := range c.Runtime.Vars {
				if v.Name == varName {
					if listVal, ok := v.Value.(variable.ListValue); ok {
						forValues = listVal.Value
						forEnabled = true
						isMatrix = false
						break
					}
					if _, ok := v.Value.(variable.MatrixValue); ok {
						forValues, _ = v.GetMatrixCombinations()
						forEnabled = true
						isMatrix = true
						break
					}
				}
			}
		}
	}
	if forEnabled {
		var displayValues []string
		if isMatrix {
			for _, val := range forValues.([][]interface{}) {
				displayValues = append(displayValues, fmt.Sprint(val))
			}
		} else {
			for _, val := range forValues.([]interface{}) {
				displayValues = append(displayValues, fmt.Sprint(val))
			}
		}
		_, _ = ctx.Colors.Cyan.Fprintf(ctx.Writer, "  🔄 Would loop over: %v\n", displayValues)
		if isMatrix {
			for key, val := range forValues.([][]interface{}) {
				_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "    → Item: %v\n", val)
				task.PushLoopStack(key, val)
				task.SetCurrentLoopIndex(key)
				c.printJobPipelinesDryRun(task, ctx)
			}
		} else {
			for key, val := range forValues.([]interface{}) {
				_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "    → Item: %v\n", val)
				task.PushLoopStack(key, val)
				task.SetCurrentLoopIndex(key)
				c.printJobPipelinesDryRun(task, ctx)
			}
		}
	} else {
		task.ClearLoopStack()
		c.printJobPipelinesDryRun(task, ctx)
	}
	_, _ = fmt.Fprintln(ctx.Writer)
}

func (c *TaskExecutor) printJobPipelinesDryRun(task *Task, ctx *ExecutionContext) {
	loopContext := task.GetActiveLoopContext()
	indent := "    "
	for _, pipeline := range task.Pipelines {
		var displayCmd string
		var isStash, isTrigger, isPrint bool

		if pipeline.Stash != nil {
			isStash = true
			if pipeline.Stash.Command != nil {
				result := pipeline.Stash.Command.GetRawResult(loopContext, ctx)
				displayCmd = fmt.Sprintf("stash %q saved: %s", pipeline.Stash.Name, strings.TrimRight(result, "\n"))
			} else {
				displayCmd = fmt.Sprintf("stash %q saved", pipeline.Stash.Name)
			}
		} else if pipeline.Function != nil {
			if loopContext != nil {
				pipeline.Function.SetLoopContext(loopContext)
			}
			result, err := pipeline.Function.Call()
			if err != nil {
				displayCmd = fmt.Sprintf("%s (%v)", pipeline.Function.Type, err)
				_, _ = ctx.Colors.Red.Fprintf(ctx.Writer, "%s❌ %s\n", indent, displayCmd)
				c.Reporter.ThrowSyntaxError(fmt.Sprintf("Error in function '%s': %v", pipeline.Function.Type, err), &pipeline.Function.Metadata)
				continue
			}
			if strings.TrimSpace(fmt.Sprint(result)) != "" {
				displayCmd = fmt.Sprint(result)
				switch pipeline.Function.Type {
				case "trigger":
					if len(pipeline.Function.Args) > 0 {
						jobName := fmt.Sprint(pipeline.Function.GetCalculatedArgsByIndex(0))
						displayCmd = fmt.Sprintf("trigger task '%s'", jobName)
						isTrigger = true
					}
				case "print", "println", "printf", "sprintf":
					displayCmd = strings.TrimRight(fmt.Sprint(result), "\n")
					isPrint = true
				default:
					displayCmd = strings.TrimRight(fmt.Sprint(result), "\n")
				}
			} else {
				continue // Skip empty output
			}
		} else if pipeline.Command != nil {
			displayCmd = pipeline.Command.GetRawResult(loopContext, ctx)
			if strings.TrimSpace(displayCmd) == "" {
				continue // Skip empty output
			}
		}

		shouldPrint := isPrint || (!isStash && !isTrigger && pipeline.Command != nil) || ctx.Config.Verbose
		if shouldPrint {
			if isTrigger {
				_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "%s🚀 %s\n", indent, displayCmd)
			} else if isStash {
				_, _ = ctx.Colors.Cyan.Fprintf(ctx.Writer, "%s▶ %s\n", indent, displayCmd)
			} else {
				_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "%s▶ %s\n", indent, displayCmd)
			}
		}
	}
}

func (c *TaskExecutor) executeTasksWithDependencies(tasks []*Task, ctx *ExecutionContext, outputChan chan TaskOutput) {
	if len(tasks) == 0 {
		return
	}
	completedTasks := make(map[string]chan bool)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, task := range tasks {
		taskId := task.GetTaskId()
		completedTasks[taskId] = make(chan bool, 1)
	}
	for _, task := range tasks {
		wg.Add(1)
		go func(t *Task) {
			defer wg.Done()
			taskId := t.GetTaskId()
			dependencies := c.Runtime.GetTaskDependencies(t)
			for _, depId := range dependencies {
				if ch, exists := completedTasks[depId]; exists {
					<-ch
				}
			}
			jobStats, taskOutput := c.executeJob(t, ctx)
			mutex.Lock()
			ctx.Stats.TotalCommands += jobStats.Total
			ctx.Stats.SuccessfulCommands += jobStats.Successful
			ctx.Stats.FailedCommands += jobStats.Failed
			ctx.Stats.JobStats[t.Name] = jobStats
			mutex.Unlock()
			if !ctx.Config.Silent {
				outputChan <- *taskOutput
			}
			completedTasks[taskId] <- true
			close(completedTasks[taskId])
		}(task)
	}
	wg.Wait()
}

func (c *TaskExecutor) executeTasksInParallel(tasks []*Task, ctx *ExecutionContext, outputChan chan TaskOutput) {
	if len(tasks) == 0 {
		return
	}
	rand.Seed(time.Now().UnixNano())
	shuffledTasks := make([]*Task, len(tasks))
	copy(shuffledTasks, tasks)
	rand.Shuffle(len(shuffledTasks), func(i, j int) {
		shuffledTasks[i], shuffledTasks[j] = shuffledTasks[j], shuffledTasks[i]
	})
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, task := range shuffledTasks {
		wg.Add(1)
		go func(t *Task) {
			defer wg.Done()
			jobStats, taskOutput := c.executeJob(t, ctx)
			mutex.Lock()
			ctx.Stats.TotalCommands += jobStats.Total
			ctx.Stats.SuccessfulCommands += jobStats.Successful
			ctx.Stats.FailedCommands += jobStats.Failed
			ctx.Stats.JobStats[t.Name] = jobStats
			mutex.Unlock()
			if !ctx.Config.Silent {
				outputChan <- *taskOutput
			}
		}(task)
	}
	wg.Wait()
}

func (c *TaskExecutor) executeJob(task *Task, ctx *ExecutionContext) (JobStats, *TaskOutput) {
	startTime := time.Now()
	output := &TaskOutput{
		TaskName:    task.Name,
		Description: task.Description,
		OutputLines: []string{},
		IsVerbose:   ctx.Config.Verbose,
	}
	if ctx.Config.Verbose && len(task.Description) > 0 {
		output.OutputLines = append(output.OutputLines, fmt.Sprintf("📦 Task: %s (%s)", task.Name, task.Description))
	} else if !ctx.Config.Silent && !ctx.Config.Short {
		output.OutputLines = append(output.OutputLines, fmt.Sprintf("📦 Task: %s", task.Name))
	}
	stats := JobStats{}
	var forValues interface{}
	var forEnabled bool
	var isMatrix bool
	for _, directive := range task.Directives {
		if directive.Type == "for" && len(directive.Params) > 0 {
			varName := fmt.Sprint(directive.Params[0])
			for _, v := range c.Runtime.Vars {
				if v.Name == varName {
					if listVal, ok := v.Value.(variable.ListValue); ok {
						forValues = listVal.Value
						forEnabled = true
						isMatrix = false
						break
					}
					if _, ok := v.Value.(variable.MatrixValue); ok {
						forValues, _ = v.GetMatrixCombinations()
						forEnabled = true
						isMatrix = true
						break
					}
				}
			}
		}
	}
	if forEnabled {
		var displayValues []string
		if isMatrix {
			for _, val := range forValues.([][]interface{}) {
				displayValues = append(displayValues, fmt.Sprint(val))
			}
		} else {
			for _, val := range forValues.([]interface{}) {
				displayValues = append(displayValues, fmt.Sprint(val))
			}
		}
		if !ctx.Config.Silent && ctx.Config.Verbose {
			output.OutputLines = append(output.OutputLines, fmt.Sprintf("  🔄 Looping over: %v", displayValues))
		}
		if isMatrix {
			for key, val := range forValues.([][]interface{}) {
				if !ctx.Config.Silent && ctx.Config.Verbose {
					output.OutputLines = append(output.OutputLines, fmt.Sprintf("    → Item: %v", val))
				}
				task.PushLoopStack(key, val)
				task.SetCurrentLoopIndex(key)
				jobStats := c.runTaskPipelines(task, ctx, output)
				stats.Total += jobStats.Total
				stats.Successful += jobStats.Successful
				stats.Failed += jobStats.Failed
			}
		} else {
			for key, val := range forValues.([]interface{}) {
				if !ctx.Config.Silent && ctx.Config.Verbose {
					output.OutputLines = append(output.OutputLines, fmt.Sprintf("    → Item: %v", val))
				}
				task.PushLoopStack(key, val)
				task.SetCurrentLoopIndex(key)
				jobStats := c.runTaskPipelines(task, ctx, output)
				stats.Total += jobStats.Total
				stats.Successful += jobStats.Successful
				stats.Failed += jobStats.Failed
			}
		}
	} else {
		task.ClearLoopStack()
		stats = c.runTaskPipelines(task, ctx, output)
	}
	stats.Duration = time.Since(startTime)
	return stats, output
}

func (c *TaskExecutor) runTaskPipelines(task *Task, ctx *ExecutionContext, output *TaskOutput) JobStats {
	stats := JobStats{}
	indent := "    "
	task.ClearUnusedStashes()
	for _, pipeline := range task.Pipelines {
		var result string
		var success bool
		var err error
		var displayResult string
		var shouldPrint bool
		stats.Total++

		if pipeline.Stash != nil {
			result, success = pipeline.Stash.GetValue(ctx, task, c.CommandExecutor)
			if success {
				c.Runtime.PushStashStack(pipeline.Stash.Name, task.GetTaskId(), result, pipeline.Stash)
				displayResult = fmt.Sprintf("%s▶ stash %q saved: %s", indent, pipeline.Stash.Name, strings.TrimRight(result, "\n"))
			} else {
				displayResult = fmt.Sprintf("%s❌ stash %q failed", indent, pipeline.Stash.Name)
				c.Reporter.ThrowRuntimeError(fmt.Sprintf("Stash %q failed", pipeline.Stash.Name), &pipeline.Stash.Metadata)
			}
			shouldPrint = ctx.Config.Verbose // Stash pipelines only print in verbose mode
		} else if pipeline.Function != nil {
			loopContext := task.GetActiveLoopContext()
			if loopContext != nil {
				pipeline.Function.SetLoopContext(loopContext)
			}
			functionResult, err := pipeline.Function.Call()
			if err != nil {
				displayResult = fmt.Sprintf("%s❌ %s (%v)", indent, pipeline.Function.Type, err)
				c.Reporter.ThrowRuntimeError(fmt.Sprintf("Function %q failed: %v", pipeline.Function.Type, err), &pipeline.Function.Metadata)
				success = false
			} else {
				success = true
				result = fmt.Sprint(functionResult)
				if strings.TrimSpace(result) != "" {
					switch pipeline.Function.Type {
					case "trigger":
						if len(pipeline.Function.Args) > 0 {
							taskName := pipeline.Function.GetCalculatedArgsByIndex(0)
							displayResult = fmt.Sprintf("%s🚀 trigger task '%s'", indent, taskName)
						}
						shouldPrint = ctx.Config.Verbose // Triggers only print in verbose mode
					case "print", "println", "printf", "sprintf":
						displayResult = fmt.Sprintf("%s▶ %s", indent, strings.TrimRight(result, "\n"))
						shouldPrint = true // Print functions always print
					default:
						displayResult = fmt.Sprintf("%s▶ %s", indent, strings.TrimRight(result, "\n"))
						shouldPrint = ctx.Config.Verbose // Other functions only print in verbose mode
					}
				}
			}
		} else if pipeline.Command != nil {
			result, success, err = pipeline.Command.GetResult(ctx, task, c.CommandExecutor)
			if strings.TrimSpace(result) != "" {
				if success {
					displayResult = fmt.Sprintf("%s▶ %s", indent, strings.TrimRight(result, "\n"))
					shouldPrint = true // Commands produce output when executed
				} else {
					displayResult = fmt.Sprintf("%s❌ %s (%v)", indent, strings.TrimRight(result, "\n"), err)
					c.Reporter.ThrowRuntimeError(fmt.Sprintf("Command failed: %v", err), &task.Metadata)
					shouldPrint = true // Errors always print
				}
			}
		}

		if !ctx.Config.Silent && displayResult != "" && shouldPrint {
			output.OutputLines = append(output.OutputLines, displayResult)
		}
		if success {
			stats.Successful++
		} else {
			stats.Failed++
		}
	}

	task.SetTaskIsFinished()

	return stats
}

func (c *TaskExecutor) printSummary(ctx *ExecutionContext) {
	if ctx.Config.Short {
		c.printShortSummary(ctx)
		return
	}
	_, _ = ctx.Colors.Cyan.Fprintln(ctx.Writer, "📊 Execution Summary:")
	_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "   ✅ Tasks executed: %d\n", ctx.Stats.TotalJobs)
	_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "   ✅ Total pipelines: %d\n", ctx.Stats.TotalCommands)
	if ctx.Stats.SuccessfulCommands == ctx.Stats.TotalCommands {
		_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "   ✅ All %d pipelines succeeded\n", ctx.Stats.SuccessfulCommands)
	} else {
		_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "   ✅ Successful: %d\n", ctx.Stats.SuccessfulCommands)
		_, _ = ctx.Colors.Red.Fprintf(ctx.Writer, "   ❌ Failed: %d\n", ctx.Stats.FailedCommands)
	}
	if ctx.Config.Profile {
		duration := ctx.Stats.EndTime.Sub(ctx.Stats.StartTime)
		_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "   ✅ Total execution time: %v\n", duration)
		if ctx.Config.Verbose {
			_, _ = ctx.Colors.White.Fprintln(ctx.Writer, "   📊 Task breakdown:")
			for jobName, jobStats := range ctx.Stats.JobStats {
				_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "     %s: %v (✅%d ❌%d)\n",
					jobName, jobStats.Duration, jobStats.Successful, jobStats.Failed)
			}
		}
	}
	if ctx.Stats.SuccessfulCommands != ctx.Stats.TotalCommands {
		_, _ = ctx.Colors.Red.Fprintln(ctx.Writer, "Execution completed with errors!")
	}
}

func (c *TaskExecutor) printShortSummary(ctx *ExecutionContext) {
	if ctx.Stats.SuccessfulCommands == ctx.Stats.TotalCommands {
		_, _ = fmt.Fprintf(ctx.Writer, "✅ %d/%d pipelines succeeded", ctx.Stats.SuccessfulCommands, ctx.Stats.TotalCommands)
	} else {
		_, _ = fmt.Fprintf(ctx.Writer, "❌ %d/%d pipelines succeeded", ctx.Stats.SuccessfulCommands, ctx.Stats.TotalCommands)
	}
	if ctx.Config.Profile {
		duration := ctx.Stats.EndTime.Sub(ctx.Stats.StartTime)
		_, _ = fmt.Fprintf(ctx.Writer, " (%v)", duration)
	}
	_, _ = fmt.Fprintln(ctx.Writer)
}

func onlyIndependentTasksWithoutDeps(tasks []*Task) bool {
	for _, task := range tasks {
		for _, directive := range task.Directives {
			if directive.Type == "depend" {
				return false
			}
		}
	}
	return true
}
