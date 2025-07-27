package functions

import (
	"fmt"
	"strings"
)

func (f *Function) split() (interface{}, error) {
	if len(f.argResults) != 2 {
		return nil, fmt.Errorf("split: requires exactly 2 arguments (separator, input)")
	}
	separator, ok := f.argResults[0].(string)
	if !ok {
		return nil, fmt.Errorf("split: first argument must be a string (separator)")
	}
	input, ok := f.argResults[1].(string)
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
