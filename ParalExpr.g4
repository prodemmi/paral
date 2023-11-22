grammar ParalExpr;
prog:   (expr NEWLINE)* ;
expr:   variable+
    |   execute?
    ;

variable: 'var' IDENT VALUE | ' ';
execute: 'exec' IDENT VALUE | ' ';

IDENT : [a-zA-Z_]+ ;
VALUE : [a-zA-Z0-9._]+ ;
NEWLINE : [\r\n]+ ;
