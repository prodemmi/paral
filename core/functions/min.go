package functions

import (
	"fmt"
)

func Min(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("min: at least one argument required")
	}

	if len(args) == 1 {
		switch v := args[0].(type) {
		case []interface{}:
			if len(v) == 0 {
				return nil, fmt.Errorf("min: empty list")
			}
			minVal, err := ToFloat64(v[0])
			if err != nil {
				return nil, fmt.Errorf("min: list elements must be numbers, got %v", v[0])
			}
			for _, item := range v[1:] {
				val, err := ToFloat64(item)
				if err != nil {
					return nil, fmt.Errorf("min: list elements must be numbers, got %v", item)
				}
				if val < minVal {
					minVal = val
				}
			}
			if IsInt(minVal) {
				return int(minVal), nil
			}
			return minVal, nil
		case [][]interface{}:
			if len(v) == 0 || len(v[0]) == 0 {
				return nil, fmt.Errorf("min: empty matrix")
			}
			minVal, err := ToFloat64(v[0][0])
			if err != nil {
				return nil, fmt.Errorf("min: matrix elements must be numbers, got %v", v[0][0])
			}
			for _, row := range v {
				for _, item := range row {
					val, err := ToFloat64(item)
					if err != nil {
						return nil, fmt.Errorf("min: matrix elements must be numbers, got %v", item)
					}
					if val < minVal {
						minVal = val
					}
				}
			}
			if IsInt(minVal) {
				return int(minVal), nil
			}
			return minVal, nil
		default:
			return nil, fmt.Errorf("min: first argument must be a list or matrix")
		}
	}

	minVal, err := ToFloat64(args[0])
	if err != nil {
		return nil, fmt.Errorf("min: arguments must be numbers, got %v", args[0])
	}
	for _, arg := range args[1:] {
		val, err := ToFloat64(arg)
		if err != nil {
			return nil, fmt.Errorf("min: arguments must be numbers, got %v", arg)
		}
		if val < minVal {
			minVal = val
		}
	}
	if IsInt(minVal) {
		return int(minVal), nil
	}
	return minVal, nil
}
