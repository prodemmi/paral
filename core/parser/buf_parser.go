package parser

import (
	"paral/antlr/antlr"
	"paral/core"
	"paral/core/metadata"
	"paral/core/runtime"
)

func (p *Parser) parseBuf(task *runtime.Task, ctx parser.IBufContext) *runtime.Buf {
	strExpr := ctx.String_expr()
	var name string

	if strExpr.STRING() != nil || strExpr.SINGLE_QUOTE_STRING() != nil {
		name = core.TrimQuotes(strExpr.STRING().GetText())
	} else {
		name = strExpr.GetText()
	}

	bufMetadata := metadata.Metadata{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}

	content := ctx.Pipeline_content()
	var cmd *runtime.Command

	for _, item := range content.AllPipeline_item() {
		if fn := item.Function(); fn != nil {
			cmd = &runtime.Command{
				Functions: []*runtime.Function{p.parseFunction(task, fn)},
				RawText:   fn.GetText(),
			}
			break
		} else if raw := item.COMMAND_RAW_TEXT(); raw != nil {
			cmd = &runtime.Command{
				Functions: nil,
				RawText:   raw.GetText(),
			}
			break
		}
	}

	buf := runtime.NewBuf(name, task.GetTaskId(), ctx.GetText(), bufMetadata, cmd)
	return buf
}
