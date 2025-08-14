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

// Fixed URL pattern - more robust and handles edge cases
URL: ('http' 's'? '://' | 'localhost:' [0-9]+) [a-zA-Z0-9._~:/?#[\]@!$&'()*+,;=-]+;

// Fixed PATH pattern - handles Windows/Unix paths properly
PATH: ([a-zA-Z] ':')? [/\\] ([a-zA-Z0-9._-]+[/\\])*[a-zA-Z0-9._-]+
    | ([a-zA-Z] ':')? [/\\] [a-zA-Z0-9._-]+
    | '.' [/\\] ([a-zA-Z0-9._-]+[/\\])*[a-zA-Z0-9._-]+
    | '..' [/\\] ([a-zA-Z0-9._-]+[/\\])*[a-zA-Z0-9._-]+;

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

ERROR: AT 'error' ;
LOOP_KEY: AT 'key' ;
LOOP_VALUE: AT 'value' ;
TRY: AT 'try' ;
CATCH: AT 'catch' ;

UNDERSCORE: '_' ;

// IMPORTANT: IDENTIFIER must come after TASK to avoid conflicts
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;

NEWLINE: ('\r'? '\n');
WS: [ \t]+ -> skip;

// Add a catch-all for debugging
UNKNOWN_DEFAULT: . -> channel(HIDDEN);

// ---------------------- PIPELINE Mode ----------------------
mode PIPELINE;

PIPELINE_NEWLINE: ('\r'? '\n') -> type(NEWLINE), popMode;  // Auto-exit pipeline mode on newline
PIPELINE_BUF: AT 'buf[' -> pushMode(BUF_MODE);
PIPELINE_STASH: AT 'stash[' -> pushMode(STASH_MODE);
PIPELINE_IF_CALL_START: AT 'if' LRBRACK -> pushMode(EXPRESSION);
PIPELINE_ELSEIF_CALL_START: AT 'elseif' LRBRACK -> pushMode(EXPRESSION);
PIPELINE_MATCH_CALL: AT 'match' LRBRACK -> pushMode(EXPRESSION);
PIPELINE_ELSE_CALL: AT 'else';
PIPELINE_FUNCTION_CALL_START: AT IDENTIFIER LRBRACK -> pushMode(FUNCTION);

PIPELINE_ERROR: AT 'error' -> type(ERROR);
PIPELINE_LOOP_KEY: AT 'key' -> type(LOOP_KEY);
PIPELINE_LOOP_VALUE: AT 'value' -> type(LOOP_VALUE);
PIPELINE_TRY_CALL: AT 'try' -> type(TRY);
PIPELINE_CATCH_CALL: AT 'catch' -> type(CATCH);

PIPELINE_LRBRACK: '(' -> type(LRBRACK);
PIPELINE_RRBRACK: ')' -> type(RRBRACK);

PIPELINE_BLOCK_START: '{' -> type(BLOCK_START);
PIPELINE_BLOCK_END: '}' -> type(BLOCK_END), popMode; // Exit pipeline mode when block ends
PIPELINE_LBRACK: '[' -> type(LBRACK);
PIPELINE_RBRACK: ']' -> type(RBRACK);
NESTED_PIPELINE_START: '->' -> type(PIPELINE_START);

// Fixed PIPELINE URL - same improvements as default mode
PIPELINE_URL: ('http' 's'? '://' | 'localhost:' [0-9]+) [a-zA-Z0-9._~:/?#[\]@!$&'()*+,;=-]+ -> type(URL);

// Fixed PIPELINE PATH - same improvements as default mode  
PIPELINE_PATH: (([a-zA-Z] ':')? [/\\] ([a-zA-Z0-9._-]+[/\\])*[a-zA-Z0-9._-]+
             | ([a-zA-Z] ':')? [/\\] [a-zA-Z0-9._-]+
             | '.' [/\\] ([a-zA-Z0-9._-]+[/\\])*[a-zA-Z0-9._-]+
             | '..' [/\\] ([a-zA-Z0-9._-]+[/\\])*[a-zA-Z0-9._-]+)
             -> type(PATH);

PIPELINE_STRING: '"' ( '\\' . | ~["\\\r\n] )* '"' -> type(STRING);
PIPELINE_SINGLE_QUOTE_STRING: '\'' ( '\\' . | ~['\\\r\n] )* '\'' -> type(SINGLE_QUOTE_STRING);
PIPELINE_FLOAT: [0-9]+ '.' [0-9]* -> type(FLOAT);
PIPELINE_NUMBER: [0-9]+ -> type(NUMBER);
PIPELINE_BOOLEAN: ('true' | 'false') -> type(BOOLEAN);
PIPELINE_ZERO_ONE: [01] -> type(ZERO_ONE);
PIPELINE_IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]* -> type(IDENTIFIER);
PIPELINE_COMMA: ',' -> type(COMMA);

PIPELINE_WS: [ \t]+ -> skip;

UNKNOWN_TEXT
    : ~[ \r\n{}@\-]+ [ \t]* // Capture trailing whitespace
    ;

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

// Add function call support in EXPRESSION mode
EXPRESSION_FUNCTION_CALL_START: AT IDENTIFIER '(' -> pushMode(FUNCTION);

EXPRESSION_ERROR: AT 'error' -> type(ERROR);
EXPRESSION_LOOP_KEY: AT 'key' -> type(LOOP_KEY);
EXPRESSION_LOOP_VALUE: AT 'value' -> type(LOOP_VALUE);

EXPRESSION_LRBRACK: '(' -> type(LRBRACK);
EXPRESSION_RRBRACK: ')' -> type(RRBRACK), popMode;

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

FUNCTION_ERROR: AT 'error' -> type(ERROR);
FUNCTION_LOOP_KEY: AT 'key' -> type(LOOP_KEY);
FUNCTION_LOOP_VALUE: AT 'value' -> type(LOOP_VALUE);

FUNCTION_LRBRACK: '(' -> type(LRBRACK);
FUNCTION_RRBRACK: ')' -> type(RRBRACK), popMode;

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