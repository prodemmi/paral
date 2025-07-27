package functions

import "fmt"

func (f *Function) sprintf() (interface{}, error) {
	if len(f.argResults) == 0 {
		return "", nil
	}
	format, ok := f.argResults[0].(string)
	if !ok {
		return "", fmt.Errorf("sprintf: first argument must be string")
	}
	args := []interface{}{}
	for _, a := range f.argResults[1:] {
		args = append(args, a)
	}
	return fmt.Sprintf(format, args...), nil
}
