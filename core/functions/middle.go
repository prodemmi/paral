package functions

import "fmt"

func (f *Function) middle() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, nil
	}
	switch v := f.argResults[0].(type) {
	case string:
		if len(v) <= 2 {
			return "", nil
		}
		midIndex := len(v) / 2
		if len(v)%2 == 0 {
			return string(v[midIndex-1 : midIndex+1]), nil
		}
		return string(v[midIndex]), nil
	case []interface{}:
		if len(v) <= 2 {
			return nil, nil
		}
		midIndex := len(v) / 2
		if len(v)%2 == 0 {
			return v[midIndex-1 : midIndex+1], nil
		}
		return v[midIndex], nil
	case [][]interface{}:
		if len(v) <= 2 {
			return nil, nil
		}
		midIndex := len(v) / 2
		if len(v)%2 == 0 {
			return v[midIndex-1 : midIndex+1], nil
		}
		return v[midIndex], nil
	}
	return nil, fmt.Errorf("middle: first argument must be a string, list, or matrix")
}
