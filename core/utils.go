package core

import (
	parser "paral/antlr/antlr"
	"strconv"
	"strings"
)

// TrimQuotes removes surrounding quotes from strings
func TrimQuotes(s string) string {
	return strings.Trim(s, "\"'")
}

// ExtractListValues extracts values from a list_expr
func ExtractListValues(list parser.IList_exprContext) []interface{} {
	var values []interface{}
	for _, val := range list.AllExpression() {
		if val.String_expr() != nil {
			values = append(values, TrimQuotes(val.GetText()))
		} else if val.Number_expr() != nil {
			if num, err := strconv.Atoi(val.GetText()); err == nil {
				values = append(values, num)
			} else {
				values = append(values, val.GetText())
			}
		} else if val.Boolean_expr() != nil {
			values = append(values, val.GetText() == "true")
		} else if val.Duration_expr() != nil {
			values = append(values, val.GetText() == "true")
		}
	}
	return values
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseList(list []interface{}) []interface{} {
	reversed := make([]interface{}, len(list))
	for i, v := range list {
		reversed[len(list)-1-i] = v
	}
	return reversed
}
