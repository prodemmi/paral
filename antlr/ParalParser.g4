parser grammar ParalParser;

options { tokenVocab=ParalLexer; }

program
    : statement* EOF
    ;

statement
    : variable_assignment
    | task_definition
    | NEWLINE
    ;

variable_assignment
    : IDENTIFIER ASSIGN expression NEWLINE
    ;

task_definition
    : (task_directive NEWLINE)* TASK IDENTIFIER NEWLINE* BLOCK_START NEWLINE* pipeline_block* NEWLINE* BLOCK_END
    ;

task_directive
    : directive
    ;

pipeline_block
    : PIPELINE_START NEWLINE* pipeline_content (NEWLINE | EOF | BLOCK_END)?
    ;
    
pipeline_content
    : pipeline_item*
    ;

pipeline_item
    : buf
    | stash
    | condition
    | try_catch
    | match_statement
    | function
    | unknown_command
    ;

unknown_command: (IDENTIFIER | UNKNOWN_TEXT | PATH | URL)+ (expression)*;

// ----------------- task directive -------------

directive
    : AT IDENTIFIER expression?
    ;

// ----------------- functions and conditions -------------

buf
    : PIPELINE_BUF string_expr RBRACK BUF_DOUBLE_BACK_ARROW pipeline_content
    ;

stash
    : PIPELINE_STASH string_expr RBRACK STASH_DOUBLE_BACK_ARROW pipeline_content
    ;

condition
    : if_condition
    ;

if_condition
    : PIPELINE_IF_CALL_START expression RRBRACK
      NEWLINE* BLOCK_START NEWLINE* pipeline_block* NEWLINE* BLOCK_END
      (NEWLINE* elseif_condition)* (NEWLINE* else_condition)?
    ;

elseif_condition
    : PIPELINE_ELSEIF_CALL_START expression RRBRACK
      NEWLINE* BLOCK_START NEWLINE* pipeline_block* NEWLINE* BLOCK_END
    ;

else_condition
    : PIPELINE_ELSE_CALL
      NEWLINE* BLOCK_START NEWLINE* pipeline_block* NEWLINE* BLOCK_END
    ;

try_catch
    : TRY NEWLINE* BLOCK_START NEWLINE* try_block NEWLINE* BLOCK_END (CATCH NEWLINE* BLOCK_START NEWLINE* catch_block NEWLINE* BLOCK_END)?
    ;

try_block
    : pipeline_block*
    ;

catch_block
    : pipeline_block*
    ;

match_statement
    : PIPELINE_MATCH_CALL expression RRBRACK
      NEWLINE* BLOCK_START NEWLINE* match_block* NEWLINE* BLOCK_END
    ;

match_block
    : match_expression pipeline_block
    ;

match_expression
    : expression 
    | UNDERSCORE
    ;

function
    : PIPELINE_FUNCTION_CALL_START argument_list? RRBRACK
    | FUNCTION_START argument_list? RRBRACK
    | EXPRESSION_FUNCTION_CALL_START argument_list? RRBRACK
    ;

nested_function
    : NESTED_FUNCTION_START argument_list? RRBRACK
    ;

// Argument list for functions
argument_list
    : expression (NEWLINE* COMMA NEWLINE* expression)*
    ;

// Expression can be a value or a nested function
expression
    : loop_variable
    | error_variable
    | function
    | nested_function
    | URL
    | PATH
    | number_expr
    | string_expr
    | boolean_expr
    | duration_expr
    | matrix_expr
    | list_expr
    | IDENTIFIER
    ;

number_expr
    : FLOAT
    | NUMBER
    ;

string_expr
    : STRING
    | SINGLE_QUOTE_STRING
    ;

boolean_expr
    : BOOLEAN
    | ZERO_ONE
    ;

duration_expr
    : DURATION
    | number_expr
    ;

matrix_expr
    : list_expr (COLONCOLON list_expr)+
    ;

list_expr
    : LBRACK (expression (NEWLINE* COMMA NEWLINE* expression)*)? RBRACK
    ;

loop_variable
    : LOOP_KEY
    | LOOP_VALUE
    ;

error_variable
    : ERROR
    ;