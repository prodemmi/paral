package parser

import (
	parser "paral/antlr/antlr"
	"paral/core/metadata"
	"paral/core/runtime"
	"strconv"
	"strings"
)

func (p *Parser) parseNestedFunction(task *runtime.Task, fnCtx parser.INested_functionContext) runtime.Function {
	fn := fnCtx.(*parser.Nested_functionContext)
	fnName := strings.Trim(fn.NESTED_FUNCTION_START().GetText(), "@(")
	args := []interface{}{}

	argList := fn.Argument_list()
	if argList != nil {
		for _, expr := range argList.AllExpression() {
			if variableValue := expr; variableValue != nil {
				// Handle identifier (e.g., filename or variable)
				if id := variableValue.IDENTIFIER(); id != nil {
					varName := id.GetText()
					if variable := p.Runtime.GetVariable(varName); variable != nil {
						args = append(args, *variable)
					} else {
						// Treat unknown identifier as raw value (e.g., file name)
						args = append(args, varName)
					}
					continue
				}

				// Handle string literals
				if strToken := variableValue.String_expr(); strToken != nil {
					quoted := strToken.GetText()
					unquoted, err := strconv.Unquote(quoted)
					if err != nil {
						unquoted = quoted
					}
					args = append(args, unquoted)
					continue
				}

				// Handle number literals
				if numToken := variableValue.Number_expr(); numToken != nil {
					text := numToken.GetText()
					if strings.Contains(text, ".") {
						if floatVal, err := strconv.ParseFloat(text, 64); err == nil {
							args = append(args, floatVal)
						}
					} else {
						if intVal, err := strconv.Atoi(text); err == nil {
							args = append(args, intVal)
						}
					}
					continue
				}

				// Handle function
				if strToken := variableValue.Function(); strToken != nil {
					nestedFunction := p.parseFunction(task, strToken)
					args = append(args, nestedFunction)
					continue
				}

				// Handle arg function
				if strToken := variableValue.Nested_function(); strToken != nil {
					nestedFunction := p.parseNestedFunction(task, strToken)
					args = append(args, nestedFunction)
					continue
				}

				if numToken := variableValue.List_expr(); numToken != nil {
					var list []interface{}
					for _, listItem := range numToken.AllExpression() {
						list = append(list, listItem.GetText())
					}
					args = append(args, list)
					continue
				}

				if numToken := variableValue.Loop_variable(); numToken != nil {
					if loopValue := numToken.PIPELINE_LOOP_VALUE(); loopValue != nil {
						args = append(args, loopValue.GetText())
					}
					if loopKey := numToken.PIPELINE_LOOP_KEY(); loopKey != nil {
						args = append(args, loopKey.GetText())
					}
					continue
				}
			}
		}
	}

	return *runtime.NewFunction(fnName, metadata.Metadata{
		Line:   fnCtx.GetStop().GetLine(),
		Column: fnCtx.GetStop().GetColumn(),
	}, task, args...)
}

func (p *Parser) parseFunction(task *runtime.Task, fnCtx parser.IFunctionContext) *runtime.Function {
	fn := fnCtx.(*parser.FunctionContext)
	var fnName string
	if fn.FUNCTION_CALL_START() != nil {
		fnName = strings.Trim(fn.FUNCTION_CALL_START().GetText(), "@(")
	} else {
		fnName = strings.Trim(fn.FUNCTION_START().GetText(), "@(")
	}
	args := []interface{}{}

	argList := fn.Argument_list()
	if argList != nil {
		for _, expr := range argList.AllExpression() {
			if variableValue := expr; variableValue != nil {
				// Handle identifier (e.g., filename or variable)
				if id := variableValue.IDENTIFIER(); id != nil {
					varName := id.GetText()
					if variable := p.Runtime.GetVariable(varName); variable != nil {
						args = append(args, *variable)
					} else {
						// Treat unknown identifier as raw value (e.g., file name)
						args = append(args, varName)
					}
					continue
				}

				// Handle string literals
				if strToken := variableValue.String_expr(); strToken != nil {
					quoted := strToken.GetText()
					unquoted, err := strconv.Unquote(quoted)
					if err != nil {
						unquoted = quoted
					}
					args = append(args, unquoted)
					continue
				}

				// Handle function
				if strToken := variableValue.Function(); strToken != nil {
					nestedFunction := p.parseFunction(task, strToken)
					args = append(args, nestedFunction)
					continue
				}

				// Handle arg function
				if strToken := variableValue.Nested_function(); strToken != nil {
					nestedFunction := p.parseNestedFunction(task, strToken)
					args = append(args, nestedFunction)
					continue
				}

				// Handle number literals
				if numToken := variableValue.Number_expr(); numToken != nil {
					text := numToken.GetText()
					if strings.Contains(text, ".") {
						if floatVal, err := strconv.ParseFloat(text, 64); err == nil {
							args = append(args, floatVal)
						}
					} else {
						if intVal, err := strconv.Atoi(text); err == nil {
							args = append(args, intVal)
						}
					}
					continue
				}

				if numToken := variableValue.List_expr(); numToken != nil {
					var list []interface{}
					for _, listItem := range numToken.AllExpression() {
						list = append(list, listItem.GetText())
					}
					args = append(args, list)
					continue
				}

				if numToken := variableValue.Loop_variable(); numToken != nil {
					if loopValue := numToken.PIPELINE_LOOP_VALUE(); loopValue != nil {
						args = append(args, loopValue.GetText())
					}
					if loopKey := numToken.PIPELINE_LOOP_KEY(); loopKey != nil {
						args = append(args, loopKey.GetText())
					}
					continue
				}
			}

			// Handle nested functions
			if innerFn := expr.Nested_function(); innerFn != nil {
				args = append(args, p.parseNestedFunction(task, innerFn))
			}
		}
	}

	return runtime.NewFunction(fnName, metadata.Metadata{
		Line:   fnCtx.GetStop().GetLine(),
		Column: fnCtx.GetStop().GetColumn(),
	}, task, args...)
}
