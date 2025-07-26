parser grammar ParalParser;

options { tokenVocab=ParalLexer; }

program
    : statement* EOF
    ;

statement
    : variable_assignment
    | job_definition
    | NEWLINE
    ;

variable_assignment
    : IDENTIFIER ASSIGN variable_value NEWLINE
    ;

job_definition
    : (job_directive NEWLINE)* IDENTIFIER COLON NEWLINE command_block+
    ;

job_directive
    : directive
    ;

command_block
    : CMD_ARROW command_content COMMAND_NEWLINE
    ;

command_content
    : command_item*
    ;

command_item
    : function
    | condition
    | loop_variable
    | COMMAND_RAW_TEXT
    ;

// ----------------- job directive -------------

directive
    : AT IDENTIFIER expression*
    ;

// ----------------- functions and conditions -------------

condition
    : if_condition
    ;

function
    : FUNCTION_CALL_START argument_list? FUNCTION_END
    ;

nested_function
    : NESTED_FUNCTION_START argument_list? FUNCTION_END
    ;


if_condition
    : AT IF LRBRACK expression RRBRACK NEWLINE* command_content NEWLINE* AT ENDIF
    ;

// Argument list for functions
argument_list
    : expression (COMMA expression)*
    ;

// Expression can be a value or a nested function
expression
    : nested_function
    | variable_value
    | loop_variable
    ;

variable_value
    : URL
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
    : LBRACK (variable_value (COMMA variable_value)*)? RBRACK
    ;

loop_variable
    : COMMAND_LOOP_KEY
    | COMMAND_LOOP_VALUE
    ;