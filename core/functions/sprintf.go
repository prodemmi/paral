package functions

import (
	"fmt"
)

func Sprintf(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return "", nil
	}

	format, ok := args[0].(string)
	if !ok {
		return "", fmt.Errorf("first argument must be string")
	}

	return fmt.Sprintf(format, args[1:]...), nil
}
