package functions

import (
	"fmt"
	"strings"
)

func (f *Function) lower() (interface{}, error) {
	if len(f.argResults) == 0 {
		return "", nil
	}
	return strings.ToLower(fmt.Sprint(f.argResults[0])), nil
}
