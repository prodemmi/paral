package functions

import (
	"fmt"
	"strings"
)

func Trim(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return "", nil
	}

	return strings.TrimSpace(fmt.Sprint(args[0])), nil
}
