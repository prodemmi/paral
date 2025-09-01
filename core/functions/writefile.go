package functions

import (
	"fmt"
	"os"
	"strings"
)

func Writefile(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("requires exactly 2 arguments (filename string, data string)")
	}

	filenameVal := args[0]
	filename, ok := filenameVal.(string)
	if !ok || strings.TrimSpace(filename) == "" {
		return nil, fmt.Errorf("first argument must be a non-empty string (filename)")
	}

	dataVal := args[1]
	data, ok := dataVal.(*string)
	if data != nil && !ok {
		return nil, fmt.Errorf("second argument must be a string (data to write)")
	}

	var content []byte
	if data != nil {
		content = []byte(*data)
	} else {
		content = make([]byte, 0)
	}

	if err := os.WriteFile(filename, content, 0644); err != nil {
		return nil, err
	}

	return "ok", nil
}
