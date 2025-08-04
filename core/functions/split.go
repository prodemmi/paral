package functions

import (
	"fmt"
	"strings"
)

func Split(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("split: requires exactly 2 arguments (separator, input)")
	}

	separator, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("split: first argument must be a string (separator)")
	}
	input, ok := args[1].(string)
	if !ok {
		return nil, fmt.Errorf("split: second argument must be a string (input)")
	}
	result := strings.Split(input, separator)
	interfaceResult := make([]interface{}, len(result))
	for i, v := range result {
		interfaceResult[i] = v
	}
	return interfaceResult, nil
}
