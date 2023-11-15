package main

import (
	"os"

	"github.com/alecthomas/kong"

	"github.com/alecthomas/repr"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Paral struct {
	Pos       lexer.Position
	Lines     []*Line `@@*`
}

type Line struct {
	Variable *Variable `@@`
	Execute  *Execute  `@@`
}

type Variable struct {
	Type    *string   `@Var`
	Name    *string   `@Ident`
	Value   *string   `@Ident`
	Options []*Option `@@*`
}

type Option struct {
	Name  *string `@Option`
	Value *string `@Ident?`
}

type Execute struct {
	Type  *string `@Exec`
	Value *string `@Ident`
}

var (
	paralLexer = lexer.MustSimple([]lexer.SimpleRule{
		{`Option`, `--([a-zA-Z_][a-zA-Z_0-9]*)` },
		{`Var`, `(?i)\b(var)\b`},
		{`Exec`, `(?i)\b(exec)\b`},
		{"Ident", `[a-zA-Z_][a-zA-Z_0-9]*`},
		{"whitespace", `\s+`},
	})

	paralParser = participle.MustBuild[Paral](
		participle.Lexer(paralLexer),
		participle.CaseInsensitive("Keyword"),
	)

	cli struct {
		File string `help:"Paral file to parse." arg:""`
	}
)

func main() {
	ctx := kong.Parse(&cli)
	r, err := os.Open(cli.File)
	ctx.FatalIfErrorf(err)
	defer r.Close()
	paral, err := paralParser.Parse(cli.File, r)
	ctx.FatalIfErrorf(err)
	repr.Println(paral)
}