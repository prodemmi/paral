// Code generated from antlr/Paral.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Paral

import "github.com/antlr4-go/antlr/v4"

// ParalListener is a complete listener for a parse tree produced by ParalParser.
type ParalListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterVariable_def is called when entering the variable_def production.
	EnterVariable_def(c *Variable_defContext)

	// EnterMatrix_def is called when entering the matrix_def production.
	EnterMatrix_def(c *Matrix_defContext)

	// EnterJob_def is called when entering the job_def production.
	EnterJob_def(c *Job_defContext)

	// EnterCmd_expr is called when entering the cmd_expr production.
	EnterCmd_expr(c *Cmd_exprContext)

	// EnterCmd_value is called when entering the cmd_value production.
	EnterCmd_value(c *Cmd_valueContext)

	// EnterDirective_expr is called when entering the directive_expr production.
	EnterDirective_expr(c *Directive_exprContext)

	// EnterDirective_args_expr is called when entering the directive_args_expr production.
	EnterDirective_args_expr(c *Directive_args_exprContext)

	// EnterFlagWithValue is called when entering the flagWithValue production.
	EnterFlagWithValue(c *FlagWithValueContext)

	// EnterFlagAlone is called when entering the flagAlone production.
	EnterFlagAlone(c *FlagAloneContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterSingleQuoteString is called when entering the singleQuoteString production.
	EnterSingleQuoteString(c *SingleQuoteStringContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// EnterDuration is called when entering the duration production.
	EnterDuration(c *DurationContext)

	// EnterList_expr is called when entering the list_expr production.
	EnterList_expr(c *List_exprContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitVariable_def is called when exiting the variable_def production.
	ExitVariable_def(c *Variable_defContext)

	// ExitMatrix_def is called when exiting the matrix_def production.
	ExitMatrix_def(c *Matrix_defContext)

	// ExitJob_def is called when exiting the job_def production.
	ExitJob_def(c *Job_defContext)

	// ExitCmd_expr is called when exiting the cmd_expr production.
	ExitCmd_expr(c *Cmd_exprContext)

	// ExitCmd_value is called when exiting the cmd_value production.
	ExitCmd_value(c *Cmd_valueContext)

	// ExitDirective_expr is called when exiting the directive_expr production.
	ExitDirective_expr(c *Directive_exprContext)

	// ExitDirective_args_expr is called when exiting the directive_args_expr production.
	ExitDirective_args_expr(c *Directive_args_exprContext)

	// ExitFlagWithValue is called when exiting the flagWithValue production.
	ExitFlagWithValue(c *FlagWithValueContext)

	// ExitFlagAlone is called when exiting the flagAlone production.
	ExitFlagAlone(c *FlagAloneContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitSingleQuoteString is called when exiting the singleQuoteString production.
	ExitSingleQuoteString(c *SingleQuoteStringContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)

	// ExitDuration is called when exiting the duration production.
	ExitDuration(c *DurationContext)

	// ExitList_expr is called when exiting the list_expr production.
	ExitList_expr(c *List_exprContext)
}
