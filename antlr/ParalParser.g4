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
    : (task_directive NEWLINE)* TASK IDENTIFIER LBRACE NEWLINE* pipeline_block* RBRACE
    ;

task_directive
    : directive
    ;

pipeline_block
    : CMD_ARROW pipeline_content PIPELINE_NEWLINE
    ;

pipeline_content
    : pipeline_item*
    ;

pipeline_item
    : buf
    | stash
    | condition
    | function
    | COMMAND_RAW_TEXT
    ;

// ----------------- task directive -------------

directive
    : AT IDENTIFIER expression*
    ;

// ----------------- functions and conditions -------------

buf
    : PIPELINE_BUF string_expr RBRACK DOUBLE_BACK_ARROW pipeline_content
    ;

stash
    : PIPELINE_STASH string_expr RBRACK DOUBLE_BACK_ARROW pipeline_content
    ;

condition
    : if_condition
    ;

function
    : FUNCTION_CALL_START argument_list? FUNCTION_END
    | FUNCTION_START argument_list? FUNCTION_END  // Add this for default mode functions
    ;

nested_function
    : NESTED_FUNCTION_START argument_list? FUNCTION_END
    ;

if_condition
    : AT IF LRBRACK expression RRBRACK NEWLINE* pipeline_content NEWLINE* AT ENDIF
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
    : PIPELINE_LOOP_KEY
    | PIPELINE_LOOP_VALUE
    ;