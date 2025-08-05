lexer grammar ParalLexer;

// ---------------------- Default Mode ----------------------

CMD_ARROW: '->' -> pushMode(PIPELINE);

AT : '@';
FUNCTION_START: '@' IDENTIFIER '(' -> pushMode(FUNCTION);
IF : 'if';
ENDIF : 'endif';
TASK : 'task';

ASSIGN : '=';

STRING: '"' ( '\\' . | ~["\\\r\n] )* '"';
SINGLE_QUOTE_STRING: '\'' ( '\\' . | ~['\\\r\n] )* '\'';

FLOAT: [0-9]+ '.' [0-9]+;
NUMBER: [0-9]+;
BOOLEAN: 'true' | 'false';
ZERO_ONE: [01];
DURATION: NUMBER [smhd];

URL : ('http' 's'? '://' | 'localhost:') URL_BODY;
fragment URL_BODY: (URL_CHAR | '.' | '/' | ':' | '?' | '=' | '&' | '%' | '-' | '_')+;
fragment URL_CHAR: [a-zA-Z0-9];

COLON : ':';
COLONCOLON : '::';
COMMA : ',';

LBRACK : '[';
RBRACK : ']';

LRBRACK: '(';
RRBRACK: ')';

LBRACE: '{';
RBRACE: '}';

DOUBLE_BACK_ARROW: '<<';

IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;

NEWLINE: ('\r'? '\n');
WS: [ \t]+ -> skip;

// ---------------------- PIPELINE Mode ----------------------

mode PIPELINE;

PIPELINE_NEWLINE: ('\r'? '\n') -> popMode;
PIPELINE_STASH: '@stash[' -> pushMode(STASH_MODE);
PIPELINE_DOUBLE_BACK_ARROW: '<<' -> type(DOUBLE_BACK_ARROW);
PIPELINE_IF: '@if(' -> mode(IF_MODE);
PIPELINE_LOOP_KEY: '@key';
PIPELINE_LOOP_VALUE: '@value';
FUNCTION_CALL_START: '@' IDENTIFIER '(' -> pushMode(FUNCTION);

PIPELINE_WS: [ \t]+ -> skip;
// Allow quoted strings within command raw text
COMMAND_RAW_TEXT: (~[@\r\n\t] | '"' ( '\\' . | ~["\\\r\n] )* '"' | '\'' ( '\\' . | ~['\\\r\n] )* '\'')+;

// ---------------------- STASH Mode ----------------------

mode STASH_MODE;

STASH_LBRACK : '[' -> type(LBRACK);
STASH_RBRACK : ']' -> type(RBRACK);
STASH_DOUBLE_BACK_ARROW: '<<' -> type(DOUBLE_BACK_ARROW), popMode;
STASH_STRING: '"' ( '\\' . | ~["\\\r\n] )* '"' -> type(STRING);
STASH_SINGLE_QUOTE_STRING: '\'' ( '\\' . | ~['\\\r\n] )* '\'' -> type(SINGLE_QUOTE_STRING);
STASH_WS: [ \t]+ -> skip;

// ---------------------- FUNCTION Mode ----------------------

mode FUNCTION;

NESTED_FUNCTION_START: '@' IDENTIFIER '(' -> pushMode(FUNCTION);
FUNCTION_END: ')' -> popMode;

// Add loop variable to FUNCTION mode
FUNCTION_LOOP_KEY: '@key' -> type(PIPELINE_LOOP_KEY);
FUNCTION_LOOP_VALUE: '@value' -> type(PIPELINE_LOOP_VALUE);

FUNCTION_LBRACK : '[' -> type(LBRACK);
FUNCTION_RBRACK : ']' -> type(RBRACK);

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

// ---------------------- IF Mode ----------------------

mode IF_MODE;

// nested functions
NESTED_IF: '@if(' -> mode(IF_MODE);
IF_END: '@endif' -> popMode;

IF_WS: [ \t]+ -> skip;