package core

import "fmt"

type Matrix struct {
	Name   string
	Values [][]interface{}
}

func (m Matrix) GetMatrixCombinations() []map[string]interface{} {
	var result []map[string]interface{}
	if len(m.Values) == 0 {
		return result
	}

	// کلیدهای پیش‌فرض: key1, key2, ...
	keys := make([]string, len(m.Values))
	for i := range m.Values {
		keys[i] = fmt.Sprint(i + 1)
	}

	var build func(index int, current map[string]interface{})
	build = func(index int, current map[string]interface{}) {
		if index == len(m.Values) {
			combination := make(map[string]interface{})
			for k, v := range current {
				combination[k] = v
			}
			result = append(result, combination)
			return
		}

		for _, val := range m.Values[index] {
			current[keys[index]] = val
			build(index+1, current)
		}
	}

	build(0, make(map[string]interface{}))
	return result
}
