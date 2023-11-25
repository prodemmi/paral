package main

import (
	"fmt"
	"log"
	"os"
	"paral/parser"

	"github.com/antlr4-go/antlr/v4"
)

type paralListener struct {
	*parser.BaseParalListener
	parser *parser.ParalParser
}

func (s *paralListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Println("VisitTerminal", node.GetText())
}

func (s *paralListener) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Println("VisitErrorNode", node.GetText())
}

func (s *paralListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("EnterEveryRule", ctx.GetText())
}

func (s *paralListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("ExitEveryRule", ctx.GetText())
}

func (s *paralListener) EnterStart(ctx *parser.StartContext) {
	fmt.Println("EnterStart", ctx.GetText())
}

func (s *paralListener) ExitStart(ctx *parser.StartContext) {
	fmt.Println("ExitStart", ctx.GetText())
}

func (s *paralListener) EnterProg(ctx *parser.ProgContext) {
	fmt.Println("EnterProg", ctx.GetText())
}

func (s *paralListener) ExitProg(ctx *parser.ProgContext) {
	fmt.Println("ExitProg", ctx.GetText())
}

func (s *paralListener) EnterVariable(ctx *parser.VariableContext) {
	fmt.Println("EnterVariable", ctx.GetText())
}

func (s *paralListener) ExitVariable(ctx *parser.VariableContext) {
	fmt.Println("ExitVariable", ctx.GetText())
}

func (s *paralListener) EnterExecute(ctx *parser.ExecuteContext) {
	fmt.Println("EnterExecute", ctx.GetText())
}

func (s *paralListener) ExitExecute(ctx *parser.ExecuteContext) {
	fmt.Println("ExitExecute", ctx.GetText())
}

func main() {
	file, err := os.ReadFile("examples/example.paral")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	input := antlr.NewInputStream(string(file))
	lexer := parser.NewParalLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewParalParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&paralListener{}, p.Start_())
}
