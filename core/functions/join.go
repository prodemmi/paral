package functions

import (
	"fmt"
	"strings"
)

func Join(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return "", nil
	}

	var result []string
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			result = append(result, v)
		case []interface{}:
			for _, item := range v {
				result = append(result, fmt.Sprint(item))
			}
		case [][]interface{}:
			for _, list := range v {
				for _, item := range list {
					result = append(result, fmt.Sprint(item))
				}
			}
		default:
			result = append(result, fmt.Sprint(v))
		}
	}
	return strings.Join(result, " "), nil
}
