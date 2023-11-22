grammar ParalExpr;

fragment SPACE: '\t' | ' ' | '\r' | '\n' | '\u000C';
fragment VAR: 'var' | 'VAR';
fragment EXEC: 'exec' | 'EXEC';

start: (variables NEWLINE)* | executeables;
variables: variable+;
executeables: execute* | ;

variable: VARIABLE IDENT VALUE;
execute: EXECUTE VALUE;

IDENT: [a-zA-Z_]+;
VALUE: ([A-Z] | [a-z] | [0-9])+;
NEWLINE: [\r\n]+;
WS: SPACE+;

// keywords
VARIABLE: VAR;
EXECUTE: EXEC;