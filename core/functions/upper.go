package functions

import (
	"fmt"
	"strings"
)

func Upper(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return "", nil
	}
	return strings.ToUpper(fmt.Sprint(args[0])), nil
}
