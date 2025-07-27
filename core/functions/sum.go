package functions

import "fmt"

func (f *Function) sum() (interface{}, error) {
	if len(f.argResults) != 1 {
		return nil, fmt.Errorf("sum: requires exactly 1 argument (list or matrix)")
	}
	switch v := f.argResults[0].(type) {
	case []interface{}:
		if len(v) == 0 {
			return 0, nil
		}
		sumVal, err := ToFloat64(v[0])
		if err != nil {
			return nil, fmt.Errorf("sum: list elements must be numbers, got %v", v[0])
		}
		for _, item := range v[1:] {
			val, err := ToFloat64(item)
			if err != nil {
				return nil, fmt.Errorf("sum: list elements must be numbers, got %v", item)
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
			return nil, fmt.Errorf("sum: matrix elements must be numbers, got %v", v[0][0])
		}
		for _, row := range v {
			for _, item := range row {
				val, err := ToFloat64(item)
				if err != nil {
					return nil, fmt.Errorf("sum: matrix elements must be numbers, got %v", item)
				}
				sumVal += val
			}
		}
		if IsInt(sumVal) {
			return int(sumVal), nil
		}
		return sumVal, nil
	default:
		return nil, fmt.Errorf("sum: argument must be a list or matrix")
	}
}
