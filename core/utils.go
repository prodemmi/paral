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
	for _, val := range list.AllVariable_value() {
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
