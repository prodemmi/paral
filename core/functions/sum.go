package functions

import (
	"fmt"
)

func Sum(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("requires exactly 1 argument (list or matrix)")
	}

	switch v := args[0].(type) {
	case []interface{}:
		if len(v) == 0 {
			return 0, nil
		}
		sumVal, err := ToFloat64(v[0])
		if err != nil {
			return nil, fmt.Errorf("list elements must be numbers, got %v", v[0])
		}
		for _, item := range v[1:] {
			val, err := ToFloat64(item)
			if err != nil {
				return nil, fmt.Errorf("list elements must be numbers, got %v", item)
			}
			sumVal += val
		}
		if IsInt(sumVal) {
			return int(sumVal), nil
		}
		return sumVal, nil
	case [][]interface{}:
		if len(v) == 0 || len(v[0]) == 0 {
			return 0, nil
		}
		sumVal, err := ToFloat64(v[0][0])
		if err != nil {
			return nil, fmt.Errorf("matrix elements must be numbers, got %v", v[0][0])
		}
		for _, row := range v {
			for _, item := range row {
				val, err := ToFloat64(item)
				if err != nil {
					return nil, fmt.Errorf("matrix elements must be numbers, got %v", item)
				}
				sumVal += val
			}
		}
		if IsInt(sumVal) {
			return int(sumVal), nil
		}
		return sumVal, nil
	default:
		return nil, fmt.Errorf("argument must be a list or matrix")
	}
}
