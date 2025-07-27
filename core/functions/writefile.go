package functions

import (
	"fmt"
	"os"
	"strings"
)

func (f *Function) writefile() (interface{}, error) {
	if len(f.argResults) != 2 {
		return nil, fmt.Errorf("@writefile requires exactly 2 arguments (filename string, data string)")
	}

	filenameVal := f.argResults[0]
	filename, ok := filenameVal.(string)
	if !ok || strings.TrimSpace(filename) == "" {
		return nil, fmt.Errorf("@writefile: first argument must be a non-empty string (filename)")
	}

	dataVal := f.argResults[1]
	data, ok := dataVal.(string)
	if !ok {
		return nil, fmt.Errorf("@writefile: second argument must be a string (data to write)")
	}

	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return nil, err
	}

	return "ok", nil
}
