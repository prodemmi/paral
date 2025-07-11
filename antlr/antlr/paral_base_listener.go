// Code generated from antlr/Paral.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Paral

import "github.com/antlr4-go/antlr/v4"

// BaseParalListener is a complete listener for a parse tree produced by ParalParser.
type BaseParalListener struct{}

var _ ParalListener = &BaseParalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseParalListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseParalListener) ExitProgram(ctx *ProgramContext) {}

// EnterLine is called when production line is entered.
func (s *BaseParalListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseParalListener) ExitLine(ctx *LineContext) {}

// EnterVariable_def is called when production variable_def is entered.
func (s *BaseParalListener) EnterVariable_def(ctx *Variable_defContext) {}

// ExitVariable_def is called when production variable_def is exited.
func (s *BaseParalListener) ExitVariable_def(ctx *Variable_defContext) {}

// EnterMatrix_def is called when production matrix_def is entered.
func (s *BaseParalListener) EnterMatrix_def(ctx *Matrix_defContext) {}

// ExitMatrix_def is called when production matrix_def is exited.
func (s *BaseParalListener) ExitMatrix_def(ctx *Matrix_defContext) {}

// EnterList_expr is called when production list_expr is entered.
func (s *BaseParalListener) EnterList_expr(ctx *List_exprContext) {}

// ExitList_expr is called when production list_expr is exited.
func (s *BaseParalListener) ExitList_expr(ctx *List_exprContext) {}

// EnterComment_def is called when production comment_def is entered.
func (s *BaseParalListener) EnterComment_def(ctx *Comment_defContext) {}

// ExitComment_def is called when production comment_def is exited.
func (s *BaseParalListener) ExitComment_def(ctx *Comment_defContext) {}

// EnterJob_def is called when production job_def is entered.
func (s *BaseParalListener) EnterJob_def(ctx *Job_defContext) {}

// ExitJob_def is called when production job_def is exited.
func (s *BaseParalListener) ExitJob_def(ctx *Job_defContext) {}

// EnterJob_directive_def is called when production job_directive_def is entered.
func (s *BaseParalListener) EnterJob_directive_def(ctx *Job_directive_defContext) {}

// ExitJob_directive_def is called when production job_directive_def is exited.
func (s *BaseParalListener) ExitJob_directive_def(ctx *Job_directive_defContext) {}

// EnterJob_directive_value is called when production job_directive_value is entered.
func (s *BaseParalListener) EnterJob_directive_value(ctx *Job_directive_valueContext) {}

// ExitJob_directive_value is called when production job_directive_value is exited.
func (s *BaseParalListener) ExitJob_directive_value(ctx *Job_directive_valueContext) {}

// EnterCmd_expr is called when production cmd_expr is entered.
func (s *BaseParalListener) EnterCmd_expr(ctx *Cmd_exprContext) {}

// ExitCmd_expr is called when production cmd_expr is exited.
func (s *BaseParalListener) ExitCmd_expr(ctx *Cmd_exprContext) {}

// EnterCmd_directive is called when production cmd_directive is entered.
func (s *BaseParalListener) EnterCmd_directive(ctx *Cmd_directiveContext) {}

// ExitCmd_directive is called when production cmd_directive is exited.
func (s *BaseParalListener) ExitCmd_directive(ctx *Cmd_directiveContext) {}

// EnterCmd_directive_iden is called when production cmd_directive_iden is entered.
func (s *BaseParalListener) EnterCmd_directive_iden(ctx *Cmd_directive_idenContext) {}

// ExitCmd_directive_iden is called when production cmd_directive_iden is exited.
func (s *BaseParalListener) ExitCmd_directive_iden(ctx *Cmd_directive_idenContext) {}

// EnterCmd_directive_value is called when production cmd_directive_value is entered.
func (s *BaseParalListener) EnterCmd_directive_value(ctx *Cmd_directive_valueContext) {}

// ExitCmd_directive_value is called when production cmd_directive_value is exited.
func (s *BaseParalListener) ExitCmd_directive_value(ctx *Cmd_directive_valueContext) {}

// EnterDirective_arg is called when production directive_arg is entered.
func (s *BaseParalListener) EnterDirective_arg(ctx *Directive_argContext) {}

// ExitDirective_arg is called when production directive_arg is exited.
func (s *BaseParalListener) ExitDirective_arg(ctx *Directive_argContext) {}

// EnterString is called when production string is entered.
func (s *BaseParalListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseParalListener) ExitString(ctx *StringContext) {}

// EnterSingleQuoteString is called when production singleQuoteString is entered.
func (s *BaseParalListener) EnterSingleQuoteString(ctx *SingleQuoteStringContext) {}

// ExitSingleQuoteString is called when production singleQuoteString is exited.
func (s *BaseParalListener) ExitSingleQuoteString(ctx *SingleQuoteStringContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseParalListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseParalListener) ExitNumber(ctx *NumberContext) {}

// EnterBool is called when production bool is entered.
func (s *BaseParalListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BaseParalListener) ExitBool(ctx *BoolContext) {}

// EnterDuration is called when production duration is entered.
func (s *BaseParalListener) EnterDuration(ctx *DurationContext) {}

// ExitDuration is called when production duration is exited.
func (s *BaseParalListener) ExitDuration(ctx *DurationContext) {}
