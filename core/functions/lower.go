package functions

import (
	"fmt"
	"strings"
)

func Lower(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return "", nil
	}

	return strings.ToLower(fmt.Sprint(args[0])), nil
}
