package functions

import (
	"errors"
	"fmt"
	"paral/core/variable"
)

func Setvar(vars []*variable.Variable, varName string, args ...interface{}) (*variable.Variable, error) {
	if len(args) != 1 {
		return nil, errors.New("expected exactly 2 arguments")
	}

	value := args[0]

	for _, v := range vars {
		if v.Name == varName {
			switch v.Value.(type) {
			case variable.StringValue:
				strVal, ok := value.(string)
				if !ok {
					return nil, errors.New("expected string value for string variable")
				}
				v.Value = variable.StringValue{Value: strVal}

			case variable.IntValue:
				intVal, ok := value.(int)
				if !ok {
					return nil, errors.New("expected int value for int variable")
				}
				v.Value = variable.IntValue{Value: intVal}

			case variable.FloatValue:
				floatVal, ok := value.(float64)
				if !ok {
					return nil, errors.New("expected float64 value for float variable")
				}
				v.Value = variable.FloatValue{Value: floatVal}

			case variable.BoolValue:
				boolVal, ok := value.(bool)
				if !ok {
					return nil, errors.New("expected bool value for bool variable")
				}
				v.Value = variable.BoolValue{Value: boolVal}

			case variable.ListValue:
				listVal, ok := value.([]interface{})
				fmt.Println("listVal", listVal)
				if !ok {
					return nil, errors.New("expected []interface{} for list variable")
				}
				v.Value = variable.ListValue{Value: listVal}

			case variable.MatrixValue:
				matrixVal, ok := value.([][]interface{})
				if !ok {
					return nil, errors.New("expected [][]interface{} for matrix variable")
				}
				v.Value = variable.MatrixValue{Value: matrixVal}
			}
			return v, nil
		}
	}

	return nil, errors.New("variable not found: " + varName)
}
