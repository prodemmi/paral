package functions

import (
	"fmt"
)

func Buf(bufValue interface{}, args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("@buf(): missing buf name argument")
	}

	_, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("@buf(): name must be a string")
	}

	return bufValue, nil
}
