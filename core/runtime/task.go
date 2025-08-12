package runtime

import (
	"fmt"
	"paral/core"
	"paral/core/metadata"
	"paral/core/variable"
	"strings"
	"time"
)

var require_error = func(name string) error {
	return fmt.Errorf("@%s requires at least one value", name)
}
var not_found_error = func(name string) error {
	return fmt.Errorf("stash %s not exists", name)
}

type Task struct {
	Filename         string
	Name             string
	Description      string
	Directives       []*Directive
	Pipelines        []*TaskPipeline
	Metadata         metadata.Metadata
	Runtime          *Runtime
	LoopStack        []*TaskLoopContext
	currentLoopIndex int

	isFinished bool
}

type TaskPipeline struct {
	Buf       *Buf
	Stash     *Stash
	Command   *Command
	Function  *Function
	Condition *Condition
	Block     int
}

type TaskLoopContext struct {
	Value interface{}
	Key   int
}

type TaskForDirectiveContext struct {
	Value    interface{}
	Key      int
	IsMatrix bool
}

type TaskScheduleDirectiveContext struct {
	Value            string
	IsLinuxFormat    bool
	IsDurationFormat bool
}

func NewTask(runtime *Runtime, name, description string, filename string, metadata metadata.Metadata) *Task {
	directives := make([]*Directive, 0)
	pipelines := make([]*TaskPipeline, 0)
	return &Task{
		Name:        name,
		Description: description,
		Directives:  directives,
		Pipelines:   pipelines,
		Filename:    filename,
		Metadata:    metadata,
		Runtime:     runtime,
	}
}

type Directive struct {
	Type     string
	Params   []interface{}
	Metadata metadata.Metadata
}

func NewTaskDirective() *Directive {
	values := make([]interface{}, 0)
	return &Directive{
		Params: values,
	}
}

// AddTaskDirective validates and adds a TaskDirective
func (t *Task) AddTaskDirective(directive *Directive) error {
	name := directive.Type
	args := directive.Params
	var dirValue []interface{}

	switch name {
	case "description", "envfile", "for", "schedule":
		if len(args) != 1 {
			return fmt.Errorf("@%s requires exactly one string value, got %d arguments", name, len(args))
		}
		if _, ok := args[0].(string); !ok {
			return fmt.Errorf("@%s requires a string value, got %v", name, args[0])
		}
		dirValue = args
	case "args", "depend":
		if len(args) == 0 {
			return require_error(name)
		}
		for i, arg := range args {
			if _, ok := arg.(string); !ok {
				return fmt.Errorf("@%s requires string arguments, got %v at index %d", name, arg, i)
			}
		}
		dirValue = args
	case "defer":
		if len(args) > 0 {
			return fmt.Errorf("@defer does not accept arguments, got %v", args)
		}
		dirValue = nil
	case "manual":
		if len(args) > 1 {
			return fmt.Errorf("@manual accepts at most one boolean value, got %d arguments", len(args))
		}
		if len(args) == 1 {
			if boolStr, ok := args[0].(string); ok && (boolStr == "true" || boolStr == "false") {
				dirValue = []interface{}{boolStr == "true"}
			} else {
				return fmt.Errorf("@manual expects a boolean value (true/false), got %v", args[0])
			}
		} else {
			dirValue = []interface{}{true}
		}
		dirValue = []interface{}{dirValue[0].(bool)}
	default:
		t.Runtime.Reporter.ThrowRuntimeError(fmt.Sprintf("Unknown task directive @%s", name), &directive.Metadata)
		dirValue = args
	}

	t.Directives = append(t.Directives, &Directive{
		Type:     name,
		Params:   dirValue,
		Metadata: directive.Metadata,
	})

	return nil
}

// AddTaskPipeline creates a TaskPipeline
func (t *Task) AddTaskPipeline(pipeline *TaskPipeline) {
	t.Pipelines = append(t.Pipelines, pipeline)
}

func (t *Task) IsManual() bool {
	for _, directive := range t.Directives {
		if directive.Type == "manual" && directive.Params[0].(bool) == true {
			return true
		}
	}
	return false
}

func (t *Task) IsScheduled() bool {
	for _, directive := range t.Directives {
		if directive.Type == "schedule" && directive.Params[0].(string) != "" {
			return true
		}
	}
	return false
}

func (t *Task) PushLoopStack(key int, value interface{}) {
	loopContext := TaskLoopContext{
		Key:   key,
		Value: value,
	}
	t.LoopStack = append(t.LoopStack, &loopContext)
}

func (t *Task) ClearLoopStack() {
	t.LoopStack = nil
}

func (t *Task) SetCurrentLoopIndex(key int) {
	t.currentLoopIndex = key
}

func (t *Task) GetActiveLoopContext() *TaskLoopContext {
	if len(t.LoopStack) == 0 || t.currentLoopIndex >= len(t.LoopStack) {
		return nil
	}
	return t.LoopStack[t.currentLoopIndex]
}

func (t *Task) ClearUnusedStashes() {
	runtime := t.Runtime

	for i := range runtime.StashStack {
		stack := runtime.StashStack[i]
		filteredStack := stack[:0]

		for _, stashCtx := range stack {
			stillNeeded := false
			for _, task := range runtime.Tasks {
				if runtime.IsTaskDependent(task.GetTaskId(), stashCtx.TaskID) && !task.isFinished {
					stillNeeded = true
					break
				}
			}
			if stillNeeded {
				filteredStack = append(filteredStack, stashCtx)
			}
		}

		runtime.StashStack[i] = filteredStack
	}
}

func (t *Task) SetTaskIsFinished() {
	t.isFinished = true
}

func (t *Task) GetTaskForDirectiveContext() *TaskForDirectiveContext {
	for _, directive := range t.Directives {
		if directive.Type == "for" && len(directive.Params) > 0 {
			varName := fmt.Sprint(directive.Params[0])
			for _, v := range t.Runtime.Vars {
				if v.Name == varName {
					if listVal, ok := v.Value.(variable.ListValue); ok {
						return &TaskForDirectiveContext{
							Value:    listVal.Value,
							IsMatrix: false,
						}
					}
					if _, ok := v.Value.(variable.MatrixValue); ok {
						value, _ := v.GetMatrixCombinations()
						return &TaskForDirectiveContext{
							Value:    value,
							IsMatrix: true,
						}
					}
				}
			}
		}
	}
	return nil
}

func (t *Task) GetTaskScheduleDirectiveContext() *TaskScheduleDirectiveContext {
	for _, directive := range t.Directives {
		if directive.Type == "schedule" && len(directive.Params) == 1 {
			scheduleValue, ok := directive.Params[0].(string)
			if !ok {
				t.Runtime.Reporter.ThrowRuntimeError("@schedule directive has not value", &directive.Metadata)
			}
			if core.HasQuotes(scheduleValue) {
				if strings.Contains(scheduleValue, " ") {
					return &TaskScheduleDirectiveContext{
						Value:         core.TrimQuotes(scheduleValue),
						IsLinuxFormat: true,
					}
				} else {
					durationFormatValue, _ := time.ParseDuration(core.TrimQuotes(scheduleValue))
					return &TaskScheduleDirectiveContext{
						Value:            durationFormatValue.String(),
						IsDurationFormat: true,
					}
				}
			} else if duration, err := time.ParseDuration(scheduleValue); err == nil {
				return &TaskScheduleDirectiveContext{
					Value:            duration.String(),
					IsDurationFormat: true,
				}
			} else {
				scheduleVarValue := t.Runtime.GetVariable(scheduleValue)
				if scheduleVarValue != nil {
					switch v := scheduleVarValue.Value.(type) {
					case variable.StringValue:
						return &TaskScheduleDirectiveContext{
							Value:         v.Value,
							IsLinuxFormat: true,
						}
					case variable.DurationValue:
						return &TaskScheduleDirectiveContext{
							Value:            v.Value,
							IsDurationFormat: true,
						}
					}
				} else {
					msg := fmt.Sprintf("@schedule directive has an unknown value %q", scheduleValue)
					t.Runtime.Reporter.ThrowSyntaxError(msg, &directive.Metadata)
				}
			}
		}
	}
	return nil
}

func (t *Task) GetScheduleDirective() *Directive {
	for _, directive := range t.Directives {
		if directive.Type == "schedule" {
			return directive
		}
	}
	return nil
}

func (t *Task) GetTaskId() string {
	return t.Name
}

func (tp *TaskPipeline) GetResult(ctx *ExecutionContext, task *Task, cmdExecutor *CommandExecutor) (bool, string, bool) {
	indent := "    "
	success := true
	displayResult := ""
	shouldPrint := false

	if tp.Buf != nil {
		result, ok := tp.Buf.GetValue(ctx, task, cmdExecutor)
		if ok {
			cmdExecutor.Runtime.PushBufStack(tp.Buf.Name, task.GetTaskId(), result, tp.Buf)
			displayResult = fmt.Sprintf("%s ▶ buf %q saved: %s", indent, tp.Buf.Name, strings.TrimRight(result, "\n"))
		} else {
			success = false
			displayResult = fmt.Sprintf("%s ❌ buf %q failed", indent, tp.Buf.Name)
			cmdExecutor.Reporter.ThrowRuntimeError(fmt.Sprintf("buf %q failed", tp.Buf.Name), &tp.Buf.Metadata)
		}
		shouldPrint = ctx.Config.Verbose

	} else if tp.Stash != nil {
		result, ok := tp.Stash.GetValue(ctx, task, cmdExecutor)
		if ok {
			cmdExecutor.Runtime.PushStashStack(tp.Stash.Name, task.GetTaskId(), result, tp.Stash)
			displayResult = fmt.Sprintf("%s ▶ stash %q saved: %s", indent, tp.Stash.Name, strings.TrimRight(result, "\n"))
		} else {
			success = false
			displayResult = fmt.Sprintf("%s ❌ stash %q failed", indent, tp.Stash.Name)
			cmdExecutor.Reporter.ThrowRuntimeError(fmt.Sprintf("Stash %q failed", tp.Stash.Name), &tp.Stash.Metadata)
		}
		shouldPrint = ctx.Config.Verbose

	} else if tp.Condition != nil && tp.Condition.IfCondition != nil {
		successes, outputs, prints := tp.Condition.IfCondition.Execute(ctx, task, cmdExecutor)
		for _, s := range successes {
			if !s {
				success = false
				break
			}
		}
		var results []string
		for i := range outputs {
			if prints[i] && strings.TrimSpace(outputs[i]) != "" {
				results = append(results, outputs[i])
				shouldPrint = true
			}
		}
		if len(results) > 0 {
			displayResult = strings.Join(results, "\n")
		} else {
			displayResult = fmt.Sprintf("%s ▶ if condition evaluated, no output", indent)
			shouldPrint = ctx.Config.Verbose
		}

	} else if tp.Function != nil {
		functionResult, err := tp.Function.Call()
		if err != nil {
			success = false
			displayResult = fmt.Sprintf("%s❌ %s (%v)", indent, tp.Function.Type, err)
			cmdExecutor.Reporter.ThrowRuntimeError(fmt.Sprintf("Function %q failed: %v", tp.Function.Type, err), &tp.Function.Metadata)
		} else {
			result := fmt.Sprint(functionResult)
			if strings.TrimSpace(result) != "" {
				switch tp.Function.Type {
				case "trigger":
					if len(tp.Function.Args) > 0 {
						taskName := tp.Function.GetCalculatedArgsByIndex(0)
						displayResult = fmt.Sprintf("%s🚀 trigger task '%s'", indent, taskName)
					}
					shouldPrint = ctx.Config.Verbose
				case "print", "println", "printf", "sprintf":
					displayResult = fmt.Sprintf("%s ▶ %s", indent, strings.TrimRight(result, "\n"))
					shouldPrint = true
				default:
					displayResult = fmt.Sprintf("%s ▶ %s", indent, strings.TrimRight(result, "\n"))
					shouldPrint = ctx.Config.Verbose
				}
			}
		}

	} else if tp.Command != nil {
		result, ok, err := tp.Command.GetResult(ctx, task, cmdExecutor)
		if strings.TrimSpace(result) != "" {
			if ok {
				displayResult = fmt.Sprintf("%s ▶ %s", indent, strings.TrimRight(result, "\n"))
				shouldPrint = true
			} else {
				success = false
				displayResult = fmt.Sprintf("%s❌ %s (%v)", indent, strings.TrimRight(result, "\n"), err)
				cmdExecutor.Reporter.ThrowRuntimeError(fmt.Sprintf("Command failed: %v", err), &task.Metadata)
				shouldPrint = true
			}
		}
	}

	return success, displayResult, shouldPrint
}

func (tp *TaskPipeline) GetDryResult(task *Task, ctx *ExecutionContext) (displayResult string, shouldPrint bool, isTrigger bool, isStash bool, isBuf bool, isPrint bool) {
	if tp.Buf != nil {
		isBuf = true
		if tp.Buf.Command != nil {
			result := tp.Buf.Command.GetRawResult()
			displayResult = fmt.Sprintf("buf %q saved: %s", tp.Buf.Name, strings.TrimRight(result, "\n"))
		} else {
			displayResult = fmt.Sprintf("buf %q saved", tp.Buf.Name)
		}
		shouldPrint = ctx.Config.Verbose

	} else if tp.Stash != nil {
		isStash = true
		result, _ := tp.Stash.GetRawValue()
		displayResult = fmt.Sprintf("stash %q saved: %s", tp.Stash.Name, strings.TrimRight(result, "\n"))
		shouldPrint = ctx.Config.Verbose

	} else if tp.Condition != nil {
		// Dry run for IfCondition — show which branch would execute
		if tp.Condition.IfCondition != nil {
			for _, branch := range tp.Condition.IfCondition.Branches {
				if branch.Expression == nil || branch.Expression.IsTrue() {
					if len(branch.Pipelines) > 0 {
						var dryOutputs []string
						for _, pipe := range branch.Pipelines {
							out, print, trig, stash, buf, printFn := pipe.GetDryResult(task, ctx)
							if print {
								dryOutputs = append(dryOutputs, out)
							}
							// bubble up flags if needed
							isTrigger = isTrigger || trig
							isStash = isStash || stash
							isBuf = isBuf || buf
							isPrint = isPrint || printFn
						}
						displayResult = strings.Join(dryOutputs, "\n")
						shouldPrint = shouldPrint || ctx.Config.Verbose
					}
					break
				}
			}
		}

	} else if tp.Function != nil {
		result, err := tp.Function.Call()
		if err != nil {
			displayResult = fmt.Sprintf("%s (%v)", tp.Function.Type, err)
			return
		}
		resStr := strings.TrimRight(fmt.Sprint(result), "\n")
		if strings.TrimSpace(resStr) == "" {
			return
		}
		displayResult = resStr

		switch tp.Function.Type {
		case "trigger":
			isTrigger = true
			if len(tp.Function.Args) > 0 {
				taskName := fmt.Sprint(tp.Function.GetCalculatedArgsByIndex(0))
				displayResult = fmt.Sprintf("trigger task '%s'", taskName)
			}
			shouldPrint = ctx.Config.Verbose
		case "print", "println", "printf", "sprintf":
			isPrint = true
			shouldPrint = true
		default:
			shouldPrint = ctx.Config.Verbose
		}

	} else if tp.Command != nil {
		displayResult = tp.Command.GetRawResult()
		if strings.TrimSpace(displayResult) != "" {
			shouldPrint = true
		}
	}

	return
}
