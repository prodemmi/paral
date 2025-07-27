package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func (f *Function) print() (interface{}, error) {
	parts := []string{}
	for _, a := range f.argResults {
		str := fmt.Sprint(a)
		if unescaped, err := strconv.Unquote(`"` + str + `"`); err == nil {
			parts = append(parts, unescaped)
		} else {
			parts = append(parts, str)
		}
	}
	f.result = strings.Join(parts, "")
	return f.result, nil
}
