package functions

import "fmt"

func (f *Function) det() (interface{}, error) {
	if len(f.argResults) != 1 {
		return nil, fmt.Errorf("det: requires exactly 1 argument (matrix)")
	}
	matrix, ok := f.argResults[0].([][]interface{})
	if !ok {
		return nil, fmt.Errorf("det: argument must be a matrix")
	}
	if len(matrix) == 0 {
		return nil, fmt.Errorf("det: empty matrix")
	}
	if len(matrix) != len(matrix[0]) {
		return nil, fmt.Errorf("det: matrix must be square")
	}
	n := len(matrix)
	m := make([][]float64, n)
	for i, row := range matrix {
		m[i] = make([]float64, n)
		for j, val := range row {
			fVal, err := ToFloat64(val)
			if err != nil {
				return nil, fmt.Errorf("det: matrix elements must be numbers, got %v", val)
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
