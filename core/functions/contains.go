package functions

import (
	"fmt"
	"strings"
)

func Contains(args ...interface{}) (interface{}, error) {
	if len(args) < 2 {
		return false, fmt.Errorf("contains: requires 2 arguments")
	}

	str := fmt.Sprint(args[0])
	substr := fmt.Sprint(args[1])
	return strings.Contains(str, substr), nil
}
