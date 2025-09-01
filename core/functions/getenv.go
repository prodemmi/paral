package functions

import (
	"fmt"
	"os"
)

func Getenv(args ...interface{}) (interface{}, error) {
	if len(args) < 1 || len(args) > 2 {
		return nil, fmt.Errorf("expects 1 or 2 arguments, got %d", len(args))
	}

	envName, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("first argument must be a string (env name)")
	}

	var defaultValue *string
	if len(args) == 2 {
		defaultValue, ok = args[1].(*string)
		if !ok {
			return nil, fmt.Errorf("second argument must be *string (default value)")
		}
	}

	val, exists := os.LookupEnv(envName)
	if !exists {
		if defaultValue != nil {
			return *defaultValue, nil
		}
		return "", nil
	}
	return val, nil
}
