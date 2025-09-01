package functions

import (
	"fmt"
)

func Array(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("requires at least 1 argument (array)")
	}

	// Unwrap first arg if it's a wrapper type
	first := ExtractPureValue(args[0])

	// Case 1: if only one argument and it's already a slice → just return it
	if len(args) == 1 {
		switch v := first.(type) {
		case []interface{}:
			return v, nil
		default:
			return []interface{}{first}, nil
		}
	}

	// Case 2: multiple args → return all as []interface{}
	unwrapped := make([]interface{}, len(args))
	for i, a := range args {
		unwrapped[i] = ExtractPureValue(a)
	}
	return unwrapped, nil
}
