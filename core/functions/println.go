package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func Println(args ...interface{}) (interface{}, error) {
	parts := []string{}

	for _, a := range args {
		str := fmt.Sprint(a)
		if unescaped, err := strconv.Unquote(`"` + str + `"`); err == nil {
			parts = append(parts, unescaped)
		} else {
			parts = append(parts, str)
		}
	}
	return strings.Join(parts, "") + "\n", nil
}
