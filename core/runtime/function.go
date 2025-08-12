package runtime

import (
	"fmt"
	"paral/core"
	"paral/core/functions"
	"paral/core/metadata"
	"paral/core/variable"
	"strings"
)

type Function struct {
	Type     string
	Args     []interface{}
	Metadata metadata.Metadata
	RawText  string
	TaskID   string
	Runtime  *Runtime
}

func NewFunction(name string, mt metadata.Metadata, rawText string, taskID string, runtime *Runtime, args ...interface{}) *Function {
	f := &Function{
		Type:     name,
		Args:     args,
		Metadata: mt,
		TaskID:   taskID,
		Runtime:  runtime,
		RawText:  rawText,
	}
	return f
}

func NewTestFunction(name string, args ...interface{}) *Function {
	f := &Function{
		Type: name,
		Args: args,
		Metadata: metadata.Metadata{
			Filename: "test.paral",
			Content:  "",
		},
	}
	return f
}

func (f *Function) Call() (interface{}, error) {
	args, err := f.CallArgs(f.Args...)
	// TODO: fix get raw to return original raw text
	// f.RawText = f.GetRaw()

	if err != nil {
		return nil, err
	}

	switch f.Type {
	case "sprintf":
		return functions.Sprintf(args...)
	case "print":
		return functions.Print(args...)
	case "println":
		return functions.Println(args...)
	case "printf":
		return functions.Printf(args...)
	case "join":
		return functions.Join(args...)
	case "upper":
		return functions.Upper(args...)
	case "lower":
		return functions.Lower(args...)
	case "len":
		return functions.Len(args...)
	case "trim":
		return functions.Trim(args...)
	case "replace":
		return functions.Replace(args...)
	case "contains":
		return functions.Contains(args...)
	case "first":
		return functions.First(args...)
	case "last":
		return functions.Last(args...)
	case "unique":
		return functions.Unique(args...)
	case "reverse":
		return functions.Reverse(args...)
	case "writefile":
		return functions.Writefile(args...)
	case "split":
		return functions.Split(args...)
	case "append":
		return functions.Append(args...)
	case "regex":
		return functions.Regex(args...)
	case "slice":
		return functions.Slice(args...)
	case "time":
		return functions.Time(args...)
	case "middle":
		return functions.Middle(args...)
	case "min":
		return functions.Min(args...)
	case "max":
		return functions.Max(args...)
	case "transpose":
		return functions.Transpose(args...)
	case "det":
		return functions.Det(args...)
	case "sum":
		return functions.Sum(args...)
	case "trigger":
		return functions.Trigger(args...)
	case "getvar":
		return functions.Getvar(args...)
	case "stash":
		stashName := args[0].(string)
		stashValue := f.GetActiveStashValue(stashName)
		return functions.Stash(stashValue, args...)
	case "buf":
		bufName := args[0].(string)
		bufValue := f.GetBufValue(bufName)
		return functions.Buf(bufValue, args...)
	case "getenv":
		return functions.Getenv(args...)
	case "setenv":
		return functions.Setenv(args...)
	}

	return nil, fmt.Errorf("@%s: unknown function", f.Type)
}

func (f *Function) GetRaw() string {
	loopContext := f.GetActiveLoopContext()

	parts := []string{}
	for _, a := range f.Args {
		switch v := a.(type) {
		case Function:
			parts = append(parts, v.GetRaw())
		case *Function:
			parts = append(parts, v.GetRaw())
		case string:
			switch v {
			case "@value":
				if loopContext != nil {
					parts = append(parts, fmt.Sprintf("%v", loopContext.Value))
				} else {
					f.Runtime.Reporter.ThrowRuntimeError("loop context is nil: cannot resolve @value", &f.Metadata)
				}
			case "@key":
				if loopContext != nil {
					parts = append(parts, fmt.Sprintf("%d", loopContext.Key))
				} else {
					f.Runtime.Reporter.ThrowRuntimeError("loop context is nil: cannot resolve @key", &f.Metadata)
				}
			default:
				parts = append(parts, fmt.Sprintf(`"%s"`, v))
			}
		default:
			parts = append(parts, fmt.Sprintf("%v", v))
		}
	}
	return fmt.Sprintf("@%s(%s)", f.Type, strings.Join(parts, ","))
}

func (f *Function) GetResult() string {
	result, _ := f.Call()
	return fmt.Sprint(result)
}

func (f *Function) GetCalculatedArgsByIndex(index int) interface{} {
	args, _ := f.CallArgs(f.Args...)
	if index < len(args) {
		return args[index]
	}
	return nil
}

func (f *Function) CallArgs(args ...interface{}) ([]interface{}, error) {
	argResults := make([]interface{}, len(args))
	var err error
	for i := len(args) - 1; i >= 0; i-- {
		argResults[i], err = f.ResolveValue(args[i])
		if err != nil {
			return nil, err
		}
	}
	return argResults, nil
}

func (f *Function) ResolveValue(arg interface{}) (interface{}, error) {
	loopContext := f.GetActiveLoopContext()
	switch v := arg.(*Expression).Result.(type) {
	case string:
		if loopContext != nil {
			if v == "@value" {
				return loopContext.Value, nil
			}
			if v == "@key" {
				return loopContext.Key, nil
			}
		}
		return core.TrimQuotes(v), nil
	case variable.Variable:
		switch val := v.Value.(type) {
		case variable.StringValue:
			return val.Value, nil
		case variable.IntValue:
			return val.Value, nil
		case variable.FloatValue:
			return val.Value, nil
		case variable.BoolValue:
			return val.Value, nil
		case variable.ListValue:
			return val.Value, nil
		case variable.MatrixValue:
			return val.Value, nil
		default:
			return val, nil
		}
	case Function:
		return v.Call()
	case *Function:
		return v.Call()
	case []interface{}:
		resolved := make([]interface{}, len(v))
		for i, item := range v {
			val, err := f.ResolveValue(item)
			if err != nil {
				return nil, err
			}
			resolved[i] = val
		}
		return resolved, nil
	case [][]interface{}:
		resolved := make([][]interface{}, len(v))
		for i, inner := range v {
			resolved[i] = make([]interface{}, len(inner))
			for j, item := range inner {
				val, err := f.ResolveValue(item)
				if err != nil {
					return nil, err
				}
				resolved[i][j] = val
			}
		}
		return resolved, nil
	default:
		return v, nil
	}
}

func (f *Function) GetActiveStashValue(stashName string) interface{} {
	return f.Runtime.GetActiveStashValue(stashName, f.TaskID)
}

func (f *Function) GetActiveLoopContext() *TaskLoopContext {
	task := f.Runtime.GetTaskByID(f.TaskID)
	if task != nil {
		return task.GetActiveLoopContext()
	}
	return nil
}

func (f *Function) GetBufValue(bufName string) interface{} {
	return f.Runtime.GetBufValue(bufName)
}

func (f *Function) String() string {
	return f.GetRaw()
}
