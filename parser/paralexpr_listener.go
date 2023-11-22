// Code generated from /home/emad/Documents/paral/ParalExpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ParalExpr
import "github.com/antlr4-go/antlr/v4"

// ParalExprListener is a complete listener for a parse tree produced by ParalExprParser.
type ParalExprListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)
}
