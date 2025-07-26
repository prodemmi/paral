package core

import (
	"fmt"
	"strings"
)

type VarBase struct {
	Name string
}

type Var struct {
	VarBase
	Value interface{}
}

// Type-specific structs
type MatrixValue struct {
	VarBase
	Value [][]interface{}
}

type ListValue struct {
	VarBase
	Value []interface{}
}

type IntValue struct {
	VarBase
	Value int
}

type FloatValue struct {
	VarBase
	Value float64
}

type StringValue struct {
	VarBase
	Value string
}

type BoolValue struct {
	VarBase
	Value bool
}

type DurationValue struct {
	VarBase
	Value string
}

// Type casting methods for Var
func (v *Var) AsMatrix() (*MatrixValue, bool) {
	if mv, ok := v.Value.(MatrixValue); ok {
		return &mv, true
	}
	return nil, false
}

func (v *Var) AsList() (*ListValue, bool) {
	if lv, ok := v.Value.(ListValue); ok {
		return &lv, true
	}
	return nil, false
}

func (v *Var) AsInt() (*IntValue, bool) {
	if iv, ok := v.Value.(IntValue); ok {
		return &iv, true
	}
	return nil, false
}

func (v *Var) AsFloat() (*FloatValue, bool) {
	if fv, ok := v.Value.(FloatValue); ok {
		return &fv, true
	}
	return nil, false
}

func (v *Var) AsString() (*StringValue, bool) {
	if sv, ok := v.Value.(StringValue); ok {
		return &sv, true
	}
	return nil, false
}

func (v *Var) AsBool() (*BoolValue, bool) {
	if bv, ok := v.Value.(BoolValue); ok {
		return &bv, true
	}
	return nil, false
}

func (v *Var) AsDuration() (*DurationValue, bool) {
	if dv, ok := v.Value.(DurationValue); ok {
		return &dv, true
	}
	return nil, false
}

// GetMatrixCombinations method with type checking
func (v *Var) GetMatrixCombinations() ([][]interface{}, error) {
	mv, ok := v.AsMatrix()
	if !ok {
		return nil, fmt.Errorf("value is not a MatrixValue")
	}

	matrixValues := mv.Value
	if len(matrixValues) == 0 {
		return nil, nil
	}

	result := [][]interface{}{{}}

	for _, list := range matrixValues {
		var newResult [][]interface{}
		for _, combo := range result {
			for _, val := range list {
				newCombo := append([]interface{}{}, combo...)
				newCombo = append(newCombo, val)
				newResult = append(newResult, newCombo)
			}
		}
		result = newResult
	}
	return result, nil
}

func (v *Var) Format() (typeStr string, valueStr string) {
	switch val := v.Value.(type) {
	case IntValue:
		return "int", fmt.Sprintf("%d", val.Value)
	case FloatValue:
		return "float", fmt.Sprintf("%f", val.Value)
	case StringValue:
		return "string", fmt.Sprint(val.Value)
	case BoolValue:
		return "bool", fmt.Sprintf("%t", val.Value)
	case DurationValue:
		return "duration", val.Value
	case ListValue:
		return "list", fmt.Sprintf("%v", val.Value)
	case MatrixValue:
		combinations, err := v.GetMatrixCombinations()
		if err != nil {
			return "matrix", fmt.Sprintf("error: %v", err)
		}
		if len(combinations) == 0 {
			return "matrix", "[]"
		}
		var b strings.Builder
		//b.WriteString("[\n")
		//for i, combo := range combinations {
		//	b.WriteString(fmt.Sprintf("      • %v", combo))
		//	if i < len(combinations)-1 {
		//		b.WriteString("\n")
		//	}
		//}
		//b.WriteString("\n    ]")
		return "matrix", fmt.Sprintf("%v\n    %s", val.Value, b.String())
	default:
		return "unknown", fmt.Sprintf("%v", v.Value)
	}
}
