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
	Entries   *Entry `@@`
}

type Entry struct {
	Variables []*Variable `@@*`
	Executes  []*Execute  `@@*`
}

type Variable struct {
	Type    *string `"var" `
	Name    *string `@Ident`
	Value   *Value  `@@`
}

type Execute struct {
	Type    *string `"exec" `
	Command *string  `@Command`
}

type Value interface{ value() }

type String struct {
	String *string `@String`
}

func (String) value() {}

type Number struct {
	Number *float64 `@Float`
}

func (Number) value() {}

var (
	paralLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"DateTime", `\d\d\d\d-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)?(-\d\d:\d\d)?`},
		{"Date", `\d\d\d\d-\d\d-\d\d`},
		{"Time", `\d\d:\d\d:\d\d(\.\d+)?`},
		{"Ident", `[a-zA-Z_][a-zA-Z_0-9]*`},
		{"Command", `\s+(.*)`},
		{"String", `"[^"]*"`},
		{`Float`, `\d+(?:\.\d+)?`},
		{"comment", `#[^\n]+`},
		{"whitespace", `\s+`},
	})

	paralParser = participle.MustBuild[Paral](
		participle.Lexer(paralLexer),
		participle.Unquote("String"),
		participle.Union[Value](String{}, Number{}),
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