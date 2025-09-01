package functions

import (
	"errors"
	"paral/core/variable"
)

func Getvar(vars []*variable.Variable, args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, errors.New("expected exactly one argument (variable name)")
	}

	name, ok := args[0].(string)
	if !ok {
		return nil, errors.New("argument must be a string (variable name)")
	}

	for _, v := range vars {
		if v.Name == name {
			return v.Value, nil
		}
	}

	return nil, errors.New("variable not found: " + name)
}
