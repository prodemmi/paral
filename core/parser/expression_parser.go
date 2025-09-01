package parser

import (
	"fmt"
	parser "paral/antlr/antlr"
	"paral/core"
	"paral/core/metadata"
	"paral/core/runtime"
	"paral/core/variable"
	"strconv"
)

func (p *Parser) parseExpression(task *runtime.Task, ctx parser.IExpressionContext) *runtime.Expression {
	rawText := ctx.GetText()
	var value interface{}
	mt := metadata.Metadata{
		Content: ctx.GetText(),
		Line:    ctx.GetStop().GetLine(),
		Column:  ctx.GetStop().GetColumn(),
	}

	switch {
	case ctx.Hich_expr() != nil:
		// Handle nil
		value = nil
		rawText = ctx.Hich_expr().GetText()
	case ctx.Loop_variable() != nil:
		// Handle LOOP_KEY (@key) or LOOP_VALUE (@value)
		value = variable.StringValue{
			Value: ctx.Loop_variable().GetText(),
		}
		rawText = ctx.Loop_variable().GetText()
	case ctx.Error_variable() != nil:
		// Handle ERROR (@error)
		value = variable.StringValue{
			Value: ctx.Error_variable().GetText(),
		}
		rawText = ctx.Error_variable().GetText()
	case ctx.Function() != nil:
		// Handle function or nested function
		fnCtx := ctx.Function()
		if fnCtx.FUNCTION_START() != nil || fnCtx.EXPRESSION_FUNCTION_CALL_START() != nil || fnCtx.PIPELINE_FUNCTION_CALL_START() != nil {
			value = p.parseFunction(task, fnCtx)
		}
	case ctx.Nested_function() != nil:
		// Handle function or nested function
		nestedFnCtx := ctx.Nested_function()
		if nestedFnCtx.NESTED_FUNCTION_START() != nil {
			value = p.parseNestedFunction(task, nestedFnCtx)
		}
	case ctx.URL() != nil:
		value = variable.StringValue{
			Value: ctx.URL().GetText(),
		}
	case ctx.Number_expr() != nil:
		// Handle FLOAT or NUMBER
		numCtx := ctx.Number_expr()
		if numCtx.FLOAT() != nil {
			f, err := strconv.ParseFloat(numCtx.GetText(), 64)
			if err == nil {
				value = variable.FloatValue{
					Value: f,
				}
			} else {
				stringNumber := numCtx.GetText()
				value = variable.StringValue{
					Value: stringNumber,
				}
			}
		} else if numCtx.NUMBER() != nil {
			i, err := strconv.Atoi(numCtx.GetText())
			if err == nil {
				value = variable.IntValue{
					Value: i,
				}
			} else {
				stringNumber := numCtx.GetText()
				value = variable.StringValue{
					Value: core.TrimQuotes(stringNumber),
				}
			}
		}
	case ctx.String_expr() != nil:
		strCtx := ctx.String_expr()
		strText := strCtx.GetText()
		value = variable.StringValue{
			Value: core.TrimQuotes(strText),
		}
	case ctx.Boolean_expr() != nil:
		boolCtx := ctx.Boolean_expr()
		if boolCtx.BOOLEAN() != nil {
			b, err := strconv.ParseBool(boolCtx.GetText())
			if err == nil {
				value = variable.BoolValue{
					Value: b,
				}
			} else {
				value = boolCtx.GetText()
			}
		} else if boolCtx.ZERO_ONE() != nil {
			value = boolCtx.GetText() == "1"
		}
		value = variable.BoolValue{
			Value: value.(bool),
		}
	case ctx.Duration_expr() != nil:
		durCtx := ctx.Duration_expr()
		if durCtx.DURATION() != nil {
			value = variable.DurationValue{
				Value: durCtx.GetText(),
			}
		} else if durCtx.Number_expr() != nil {
			numCtx := durCtx.Number_expr()
			if numCtx.FLOAT() != nil {
				f, err := strconv.ParseFloat(numCtx.GetText(), 64)
				if err == nil {
					value = variable.DurationValue{
						Value: f,
					}
				} else {
					value = variable.DurationValue{
						Value: numCtx.GetText(),
					}
				}

			} else if numCtx.NUMBER() != nil {
				i, err := strconv.Atoi(numCtx.GetText())
				if err == nil {
					value = variable.DurationValue{
						Value: i,
					}
				} else {
					value = variable.DurationValue{
						Value: numCtx.GetText(),
					}
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
		value = variable.MatrixValue{
			Value: matrix,
		}
	case ctx.List_expr() != nil:
		value = variable.ListValue{
			Value: p.parseListExpr(task, ctx.List_expr()),
		}
	case ctx.IDENTIFIER() != nil:
		idText := ctx.IDENTIFIER().GetText()
		// Check if it's @error first

		task := p.Runtime.GetTaskByID(idText)
		if task != nil {
			value = variable.StringValue{
				Value: task.GetTaskId(),
			}
			rawText = idText
			break
		}

		variableValue := p.Runtime.GetVariable(idText)

		if variableValue != nil {
			value = variableValue
			rawText = idText
			break
		}
		if variableValue == nil || task == nil {
			p.Runtime.Reporter.ThrowSyntaxError(fmt.Sprintf("undefined variable: %s", idText), &mt)
		}
	default:
		value = variable.StringValue{
			Value: rawText,
		}
	}

	return &runtime.Expression{
		Value:    value,
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
		var result interface{}
		var err error
		if task != nil {
			result, err = expr.EvaluateValue(task.GetActiveLoopContext())
		} else {
			result, err = expr.EvaluateValue(nil)
		}
		if err != nil {
			p.Runtime.Reporter.ThrowRuntimeError("failed to evaluate expression", &metadata.Metadata{
				Content: ctx.GetText(),
				Line:    ctx.GetStop().GetLine(),
				Column:  ctx.GetStop().GetColumn(),
			})
		}
		list = append(list, result)
	}
	return list
}
