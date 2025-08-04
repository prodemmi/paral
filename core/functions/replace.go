package functions

import (
	"fmt"
	"strings"
)

func Replace(args ...interface{}) (interface{}, error) {
	if len(args) < 3 {
		return "", fmt.Errorf("replace: requires 3 arguments")
	}

	old := fmt.Sprint(args[0])
	newVal := fmt.Sprint(args[1])
	str := fmt.Sprint(args[2])
	return strings.ReplaceAll(str, old, newVal), nil
}
