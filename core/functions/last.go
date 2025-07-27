package functions

func (f *Function) last() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, nil
	}
	switch v := f.argResults[0].(type) {
	case string:
		if len(v) > 0 {
			return string(v[len(v)-1]), nil
		}
	case []interface{}:
		if len(v) > 0 {
			return v[len(v)-1], nil
		}
	case [][]interface{}:
		if len(v) > 0 && len(v[len(v)-1]) > 0 {
			lastList := v[len(v)-1]
			return lastList[len(lastList)-1], nil
		}
	}
	return nil, nil
}
