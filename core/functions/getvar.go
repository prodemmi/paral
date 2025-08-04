package functions

import (
	"errors"
)

func Getvar(args ...interface{}) (interface{}, error) {
	if len(args) > 1 || len(args) == 0 {
		return nil, errors.New("expected exactly one argument: variable name")
	}

	result, ok := args[0].(string)
	if !ok {
		return nil, errors.New("first argument must be a variable name")
	}
	return result, nil
}
