#!/bin/bash

clear && \
rm -rf antlr/antlr && \
java -jar ~/antlr4/antlr-4.13.2-complete.jar -Dlanguage=Go -o antlr antlr/Paral.g4 && \
go build -o paral main.go && \
./paral examples/simple.paral
