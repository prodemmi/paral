// Code generated from antlr/ParalParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

	// EnterTask_definition is called when entering the task_definition production.
	EnterTask_definition(c *Task_definitionContext)

	// EnterTask_directive is called when entering the task_directive production.
	EnterTask_directive(c *Task_directiveContext)

	// EnterPipeline_block is called when entering the pipeline_block production.
	EnterPipeline_block(c *Pipeline_blockContext)

	// EnterPipeline_content is called when entering the pipeline_content production.
	EnterPipeline_content(c *Pipeline_contentContext)

	// EnterPipeline_item is called when entering the pipeline_item production.
	EnterPipeline_item(c *Pipeline_itemContext)

	// EnterUnknown_command is called when entering the unknown_command production.
	EnterUnknown_command(c *Unknown_commandContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterBuf is called when entering the buf production.
	EnterBuf(c *BufContext)

	// EnterStash is called when entering the stash production.
	EnterStash(c *StashContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterIf_condition is called when entering the if_condition production.
	EnterIf_condition(c *If_conditionContext)

	// EnterElseif_condition is called when entering the elseif_condition production.
	EnterElseif_condition(c *Elseif_conditionContext)

	// EnterElse_condition is called when entering the else_condition production.
	EnterElse_condition(c *Else_conditionContext)

	// EnterTry_catch is called when entering the try_catch production.
	EnterTry_catch(c *Try_catchContext)

	// EnterTry_block is called when entering the try_block production.
	EnterTry_block(c *Try_blockContext)

	// EnterCatch_block is called when entering the catch_block production.
	EnterCatch_block(c *Catch_blockContext)

	// EnterMatch_statement is called when entering the match_statement production.
	EnterMatch_statement(c *Match_statementContext)

	// EnterMatch_block is called when entering the match_block production.
	EnterMatch_block(c *Match_blockContext)

	// EnterMatch_expression is called when entering the match_expression production.
	EnterMatch_expression(c *Match_expressionContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterNested_function is called when entering the nested_function production.
	EnterNested_function(c *Nested_functionContext)

	// EnterArgument_list is called when entering the argument_list production.
	EnterArgument_list(c *Argument_listContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

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

	// EnterError_variable is called when entering the error_variable production.
	EnterError_variable(c *Error_variableContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVariable_assignment is called when exiting the variable_assignment production.
	ExitVariable_assignment(c *Variable_assignmentContext)

	// ExitTask_definition is called when exiting the task_definition production.
	ExitTask_definition(c *Task_definitionContext)

	// ExitTask_directive is called when exiting the task_directive production.
	ExitTask_directive(c *Task_directiveContext)

	// ExitPipeline_block is called when exiting the pipeline_block production.
	ExitPipeline_block(c *Pipeline_blockContext)

	// ExitPipeline_content is called when exiting the pipeline_content production.
	ExitPipeline_content(c *Pipeline_contentContext)

	// ExitPipeline_item is called when exiting the pipeline_item production.
	ExitPipeline_item(c *Pipeline_itemContext)

	// ExitUnknown_command is called when exiting the unknown_command production.
	ExitUnknown_command(c *Unknown_commandContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitBuf is called when exiting the buf production.
	ExitBuf(c *BufContext)

	// ExitStash is called when exiting the stash production.
	ExitStash(c *StashContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitIf_condition is called when exiting the if_condition production.
	ExitIf_condition(c *If_conditionContext)

	// ExitElseif_condition is called when exiting the elseif_condition production.
	ExitElseif_condition(c *Elseif_conditionContext)

	// ExitElse_condition is called when exiting the else_condition production.
	ExitElse_condition(c *Else_conditionContext)

	// ExitTry_catch is called when exiting the try_catch production.
	ExitTry_catch(c *Try_catchContext)

	// ExitTry_block is called when exiting the try_block production.
	ExitTry_block(c *Try_blockContext)

	// ExitCatch_block is called when exiting the catch_block production.
	ExitCatch_block(c *Catch_blockContext)

	// ExitMatch_statement is called when exiting the match_statement production.
	ExitMatch_statement(c *Match_statementContext)

	// ExitMatch_block is called when exiting the match_block production.
	ExitMatch_block(c *Match_blockContext)

	// ExitMatch_expression is called when exiting the match_expression production.
	ExitMatch_expression(c *Match_expressionContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitNested_function is called when exiting the nested_function production.
	ExitNested_function(c *Nested_functionContext)

	// ExitArgument_list is called when exiting the argument_list production.
	ExitArgument_list(c *Argument_listContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

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

	// ExitError_variable is called when exiting the error_variable production.
	ExitError_variable(c *Error_variableContext)
}
