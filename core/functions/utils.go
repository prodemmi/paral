package functions

import (
	"fmt"
	"math"
)

func IsInt(val float64) bool {
	return math.Mod(val, 1) == 0
}

func Determinant(m [][]float64) float64 {
	n := len(m)
	if n == 1 {
		return m[0][0]
	}
	if n == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}
	det := 0.0
	for j := 0; j < n; j++ {
		det += m[0][j] * float64(Cofactor(m, 0, j))
	}
	return det
}

func Cofactor(m [][]float64, row, col int) float64 {
	subMatrix := make([][]float64, len(m)-1)
	for i := 0; i < len(m)-1; i++ {
		subMatrix[i] = make([]float64, len(m)-1)
		for j := 0; j < len(m)-1; j++ {
			r := i
			if i >= row {
				r++
			}
			c := j
			if j >= col {
				c++
			}
			subMatrix[i][j] = m[r][c]
		}
	}
	sign := 1.0
	if (row+col)%2 != 0 {
		sign = -1.0
	}
	return sign * Determinant(subMatrix)
}

func ToFloat64(val interface{}) (float64, error) {
	switch v := val.(type) {
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		return 0, fmt.Errorf("cannot convert string to number: %v", v)
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}
