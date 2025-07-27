package functions

import (
	"fmt"
	"strconv"
)

func (f *Function) printf() (interface{}, error) {
	if len(f.argResults) == 0 {
		f.result = ""
		return f.result, nil
	}
	format, ok := f.argResults[0].(string)
	if !ok {
		f.result = ""
		return f.result, fmt.Errorf("printf: first argument must be string")
	}
	if unescaped, err := strconv.Unquote(`"` + format + `"`); err == nil {
		format = unescaped
	}
	args := f.argResults[1:]
	for i, arg := range args {
		if str, ok := arg.(string); ok {
			if unescaped, err := strconv.Unquote(`"` + str + `"`); err == nil {
				args[i] = unescaped
			}
		}
	}
	f.result = fmt.Sprintf(format, args...)
	return f.result, nil
}
