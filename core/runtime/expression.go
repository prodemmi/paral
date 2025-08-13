package runtime

import (
	"paral/core/metadata"
)

type Expression struct {
	Result   interface{}
	RawText  string
	Metadata *metadata.Metadata
}

func (e *Expression) IsTrue() bool {
	if e == nil || e.Result == nil {
		return false
	}

	switch v := e.Result.(type) {
	case *Function:
		result, err := v.Call()
		if err != nil {
			return false
		}
		return resolveBoolean(result)
	case Function:
		result, err := v.Call()
		if err != nil {
			return false
		}
		return resolveBoolean(result)
	}

	return resolveBoolean(e.Result)
}

func resolveBoolean(input interface{}) bool {
	switch v := input.(type) {
	case bool:
		return v

	case string:
		if v == "1" || v == "0" {
			return v == "1"
		}
		if v == "true" || v == "false" {
			return v == "true"
		}
		return len(v) > 0

	case int:
		return v != 0

	case float64:
		return v != 0.0

	case []interface{}:
		return len(v) > 0

	case [][]interface{}:
		return len(v) > 0

	case map[string]interface{}:
		return true

	case *Expression:
		return v.IsTrue()

	default:
		return v != nil
	}
}
