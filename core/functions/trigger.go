package functions

import (
	"errors"
	"fmt"
)

func Trigger(args ...interface{}) (interface{}, error) {
	if len(args) < 1 {
		return nil, errors.New("no arguments provided")
	}

	jobName, ok := args[0].(string)
	if !ok || jobName == "" {
		return nil, errors.New("task id not provided or invalid")
	}
	return fmt.Sprintf("task '%s' executed", jobName), nil
}
