package main

import (
	"os"
	parser "paral/antlr/antlr"
	"paral/core"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) < 2 {
		core.ThrowSyntaxErrorSimple("No input file provided")
	}

	inputFile := os.Args[1]
	input, err := antlr.NewFileStream(inputFile)
	if err != nil {
		core.ThrowSyntaxErrorSimple("Failed to read input file: " + err.Error())
	}

	lexer := parser.NewParalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewParalParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	coreInstance := core.NewCore()
	visitor := &core.CoreVisitor{Core: coreInstance}
	antlr.ParseTreeWalkerDefault.Walk(visitor, parser.Program())

	coreInstance.PrintValues()
}
