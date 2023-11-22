grammar ParalExpr;

fragment SPACE: '\t' | ' ' | '\r' | '\n' | '\u000C';
fragment VAR: 'var' | 'VAR';

prog: (expr NEWLINE)*;
expr: variable?;

variable: VARIABLE IDENT VALUE;

IDENT: [a-zA-Z_]+;
VALUE: ([A-Z] | [a-z] | [0-9])+;
NEWLINE: [\r\n]+;
WS: SPACE+;

// keywords
VARIABLE: VAR;