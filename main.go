package main

import (
	"fmt"
	"log"
	"os"
	"paral/parser"

	"github.com/antlr4-go/antlr/v4"
)

type paralExpr struct {
	*parser.BaseParalExprListener
}

func main() {
	file, err := os.ReadFile("examples/example.paral")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	input := antlr.NewInputStream(string(file))
	lexer := parser.NewParalExprLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewParalExprParser(stream)

	newparalExpr := paralExpr{}

	antlr.ParseTreeWalkerDefault.Walk(&newparalExpr, p.Expr())

	fmt.Println(newparalExpr)
}
