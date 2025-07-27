lexer grammar ParalLexer;

// ---------------------- Default Mode ----------------------

CMD_ARROW: '->' -> pushMode(COMMAND);

AT : '@';
FUNCTION_START: '@' IDENTIFIER '(' -> pushMode(FUNCTION);
IF : 'if';
ENDIF : 'endif';

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

LBRACK : '[';  // Move these up
RBRACK : ']';

LRBRACK: '(';
RRBRACK: ')';

IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;  // Keep this after brackets

NEWLINE: ('\r'? '\n');
WS: [ \t]+ -> skip;

// ---------------------- COMMAND Mode ----------------------

mode COMMAND;

COMMAND_NEWLINE: ('\r'? '\n') -> popMode;
FUNCTION_CALL_START: '@' IDENTIFIER '(' -> pushMode(FUNCTION);
COMMAND_IF: '@if(' -> mode(IF_MODE);
COMMAND_LOOP_KEY: '@key';
COMMAND_LOOP_VALUE: '@value';
COMMAND_STRING: STRING -> type(STRING);

COMMAND_LBRACK : '[' -> type(LBRACK);
COMMAND_RBRACK : ']' -> type(RBRACK);

COMMAND_WS: [ \t]+ -> skip;
COMMAND_RAW_TEXT: ~[@\r\n\t(),]+;

// ---------------------- FUNCTION Mode ----------------------

mode FUNCTION;

NESTED_FUNCTION_START: '@' IDENTIFIER '(' -> pushMode(FUNCTION);
FUNCTION_END: ')' -> popMode;

// Add loop variable to FUNCTION mode
FUNCTION_LOOP_KEY: '@key' -> type(COMMAND_LOOP_KEY);
FUNCTION_LOOP_VALUE: '@value' -> type(COMMAND_LOOP_VALUE);

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