package parser

import (
	parser "paral/antlr/antlr"
	"paral/core/functions"
	"paral/core/runtime"
	"strings"
)

func (p *Parser) parseCommand(ctx parser.ICommand_blockContext) *runtime.Command {
	rawText, _ := strings.CutPrefix(ctx.GetText(), "->")
	cmd := runtime.Command{
		RawText:   rawText,
		Functions: []*functions.Function{},
		Block:     ctx.GetStart().GetLine(),
	}
	if commandContent := ctx.Command_content(); commandContent != nil {
		for _, commandItem := range commandContent.AllCommand_item() {
			if fnCtx := commandItem.Function(); fnCtx != nil {
				function := p.parseFunction(fnCtx)
				cmd.Functions = append(cmd.Functions, &function)
			}
		}
	}
	return &cmd
}
