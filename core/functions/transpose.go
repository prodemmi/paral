package functions

import "fmt"

func (f *Function) transpose() (interface{}, error) {
	if len(f.argResults) != 1 {
		return nil, fmt.Errorf("transpose: requires exactly 1 argument (matrix)")
	}
	matrix, ok := f.argResults[0].([][]interface{})
	if !ok {
		return nil, fmt.Errorf("transpose: argument must be a matrix")
	}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]interface{}{}, nil
	}
	cols := len(matrix[0])
	for _, row := range matrix[1:] {
		if len(row) != cols {
			return nil, fmt.Errorf("transpose: matrix rows must have equal length")
		}
	}
	rows := len(matrix)
	result := make([][]interface{}, cols)
	for i := range result {
		result[i] = make([]interface{}, rows)
		for j := range matrix {
			result[i][j] = matrix[j][i]
		}
	}
	return result, nil
}
