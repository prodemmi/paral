package runtime

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"paral/core/variable"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/fatih/color"
)

type TaskExecutor struct {
	Runtime         *Runtime
	Reporter        *Reporter
	CommandExecutor *CommandExecutor
	scheduler       *cron.Cron
	scheduledTasks  []*Task
	regularTasks    []*Task
	ctx             context.Context
	cancel          context.CancelFunc
	wg              sync.WaitGroup
}

type TaskOutput struct {
	TaskName    string
	Description string
	OutputLines []string
	IsVerbose   bool
	Timestamp   time.Time
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
	LogLevel    string // @TODO: implement
	MaxWorkers  int
	Timeout     int
	OutputFile  string
	Writer      io.Writer
	KeepAlive   bool // New flag to keep app running for scheduled tasks
}

type ExecutionStats struct {
	TotalJobs          int
	TotalCommands      int
	SuccessfulCommands int
	FailedCommands     int
	StartTime          time.Time
	EndTime            time.Time
	JobStats           map[string]JobStats
	ScheduledJobStats  map[string]*ScheduledJobStats
	mutex              sync.RWMutex
}

type JobStats struct {
	Total      int
	Successful int
	Failed     int
	Duration   time.Duration
}

type ScheduledJobStats struct {
	TotalRuns       int
	SuccessfulRuns  int
	FailedRuns      int
	LastRun         time.Time
	NextRun         time.Time
	TotalDuration   time.Duration
	AverageDuration time.Duration
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
	ctx, cancel := context.WithCancel(context.Background())

	return &TaskExecutor{
		Runtime:         runtime,
		Reporter:        runtime.Reporter,
		CommandExecutor: NewCommandExecutor(runtime),
		scheduler:       cron.New(cron.WithSeconds()), // Enable seconds precision
		ctx:             ctx,
		cancel:          cancel,
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
	if t == nil || t.writer == nil {
		return 0, fmt.Errorf("writer is nil")
	}
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

	if !config.Silent && !config.SummaryOnly {
		c.printVariables(ctx)
	}

	// Separate scheduled and regular tasks
	c.separateTasks()

	if config.DryRun {
		_, _ = fmt.Fprintln(ctx.Writer, "DRY RUN - Commands that would be executed:")
		c.executeDryRun(ctx)
		return
	}

	// Setup signal handling for graceful shutdown
	c.setupSignalHandling(ctx)

	// Start output processor
	outputChan := make(chan TaskOutput, 100)
	var printWg sync.WaitGroup

	if !ctx.Config.Silent && !config.SummaryOnly {
		printWg.Add(1)
		go func() {
			defer printWg.Done()
			c.processOutput(outputChan, ctx)
		}()
	}

	// Execute regular tasks first
	if len(c.regularTasks) > 0 {
		c.executeRegularTasks(ctx, outputChan)
	}

	// Setup and start scheduled tasks
	hasScheduledTasks := len(c.scheduledTasks) > 0
	if hasScheduledTasks {
		c.setupScheduledTasks(ctx, outputChan)
		c.scheduler.Start()

		if !config.Silent {
			c.printScheduledTasksInfo(ctx)
		}
	}

	// Keep the application running if there are scheduled tasks or keep-alive is requested
	if hasScheduledTasks && (config.KeepAlive || hasScheduledTasks) {
		c.waitForScheduledTasks(ctx)
	}

	// Cleanup
	c.shutdown(ctx)

	// Close output channel and wait for processing to complete
	close(outputChan)
	if !ctx.Config.Silent {
		printWg.Wait()
	}

	if config.Profile {
		ctx.Stats.EndTime = time.Now()
	}

	if !config.Silent || config.SummaryOnly {
		c.printSummary(ctx)
	}
}

func (c *TaskExecutor) separateTasks() {
	c.scheduledTasks = []*Task{}
	c.regularTasks = []*Task{}

	for _, task := range c.Runtime.Tasks {
		if task.IsManual() {
			continue // Skip manual tasks
		}

		if task.IsScheduled() {
			c.scheduledTasks = append(c.scheduledTasks, task)
		} else {
			c.regularTasks = append(c.regularTasks, task)
		}
	}
}

func (c *TaskExecutor) setupSignalHandling(ctx *ExecutionContext) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		if ctx.Config.Verbose {
			_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "\nreceived %s signal, stopping scheduled tasks...\n", sig)
		}
		c.cancel()

		// Force exit after timeout if graceful shutdown fails
		go func() {
			time.Sleep(10 * time.Second)
			if ctx.Config.Verbose {
				_, _ = ctx.Colors.Red.Fprintln(ctx.Writer, "force exit after timeout")
			}
			os.Exit(1)
		}()
	}()
}

func (c *TaskExecutor) setupScheduledTasks(ctx *ExecutionContext, outputChan chan TaskOutput) {
	for _, task := range c.scheduledTasks {
		scheduleDirective := task.GetScheduleDirective()
		if scheduleDirective == nil {
			continue
		}

		taskScheduleDirectiveContext := task.GetTaskScheduleDirectiveContext()
		if taskScheduleDirectiveContext == nil {
			continue
		}

		// Create a wrapper function that handles the execution context
		taskFunc := c.createScheduledTaskFunc(task, ctx, outputChan)

		// Add to cron scheduler
		value := taskScheduleDirectiveContext.Value
		var cronSpec string

		if taskScheduleDirectiveContext.IsLinuxFormat {
			cronSpec = value
		} else if taskScheduleDirectiveContext.IsDurationFormat {
			cronSpec = "@every " + value
		} else {
			cronSpec = value
		}

		entryID, err := c.scheduler.AddFunc(cronSpec, taskFunc)
		if err != nil {
			c.Reporter.ThrowRuntimeError(fmt.Sprintf("failed to schedule task '%s': %v", task.Name, err), &task.Metadata)
			continue
		}

		// Initialize stats for this scheduled task
		ctx.Stats.mutex.Lock()
		if ctx.Stats.ScheduledJobStats == nil {
			ctx.Stats.ScheduledJobStats = make(map[string]*ScheduledJobStats)
		}
		ctx.Stats.ScheduledJobStats[task.Name] = &ScheduledJobStats{
			NextRun: c.scheduler.Entry(entryID).Next,
		}
		ctx.Stats.mutex.Unlock()
	}
}

func (c *TaskExecutor) createScheduledTaskFunc(task *Task, ctx *ExecutionContext, outputChan chan TaskOutput) func() {
	return func() {
		select {
		case <-c.ctx.Done():
			return // Context cancelled, stop execution
		default:
			// Execute the task
			startTime := time.Now()

			// Update stats - mark as started
			ctx.Stats.mutex.Lock()
			if stats, exists := ctx.Stats.ScheduledJobStats[task.Name]; exists {
				stats.TotalRuns++
				stats.LastRun = startTime
				// Update next run time
				for _, entry := range c.scheduler.Entries() {
					stats.NextRun = entry.Next
					break
				}
			}
			ctx.Stats.mutex.Unlock()

			// Execute the task with proper dependency handling
			jobStats, taskOutput := c.executeScheduledJob(task, ctx)

			duration := time.Since(startTime)

			// Update final stats
			ctx.Stats.mutex.Lock()
			if stats, exists := ctx.Stats.ScheduledJobStats[task.Name]; exists {
				if jobStats.Failed == 0 {
					stats.SuccessfulRuns++
				} else {
					stats.FailedRuns++
				}
				stats.TotalDuration += duration
				if stats.TotalRuns > 0 {
					stats.AverageDuration = stats.TotalDuration / time.Duration(stats.TotalRuns)
				}
			}

			ctx.Stats.TotalCommands += jobStats.Total
			ctx.Stats.SuccessfulCommands += jobStats.Successful
			ctx.Stats.FailedCommands += jobStats.Failed
			ctx.Stats.mutex.Unlock()

			// Send output for live display - FIX: Always send output, even if silent
			if taskOutput != nil {
				taskOutput.Timestamp = startTime
				select {
				case outputChan <- *taskOutput:
				case <-time.After(100 * time.Millisecond):
					// Timeout to prevent blocking if channel is full
				}
			}
		}
	}
}

func (c *TaskExecutor) executeRegularTasks(ctx *ExecutionContext, outputChan chan TaskOutput) {
	independentTasks, deferTasks, err := c.Runtime.GetExecutionOrder()
	if err != nil {
		c.Reporter.ThrowRuntimeError(err.Error(), c.Runtime.Metadata)
		return
	}

	// Filter out scheduled tasks from regular execution
	independentTasks = c.filterScheduledTasks(independentTasks)
	deferTasks = c.filterScheduledTasks(deferTasks)

	if onlyIndependentTasksWithoutDeps(independentTasks) {
		c.executeTasksInParallel(independentTasks, ctx, outputChan)
	} else {
		c.executeTasksWithDependencies(independentTasks, ctx, outputChan)
	}

	if len(deferTasks) > 0 {
		c.executeTasksInParallel(deferTasks, ctx, outputChan)
	}
}

func (c *TaskExecutor) filterScheduledTasks(tasks []*Task) []*Task {
	filtered := []*Task{}
	for _, task := range tasks {
		if !task.IsScheduled() {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func (c *TaskExecutor) waitForScheduledTasks(ctx *ExecutionContext) {
	if !ctx.Config.Silent {
		_, _ = ctx.Colors.Cyan.Fprintln(ctx.Writer, "🕐 Waiting for scheduled tasks... (Press Ctrl+C to stop)")
	}

	<-c.ctx.Done()
}

func (c *TaskExecutor) shutdown(ctx *ExecutionContext) {
	if c.scheduler != nil {
		// Stop the scheduler with context
		shutdownCtx := c.scheduler.Stop()

		select {
		case <-shutdownCtx.Done():
		case <-time.After(3 * time.Second):
			if ctx.Config.Verbose {
				_, _ = ctx.Colors.Yellow.Fprintln(ctx.Writer, "scheduler shutdown timeout, forcing exit")
			}
		}
	}

	// Wait for goroutines with timeout
	done := make(chan struct{})
	go func() {
		c.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// All goroutines finished
	case <-time.After(2 * time.Second):
		if !ctx.Config.Silent {
			_, _ = ctx.Colors.Yellow.Fprintln(ctx.Writer, "⚠️  Timeout waiting for goroutines")
		}
	}
}

func (c *TaskExecutor) printScheduledTasksInfo(ctx *ExecutionContext) {
	if len(c.scheduledTasks) == 0 {
		return
	}

	_, _ = ctx.Colors.Cyan.Fprintln(ctx.Writer, "  Scheduled Tasks:")
	for _, task := range c.scheduledTasks {
		scheduleDirective := task.GetScheduleDirective()
		if scheduleDirective != nil && len(scheduleDirective.Params) > 0 {
			schedule := fmt.Sprint(scheduleDirective.Params[0])
			_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "   🕐 %s: %s\n", task.Name, schedule)

			ctx.Stats.mutex.RLock()
			if stats, exists := ctx.Stats.ScheduledJobStats[task.Name]; exists && !stats.NextRun.IsZero() {
				_, _ = ctx.Colors.Gray.Fprintf(ctx.Writer, "      Next run: %s\n", stats.NextRun.Format("2006-01-02 15:04:05"))
			}
			ctx.Stats.mutex.RUnlock()
		}
	}
	_, _ = fmt.Fprintln(ctx.Writer)
}

func (c *TaskExecutor) executeScheduledJob(task *Task, ctx *ExecutionContext) (JobStats, *TaskOutput) {
	// Check dependencies first
	dependencies := c.Runtime.GetTaskDependencies(task)
	for _, depId := range dependencies {
		depTask := c.Runtime.GetTaskByID(depId)
		if depTask != nil && !depTask.isFinished {
			// Execute dependency first (this could be improved with better dependency resolution)
			c.executeJob(depTask, ctx)
		}
	}

	// Execute the scheduled task
	return c.executeJob(task, ctx)
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

	nonAutoStartTasks := 0
	for _, task := range c.Runtime.Tasks {
		if !(task.IsManual()) {
			nonAutoStartTasks++
		}
	}

	stats := &ExecutionStats{
		TotalJobs:         nonAutoStartTasks,
		JobStats:          make(map[string]JobStats),
		ScheduledJobStats: make(map[string]*ScheduledJobStats),
		StartTime:         time.Now(),
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
}

func (c *TaskExecutor) printVerboseLine(line string, ctx *ExecutionContext) {
	switch {
	case strings.HasPrefix(line, "  Task"):
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
		_, _ = fmt.Fprintf(ctx.Writer, "  Variables: ")
		for i, v := range c.Runtime.Vars {
			if i > 0 {
				_, _ = fmt.Fprint(ctx.Writer, ", ")
			}
			_, valueStr := v.Format()
			_, _ = fmt.Fprintf(ctx.Writer, "%s=%s", v.Name, valueStr)
		}
		return
	}
	_, _ = ctx.Colors.Magenta.Fprint(ctx.Writer, "  Variables: ")
	for i, v := range c.Runtime.Vars {
		if i > 0 {
			_, _ = ctx.Colors.Magenta.Fprint(ctx.Writer, ", ")
		}
		_, valueStr := v.Format()
		_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "%s=%s", v.Name, valueStr)
	}
	_, _ = fmt.Fprintln(ctx.Writer)
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
		_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "  Task %q (%s):\n", taskId, task.Description)
	} else {
		_, _ = ctx.Colors.Yellow.Fprintf(ctx.Writer, "  Task %q\n", taskId)
	}
	if len(dependencies) > 0 {
		_, _ = ctx.Colors.Magenta.Fprintf(ctx.Writer, "  🕐 Depends on: %v\n", dependencies)
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
	indent := "    "
	for _, pipeline := range task.Pipelines {
		displayCmd, _, isTrigger, isStash, isBuf, isPrint := pipeline.GetDryResult(task, ctx)

		shouldPrintOutput := isPrint || (!isBuf && !isStash && !isTrigger && pipeline.Command != nil) || ctx.Config.Verbose
		if shouldPrintOutput {
			if isTrigger {
				_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "%s🚀 %s\n", indent, displayCmd)
			} else if isStash {
				_, _ = ctx.Colors.Cyan.Fprintf(ctx.Writer, "%s ▶ %s\n", indent, displayCmd)
			} else {
				_, _ = ctx.Colors.Blue.Fprintf(ctx.Writer, "%s ▶ %s\n", indent, displayCmd)
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

	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, task := range tasks {
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

	dependencies := c.Runtime.GetTaskDependencies(task)
	if ctx.Config.Verbose && len(dependencies) > 0 {
		taskDirectivesPrefix := ctx.Colors.Gray.Sprintf("\n  @depends on: %v", dependencies)
		output.OutputLines = append(output.OutputLines, taskDirectivesPrefix)
	}

	// Add timestamp only for scheduled tasks
	var taskPrefix string
	if task.IsScheduled() {
		timestamp := startTime.Format("15:04:05")
		taskPrefix = ctx.Colors.Magenta.Sprintf("  Task %q [%s]:", task.Name, timestamp)
	} else {
		taskPrefix = ctx.Colors.Magenta.Sprintf("  Task %q", task.Name)

	}

	if ctx.Config.Verbose && len(task.Description) > 0 {
		output.OutputLines = append(output.OutputLines, fmt.Sprintf("%s (%s)", taskPrefix, task.Description))
	} else if !ctx.Config.Silent && !ctx.Config.Short {
		output.OutputLines = append(output.OutputLines, taskPrefix)
	}

	stats := JobStats{}

	taskForDirectiveContext := task.GetTaskForDirectiveContext()
	taskScheduleDirectiveContext := task.GetTaskScheduleDirectiveContext()

	if taskForDirectiveContext != nil {
		var displayValues []string
		if taskForDirectiveContext.IsMatrix {
			for _, val := range taskForDirectiveContext.Value.([][]interface{}) {
				displayValues = append(displayValues, fmt.Sprint(val))
			}
		} else {
			for _, val := range taskForDirectiveContext.Value.([]interface{}) {
				displayValues = append(displayValues, fmt.Sprint(val))
			}
		}
		if !ctx.Config.Silent && ctx.Config.Verbose {
			output.OutputLines = append(output.OutputLines, fmt.Sprintf("  🔄 Looping over: %v", displayValues))
		}
		if taskForDirectiveContext.IsMatrix {
			for key, val := range taskForDirectiveContext.Value.([][]interface{}) {
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
			for key, val := range taskForDirectiveContext.Value.([]interface{}) {
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
	} else if taskScheduleDirectiveContext != nil {
		// FIX: Don't create a new cron scheduler here - scheduled tasks are handled elsewhere
		// This was causing the blocking issue
		task.ClearLoopStack()
		stats = c.runTaskPipelines(task, ctx, output)
	} else {
		task.ClearLoopStack()
		stats = c.runTaskPipelines(task, ctx, output)
	}

	stats.Duration = time.Since(startTime)
	return stats, output
}

func (c *TaskExecutor) runTaskPipelines(task *Task, ctx *ExecutionContext, output *TaskOutput) JobStats {
	stats := JobStats{}
	task.ClearUnusedStashes()
	for _, pipeline := range task.Pipelines {
		stats.Total++
		success, displayResult, shouldPrint := pipeline.GetResult(ctx, task, c.CommandExecutor)

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

	if ctx.Config.SummaryOnly {
		_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "  Regular tasks executed: %d\n", len(c.regularTasks))
		_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "  Scheduled tasks: %d\n", len(c.scheduledTasks))
		_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "  Total pipelines: %d\n", ctx.Stats.TotalCommands)
	} else {
		_, _ = ctx.Colors.Cyan.Fprintln(ctx.Writer, "\n  Execution Summary:")
		_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "   ✅ Regular tasks executed: %d\n", len(c.regularTasks))
		_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "   ✅ Scheduled tasks: %d\n", len(c.scheduledTasks))
		_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "   ✅ Total pipelines: %d\n", ctx.Stats.TotalCommands)
	}

	if ctx.Stats.SuccessfulCommands == ctx.Stats.TotalCommands {
		if ctx.Config.SummaryOnly {
			_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "  All %d pipelines succeeded\n", ctx.Stats.SuccessfulCommands)
		} else {
			_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "   ✅ All %d pipelines succeeded\n", ctx.Stats.SuccessfulCommands)
		}
	} else {
		_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "   ✅ Successful: %d\n", ctx.Stats.SuccessfulCommands)
		_, _ = ctx.Colors.Red.Fprintf(ctx.Writer, "   ❌ Failed: %d\n", ctx.Stats.FailedCommands)
	}

	// Print scheduled task statistics
	ctx.Stats.mutex.RLock()
	if len(ctx.Stats.ScheduledJobStats) > 0 {
		_, _ = ctx.Colors.Cyan.Fprintln(ctx.Writer, "   Scheduled Task Statistics:")
		for taskName, stats := range ctx.Stats.ScheduledJobStats {
			_, _ = ctx.Colors.White.Fprintf(ctx.Writer, "     %s: %d runs (✅%d ❌%d) avg: %v\n",
				taskName, stats.TotalRuns, stats.SuccessfulRuns, stats.FailedRuns, stats.AverageDuration)
		}
	}
	ctx.Stats.mutex.RUnlock()

	if ctx.Config.Profile {
		duration := ctx.Stats.EndTime.Sub(ctx.Stats.StartTime)
		_, _ = ctx.Colors.Green.Fprintf(ctx.Writer, "   ✅ Total execution time: %v\n", duration)
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
