package core

import (
	"fmt"
	parser "paral/antlr/antlr"
	"strconv"
	"strings"

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
func (s *CoreVisitor) EnterProgram(ctx *parser.ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *CoreVisitor) ExitProgram(ctx *parser.ProgramContext) {}

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

// ExitComment_def is called when production comment_def is exited.
func (s *CoreVisitor) ExitComment_def(ctx *parser.Comment_defContext) {}

// ExitString is called when production string is exited.
func (s *CoreVisitor) ExitString(ctx *parser.StringContext) {}

// ExitNumber is called when production number is exited.
func (s *CoreVisitor) ExitNumber(ctx *parser.NumberContext) {}

// ExitBool is called when production bool is exited.
func (s *CoreVisitor) ExitBool(ctx *parser.BoolContext) {}

// ExitDuration is called when production duration is exited.
func (s *CoreVisitor) ExitDuration(ctx *parser.DurationContext) {}

// ExitJob_def is called when production job_def is exited.
func (s *CoreVisitor) ExitJob_def(ctx *parser.Job_defContext) {
	name := ctx.IDENTIFIER().GetText()
	filename := ctx.GetStop().GetInputStream().GetSourceName()
	job := NewJob(name)
	var commands []Command

	// Process job-level directives (@id, @name, @depend, @for, etc.)
	for _, dirDefCtx := range ctx.AllJob_directive_def() {
		jobDirName := dirDefCtx.IDENTIFIER().GetText()
		var jobDirArgs []interface{}
		line := dirDefCtx.GetStop().GetLine()
		column := dirDefCtx.GetStop().GetColumn()

		for _, dirValue := range dirDefCtx.AllJob_directive_value() {
			if dirValue.Value_expr() != nil {
				jobDirArgs = append(jobDirArgs, ExtractValue(dirValue.Value_expr()))
			} else if dirValue.IDENTIFIER() != nil {
				jobDirArgs = append(jobDirArgs, dirValue.IDENTIFIER().GetText())
			} else if dirValue.MATRIX_REF() != nil {
				jobDirArgs = append(jobDirArgs, dirValue.MATRIX_REF().GetText())
			} else if dirValue.REF() != nil {
				jobDirArgs = append(jobDirArgs, dirValue.REF().GetText())
			} else if dirValue.VALUE() != nil {
				jobDirArgs = append(jobDirArgs, dirValue.VALUE().GetText())
			}
		}

		if err := job.AddJobDirective(jobDirName, jobDirArgs); err != nil {
			ThrowSyntaxError(err.Error(), filename, line, column)
		}
	}

	// Process commands
	for _, cmdCtx := range ctx.AllCmd_expr() {
		var cmdTokens []string
		var cmdDirectives []CMDDirective
		line := cmdCtx.GetStop().GetLine()
		column := cmdCtx.GetStop().GetColumn()

		// Process command directive (@once, @try, @concat, etc.)
		if cmdCtx.Cmd_directive() != nil {
			directive := cmdCtx.Cmd_directive()
			dirName := directive.Cmd_directive_iden().IDENTIFIER().GetText()
			var args []interface{}

			for _, directiveValue := range directive.AllCmd_directive_value() {
				for _, arg := range directiveValue.AllDirective_arg() {
					if arg.Value_expr() != nil {
						args = append(args, ExtractValue(arg.Value_expr()))
					} else if arg.IDENTIFIER() != nil {
						args = append(args, arg.IDENTIFIER().GetText())
					} else if arg.MATRIX_REF() != nil {
						args = append(args, arg.MATRIX_REF().GetText())
					} else if arg.REF() != nil {
						args = append(args, arg.REF().GetText())
					} else if arg.VALUE() != nil {
						args = append(args, arg.VALUE().GetText())
					}
				}
			}

			cmdDirective, err := job.AddCMDDirective(dirName, args)
			if err != nil {
				ThrowSyntaxError(err.Error(), filename, line, column)
			}
			cmdDirectives = append(cmdDirectives, *cmdDirective)
		}

		// Collect command tokens
		for _, token := range cmdCtx.AllIDENTIFIER() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for _, token := range cmdCtx.AllSTRING() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for _, token := range cmdCtx.AllSINGLE_QUOTE_STRING() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for _, token := range cmdCtx.AllNUMBER() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for _, token := range cmdCtx.AllMATRIX_REF() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for _, token := range cmdCtx.AllARG() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for _, token := range cmdCtx.AllSINGLE_DASH_ARG() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for _, token := range cmdCtx.AllEQUALS() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for _, token := range cmdCtx.AllREF() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for _, token := range cmdCtx.AllURL() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for _, token := range cmdCtx.AllVALUE() {
			cmdTokens = append(cmdTokens, token.GetText())
		}
		for range cmdCtx.AllBSL_NEWLINE() {
			cmdTokens = append(cmdTokens, " ") // Replace backslash-newline with space
		}

		// Join tokens into a single command string
		cmdString := strings.Join(cmdTokens, " ")

		// If the job has a @for directive, resolve matrix placeholders
		for _, dir := range job.Directives {
			if dir.Type == "for" {
				matrixName, ok := dir.Value.(string)
				if !ok {
					ThrowRuntimeError(fmt.Sprintf("Invalid value for @for: %v", dir.Value), filename, line, column)
					continue
				}
				if matrixValues, exists := s.Core.ResolveMatrix(matrixName); exists {
					for _, values := range matrixValues {
						tempCmd := cmdString
						tempDirectives := make([]CMDDirective, len(cmdDirectives))
						copy(tempDirectives, cmdDirectives)

						// Replace matrix placeholders in command and directive arguments
						for i, val := range values {
							placeholder := "@value(" + fmt.Sprintf("%d", i+1) + ")"
							tempCmd = strings.ReplaceAll(tempCmd, placeholder, val)
							for j, dir := range tempDirectives {
								if args, ok := dir.Value.([]interface{}); ok && (dir.Type == "concat" || dir.Type == "print") {
									newArgs := make([]interface{}, len(args))
									for k, arg := range args {
										newArgs[k] = strings.ReplaceAll(fmt.Sprintf("%v", arg), placeholder, val)
									}
									tempDirectives[j].Value = newArgs
								}
							}
						}
						// Replace @value with the first value (if used without index)
						if len(values) > 0 {
							tempCmd = strings.ReplaceAll(tempCmd, "@value", values[0])
							for j, dir := range tempDirectives {
								if args, ok := dir.Value.([]interface{}); ok && (dir.Type == "concat" || dir.Type == "print") {
									newArgs := make([]interface{}, len(args))
									for k, arg := range args {
										newArgs[k] = strings.ReplaceAll(fmt.Sprintf("%v", arg), "@value", values[0])
									}
									tempDirectives[j].Value = newArgs
								}
							}
						}

						commands = append(commands, Command{
							Directive: tempDirectives,
							CMD:       tempCmd,
						})
					}
					break
				} else {
					ThrowRuntimeError(fmt.Sprintf("Matrix or variable %s not found for @for", matrixName), filename, line, column)
				}
			}
		}

		// If no @for directive, add the command as-is
		if len(commands) == 0 {
			commands = append(commands, Command{
				Directive: cmdDirectives,
				CMD:       cmdString,
			})
		}
	}

	job.Commands = commands
	s.Core.AddJob(*job)
}
