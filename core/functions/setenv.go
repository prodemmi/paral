package functions

import (
	"fmt"
	"os"
)

func Setenv(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("@setenv() expects exactly 2 arguments (key, value), got %d", len(args))
	}

	key, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("first argument (key) must be a string")
	}

	value := args[1]
	strVal := fmt.Sprint(value)

	return nil, os.Setenv(key, strVal)
}
