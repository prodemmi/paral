package main

import (
	"fmt"
	"log"
	"os"
	"paral/parser"

	"github.com/antlr4-go/antlr/v4"
)

type paralExprListener struct {
	*parser.BaseParalExprListener
	parser *parser.ParalExprParser
}

func main() {
	file, err := os.ReadFile("examples/example.paral")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	input := antlr.NewInputStream(string(file))
	lexer := parser.NewParalExprLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewParalExprParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := paralExprListener{}
	listener.parser = p
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start_())

	fmt.Println(listener)
}
