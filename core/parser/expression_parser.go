package parser

import (
	"fmt"
	parser "paral/antlr/antlr"
	"paral/core/metadata"
	"paral/core/runtime"
	"strconv"
)

func (p *Parser) parseExpression(task *runtime.Task, ctx parser.IExpressionContext) *runtime.Expression {
	rawText := ctx.GetText()
	mt := metadata.Metadata{
		Content: ctx.GetText(),
		Line:    ctx.GetStop().GetLine(),
		Column:  ctx.GetStop().GetColumn(),
	}

	var result interface{}
	switch {
	case ctx.Loop_variable() != nil:
		// Handle LOOP_KEY (@key) or LOOP_VALUE (@value)
		result = ctx.Loop_variable().GetText()
	case ctx.Error_variable() != nil:
		// Handle ERROR (@error)
		result = p.Runtime.GetTryCatchError()
	case ctx.Function() != nil:
		// Handle function or nested function
		fnCtx := ctx.Function()
		if fnCtx.FUNCTION_START() != nil || fnCtx.EXPRESSION_FUNCTION_CALL_START() != nil || fnCtx.PIPELINE_FUNCTION_CALL_START() != nil {
			result = p.parseFunction(task, fnCtx)
		}
	case ctx.Nested_function() != nil:
		// Handle function or nested function
		nestedFnCtx := ctx.Nested_function()
		if nestedFnCtx.NESTED_FUNCTION_START() != nil {
			result = p.parseNestedFunction(task, nestedFnCtx)
		}
	case ctx.URL() != nil:
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
		strCtx := ctx.String_expr()
		strText := strCtx.GetText()
		result = strText
	case ctx.Boolean_expr() != nil:
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
		matrixCtx := ctx.Matrix_expr()
		lists := matrixCtx.AllList_expr()
		matrix := make([][]interface{}, 0)
		for _, listCtx := range lists {
			matrix = append(matrix, p.parseListExpr(task, listCtx))
		}
		result = matrix

	case ctx.List_expr() != nil:
		result = p.parseListExpr(task, ctx.List_expr())
	case ctx.IDENTIFIER() != nil:
		variableName := ctx.IDENTIFIER().GetText()
		// Check if it's @error first
		if variableName == "@error" {
			result = "@error"
			rawText = "@error"
		} else {
			variableValue := p.Runtime.GetVariable(variableName)
			if variableValue == nil {
				p.Runtime.Reporter.ThrowSyntaxError(fmt.Sprintf("Undefined variable: %s", variableName), &mt)
			}
			_, result = variableValue.Format()
			rawText = variableName
		}
	default:
		// Check if the raw text is @error
		if rawText == "@error" {
			result = "@error"
		} else {
			result = rawText
		}
	}

	return &runtime.Expression{
		Result:   result,
		RawText:  rawText,
		Metadata: &mt,
	}
}

func (p *Parser) parseListExpr(task *runtime.Task, ctx parser.IList_exprContext) []interface{} {
	if ctx == nil {
		return nil
	}

	list := make([]interface{}, 0)
	for _, exprCtx := range ctx.AllExpression() {
		expr := p.parseExpression(task, exprCtx)
		list = append(list, expr.Result)
	}
	return list
}
