package functions

func (f *Function) len() (interface{}, error) {
	if len(f.argResults) == 0 {
		return 0, nil
	}
	switch v := f.argResults[0].(type) {
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
