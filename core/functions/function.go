package functions

import (
	"fmt"
	"paral/core/metadata"
	"strings"
)

type Function struct {
	Type       string
	Args       []interface{}
	Report     string
	Metadata   metadata.Metadata
	raw        string
	result     interface{}
	argResults []interface{}
	forValues  interface{}
	forIndex   int
}

func NewFunction(name string, mt metadata.Metadata, args ...interface{}) *Function {
	f := &Function{
		Type:     name,
		Args:     args,
		Metadata: mt,
	}
	f.setRaw()
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
	f.setRaw()
	return f
}

func (f *Function) Call() (interface{}, error) {
	var err error

	err = f.callArgs()
	if err != nil {
		return nil, err
	}

	switch f.Type {
	case "sprintf":
		f.result, err = f.sprintf()
	case "print":
		f.result, err = f.print()
	case "println":
		f.result, err = f.println()
	case "printf":
		f.result, err = f.printf()
	case "join":
		f.result, err = f.join()
	case "upper":
		f.result, err = f.upper()
	case "lower":
		f.result, err = f.lower()
	case "len":
		f.result, err = f.len()
	case "trim":
		f.result, err = f.trim()
	case "replace":
		f.result, err = f.replace()
	case "contains":
		f.result, err = f.contains()
	case "first":
		f.result, err = f.first()
	case "last":
		f.result, err = f.last()
	case "unique":
		f.result, err = f.unique()
	case "reverse":
		f.result, err = f.reverse()
	case "writefile":
		f.result, err = f.writefile()
	case "split":
		f.result, err = f.split()
	case "append":
		f.result, err = f.append()
	case "regex":
		f.result, err = f.regex()
	case "slice":
		f.result, err = f.slice()
	case "time":
		f.result, err = f.time()
	case "middle":
		f.result, err = f.middle()
	case "min":
		f.result, err = f.min()
	case "max":
		f.result, err = f.max()
	case "transpose":
		f.result, err = f.transpose()
	case "det":
		f.result, err = f.det()
	case "sum":
		f.result, err = f.sum()
	case "trigger":
		f.result, err = f.trigger()
	case "getenv":
		var envName string
		var defaultVal interface{} = nil

		if len(f.argResults) == 0 {
			return nil, fmt.Errorf("getenv: missing environment variable name argument")
		}

		if s, ok := f.argResults[0].(string); ok {
			envName = s
		} else {
			return nil, fmt.Errorf("getenv: first argument must be a string")
		}

		if len(f.argResults) > 1 {
			defaultVal = f.argResults[1]
		}

		f.result, err = f.getenv(envName, defaultVal)
	case "setenv":
		if len(f.argResults) < 2 {
			return nil, fmt.Errorf("setenv: requires at least 2 arguments")
		}
		key, ok := f.argResults[0].(string)
		if !ok {
			return nil, fmt.Errorf("setenv: first argument must be a string")
		}
		f.result, err = f.setenv(key, f.argResults[1])
	default:
		f.result = f.raw
		err = fmt.Errorf("@%s: unknown function", f.Type)
	}

	return f.result, err
}

func (f *Function) callArgs() error {
	f.argResults = make([]interface{}, len(f.Args))
	var err error
	for i := len(f.Args) - 1; i >= 0; i-- {
		f.argResults[i], err = ResolveValue(f.Args[i], f.forValues, f.forIndex)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Function) GetRaw() string {
	parts := []string{}
	for _, a := range f.Args {
		switch v := a.(type) {
		case *Function:
			parts = append(parts, v.GetRaw())
		case string:
			if v == "@value" && f.forValues != nil {
				parts = append(parts, fmt.Sprintf("%v", f.forValues))
			} else if v == "@key" && f.forIndex >= 0 {
				parts = append(parts, fmt.Sprintf("%d", f.forIndex))
			} else {
				parts = append(parts, fmt.Sprintf(`"%s"`, v))
			}
		default:
			parts = append(parts, fmt.Sprintf("%v", v))
		}
	}
	return fmt.Sprintf("@%s(%s)", f.Type, strings.Join(parts, ","))
}

func (f *Function) setRaw() {
	f.raw = f.GetRaw()
}

func (f *Function) GetResult() interface{} {
	return f.result
}

func (f *Function) GetCalculatedArgs() []interface{} {
	return f.argResults
}

func (f *Function) GetCalculatedArgsByIndex(index int) interface{} {
	return f.argResults[index]
}

func (f *Function) SetLoopData(forValue interface{}, forIndex int) {
	f.forIndex = forIndex
	f.forValues = forValue
}

func (f *Function) String() string {
	if f.result != nil {
		return fmt.Sprint(f.result)
	}
	return f.GetRaw()
}
