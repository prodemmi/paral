package functions

import (
	"fmt"
)

func Slice(args ...interface{}) (interface{}, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("slice: requires exactly 3 arguments (input, from, to)")
	}

	from, ok := args[1].(int)
	if !ok {
		return nil, fmt.Errorf("slice: second argument must be an integer (from index)")
	}
	to, ok := args[2].(int)
	if !ok {
		return nil, fmt.Errorf("slice: third argument must be an integer (to index)")
	}
	if from < 0 || from > to {
		return nil, fmt.Errorf("slice: invalid indices (from=%d, to=%d)", from, to)
	}
	switch input := args[0].(type) {
	case string:
		if to > len(input) {
			return nil, fmt.Errorf("slice: to index (%d) exceeds string length (%d)", to, len(input))
		}
		return input[from:to], nil
	case []interface{}:
		if to > len(input) {
			return nil, fmt.Errorf("slice: to index (%d) exceeds list length (%d)", to, len(input))
		}
		return input[from:to], nil
	case [][]interface{}:
		if to > len(input) {
			return nil, fmt.Errorf("slice: to index (%d) exceeds matrix length (%d)", to, len(input))
		}
		return input[from:to], nil
	default:
		return nil, fmt.Errorf("slice: first argument must be a string, list, or matrix")
	}
}
