package runtime

import (
	"fmt"
	"paral/core/metadata"
	"paral/core/variable"
)

type Expression struct {
	Value    interface{}
	RawText  string
	Metadata *metadata.Metadata
	runtime  *Runtime
}

func NewLiteralExpression(literal string, runtime *Runtime, mt *metadata.Metadata) *Expression {
	return &Expression{
		Value:    literal,
		RawText:  literal,
		runtime:  runtime,
		Metadata: mt,
	}
}

func (e *Expression) EvaluateValue(loopContext *TaskLoopContext) (interface{}, error) {
	switch e.RawText {
	case "hich", "HICH":
		return nil, nil
	case "@error":
		if e.runtime.HasTryCatchError() {
			return e.runtime.GetTryCatchError(), nil
		}
		return "", nil
	case "@value":
		if loopContext != nil {
			return loopContext.Value, nil
		}
		return nil, fmt.Errorf("failed to find loop context %q", e.RawText)
	case "@key":
		if loopContext != nil {
			return loopContext.Key, nil
		}
		return nil, fmt.Errorf("failed to find loop context %q", e.RawText)
	}

	switch v := e.Value.(type) {
	case *Function:
		result, err := v.Call()
		if err != nil {
			return nil, err
		}
		return result, nil
	case Function:
		result, err := v.Call()
		if err != nil {
			return nil, err
		}
		return result, nil
	case variable.Variable:
		return e.extractPureValue(v.Value), nil
	case *variable.Variable:
		return e.extractPureValue(v.Value), nil
	}
	return e.extractPureValue(e.Value), nil
}

func (e *Expression) extractPureValue(value interface{}) interface{} {
	switch v := value.(type) {
	case variable.IntValue:
		return v.Value
	case variable.FloatValue:
		return v.Value
	case variable.StringValue:
		return v.Value
	case variable.BoolValue:
		return v.Value
	case variable.DurationValue:
		return v.Value
	case variable.ListValue:
		return v.Value
	case variable.MatrixValue:
		return v.Value
	case *variable.IntValue:
		return v.Value
	case *variable.FloatValue:
		return v.Value
	case *variable.StringValue:
		return v.Value
	case *variable.BoolValue:
		return v.Value
	case *variable.DurationValue:
		return v.Value
	case *variable.ListValue:
		return v.Value
	case *variable.MatrixValue:
		return v.Value
	default:
		return value
	}
}

func (e *Expression) IsLiteral(literal string) bool {
	return e.RawText == literal || e.Value == literal
}

func (e *Expression) IsTrue(loopContext *TaskLoopContext) (bool, error) {
	if e == nil || e.Value == nil {
		return false, nil
	}
	evaluated, err := e.EvaluateValue(loopContext)
	if err != nil {
		return false, err
	}
	return resolveBoolean(evaluated, loopContext)
}

func resolveBoolean(input interface{}, loopContext *TaskLoopContext) (bool, error) {
	switch v := input.(type) {
	case bool:
		return v, nil
	case string:
		if v == "1" || v == "0" {
			return v == "1", nil
		}
		if v == "true" || v == "false" {
			return v == "true", nil
		}
		return len(v) > 0, nil

	case int:
		return v != 0, nil

	case float64:
		return v != 0.0, nil

	case []interface{}:
		return len(v) > 0, nil

	case [][]interface{}:
		return len(v) > 0, nil

	case map[string]interface{}:
		return true, nil

	case *Expression:
		return v.IsTrue(loopContext)

	default:
		return v != nil, nil
	}
}
