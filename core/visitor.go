package core

import (
	"fmt"
	"strconv"
	"strings"

	parser "paral/antlr/antlr"

	"github.com/antlr4-go/antlr/v4"
)

type CoreVisitor struct {
	*parser.BaseParalListener
	Core *Core
}

// VisitTerminal is called when a terminal node is visited.
func (s *CoreVisitor) VisitTerminal(node antlr.TerminalNode) {}

// ExitEveryRule is called when any rule is exited.
func (s *CoreVisitor) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *CoreVisitor) EnterProgram(ctx *parser.ProgramContext) {
	Info("Starting to parse .paral file", ctx.GetStop().GetInputStream().GetSourceName(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
}

// ExitProgram is called when production program is exited.
func (s *CoreVisitor) ExitProgram(ctx *parser.ProgramContext) {
	Info("Finished parsing .paral file", ctx.GetStop().GetInputStream().GetSourceName(), ctx.GetStop().GetLine(), ctx.GetStop().GetColumn())
}

// ExitLine is called when production line is exited.
func (s *CoreVisitor) ExitLine(ctx *parser.LineContext) {}

// ExitVariable_def is called when production variable_def is exited.
func (s *CoreVisitor) ExitVariable_def(ctx *parser.Variable_defContext) {
	name := ctx.IDENTIFIER().GetText()
	filename := ctx.GetStop().GetInputStream().GetSourceName()
	line := ctx.GetStop().GetLine()
	column := ctx.GetStop().GetColumn()

	var value interface{}
	var _type string

	if ctx.Value_expr() != nil {
		switch val := ctx.Value_expr().(type) {
		case *parser.StringContext:
			value = TrimQuotes(val.GetText())
			_type = "string"
		case *parser.SingleQuoteStringContext:
			value = TrimQuotes(val.GetText())
			_type = "string"
		case *parser.NumberContext:
			num, err := strconv.Atoi(val.GetText())
			if err != nil {
				ThrowSyntaxError(fmt.Sprintf("Invalid number: %s", val.GetText()), filename, line, column)
			}
			value = num
			_type = "number"
		case *parser.BoolContext:
			value = val.GetText() == "true"
			_type = "bool"
		case *parser.DurationContext:
			value = val.GetText()
			_type = "duration"
		default:
			ThrowSyntaxError("Unknown value type in variable definition", filename, line, column)
		}
	} else if ctx.List_expr() != nil {
		value = ExtractListValues(ctx.List_expr())
		_type = "list"
	}

	s.Core.AddVar(Var{Name: name, Value: value, Type: _type})
}

// ExitMatrix_def is called when production matrix_def is exited.
func (s *CoreVisitor) ExitMatrix_def(ctx *parser.Matrix_defContext) {
	name := ctx.IDENTIFIER().GetText()
	filename := ctx.GetStop().GetInputStream().GetSourceName()
	line := ctx.GetStop().GetLine()
	column := ctx.GetStop().GetColumn()

	var matrixValues [][]interface{}
	for _, list := range ctx.AllList_expr() {
		row := ExtractListValues(list)
		if len(row) == 0 {
			Warn(fmt.Sprintf("Empty list in matrix %s", name), filename, line, column)
		}
		matrixValues = append(matrixValues, row)
	}

	s.Core.AddMatrix(Matrix{Name: name, Values: matrixValues})
}

// ExitList_expr is called when production list_expr is exited.
func (s *CoreVisitor) ExitList_expr(ctx *parser.List_exprContext) {}

// ExitJob_def is called when production job_def is exited.
func (s *CoreVisitor) ExitJob_def(ctx *parser.Job_defContext) {
	name := ctx.IDENTIFIER().GetText()
	filename := ctx.GetStop().GetInputStream().GetSourceName()
	job := NewJob(name, filename)

	for _, command := range ctx.AllCmd_expr() {
		cmdText := ""
		var directives []CMDDirective

		for _, expr := range command.AllCmd_value() {
			directive := NewCMDDirective()
			if directiveExpr := expr.Directive_expr(); directiveExpr != nil {
				directive.Type = strings.Replace(directiveExpr.DIRECTIVE().GetText(), "@", "", 1)
				if directiveExprArgs := directiveExpr.AllDirective_args_expr(); directiveExprArgs != nil {
					for _, arg := range directiveExprArgs {
						args := ""
						for i, value := range arg.AllValue_expr() {
							if i < len(arg.AllValue_expr())-1 {
								args += value.GetText() + ", "
							} else {
								args += value.GetText()
							}
						}
						directive.Params = append(directive.Params, args)
					}
				}
				directives = append(directives, *directive)
			}
			cmdText += expr.GetText() + " "

		}

		cmdText = strings.TrimSpace(cmdText)
		if cmdText == "" {
			ThrowSyntaxError("job %s has no command", filename, command.GetStop().GetLine(), command.GetStop().GetColumn())
		}

		job.AddJobCommand(cmdText, directives)
	}

	s.Core.AddJob(*job)
}
