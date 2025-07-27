package functions

import (
	"fmt"
	"strings"
)

func (f *Function) unique() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, nil
	}
	seen := make(map[string]bool)
	switch v := f.argResults[0].(type) {
	case string:
		var result []string
		for _, r := range v {
			s := string(r)
			if !seen[s] {
				seen[s] = true
				result = append(result, s)
			}
		}
		return strings.Join(result, ""), nil
	case []interface{}:
		var result []interface{}
		for _, item := range v {
			s := fmt.Sprint(item)
			if !seen[s] {
				seen[s] = true
				result = append(result, item)
			}
		}
		return result, nil
	}
	return nil, nil
}
