package functions

import (
	"fmt"
)

func Det(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("requires exactly 1 argument (matrix)")
	}

	matrix, ok := args[0].([][]interface{})
	if !ok {
		return nil, fmt.Errorf("argument must be a matrix")
	}
	if len(matrix) == 0 {
		return nil, fmt.Errorf("empty matrix")
	}
	if len(matrix) != len(matrix[0]) {
		return nil, fmt.Errorf("matrix must be square")
	}
	n := len(matrix)
	m := make([][]float64, n)
	for i, row := range matrix {
		m[i] = make([]float64, n)
		for j, val := range row {
			fVal, err := ToFloat64(val)
			if err != nil {
				return nil, fmt.Errorf("matrix elements must be numbers, got %v", val)
			}
			m[i][j] = fVal
		}
	}
	detVal := Determinant(m)
	if IsInt(detVal) {
		return int(detVal), nil
	}
	return detVal, nil
}
