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

	// EnterList_expr is called when entering the list_expr production.
	EnterList_expr(c *List_exprContext)

	// EnterComment_def is called when entering the comment_def production.
	EnterComment_def(c *Comment_defContext)

	// EnterJob_def is called when entering the job_def production.
	EnterJob_def(c *Job_defContext)

	// EnterJob_directive_def is called when entering the job_directive_def production.
	EnterJob_directive_def(c *Job_directive_defContext)

	// EnterJob_directive_value is called when entering the job_directive_value production.
	EnterJob_directive_value(c *Job_directive_valueContext)

	// EnterCmd_expr is called when entering the cmd_expr production.
	EnterCmd_expr(c *Cmd_exprContext)

	// EnterCmd_directive is called when entering the cmd_directive production.
	EnterCmd_directive(c *Cmd_directiveContext)

	// EnterCmd_directive_iden is called when entering the cmd_directive_iden production.
	EnterCmd_directive_iden(c *Cmd_directive_idenContext)

	// EnterCmd_directive_value is called when entering the cmd_directive_value production.
	EnterCmd_directive_value(c *Cmd_directive_valueContext)

	// EnterDirective_arg is called when entering the directive_arg production.
	EnterDirective_arg(c *Directive_argContext)

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

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitVariable_def is called when exiting the variable_def production.
	ExitVariable_def(c *Variable_defContext)

	// ExitMatrix_def is called when exiting the matrix_def production.
	ExitMatrix_def(c *Matrix_defContext)

	// ExitList_expr is called when exiting the list_expr production.
	ExitList_expr(c *List_exprContext)

	// ExitComment_def is called when exiting the comment_def production.
	ExitComment_def(c *Comment_defContext)

	// ExitJob_def is called when exiting the job_def production.
	ExitJob_def(c *Job_defContext)

	// ExitJob_directive_def is called when exiting the job_directive_def production.
	ExitJob_directive_def(c *Job_directive_defContext)

	// ExitJob_directive_value is called when exiting the job_directive_value production.
	ExitJob_directive_value(c *Job_directive_valueContext)

	// ExitCmd_expr is called when exiting the cmd_expr production.
	ExitCmd_expr(c *Cmd_exprContext)

	// ExitCmd_directive is called when exiting the cmd_directive production.
	ExitCmd_directive(c *Cmd_directiveContext)

	// ExitCmd_directive_iden is called when exiting the cmd_directive_iden production.
	ExitCmd_directive_iden(c *Cmd_directive_idenContext)

	// ExitCmd_directive_value is called when exiting the cmd_directive_value production.
	ExitCmd_directive_value(c *Cmd_directive_valueContext)

	// ExitDirective_arg is called when exiting the directive_arg production.
	ExitDirective_arg(c *Directive_argContext)

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
}
