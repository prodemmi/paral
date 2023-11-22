// Code generated from ParalExpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ParalExpr

import "github.com/antlr4-go/antlr/v4"

// ParalExprListener is a complete listener for a parse tree produced by ParalExprParser.
type ParalExprListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterVariables is called when entering the variables production.
	EnterVariables(c *VariablesContext)

	// EnterExecutables is called when entering the executables production.
	EnterExecutables(c *ExecutablesContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitVariables is called when exiting the variables production.
	ExitVariables(c *VariablesContext)

	// ExitExecutables is called when exiting the executables production.
	ExitExecutables(c *ExecutablesContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)
}
