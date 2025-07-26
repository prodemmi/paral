package core

import (
	"fmt"
	"strconv"
	"strings"

	parser "paral/antlr/antlr"
)

type CoreVisitor struct {
	*parser.BaseParalParserListener
	Core *Core
}

// ExitVariable_assignment is called when production variable_def is exited.
func (s *CoreVisitor) ExitVariable_assignment(ctx *parser.Variable_assignmentContext) {
	name := ctx.IDENTIFIER().GetText()
	filename := ctx.GetStop().GetInputStream().GetSourceName()
	line := ctx.GetStop().GetLine()
	column := ctx.GetStop().GetColumn()

	val := ctx.Variable_value()
	var variable Var

	if val.Number_expr() != nil {
		numText := val.GetText()
		num, err := strconv.Atoi(numText)
		if err != nil {
			fnum, err := strconv.ParseFloat(numText, 64)
			if err != nil {
				ThrowSyntaxError(fmt.Sprintf("Invalid number: %s", numText), filename, line, column)
			}
			variable = Var{
				VarBase: VarBase{Name: name},
				Value:   FloatValue{VarBase: VarBase{Name: name}, Value: fnum},
			}
		} else {
			variable = Var{
				VarBase: VarBase{Name: name},
				Value:   IntValue{VarBase: VarBase{Name: name}, Value: num},
			}
		}

	} else if val.String_expr() != nil {
		variable = Var{
			VarBase: VarBase{Name: name},
			Value:   StringValue{VarBase: VarBase{Name: name}, Value: TrimQuotes(val.GetText())},
		}

	} else if val.Boolean_expr() != nil {
		boolText := val.GetText()
		variable = Var{
			VarBase: VarBase{Name: name},
			Value:   BoolValue{VarBase: VarBase{Name: name}, Value: boolText == "true"},
		}

	} else if val.Duration_expr() != nil {
		variable = Var{
			VarBase: VarBase{Name: name},
			Value:   DurationValue{VarBase: VarBase{Name: name}, Value: val.GetText()},
		}

	} else if val.URL() != nil {
		variable = Var{
			VarBase: VarBase{Name: name},
			Value:   StringValue{VarBase: VarBase{Name: name}, Value: val.GetText()},
		}

	} else if val.Matrix_expr() != nil {
		var matrixValues [][]interface{}
		for _, list := range val.Matrix_expr().AllList_expr() {
			row := ExtractListValues(list)
			if len(row) == 0 {
				Warn(fmt.Sprintf("Empty list in matrix %s", name), filename, line, column)
			}
			matrixValues = append(matrixValues, row)
		}
		variable = Var{
			VarBase: VarBase{Name: name},
			Value:   MatrixValue{VarBase: VarBase{Name: name}, Value: matrixValues},
		}

	} else if val.List_expr() != nil {
		listValues := ExtractListValues(val.List_expr())
		variable = Var{
			VarBase: VarBase{Name: name},
			Value:   ListValue{VarBase: VarBase{Name: name}, Value: listValues},
		}

	} else if val.IDENTIFIER() != nil {
		// Variable reference by name
		variable = Var{
			VarBase: VarBase{Name: name},
			Value:   StringValue{VarBase: VarBase{Name: name}, Value: val.IDENTIFIER().GetText()},
		}

	} else {
		ThrowSyntaxError("Unknown value type in variable definition", filename, line, column)
	}

	s.Core.AddVar(variable)
}

func (s *CoreVisitor) ExitFunction(ctx *parser.FunctionContext) {
	// Functions are handled within ExitJob_definition
}

func (s *CoreVisitor) parseNestedFunction(fnCtx parser.INested_functionContext) Function {
	fn := fnCtx.(*parser.Nested_functionContext)
	fnName := strings.Trim(fn.NESTED_FUNCTION_START().GetText(), "@(")
	args := []interface{}{}

	argList := fn.Argument_list()
	if argList != nil {
		for _, expr := range argList.AllExpression() {
			if variableValue := expr.Variable_value(); variableValue != nil {
				// Handle identifier (e.g., filename or variable)
				if id := variableValue.IDENTIFIER(); id != nil {
					varName := id.GetText()
					if variable := s.Core.findVarByName(varName); variable != nil {
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

				if numToken := variableValue.List_expr(); numToken != nil {
					var list []interface{}
					for _, listItem := range numToken.AllVariable_value() {
						list = append(list, listItem.GetText())
					}
					args = append(args, list)
					continue
				}
			}
		}
	}

	return *NewFunction(fnName, s.Core.Filename, fnCtx.GetStop().GetLine(), fnCtx.GetStop().GetColumn(), args...)
}

func (s *CoreVisitor) parseFunction(fnCtx parser.IFunctionContext) Function {
	fn := fnCtx.(*parser.FunctionContext)
	fnName := strings.Trim(fn.FUNCTION_CALL_START().GetText(), "@(")
	args := []interface{}{}

	argList := fn.Argument_list()
	if argList != nil {
		for _, expr := range argList.AllExpression() {
			if variableValue := expr.Variable_value(); variableValue != nil {
				// Handle identifier (e.g., filename or variable)
				if id := variableValue.IDENTIFIER(); id != nil {
					varName := id.GetText()
					if variable := s.Core.findVarByName(varName); variable != nil {
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

				if numToken := variableValue.List_expr(); numToken != nil {
					var list []interface{}
					for _, listItem := range numToken.AllVariable_value() {
						list = append(list, listItem.GetText())
					}
					args = append(args, list)
					continue
				}
			}

			// Handle nested functions
			if innerFn := expr.Nested_function(); innerFn != nil {
				args = append(args, s.parseNestedFunction(innerFn))
			}
		}
	}

	return *NewFunction(fnName, s.Core.Filename, fnCtx.GetStop().GetLine(), fnCtx.GetStop().GetColumn(), args...)
}

// ExitJob_definition is called when production job_def is exited.
func (s *CoreVisitor) ExitJob_definition(ctx *parser.Job_definitionContext) {
	name := ctx.IDENTIFIER().GetText()
	filename := ctx.GetStop().GetInputStream().GetSourceName()
	job := NewJob(name, filename)

	// Handle job directives
	for _, directiveExpr := range ctx.AllJob_directive() {
		directive := NewJobDirective()

		if dcr := directiveExpr.Directive(); dcr != nil {
			directive.Type = dcr.IDENTIFIER().GetText()
			for _, dcrParam := range dcr.AllExpression() {
				// Handle directive parameters, removing quotes from string literals
				text := dcrParam.GetText()
				if strings.HasPrefix(text, "\"") && strings.HasSuffix(text, "\"") {
					text = strings.Trim(text, "\"")
				}
				directive.Params = append(directive.Params, text)
			}
		}
		_ = job.AddJobDirective(directive)
	}

	// Handle command blocks
	for _, commandBlock := range ctx.AllCommand_block() {
		rawText, _ := strings.CutPrefix(commandBlock.GetText(), "->")
		cmd := Command{
			RawText:   rawText,
			Functions: []Function{},
			Block:     commandBlock.GetStart().GetLine(),
		}

		if commandContent := commandBlock.Command_content(); commandContent != nil {
			for _, commandItem := range commandContent.AllCommand_item() {
				if fnCtx := commandItem.Function(); fnCtx != nil {
					function := s.parseFunction(fnCtx)
					cmd.Functions = append(cmd.Functions, function)
				}
			}
		}

		job.AddJobCommand(cmd)
	}

	s.Core.AddJob(*job)
}
