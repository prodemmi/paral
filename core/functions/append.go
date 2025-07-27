package functions

import "fmt"

func (f *Function) append() (interface{}, error) {
	if len(f.argResults) != 2 {
		return nil, fmt.Errorf("append: requires exactly 2 arguments (input, appendTo)")
	}
	appendTo := f.argResults[1]
	switch input := f.argResults[0].(type) {
	case string:
		appendStr, ok := appendTo.(string)
		if !ok {
			return nil, fmt.Errorf("append: second argument must be a string when input is a string")
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
			return nil, fmt.Errorf("append: second argument must be a list when input is a matrix")
		}
		result := make([][]interface{}, len(input)+1)
		copy(result, input)
		result[len(input)] = newRow
		return result, nil
	default:
		return nil, fmt.Errorf("append: first argument must be a string, list, or matrix")
	}
}
