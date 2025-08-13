package parser

import (
	parser "paral/antlr/antlr"
	"paral/core"
	"paral/core/metadata"
	"paral/core/runtime"
)

func (p *Parser) parseStash(task *runtime.Task, ctx parser.IStashContext, parent *runtime.TaskPipeline) *runtime.Stash {
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

	// Create a stash pipeline that has the current pipeline as parent
	stashPipeline := &runtime.TaskPipeline{
		Parent: parent,
	}

	// Parse the pipeline content with the stash pipeline as parent
	contentPipeline := p.parsePipelineContent(task, ctx.Pipeline_content(), stashPipeline)

	stash := runtime.NewStash(name, task.GetTaskId(), ctx.GetText(), stashMetadata, contentPipeline)

	return stash
}
