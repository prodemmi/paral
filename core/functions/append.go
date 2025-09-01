package functions

import (
	"fmt"
)

func Append(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("requires exactly 2 arguments (input, appendTo)")
	}
	appendTo := args[1]
	switch input := args[0].(type) {
	case string:
		appendStr, ok := appendTo.(string)
		if !ok {
			return nil, fmt.Errorf("second argument must be a string when input is a string")
		}
		return input + appendStr, nil
	case []interface{}:
		result := make([]interface{}, len(input)+1)
		copy(result, input)
		result[len(input)] = appendTo
		return result, nil
	case [][]interface{}:
		newRow, ok := appendTo.([]interface{})
		if !ok {
			return nil, fmt.Errorf("second argument must be a list when input is a matrix")
		}
		result := make([][]interface{}, len(input)+1)
		copy(result, input)
		result[len(input)] = newRow
		return result, nil
	default:
		return nil, fmt.Errorf("first argument must be a string, list, or matrix")
	}
}
