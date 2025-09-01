package functions

import (
	"fmt"
	"regexp"
)

func Regex(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("requires exactly 2 arguments (regex, input)")
	}

	pattern, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("first argument must be a string (regex pattern)")
	}
	input, ok := args[1].(string)
	if !ok {
		return nil, fmt.Errorf("second argument must be a string (input)")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid pattern %s: %v", pattern, err)
	}
	matches := re.FindAllString(input, -1)
	if len(matches) == 0 {
		return nil, nil
	}
	result := make([]interface{}, len(matches))
	for i, match := range matches {
		result[i] = match
	}
	return result, nil
}
