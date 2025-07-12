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
	Filename   string
}

func NewJob(name string, filename string) *Job {
	directives := new([]JobDirective)
	commands := new([]Command)

	return &Job{
		Name:       name,
		Directives: *directives,
		Commands:   *commands,
		Filename:   filename,
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
	Type   string
	Params []interface{}
}

func NewCMDDirective() *CMDDirective {
	values := new([]interface{})
	return &CMDDirective{
		Params: *values,
	}
}

// AddCMDDirective creates a CMDDirective based on the directive name and arguments
func (j *Job) AddCMDDirective(name string, args []interface{}) (*CMDDirective, error) {
	var dirValue []interface{}

	switch name {
	case "once":
		dirValue = []interface{}{true}

	case "try":
		if len(args) > 0 {
			if str, ok := args[0].(string); ok && strings.Trim(str, "0123456789") == "" {
				dirValue = []interface{}{str}
			} else {
				return nil, fmt.Errorf("@try requires a NUMBER argument, got %v", args[0])
			}
		} else {
			dirValue = []interface{}{"1"} // default: try(1)
		}

	case "stash", "buf", "output", "trigger":
		if len(args) > 0 {
			dirValue = []interface{}{args[0]}
		} else {
			return nil, fmt.Errorf("@%s requires an IDENTIFIER or REF argument", name)
		}

	case "concat", "print":
		dirValue = args // may be empty

	default:
		Warn(fmt.Sprintf("Unknown command directive @%s", name), j.Filename, 0, 0)
		dirValue = args
	}

	return &CMDDirective{
		Type:   name,
		Params: dirValue,
	}, nil
}

// AddJobDirective creates a JobDirective based on the directive name and arguments
func (j *Job) AddJobDirective(name string, args ...interface{}) error {
	var dirValue interface{}

	switch name {
	case "id", "name", "description", "working_dir", "timeout", "depend", "repeat", "for":
		if len(args) == 0 {
			return require_error(name)
		}
		dirValue = args // Store all values for multi-value directives like @depend, @for
	case "manual":
		if len(args) > 0 {
			if boolVal, ok := args[0].(string); ok && (boolVal == "true" || boolVal == "false") {
				dirValue = boolVal == "true"
			} else {
				return fmt.Errorf("@manual expects a boolean value, got %v", args[0])
			}
		} else {
			dirValue = true
		}
	default:
		Warn(fmt.Sprintf("Unknown job directive @%s", name), j.Filename, 0, 0)
		dirValue = args
	}

	j.Directives = append(j.Directives, JobDirective{
		Type:  name,
		Value: dirValue,
	})

	return nil
}

// AddJobCommand creates a Command
func (j *Job) AddJobCommand(cmd string, directives []CMDDirective) {
	j.Commands = append(j.Commands, Command{
		CMD:       cmd,
		Directive: directives,
	})
}
