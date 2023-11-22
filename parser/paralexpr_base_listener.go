// Code generated from /home/emad/Documents/paral/ParalExpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ParalExpr
import "github.com/antlr4-go/antlr/v4"

// BaseParalExprListener is a complete listener for a parse tree produced by ParalExprParser.
type BaseParalExprListener struct{}

var _ ParalExprListener = &BaseParalExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParalExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParalExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParalExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParalExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseParalExprListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseParalExprListener) ExitStart(ctx *StartContext) {}

// EnterVariables is called when production variables is entered.
func (s *BaseParalExprListener) EnterVariables(ctx *VariablesContext) {}

// ExitVariables is called when production variables is exited.
func (s *BaseParalExprListener) ExitVariables(ctx *VariablesContext) {}

// EnterExecuteables is called when production executeables is entered.
func (s *BaseParalExprListener) EnterExecuteables(ctx *ExecuteablesContext) {}

// ExitExecuteables is called when production executeables is exited.
func (s *BaseParalExprListener) ExitExecuteables(ctx *ExecuteablesContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseParalExprListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseParalExprListener) ExitVariable(ctx *VariableContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseParalExprListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseParalExprListener) ExitExecute(ctx *ExecuteContext) {}
