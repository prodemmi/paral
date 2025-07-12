grammar Paral;

program
    : line* EOF
    ;

line
    : variable_def 
    | matrix_def 
    | job_def
    | NEWLINE
    ;

variable_def
    : IDENTIFIER '=' (value_expr | list_expr) NEWLINE
    ;

matrix_def
    : IDENTIFIER '=' list_expr (COLONCOLON list_expr)* NEWLINE
    ;

job_def
    : IDENTIFIER ':' NEWLINE cmd_expr+
    ;

cmd_expr
    : ARROW cmd_value+ NEWLINE
    ;

cmd_value
    : value_expr                       
    | directive_expr
    | cmd_value '[' cmd_value ']'               
    | '(' cmd_value ')'                    
    | flag_expr
    ;

directive_expr
    : DIRECTIVE directive_args_expr*
    ;

directive_args_expr
    :  LRBRACK value_expr? (COMMA value_expr)* RRBRACK  
    ;

flag_expr
    : FLAG_NAME '=' value_expr  # flagWithValue
    | FLAG_NAME                 # flagAlone
    ;

value_expr
    : IDENTIFIER          # identifier         
    | STRING              # string
    | SINGLE_QUOTE_STRING # singleQuoteString
    | NUMBER              # number
    | BOOLEAN             # bool
    | DURATION            # duration
    ;

list_expr
    : LBRACK (value_expr (COMMA value_expr)*)? RBRACK
    ;

// === Lexer rules ===

ARROW: '>' ;

STRING
    : '"' ( '\\' . | ~["\\\r\n] )* '"'
    ;

SINGLE_QUOTE_STRING
    : '\'' ( '\\' . | ~['\\\r\n] )* '\''
    ;

NUMBER
    : [0-9]+
    ;

BOOLEAN
    : 'true'
    | 'false'
    ;

DURATION
    : NUMBER [smh]
    ;

IDENTIFIER
    : [a-zA-Z_][a-zA-Z0-9_]*
    ;

FLAG_NAME
    : '-' '-'? [a-zA-Z_][a-zA-Z0-9_-]*
    ;

DIRECTIVE
    : '@' [a-zA-Z_][a-zA-Z0-9_-]*
    ;

LRBRACK
    : '('
    ;

RRBRACK
    : ')'
    ;

LBRACK
    : '['
    ;

RBRACK
    : ']'
    ;

COMMA
    : ','
    ;

COLONCOLON
    : '::'
    ;

NEWLINE
    : ('\r'? '\n')
    ;

WS
    : [ \t]+ -> skip
    ;
