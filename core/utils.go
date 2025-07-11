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

// ExtractValue extracts the value from a value_expr
func ExtractValue(val parser.IValue_exprContext) interface{} {
	switch v := val.(type) {
	case *parser.StringContext:
		return TrimQuotes(v.GetText())
	case *parser.SingleQuoteStringContext:
		return TrimQuotes(v.GetText())
	case *parser.NumberContext:
		if num, err := strconv.Atoi(v.GetText()); err == nil {
			return num
		}
		return v.GetText()
	case *parser.BoolContext:
		return v.GetText() == "true"
	case *parser.DurationContext:
		return v.GetText()
	default:
		return v.GetText()
	}
}

// ExtractListValues extracts values from a list_expr
func ExtractListValues(list parser.IList_exprContext) []interface{} {
	var values []interface{}
	for _, val := range list.AllValue_expr() {
		switch v := val.(type) {
		case *parser.StringContext:
			values = append(values, TrimQuotes(v.GetText()))
		case *parser.SingleQuoteStringContext:
			values = append(values, TrimQuotes(v.GetText()))
		case *parser.NumberContext:
			if num, err := strconv.Atoi(v.GetText()); err == nil {
				values = append(values, num)
			} else {
				values = append(values, v.GetText())
			}
		case *parser.BoolContext:
			values = append(values, v.GetText() == "true")
		case *parser.DurationContext:
			values = append(values, v.GetText())
		}
	}
	return values
}
