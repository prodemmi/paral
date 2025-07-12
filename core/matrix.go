package core

type Matrix struct {
	Name   string
	Values [][]interface{}
}

func (m *Matrix) GetMatrixCombinations() [][]interface{} {
	var result [][]interface{}
	if len(m.Values) == 0 {
		return result
	}
	if len(m.Values) == 1 {
		for _, v := range m.Values[0] {
			result = append(result, []interface{}{v})
		}
		return result
	}
	result = append(result, []interface{}{})
	for _, list := range m.Values {
		var newResult [][]interface{}
		for _, combo := range result {
			for _, val := range list {
				newCombo := append([]interface{}{}, combo...)
				newCombo = append(newCombo, val)
				newResult = append(newResult, newCombo)
			}
		}
		result = newResult
	}
	return result
}
