package runtime

type IfCondition struct {
	Type          string
	Expression    *Expression
	ThenPipelines []*TaskPipeline
	ElsePipelines []*TaskPipeline
}

func (i *IfCondition) IsTrue() bool {
	if i.Expression == nil || i.Expression.Result == nil {
		return false
	}

	switch v := i.Expression.Result.(type) {
	case bool:
		// Direct boolean value (from BOOLEAN or ZERO_ONE)
		return v

	case string:
		if v == "true" || v == "false" {
			return v == "true"
		}

		// Handle string_expr, URL, IDENTIFIER, LOOP_KEY, LOOP_VALUE, or DURATION
		// Consider non-empty strings as true
		return len(v) > 0

	case int:
		// Handle NUMBER from number_expr or duration_expr
		return v != 0

	case float64:
		// Handle FLOAT from number_expr or duration_expr
		return v != 0.0

	case []interface{}:
		// Handle list_expr
		return len(v) > 0

	case [][]interface{}:
		// Handle matrix_expr
		return len(v) > 0

	case map[string]interface{}:
		// Handle function or nested_function
		// Consider any function call as true (since it exists)
		return true

	default:
		// Fallback: consider non-nil values as true
		return true
	}
}
