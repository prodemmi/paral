#!/bin/bash
java -jar ~/antlr4/antlr-4.13.1-complete.jar -Dlanguage=Go -o antlr antlr/ParalLexer.g4
java -jar ~/antlr4/antlr-4.13.1-complete.jar -Dlanguage=Go -lib antlr/antlr -o antlr antlr/ParalParser.g4
go build -o paral cmd/main.go
./paral examples/simple.paral