// Code generated from /home/emad/Documents/paral/ParalExpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ParalExpr
import "github.com/antlr4-go/antlr/v4"

// ParalExprListener is a complete listener for a parse tree produced by ParalExprParser.
type ParalExprListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterVariables is called when entering the variables production.
	EnterVariables(c *VariablesContext)

	// EnterExecuteables is called when entering the executeables production.
	EnterExecuteables(c *ExecuteablesContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitVariables is called when exiting the variables production.
	ExitVariables(c *VariablesContext)

	// ExitExecuteables is called when exiting the executeables production.
	ExitExecuteables(c *ExecuteablesContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)
}
