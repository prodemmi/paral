package parser

import (
	"paral/antlr/antlr"
	"paral/core"
	"paral/core/metadata"
	"paral/core/runtime"
)

func (p *Parser) parseStash(task *runtime.Task, ctx parser.IStashContext) *runtime.Stash {
	strExpr := ctx.String_expr()
	var name string

	if strExpr.STRING() != nil || strExpr.SINGLE_QUOTE_STRING() != nil {
		name = core.TrimQuotes(strExpr.STRING().GetText())
	} else {
		name = strExpr.GetText()
	}

	stashMetadata := metadata.Metadata{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}

	pipeline := p.parsePipelineContent(task, ctx.Pipeline_content())

	stash := runtime.NewStash(name, task.GetTaskId(), ctx.GetText(), stashMetadata, pipeline.Command)
	return stash
}
