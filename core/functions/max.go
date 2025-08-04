package functions

import (
	"fmt"
)

func Max(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("max: at least one argument required")
	}

	if len(args) == 1 {
		switch v := args[0].(type) {
		case []interface{}:
			if len(v) == 0 {
				return nil, fmt.Errorf("max: empty list")
			}
			maxVal, err := ToFloat64(v[0])
			if err != nil {
				return nil, fmt.Errorf("max: list elements must be numbers, got %v", v[0])
			}
			for _, item := range v[1:] {
				val, err := ToFloat64(item)
				if err != nil {
					return nil, fmt.Errorf("max: list elements must be numbers, got %v", item)
				}
				if val > maxVal {
					maxVal = val
				}
			}
			if IsInt(maxVal) {
				return int(maxVal), nil
			}
			return maxVal, nil
		case [][]interface{}:
			if len(v) == 0 || len(v[0]) == 0 {
				return nil, fmt.Errorf("max: empty matrix")
			}
			maxVal, err := ToFloat64(v[0][0])
			if err != nil {
				return nil, fmt.Errorf("max: matrix elements must be numbers, got %v", v[0][0])
			}
			for _, row := range v {
				for _, item := range row {
					val, err := ToFloat64(item)
					if err != nil {
						return nil, fmt.Errorf("max: matrix elements must be numbers, got %v", item)
					}
					if val > maxVal {
						maxVal = val
					}
				}
			}
			if IsInt(maxVal) {
				return int(maxVal), nil
			}
			return maxVal, nil
		default:
			return nil, fmt.Errorf("max: first argument must be a list or matrix")
		}
	}

	maxVal, err := ToFloat64(args[0])
	if err != nil {
		return nil, fmt.Errorf("max: arguments must be numbers, got %v", args[0])
	}
	for _, arg := range args[1:] {
		val, err := ToFloat64(arg)
		if err != nil {
			return nil, fmt.Errorf("max: arguments must be numbers, got %v", arg)
		}
		if val > maxVal {
			maxVal = val
		}
	}
	if IsInt(maxVal) {
		return int(maxVal), nil
	}
	return maxVal, nil
}
