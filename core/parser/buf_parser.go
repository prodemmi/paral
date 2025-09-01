package parser

import (
	parser "paral/antlr/antlr"
	"paral/core"
	"paral/core/metadata"
	"paral/core/runtime"
)

func (p *Parser) parseBuf(task *runtime.Task, ctx parser.IBufContext, parent *runtime.TaskPipeline) *runtime.Buf {
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

	// Create a buf pipeline that has the current pipeline as parent
	bufPipeline := &runtime.TaskPipeline{
		Parent: parent,
	}

	// Parse the pipeline content with the buf pipeline as parent
	contentPipeline := p.parsePipelineContent(task, ctx.Pipeline_content(), bufPipeline)

	buf := runtime.NewBuf(name, task.GetTaskId(), ctx.GetText(), bufMetadata, contentPipeline)
	return buf
}
