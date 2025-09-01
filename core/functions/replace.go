package functions

import (
	"fmt"
	"strings"
)

func Replace(args ...interface{}) (interface{}, error) {
	if len(args) < 3 {
		return "", fmt.Errorf("requires 3 arguments")
	}

	str := fmt.Sprint(args[0])
	old := fmt.Sprint(args[1])
	newVal := fmt.Sprint(args[2])
	return strings.ReplaceAll(str, old, newVal), nil
}
