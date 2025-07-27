package functions

import (
	"fmt"
	"strings"
)

func (f *Function) upper() (interface{}, error) {
	if len(f.argResults) == 0 {
		return "", nil
	}
	return strings.ToUpper(fmt.Sprint(f.argResults[0])), nil
}
