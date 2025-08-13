// Code generated from antlr/ParalParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ParalParser

import "github.com/antlr4-go/antlr/v4"

// BaseParalParserListener is a complete listener for a parse tree produced by ParalParser.
type BaseParalParserListener struct{}

var _ ParalParserListener = &BaseParalParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParalParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParalParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParalParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParalParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseParalParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseParalParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseParalParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseParalParserListener) ExitStatement(ctx *StatementContext) {}

// EnterVariable_assignment is called when production variable_assignment is entered.
func (s *BaseParalParserListener) EnterVariable_assignment(ctx *Variable_assignmentContext) {}

// ExitVariable_assignment is called when production variable_assignment is exited.
func (s *BaseParalParserListener) ExitVariable_assignment(ctx *Variable_assignmentContext) {}

// EnterTask_definition is called when production task_definition is entered.
func (s *BaseParalParserListener) EnterTask_definition(ctx *Task_definitionContext) {}

// ExitTask_definition is called when production task_definition is exited.
func (s *BaseParalParserListener) ExitTask_definition(ctx *Task_definitionContext) {}

// EnterTask_directive is called when production task_directive is entered.
func (s *BaseParalParserListener) EnterTask_directive(ctx *Task_directiveContext) {}

// ExitTask_directive is called when production task_directive is exited.
func (s *BaseParalParserListener) ExitTask_directive(ctx *Task_directiveContext) {}

// EnterPipeline_block is called when production pipeline_block is entered.
func (s *BaseParalParserListener) EnterPipeline_block(ctx *Pipeline_blockContext) {}

// ExitPipeline_block is called when production pipeline_block is exited.
func (s *BaseParalParserListener) ExitPipeline_block(ctx *Pipeline_blockContext) {}

// EnterPipeline_content is called when production pipeline_content is entered.
func (s *BaseParalParserListener) EnterPipeline_content(ctx *Pipeline_contentContext) {}

// ExitPipeline_content is called when production pipeline_content is exited.
func (s *BaseParalParserListener) ExitPipeline_content(ctx *Pipeline_contentContext) {}

// EnterPipeline_item is called when production pipeline_item is entered.
func (s *BaseParalParserListener) EnterPipeline_item(ctx *Pipeline_itemContext) {}

// ExitPipeline_item is called when production pipeline_item is exited.
func (s *BaseParalParserListener) ExitPipeline_item(ctx *Pipeline_itemContext) {}

// EnterUnknown_command is called when production unknown_command is entered.
func (s *BaseParalParserListener) EnterUnknown_command(ctx *Unknown_commandContext) {}

// ExitUnknown_command is called when production unknown_command is exited.
func (s *BaseParalParserListener) ExitUnknown_command(ctx *Unknown_commandContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseParalParserListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseParalParserListener) ExitDirective(ctx *DirectiveContext) {}

// EnterBuf is called when production buf is entered.
func (s *BaseParalParserListener) EnterBuf(ctx *BufContext) {}

// ExitBuf is called when production buf is exited.
func (s *BaseParalParserListener) ExitBuf(ctx *BufContext) {}

// EnterStash is called when production stash is entered.
func (s *BaseParalParserListener) EnterStash(ctx *StashContext) {}

// ExitStash is called when production stash is exited.
func (s *BaseParalParserListener) ExitStash(ctx *StashContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseParalParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseParalParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterIf_condition is called when production if_condition is entered.
func (s *BaseParalParserListener) EnterIf_condition(ctx *If_conditionContext) {}

// ExitIf_condition is called when production if_condition is exited.
func (s *BaseParalParserListener) ExitIf_condition(ctx *If_conditionContext) {}

// EnterElseif_condition is called when production elseif_condition is entered.
func (s *BaseParalParserListener) EnterElseif_condition(ctx *Elseif_conditionContext) {}

// ExitElseif_condition is called when production elseif_condition is exited.
func (s *BaseParalParserListener) ExitElseif_condition(ctx *Elseif_conditionContext) {}

// EnterElse_condition is called when production else_condition is entered.
func (s *BaseParalParserListener) EnterElse_condition(ctx *Else_conditionContext) {}

// ExitElse_condition is called when production else_condition is exited.
func (s *BaseParalParserListener) ExitElse_condition(ctx *Else_conditionContext) {}

// EnterTry_catch is called when production try_catch is entered.
func (s *BaseParalParserListener) EnterTry_catch(ctx *Try_catchContext) {}

// ExitTry_catch is called when production try_catch is exited.
func (s *BaseParalParserListener) ExitTry_catch(ctx *Try_catchContext) {}

// EnterTry_block is called when production try_block is entered.
func (s *BaseParalParserListener) EnterTry_block(ctx *Try_blockContext) {}

// ExitTry_block is called when production try_block is exited.
func (s *BaseParalParserListener) ExitTry_block(ctx *Try_blockContext) {}

// EnterCatch_block is called when production catch_block is entered.
func (s *BaseParalParserListener) EnterCatch_block(ctx *Catch_blockContext) {}

// ExitCatch_block is called when production catch_block is exited.
func (s *BaseParalParserListener) ExitCatch_block(ctx *Catch_blockContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseParalParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseParalParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterNested_function is called when production nested_function is entered.
func (s *BaseParalParserListener) EnterNested_function(ctx *Nested_functionContext) {}

// ExitNested_function is called when production nested_function is exited.
func (s *BaseParalParserListener) ExitNested_function(ctx *Nested_functionContext) {}

// EnterArgument_list is called when production argument_list is entered.
func (s *BaseParalParserListener) EnterArgument_list(ctx *Argument_listContext) {}

// ExitArgument_list is called when production argument_list is exited.
func (s *BaseParalParserListener) ExitArgument_list(ctx *Argument_listContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseParalParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseParalParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNumber_expr is called when production number_expr is entered.
func (s *BaseParalParserListener) EnterNumber_expr(ctx *Number_exprContext) {}

// ExitNumber_expr is called when production number_expr is exited.
func (s *BaseParalParserListener) ExitNumber_expr(ctx *Number_exprContext) {}

// EnterString_expr is called when production string_expr is entered.
func (s *BaseParalParserListener) EnterString_expr(ctx *String_exprContext) {}

// ExitString_expr is called when production string_expr is exited.
func (s *BaseParalParserListener) ExitString_expr(ctx *String_exprContext) {}

// EnterBoolean_expr is called when production boolean_expr is entered.
func (s *BaseParalParserListener) EnterBoolean_expr(ctx *Boolean_exprContext) {}

// ExitBoolean_expr is called when production boolean_expr is exited.
func (s *BaseParalParserListener) ExitBoolean_expr(ctx *Boolean_exprContext) {}

// EnterDuration_expr is called when production duration_expr is entered.
func (s *BaseParalParserListener) EnterDuration_expr(ctx *Duration_exprContext) {}

// ExitDuration_expr is called when production duration_expr is exited.
func (s *BaseParalParserListener) ExitDuration_expr(ctx *Duration_exprContext) {}

// EnterMatrix_expr is called when production matrix_expr is entered.
func (s *BaseParalParserListener) EnterMatrix_expr(ctx *Matrix_exprContext) {}

// ExitMatrix_expr is called when production matrix_expr is exited.
func (s *BaseParalParserListener) ExitMatrix_expr(ctx *Matrix_exprContext) {}

// EnterList_expr is called when production list_expr is entered.
func (s *BaseParalParserListener) EnterList_expr(ctx *List_exprContext) {}

// ExitList_expr is called when production list_expr is exited.
func (s *BaseParalParserListener) ExitList_expr(ctx *List_exprContext) {}

// EnterLoop_variable is called when production loop_variable is entered.
func (s *BaseParalParserListener) EnterLoop_variable(ctx *Loop_variableContext) {}

// ExitLoop_variable is called when production loop_variable is exited.
func (s *BaseParalParserListener) ExitLoop_variable(ctx *Loop_variableContext) {}

// EnterError_variable is called when production error_variable is entered.
func (s *BaseParalParserListener) EnterError_variable(ctx *Error_variableContext) {}

// ExitError_variable is called when production error_variable is exited.
func (s *BaseParalParserListener) ExitError_variable(ctx *Error_variableContext) {}
