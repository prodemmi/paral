package main

import (
	"fmt"
	"log"
	"os"
	"paral/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	file, err := os.ReadFile("examples/example.paral")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	input := antlr.NewInputStream(string(file))

	result := parser.NewParalExprLexer(input)

	fmt.Println(result)
}
