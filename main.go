package main

import (
	"fmt"
	"os"
	parser "paral/antlr/antlr"
	"paral/core"

	"github.com/antlr4-go/antlr/v4"
)

type errorListener struct {
	*antlr.DefaultErrorListener
	Filename  string
	hasErrors bool
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.hasErrors = true
	fmt.Printf("Syntax error at line %s:%d:%d:\n\t%s (offending symbol: %v)\n", l.Filename, line, column, msg, offendingSymbol)
	core.ThrowSyntaxError(msg, l.Filename, line, column)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./paral_interpreter file.paral")
		return
	}

	inputFile := os.Args[1]
	input, err := antlr.NewFileStream(inputFile)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}

	// Enable debug output
	lexer := parser.NewParalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Reset stream for parsing
	stream.Seek(0)

	p := parser.NewParalParser(stream)

	// Create and add error listener
	errListener := &errorListener{Filename: inputFile}
	p.AddErrorListener(errListener)

	// Create listener with debug output
	c := core.NewCore(inputFile)
	listener := &core.CoreVisitor{Core: c}
	p.AddParseListener(listener)

	// Enable parse tree building
	p.BuildParseTrees = true

	// Parse and get the tree
	p.Program()

	// Print the parse tree
	// Check for parsing errors
	if errListener.hasErrors {
		fmt.Println("Parsing failed with syntax errors")
		return
	}

	c.PrintValues()

	fmt.Println("Program executed successfully")
}
