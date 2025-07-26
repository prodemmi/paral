package core

import (
	"fmt"
	"strings"
)

var require_error = func(name string) error {
	return fmt.Errorf("@%s requires at least one value", name)
}

type Job struct {
	Filename   string
	Name       string
	Directives []Directive
	Commands   []Command
}

func NewJob(name string, filename string) *Job {
	directives := make([]Directive, 0)
	commands := make([]Command, 0)
	return &Job{
		Name:       name,
		Directives: directives,
		Commands:   commands,
		Filename:   filename,
	}
}

type Directive struct {
	Type   string
	Params []interface{}
}

func NewJobDirective() *Directive {
	values := make([]interface{}, 0)
	return &Directive{
		Params: values,
	}
}

// AddJobDirective validates and adds a JobDirective
func (j *Job) AddJobDirective(directive *Directive) error {
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
	default:
		Warn(fmt.Sprintf("Unknown job directive @%s", name), j.Filename, 0, 0)
		dirValue = args
	}

	j.Directives = append(j.Directives, Directive{
		Type:   name,
		Params: dirValue,
	})

	return nil
}

// AddFunction validates and adds a CMDDirective
func (j *Job) AddFunction(name string, args []interface{}) (*Function, error) {
	var dirValue []interface{}

	switch name {
	case "print", "println", "printf":
		// Allow arbitrary arguments (strings or nested directives)
		if len(args) == 0 {
			return nil, fmt.Errorf("@%s requires at least one argument", name)
		}
		for i, arg := range args {
			if _, ok := arg.(string); !ok {
				return nil, fmt.Errorf("@%s requires string arguments, got %v at index %d", name, arg, i)
			}
		}
		dirValue = args
	case "format":
		if len(args) < 1 {
			return nil, fmt.Errorf("@format requires at least one format string, got %d arguments", len(args))
		}
		if _, ok := args[0].(string); !ok {
			return nil, fmt.Errorf("@format requires a string format as first argument, got %v", args[0])
		}
		dirValue = args
	case "arg":
		if len(args) != 1 {
			return nil, fmt.Errorf("@arg requires exactly one string name, got %d arguments", len(args))
		}
		if _, ok := args[0].(string); !ok {
			return nil, fmt.Errorf("@arg requires a string name, got %v", args[0])
		}
		dirValue = args
	case "value":
		if len(args) > 1 {
			return nil, fmt.Errorf("@value accepts at most one index, got %d arguments", len(args))
		}
		if len(args) == 1 {
			if str, ok := args[0].(string); ok {
				if strings.Trim(str, "0123456789") != "" {
					return nil, fmt.Errorf("@value requires a numeric index, got %s", str)
				}
			} else {
				return nil, fmt.Errorf("@value requires a string numeric index, got %v", args[0])
			}
		}
		dirValue = args
	case "key":
		if len(args) != 0 {
			return nil, fmt.Errorf("@key does not accept arguments, got %d arguments", len(args))
		}
		dirValue = args
	case "getenv":
		if len(args) < 1 || len(args) > 2 {
			return nil, fmt.Errorf("@getenv requires one name and optional default, got %d arguments", len(args))
		}
		if _, ok := args[0].(string); !ok {
			return nil, fmt.Errorf("@getenv requires a string name, got %v", args[0])
		}
		if len(args) == 2 {
			if _, ok := args[1].(string); !ok {
				return nil, fmt.Errorf("@getenv default must be a string, got %v", args[1])
			}
		}
		dirValue = args
	case "setenv":
		if len(args) != 2 {
			return nil, fmt.Errorf("@setenv requires exactly two arguments (name, value), got %d arguments", len(args))
		}
		if _, ok := args[0].(string); !ok {
			return nil, fmt.Errorf("@setenv requires a string name, got %v", args[0])
		}
		if _, ok := args[1].(string); !ok {
			return nil, fmt.Errorf("@setenv requires a string value, got %v", args[1])
		}
		dirValue = args
	case "stash", "buf":
		if len(args) != 2 {
			return nil, fmt.Errorf("@%s requires exactly two arguments (name, value), got %d arguments", name, len(args))
		}
		if _, ok := args[0].(string); !ok {
			return nil, fmt.Errorf("@%s requires a string name, got %v", name, args[0])
		}
		dirValue = args
	case "try":
		if len(args) > 1 {
			return nil, fmt.Errorf("@try accepts at most one number, got %d arguments", len(args))
		}
		if len(args) == 1 {
			if str, ok := args[0].(string); ok {
				if strings.Trim(str, "0123456789") != "" {
					return nil, fmt.Errorf("@try requires a numeric argument, got %s", str)
				}
			} else {
				return nil, fmt.Errorf("@try requires a string numeric argument, got %v", args[0])
			}
		} else {
			dirValue = []interface{}{"1"} // Default to 1 retry
		}
		dirValue = args
	case "catch":
		if len(args) != 0 {
			return nil, fmt.Errorf("@catch does not accept arguments, got %d arguments", len(args))
		}
		dirValue = args
	case "fail":
		if len(args) > 1 {
			return nil, fmt.Errorf("@fail accepts at most one message, got %d arguments", len(args))
		}
		if len(args) == 1 {
			if _, ok := args[0].(string); !ok {
				return nil, fmt.Errorf("@fail requires a string message, got %v", args[0])
			}
		}
		dirValue = args
	case "trigger":
		if len(args) < 1 {
			return nil, fmt.Errorf("@trigger requires at least one job ID, got %d arguments", len(args))
		}
		if _, ok := args[0].(string); !ok {
			return nil, fmt.Errorf("@trigger requires a string job ID, got %v", args[0])
		}
		dirValue = args
	case "if", "ifempty", "iflen":
		if len(args) != 1 {
			return nil, fmt.Errorf("@%s requires exactly one argument, got %d arguments", name, len(args))
		}
		if _, ok := args[0].(string); !ok {
			return nil, fmt.Errorf("@%s requires a string argument, got %v", name, args[0])
		}
		dirValue = args
	case "file_read":
		if len(args) < 1 || len(args) > 2 {
			return nil, fmt.Errorf("@file_read requires one path and optional encoding, got %d arguments", len(args))
		}
		if _, ok := args[0].(string); !ok {
			return nil, fmt.Errorf("@file_read requires a string path, got %v", args[0])
		}
		if len(args) == 2 {
			if _, ok := args[1].(string); !ok {
				return nil, fmt.Errorf("@file_read encoding must be a string, got %v", args[1])
			}
		}
		dirValue = args
	case "regexmatch":
		if len(args) != 2 {
			return nil, fmt.Errorf("@regexmatch requires exactly two arguments (pattern, target), got %d arguments", len(args))
		}
		if _, ok := args[0].(string); !ok {
			return nil, fmt.Errorf("@regexmatch requires a string pattern, got %v", args[0])
		}
		if _, ok := args[1].(string); !ok {
			return nil, fmt.Errorf("@regexmatch requires a string target, got %v", args[1])
		}
		dirValue = args
	case "defer":
		if len(args) != 0 {
			return nil, fmt.Errorf("@defer does not accept arguments, got %d arguments", len(args))
		}
		dirValue = args
	case "once":
		if len(args) != 0 {
			return nil, fmt.Errorf("@once does not accept arguments, got %d arguments", len(args))
		}
		dirValue = []interface{}{true}
	case "output":
		if len(args) < 1 {
			return nil, fmt.Errorf("@output requires at least one ID, got %d arguments", len(args))
		}
		if _, ok := args[0].(string); !ok {
			return nil, fmt.Errorf("@output requires a string ID, got %v", args[0])
		}
		dirValue = args
	default:
		Warn(fmt.Sprintf("Unknown command directive @%s", name), j.Filename, 0, 0)
		dirValue = args
	}

	return &Function{
		Type: name,
		Args: dirValue,
	}, nil
}

// AddJobCommand creates a Command
func (j *Job) AddJobCommand(cmd Command) {
	j.Commands = append(j.Commands, cmd)
}
