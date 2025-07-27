package functions

func (f *Function) first() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, nil
	}
	switch v := f.argResults[0].(type) {
	case string:
		if len(v) > 0 {
			return string(v[0]), nil
		}
	case []interface{}:
		if len(v) > 0 {
			return v[0], nil
		}
	case [][]interface{}:
		if len(v) > 0 && len(v[0]) > 0 {
			return v[0][0], nil
		}
	}
	return nil, nil
}
