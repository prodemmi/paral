package parser

import (
	"fmt"
	parser "paral/antlr/antlr"
	"paral/core"
	"paral/core/runtime"
	"strconv"
)

func (p *Parser) parseExpression(ctx parser.IExpressionContext) *runtime.Expression {
	rawText := ctx.GetText()
	var result interface{}
	switch {
	case ctx.Loop_variable() != nil:
		// Handle LOOP_KEY (@key) or LOOP_VALUE (@value)
		loopVar := ctx.Loop_variable()
		result = loopVar.GetText()

	case ctx.Function() != nil:
		// Handle function or nested function
		fnCtx := ctx.Function()
		if fnCtx.PIPELINE_FUNCTION_CALL_START() != nil || fnCtx.FUNCTION_START() != nil {
			// Parse function call`
			fnName := fnCtx.GetStart().GetText()[1:] // Remove @ prefix
			args := p.parseArgumentList(fnCtx.Argument_list())
			result = map[string]interface{}{
				"type": "function",
				"name": fnName,
				"args": args,
			}
		}

	case ctx.URL() != nil:
		// Handle URL
		result = ctx.URL().GetText()

	case ctx.Number_expr() != nil:
		// Handle FLOAT or NUMBER
		numCtx := ctx.Number_expr()
		if numCtx.FLOAT() != nil {
			f, err := strconv.ParseFloat(numCtx.GetText(), 64)
			if err == nil {
				result = f
			} else {
				result = numCtx.GetText()
			}
		} else if numCtx.NUMBER() != nil {
			i, err := strconv.Atoi(numCtx.GetText())
			if err == nil {
				result = i
			} else {
				result = numCtx.GetText()
			}
		}

	case ctx.String_expr() != nil:
		// Handle STRING or SINGLE_QUOTE_STRING
		strCtx := ctx.String_expr()
		strText := strCtx.GetText()
		result = core.TrimQuotes(strText)
	case ctx.Boolean_expr() != nil:
		// Handle BOOLEAN (true/false) or ZERO_ONE (0/1)
		boolCtx := ctx.Boolean_expr()
		if boolCtx.BOOLEAN() != nil {
			b, err := strconv.ParseBool(boolCtx.GetText())
			if err == nil {
				result = b
			} else {
				result = boolCtx.GetText()
			}
		} else if boolCtx.ZERO_ONE() != nil {
			result = boolCtx.GetText() == "1"
		}

	case ctx.Duration_expr() != nil:
		// Handle DURATION or number_expr
		durCtx := ctx.Duration_expr()
		if durCtx.DURATION() != nil {
			result = durCtx.GetText()
		} else if durCtx.Number_expr() != nil {
			numCtx := durCtx.Number_expr()
			if numCtx.FLOAT() != nil {
				f, err := strconv.ParseFloat(numCtx.GetText(), 64)
				if err == nil {
					result = f
				} else {
					result = numCtx.GetText()
				}
			} else if numCtx.NUMBER() != nil {
				i, err := strconv.Atoi(numCtx.GetText())
				if err == nil {
					result = i
				} else {
					result = numCtx.GetText()
				}
			}
		}

	case ctx.Matrix_expr() != nil:
		// Handle matrix_expr (list_expr :: list_expr ...)
		matrixCtx := ctx.Matrix_expr()
		lists := matrixCtx.AllList_expr()
		matrix := make([][]interface{}, 0)
		for _, listCtx := range lists {
			matrix = append(matrix, p.parseListExpr(listCtx))
		}
		result = matrix

	case ctx.List_expr() != nil:
		// Handle list_expr
		result = p.parseListExpr(ctx.List_expr())

	case ctx.IDENTIFIER() != nil:
		// Handle IDENTIFIER
		variableName := ctx.IDENTIFIER().GetText()
		variableValue := p.Runtime.GetVariable(variableName)
		if variableValue == nil {
			p.Runtime.Reporter.ThrowSyntaxError(fmt.Sprintf("Undefined variable: %s", variableValue), variableValue.Metadata)
		}
		_, result = variableValue.Format()
		rawText = variableName
	default:
		result = rawText
	}
	return &runtime.Expression{
		Result:  result,
		RawText: rawText,
	}
}

// Helper function to parse argument lists
func (p *Parser) parseArgumentList(ctx parser.IArgument_listContext) []interface{} {
	if ctx == nil {
		return nil
	}

	args := make([]interface{}, 0)
	for _, exprCtx := range ctx.AllExpression() {
		expr := p.parseExpression(exprCtx)
		args = append(args, expr.Result)
	}
	return args
}

// Helper function to parse list expressions
func (p *Parser) parseListExpr(ctx parser.IList_exprContext) []interface{} {
	if ctx == nil {
		return nil
	}

	list := make([]interface{}, 0)
	for _, exprCtx := range ctx.AllExpression() {
		expr := p.parseExpression(exprCtx)
		list = append(list, expr.Result)
	}
	return list
}
