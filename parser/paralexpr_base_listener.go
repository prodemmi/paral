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

// EnterProg is called when production prog is entered.
func (s *BaseParalExprListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseParalExprListener) ExitProg(ctx *ProgContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseParalExprListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseParalExprListener) ExitExpr(ctx *ExprContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseParalExprListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseParalExprListener) ExitVariable(ctx *VariableContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseParalExprListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseParalExprListener) ExitExecute(ctx *ExecuteContext) {}
