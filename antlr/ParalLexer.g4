lexer grammar ParalLexer;

// ---------------------- Default Mode ----------------------

TASK: 'task';
ASSIGN: '=';

STRING: '"' ( '\\' . | ~["\\\r\n] )* '"';
SINGLE_QUOTE_STRING: '\'' ( '\\' . | ~['\\\r\n] )* '\'';
FLOAT: [0-9]+ '.' [0-9]+;
NUMBER: [0-9]+;
BOOLEAN: 'true' | 'false';
ZERO_ONE: [01];
DURATION: NUMBER [smhd];

URL: ('http' 's'? '://' | 'localhost:') URL_BODY;
fragment URL_BODY: (URL_CHAR | '.' | '/' | ':' | '?' | '=' | '&' | '%' | '-' | '_')+;
fragment URL_CHAR: [a-zA-Z0-9];

COLON: ':';
COLONCOLON: '::';
COMMA: ',';

LBRACK: '[';
RBRACK: ']';
LRBRACK: '(';
RRBRACK: ')';

BLOCK_START: '{' ;
BLOCK_END: '}' ;

AT: '@';
FUNCTION_START: AT IDENTIFIER '(' -> pushMode(FUNCTION);

PIPELINE_START: '->' -> pushMode(PIPELINE) ;

LOOP_KEY: AT 'key' ;
LOOP_VALUE: AT 'value' ;

IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;

NEWLINE: ('\r'? '\n');
WS: [ \t]+ -> skip;

// ---------------------- PIPELINE Mode ----------------------
mode PIPELINE;
PIPELINE_NEWLINE: ('\r'? '\n') -> popMode;
PIPELINE_BUF: AT 'buf[' -> pushMode(BUF_MODE);
PIPELINE_STASH: AT 'stash[' -> pushMode(STASH_MODE);
PIPELINE_IF_CALL_START: AT 'if(' -> pushMode(EXPRESSION) ;
PIPELINE_FUNCTION_CALL_START: AT IDENTIFIER '(' -> pushMode(FUNCTION);

PIPELINE_LOOP_KEY: AT 'key' -> type(LOOP_KEY);
PIPELINE_LOOP_VALUE: AT 'value' -> type(LOOP_VALUE);
PIPELINE_BLOCK_START: '{' -> type(BLOCK_START), popMode ;
PIPELINE_BLOCK_END: '}' -> type(BLOCK_END), popMode ; // Add this line
PIPELINE_LBRACK: '[' -> type(LBRACK);
PIPELINE_RBRACK: ']' -> type(RBRACK);
NESTED_PIPELINE_START: '->' -> type(PIPELINE_START) ;
PIPELINE_STRING: '"' ( '\\' . | ~["\\\r\n] )* '"' -> type(STRING);
PIPELINE_SINGLE_QUOTE_STRING: '\'' ( '\\' . | ~['\\\r\n] )* '\'' -> type(SINGLE_QUOTE_STRING);
PIPELINE_FLOAT: [0-9]+ '.' [0-9]* -> type(FLOAT);
PIPELINE_NUMBER: [0-9]+ -> type(NUMBER);
PIPELINE_BOOLEAN: ('true' | 'false') -> type(BOOLEAN);
PIPELINE_ZERO_ONE: [01] -> type(ZERO_ONE);
PIPELINE_IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]* -> type(IDENTIFIER);
PIPELINE_COMMA: ',' -> type(COMMA);

PIPELINE_WS: [ \t]+ -> skip;

// ---------------------- BUF Mode ----------------------
mode BUF_MODE;
BUF_LBRACK: '[' -> type(LBRACK);
BUF_RBRACK: ']' -> type(RBRACK);
BUF_DOUBLE_BACK_ARROW: '<<' -> popMode;
BUF_STRING: '"' ( '\\' . | ~["\\\r\n] )* '"' -> type(STRING);
BUF_SINGLE_QUOTE_STRING: '\'' ( '\\' . | ~['\\\r\n] )* '\'' -> type(SINGLE_QUOTE_STRING);
BUF_WS: [ \t]+ -> skip;

// ---------------------- STASH Mode ----------------------
mode STASH_MODE;
STASH_LBRACK: '[' -> type(LBRACK);
STASH_RBRACK: ']' -> type(RBRACK);
STASH_DOUBLE_BACK_ARROW: '<<' -> popMode;
STASH_STRING: '"' ( '\\' . | ~["\\\r\n] )* '"' -> type(STRING);
STASH_SINGLE_QUOTE_STRING: '\'' ( '\\' . | ~['\\\r\n] )* '\'' -> type(SINGLE_QUOTE_STRING);
STASH_WS: [ \t]+ -> skip;

// ---------------------- EXPRESSION Mode ----------------------
mode EXPRESSION;

IF_CONDITION_END: ')' -> popMode;

// Add function call support in EXPRESSION mode
EXPRESSION_FUNCTION_CALL_START: AT IDENTIFIER '(' -> pushMode(FUNCTION);

EXPRESSION_LOOP_KEY: AT 'key' -> type(LOOP_KEY);
EXPRESSION_LOOP_VALUE: AT 'value' -> type(LOOP_VALUE);
EXPRESSION_LBRACK: '[' -> type(LBRACK);
EXPRESSION_RBRACK: ']' -> type(RBRACK);
EXPRESSION_STRING: '"' ( '\\' . | ~["\\\r\n] )* '"' -> type(STRING);
EXPRESSION_SINGLE_QUOTE_STRING: '\'' ( '\\' . | ~['\\\r\n] )* '\'' -> type(SINGLE_QUOTE_STRING);
EXPRESSION_FLOAT: [0-9]+ '.' [0-9]* -> type(FLOAT);
EXPRESSION_NUMBER: [0-9]+ -> type(NUMBER);
EXPRESSION_BOOLEAN: ('true' | 'false') -> type(BOOLEAN);
EXPRESSION_ZERO_ONE: [01] -> type(ZERO_ONE);
EXPRESSION_IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]* -> type(IDENTIFIER);
EXPRESSION_COMMA: ',' -> type(COMMA);

EXPRESSION_WS: [ \t]+ -> skip;
EXPRESSION_NEWLINE: ('\r'? '\n') -> type(NEWLINE);

// ---------------------- FUNCTION Mode ----------------------
mode FUNCTION;
NESTED_FUNCTION_START: AT IDENTIFIER '(' -> pushMode(FUNCTION);
FUNCTION_END: ')' -> popMode;

FUNCTION_LOOP_KEY: AT 'key' -> type(LOOP_KEY);
FUNCTION_LOOP_VALUE: AT 'value' -> type(LOOP_VALUE);
FUNCTION_LBRACK: '[' -> type(LBRACK);
FUNCTION_RBRACK: ']' -> type(RBRACK);
FUNCTION_STRING: '"' ( '\\' . | ~["\\\r\n] )* '"' -> type(STRING);
FUNCTION_SINGLE_QUOTE_STRING: '\'' ( '\\' . | ~['\\\r\n] )* '\'' -> type(SINGLE_QUOTE_STRING);
FUNCTION_FLOAT: [0-9]+ '.' [0-9]* -> type(FLOAT);
FUNCTION_NUMBER: [0-9]+ -> type(NUMBER);
FUNCTION_BOOLEAN: ('true' | 'false') -> type(BOOLEAN);
FUNCTION_ZERO_ONE: [01] -> type(ZERO_ONE);
FUNCTION_IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]* -> type(IDENTIFIER);
FUNCTION_COMMA: ',' -> type(COMMA);

FUNCTION_WS: [ \t]+ -> skip;
FUNCTION_NEWLINE: ('\r'? '\n') -> type(NEWLINE);