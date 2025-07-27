package functions

import (
	"errors"
	"fmt"
)

func (f *Function) trigger() (interface{}, error) {
	if len(f.argResults) < 1 {
		return nil, errors.New("no arguments provided")
	}
	jobName, ok := f.argResults[0].(string)
	if !ok || jobName == "" {
		return nil, errors.New("task id not provided or invalid")
	}
	return fmt.Sprintf("task '%s' executed", jobName), nil
}
