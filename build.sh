#! /bin/bash

rm -rf parser/ .antlr/
antlr4 -Dlanguage=Go -o parser Paral.g4