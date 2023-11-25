// Code generated from Paral.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

// EnterStart is called when production start is entered.
func (s *BaseParalListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseParalListener) ExitStart(ctx *StartContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseParalListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseParalListener) ExitProg(ctx *ProgContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseParalListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseParalListener) ExitVariable(ctx *VariableContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseParalListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseParalListener) ExitExecute(ctx *ExecuteContext) {}
