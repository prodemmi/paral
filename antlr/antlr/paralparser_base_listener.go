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

// EnterJob_definition is called when production job_definition is entered.
func (s *BaseParalParserListener) EnterJob_definition(ctx *Job_definitionContext) {}

// ExitJob_definition is called when production job_definition is exited.
func (s *BaseParalParserListener) ExitJob_definition(ctx *Job_definitionContext) {}

// EnterJob_directive is called when production job_directive is entered.
func (s *BaseParalParserListener) EnterJob_directive(ctx *Job_directiveContext) {}

// ExitJob_directive is called when production job_directive is exited.
func (s *BaseParalParserListener) ExitJob_directive(ctx *Job_directiveContext) {}

// EnterCommand_block is called when production command_block is entered.
func (s *BaseParalParserListener) EnterCommand_block(ctx *Command_blockContext) {}

// ExitCommand_block is called when production command_block is exited.
func (s *BaseParalParserListener) ExitCommand_block(ctx *Command_blockContext) {}

// EnterCommand_content is called when production command_content is entered.
func (s *BaseParalParserListener) EnterCommand_content(ctx *Command_contentContext) {}

// ExitCommand_content is called when production command_content is exited.
func (s *BaseParalParserListener) ExitCommand_content(ctx *Command_contentContext) {}

// EnterCommand_item is called when production command_item is entered.
func (s *BaseParalParserListener) EnterCommand_item(ctx *Command_itemContext) {}

// ExitCommand_item is called when production command_item is exited.
func (s *BaseParalParserListener) ExitCommand_item(ctx *Command_itemContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseParalParserListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseParalParserListener) ExitDirective(ctx *DirectiveContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseParalParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseParalParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseParalParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseParalParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterNested_function is called when production nested_function is entered.
func (s *BaseParalParserListener) EnterNested_function(ctx *Nested_functionContext) {}

// ExitNested_function is called when production nested_function is exited.
func (s *BaseParalParserListener) ExitNested_function(ctx *Nested_functionContext) {}

// EnterIf_condition is called when production if_condition is entered.
func (s *BaseParalParserListener) EnterIf_condition(ctx *If_conditionContext) {}

// ExitIf_condition is called when production if_condition is exited.
func (s *BaseParalParserListener) ExitIf_condition(ctx *If_conditionContext) {}

// EnterArgument_list is called when production argument_list is entered.
func (s *BaseParalParserListener) EnterArgument_list(ctx *Argument_listContext) {}

// ExitArgument_list is called when production argument_list is exited.
func (s *BaseParalParserListener) ExitArgument_list(ctx *Argument_listContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseParalParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseParalParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVariable_value is called when production variable_value is entered.
func (s *BaseParalParserListener) EnterVariable_value(ctx *Variable_valueContext) {}

// ExitVariable_value is called when production variable_value is exited.
func (s *BaseParalParserListener) ExitVariable_value(ctx *Variable_valueContext) {}

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
