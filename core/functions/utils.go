package functions

import (
	"fmt"
	"math"
	"paral/core/variable"
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

func ResolveValue(arg interface{}, forValue interface{}, forIndex int) (interface{}, error) {
	if str, ok := arg.(string); ok {
		if str == "@value" && forValue != nil {
			return forValue, nil
		}
		if str == "@key" && forIndex >= 0 {
			return forIndex, nil
		}
	}

	switch v := arg.(type) {
	case variable.Variable:
		switch val := v.Value.(type) {
		case variable.StringValue:
			return val.Value, nil
		case variable.IntValue:
			return val.Value, nil
		case variable.FloatValue:
			return val.Value, nil
		case variable.BoolValue:
			return val.Value, nil
		case variable.ListValue:
			return val.Value, nil
		case variable.MatrixValue:
			return val.Value, nil
		default:
			return val, nil
		}
	case Function:
		v.forValues = forValue
		v.forIndex = forIndex
		return v.Call()
	case *Function:
		v.forValues = forValue
		v.forIndex = forIndex
		return v.Call()
	case []interface{}:
		resolved := make([]interface{}, len(v))
		for i, item := range v {
			val, err := ResolveValue(item, forValue, forIndex)
			if err != nil {
				return nil, err
			}
			resolved[i] = val
		}
		return resolved, nil
	case [][]interface{}:
		resolved := make([][]interface{}, len(v))
		for i, inner := range v {
			resolved[i] = make([]interface{}, len(inner))
			for j, item := range inner {
				val, err := ResolveValue(item, forValue, forIndex)
				if err != nil {
					return nil, err
				}
				resolved[i][j] = val
			}
		}
		return resolved, nil
	default:
		return v, nil
	}
}
