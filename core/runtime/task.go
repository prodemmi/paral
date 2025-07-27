package runtime

import (
	"fmt"
	"paral/core/metadata"
)

var require_error = func(name string) error {
	return fmt.Errorf("@%s requires at least one value", name)
}

type Task struct {
	Filename    string
	Name        string
	Description string
	Directives  []*Directive
	Commands    []*Command
	IsManual    bool
	Metadata    metadata.Metadata
	Reporter    *Reporter
}

func NewTask(name, description string, filename string, metadata metadata.Metadata) *Task {
	directives := make([]*Directive, 0)
	commands := make([]*Command, 0)
	return &Task{
		Name:        name,
		Description: description,
		Directives:  directives,
		Commands:    commands,
		Filename:    filename,
		Metadata:    metadata,
		Reporter:    NewReporter(&metadata),
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
func (j *Task) AddTaskDirective(directive *Directive) error {
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
	case "parallel":
		if len(args) > 0 {
			return fmt.Errorf("@parallel does not accept arguments, got %v", args)
		}
		dirValue = []interface{}{true}
	case "args", "depend", "dependall":
		if len(args) == 0 {
			return require_error(name)
		}
		for i, arg := range args {
			if _, ok := arg.(string); !ok {
				return fmt.Errorf("@%s requires string arguments, got %v at index %d", name, arg, i)
			}
		}
		dirValue = args
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
		j.IsManual = dirValue[0].(bool)
	default:
		j.Reporter.Warn(fmt.Sprintf("Unknown task directive @%s", name), &j.Metadata)
		dirValue = args
	}

	j.Directives = append(j.Directives, &Directive{
		Type:   name,
		Params: dirValue,
	})

	return nil
}

// AddTaskCommand creates a Command
func (j *Task) AddTaskCommand(cmd *Command) {
	j.Commands = append(j.Commands, cmd)
}
