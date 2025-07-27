package functions

import (
	"fmt"
	"strings"
)

func (f *Function) trim() (interface{}, error) {
	if len(f.argResults) == 0 {
		return "", nil
	}
	return strings.TrimSpace(fmt.Sprint(f.argResults[0])), nil
}
