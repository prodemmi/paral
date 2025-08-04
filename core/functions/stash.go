package functions

import (
	"fmt"
)

func Stash(stashValue interface{}, args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("@stash(): missing stash name argument")
	}

	_, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("@stash(): name must be a string")
	}

	return stashValue, nil
}
