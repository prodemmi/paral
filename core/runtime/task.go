package runtime

import (
	"fmt"
	"paral/core/metadata"
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
	StashStack       [][]*TaskStashContext
	currentLoopIndex int
}

type TaskPipeline struct {
	Stash    *Stash
	Command  *Command
	Function *Function
	Block    int
}

type TaskLoopContext struct {
	Value interface{}
	Key   int
}

type TaskStashContext struct {
	Stash       *Stash
	CacheResult interface{}
	Key         string
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
	Type   string
	Params []interface{}
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
	case "id", "description", "envfile", "for", "schedule":
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
		t.Runtime.Reporter.Warn(fmt.Sprintf("Unknown task directive @%s", name), &t.Metadata)
		dirValue = args
	}

	t.Directives = append(t.Directives, &Directive{
		Type:   name,
		Params: dirValue,
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

func (t *Task) PushStashStack(name string, value interface{}, stash *Stash) {
	for len(t.StashStack) <= t.currentLoopIndex {
		t.StashStack = append(t.StashStack, []*TaskStashContext{})
	}
	t.StashStack[t.currentLoopIndex] = append(t.StashStack[t.currentLoopIndex], &TaskStashContext{
		Key:         name,
		Stash:       stash,
		CacheResult: value,
	})
}

func (t *Task) GetActiveStashValue(name string) interface{} {
	for i := len(t.StashStack) - 1; i >= 0; i-- {
		for j := len(t.StashStack[i]) - 1; j >= 0; j-- {
			if t.StashStack[i][j].Key == name {
				return t.StashStack[i][j].CacheResult
			}
		}
	}
	return nil
}

// GetTaskId returns the task ID (either from @id directive or task name)
func (t *Task) GetTaskId() string {
	for _, directive := range t.Directives {
		if directive.Type == "id" && len(directive.Params) > 0 {
			return fmt.Sprint(directive.Params[0])
		}
	}
	return t.Name
}
