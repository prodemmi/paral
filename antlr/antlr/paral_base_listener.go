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

// EnterJob_def is called when production job_def is entered.
func (s *BaseParalListener) EnterJob_def(ctx *Job_defContext) {}

// ExitJob_def is called when production job_def is exited.
func (s *BaseParalListener) ExitJob_def(ctx *Job_defContext) {}

// EnterCmd_expr is called when production cmd_expr is entered.
func (s *BaseParalListener) EnterCmd_expr(ctx *Cmd_exprContext) {}

// ExitCmd_expr is called when production cmd_expr is exited.
func (s *BaseParalListener) ExitCmd_expr(ctx *Cmd_exprContext) {}

// EnterCmd_value is called when production cmd_value is entered.
func (s *BaseParalListener) EnterCmd_value(ctx *Cmd_valueContext) {}

// ExitCmd_value is called when production cmd_value is exited.
func (s *BaseParalListener) ExitCmd_value(ctx *Cmd_valueContext) {}

// EnterDirective_expr is called when production directive_expr is entered.
func (s *BaseParalListener) EnterDirective_expr(ctx *Directive_exprContext) {}

// ExitDirective_expr is called when production directive_expr is exited.
func (s *BaseParalListener) ExitDirective_expr(ctx *Directive_exprContext) {}

// EnterDirective_args_expr is called when production directive_args_expr is entered.
func (s *BaseParalListener) EnterDirective_args_expr(ctx *Directive_args_exprContext) {}

// ExitDirective_args_expr is called when production directive_args_expr is exited.
func (s *BaseParalListener) ExitDirective_args_expr(ctx *Directive_args_exprContext) {}

// EnterFlagWithValue is called when production flagWithValue is entered.
func (s *BaseParalListener) EnterFlagWithValue(ctx *FlagWithValueContext) {}

// ExitFlagWithValue is called when production flagWithValue is exited.
func (s *BaseParalListener) ExitFlagWithValue(ctx *FlagWithValueContext) {}

// EnterFlagAlone is called when production flagAlone is entered.
func (s *BaseParalListener) EnterFlagAlone(ctx *FlagAloneContext) {}

// ExitFlagAlone is called when production flagAlone is exited.
func (s *BaseParalListener) ExitFlagAlone(ctx *FlagAloneContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseParalListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseParalListener) ExitIdentifier(ctx *IdentifierContext) {}

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

// EnterList_expr is called when production list_expr is entered.
func (s *BaseParalListener) EnterList_expr(ctx *List_exprContext) {}

// ExitList_expr is called when production list_expr is exited.
func (s *BaseParalListener) ExitList_expr(ctx *List_exprContext) {}
