// Code generated from antlr/ParalParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ParalParser

import "github.com/antlr4-go/antlr/v4"

// ParalParserListener is a complete listener for a parse tree produced by ParalParser.
type ParalParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVariable_assignment is called when entering the variable_assignment production.
	EnterVariable_assignment(c *Variable_assignmentContext)

	// EnterJob_definition is called when entering the job_definition production.
	EnterJob_definition(c *Job_definitionContext)

	// EnterJob_directive is called when entering the job_directive production.
	EnterJob_directive(c *Job_directiveContext)

	// EnterCommand_block is called when entering the command_block production.
	EnterCommand_block(c *Command_blockContext)

	// EnterCommand_content is called when entering the command_content production.
	EnterCommand_content(c *Command_contentContext)

	// EnterCommand_item is called when entering the command_item production.
	EnterCommand_item(c *Command_itemContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterNested_function is called when entering the nested_function production.
	EnterNested_function(c *Nested_functionContext)

	// EnterIf_condition is called when entering the if_condition production.
	EnterIf_condition(c *If_conditionContext)

	// EnterArgument_list is called when entering the argument_list production.
	EnterArgument_list(c *Argument_listContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVariable_value is called when entering the variable_value production.
	EnterVariable_value(c *Variable_valueContext)

	// EnterNumber_expr is called when entering the number_expr production.
	EnterNumber_expr(c *Number_exprContext)

	// EnterString_expr is called when entering the string_expr production.
	EnterString_expr(c *String_exprContext)

	// EnterBoolean_expr is called when entering the boolean_expr production.
	EnterBoolean_expr(c *Boolean_exprContext)

	// EnterDuration_expr is called when entering the duration_expr production.
	EnterDuration_expr(c *Duration_exprContext)

	// EnterMatrix_expr is called when entering the matrix_expr production.
	EnterMatrix_expr(c *Matrix_exprContext)

	// EnterList_expr is called when entering the list_expr production.
	EnterList_expr(c *List_exprContext)

	// EnterLoop_variable is called when entering the loop_variable production.
	EnterLoop_variable(c *Loop_variableContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVariable_assignment is called when exiting the variable_assignment production.
	ExitVariable_assignment(c *Variable_assignmentContext)

	// ExitJob_definition is called when exiting the job_definition production.
	ExitJob_definition(c *Job_definitionContext)

	// ExitJob_directive is called when exiting the job_directive production.
	ExitJob_directive(c *Job_directiveContext)

	// ExitCommand_block is called when exiting the command_block production.
	ExitCommand_block(c *Command_blockContext)

	// ExitCommand_content is called when exiting the command_content production.
	ExitCommand_content(c *Command_contentContext)

	// ExitCommand_item is called when exiting the command_item production.
	ExitCommand_item(c *Command_itemContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitNested_function is called when exiting the nested_function production.
	ExitNested_function(c *Nested_functionContext)

	// ExitIf_condition is called when exiting the if_condition production.
	ExitIf_condition(c *If_conditionContext)

	// ExitArgument_list is called when exiting the argument_list production.
	ExitArgument_list(c *Argument_listContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVariable_value is called when exiting the variable_value production.
	ExitVariable_value(c *Variable_valueContext)

	// ExitNumber_expr is called when exiting the number_expr production.
	ExitNumber_expr(c *Number_exprContext)

	// ExitString_expr is called when exiting the string_expr production.
	ExitString_expr(c *String_exprContext)

	// ExitBoolean_expr is called when exiting the boolean_expr production.
	ExitBoolean_expr(c *Boolean_exprContext)

	// ExitDuration_expr is called when exiting the duration_expr production.
	ExitDuration_expr(c *Duration_exprContext)

	// ExitMatrix_expr is called when exiting the matrix_expr production.
	ExitMatrix_expr(c *Matrix_exprContext)

	// ExitList_expr is called when exiting the list_expr production.
	ExitList_expr(c *List_exprContext)

	// ExitLoop_variable is called when exiting the loop_variable production.
	ExitLoop_variable(c *Loop_variableContext)
}
