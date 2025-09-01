package functions

import (
	"fmt"
	"strconv"
)

func Printf(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return "", nil
	}
	format, ok := args[0].(string)
	if !ok {
		return "", fmt.Errorf("first argument must be string")
	}
	if unescaped, err := strconv.Unquote(`"` + format + `"`); err == nil {
		format = unescaped
	}
	newArgs := args[1:]
	for i, arg := range args {
		if str, ok := arg.(string); ok {
			if unescaped, err := strconv.Unquote(`"` + str + `"`); err == nil {
				args[i] = unescaped
			}
		}
	}
	return fmt.Sprintf(format, newArgs...), nil
}
