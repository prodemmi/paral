#!/bin/bash
rm -rf antlr/antlr 
java -jar ~/antlr4/antlr-4.13.2-complete.jar -Dlanguage=Go -o antlr antlr/ParalLexer.g4 
java -jar ~/antlr4/antlr-4.13.2-complete.jar -Dlanguage=Go -lib antlr/antlr -o antlr antlr/ParalParser.g4
go build -o paral main.go 
./paral --graph examples/simple.paral