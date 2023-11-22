#! /bin/bash

rm -rf parser/ .antlr/
antlr4 -Dlanguage=Go -o parser ParalExpr.g4