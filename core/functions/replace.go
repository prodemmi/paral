package functions

import (
	"fmt"
	"strings"
)

func (f *Function) replace() (interface{}, error) {
	if len(f.argResults) < 3 {
		return "", fmt.Errorf("replace: requires 3 arguments")
	}
	old := fmt.Sprint(f.argResults[0])
	newVal := fmt.Sprint(f.argResults[1])
	str := fmt.Sprint(f.argResults[2])
	return strings.ReplaceAll(str, old, newVal), nil
}
