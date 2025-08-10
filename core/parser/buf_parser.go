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

	pipeline := p.parsePipelineContent(task, ctx.Pipeline_content())

	buf := runtime.NewBuf(name, task.GetTaskId(), ctx.GetText(), bufMetadata, pipeline.Command)
	return buf
}
