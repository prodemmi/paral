package core

import (
	"fmt"
	"strings"
)

var require_error = func(name string) error {
	return fmt.Errorf("@%s requires at least one value", name)
}

type Job struct {
	Name       string
	Directives []JobDirective
	Commands   []Command
}

func NewJob(name string) *Job {
	directives := new([]JobDirective)
	commands := new([]Command)

	return &Job{
		Name:       name,
		Directives: *directives,
		Commands:   *commands,
	}
}

type JobDirective struct {
	Type  string
	Value interface{}
}

type Command struct {
	Directive []CMDDirective
	CMD       string
}

type CMDDirective struct {
	Type  string
	Value interface{}
}

// AddCMDDirective creates a CMDDirective based on the directive name and arguments
func (j *Job) AddCMDDirective(name string, args []interface{}) (*CMDDirective, error) {
	var dirValue interface{}

	switch name {
	case "once":
		dirValue = true
	case "try":
		if len(args) > 0 {
			if num, ok := args[0].(string); ok && strings.Trim(num, "0123456789") == "" {
				dirValue = num
			} else {
				return nil, fmt.Errorf("@try requires a NUMBER argument, got %v", args[0])
			}
		} else {
			dirValue = "1"
		}
	case "stash", "buf", "output", "trigger":
		if len(args) > 0 {
			dirValue = args[0]
		} else {
			return nil, fmt.Errorf("@%s requires an IDENTIFIER or REF argument", name)
		}
	case "concat", "print":
		dirValue = args
	default:
		dirValue = args
		return nil, fmt.Errorf("Unknown command directive @%s", name)
	}

	return &CMDDirective{
		Type:  name,
		Value: dirValue,
	}, nil
}

// AddJobDirective creates a JobDirective based on the directive name and arguments
func (j *Job) AddJobDirective(name string, args ...interface{}) error {
	var dirValue interface{}

	switch name {
	case "id", "name", "description", "working_dir", "timeout", "depend", "repeat":
		if len(args) == 0 {
			return require_error(name)
		}
		dirValue = args // Store all values for multi-value directives like @depend
	case "for":
		if len(args) == 0 {
			return require_error(name)
		}
		dirValue = args[0] // @for expects a single value
	case "manual":
		if len(args) > 0 {
			if arg, ok := args[0].([]interface{}); ok {
				arg := arg[0]
				if boolVal, ok := arg.(string); ok && (boolVal == "true" || boolVal == "false") {
					dirValue = boolVal == "true"
				} else {
					return fmt.Errorf("@manual expects a boolean value, got %v", arg)
				}
			}
		} else {
			dirValue = true
		}
	default:
		dirValue = args
		return fmt.Errorf("Unknown job directive @%s", name)
	}

	j.Directives = append(j.Directives, JobDirective{
		Type:  name,
		Value: dirValue,
	})

	return nil
}
