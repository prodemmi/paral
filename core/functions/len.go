package functions

func Len(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return 0, nil
	}

	switch v := args[0].(type) {
	case string:
		return len(v), nil
	case []interface{}:
		return len(v), nil
	case [][]interface{}:
		return len(v), nil
	default:
		return 0, nil
	}
}
