package functions

func Reverse(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, nil
	}

	switch v := args[0].(type) {
	case string:
		runes := []rune(v)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes), nil
	case []interface{}:
		n := len(v)
		reversed := make([]interface{}, n)
		for i := 0; i < n; i++ {
			reversed[i] = v[n-1-i]
		}
		return reversed, nil
	case [][]interface{}:
		n := len(v)
		reversed := make([][]interface{}, n)
		for i := 0; i < n; i++ {
			reversed[i] = v[n-1-i]
		}
		return reversed, nil
	}
	return nil, nil
}
