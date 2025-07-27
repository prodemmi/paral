package functions

import (
	"fmt"
	"strings"
)

func (f *Function) contains() (interface{}, error) {
	if len(f.argResults) < 2 {
		return false, fmt.Errorf("contains: requires 2 arguments")
	}
	substr := fmt.Sprint(f.argResults[0])
	str := fmt.Sprint(f.argResults[1])
	return strings.Contains(str, substr), nil
}
