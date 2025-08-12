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
    | function
    | unknown_command
    ;

unknown_command: (IDENTIFIER | UNKNOWN_TEXT | PATH | URL)+ (expression)*;

// ----------------- task directive -------------

directive
    : AT IDENTIFIER expression*
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
    : PIPELINE_IF_CALL_START expression IF_CONDITION_END
      NEWLINE* BLOCK_START NEWLINE* pipeline_block* NEWLINE* BLOCK_END
      (NEWLINE* elseif_condition)* (NEWLINE* else_condition)?
    ;

elseif_condition
    : PIPELINE_ELSEIF_CALL_START expression IF_CONDITION_END
      NEWLINE* BLOCK_START NEWLINE* pipeline_block* NEWLINE* BLOCK_END
    ;

else_condition
    : PIPELINE_ELSE_CALL
      NEWLINE* BLOCK_START NEWLINE* pipeline_block* NEWLINE* BLOCK_END
    ;

function
    : PIPELINE_FUNCTION_CALL_START argument_list? FUNCTION_END
    | FUNCTION_START argument_list? FUNCTION_END
    | EXPRESSION_FUNCTION_CALL_START argument_list? FUNCTION_END
    ;

nested_function
    : NESTED_FUNCTION_START argument_list? FUNCTION_END
    ;

// Argument list for functions
argument_list
    : expression (NEWLINE* COMMA NEWLINE* expression)*
    ;

// Expression can be a value or a nested function
expression
    : loop_variable
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