package parser

import (
	parser "paral/antlr/antlr"
	"paral/core/metadata"
	"paral/core/runtime"
)

func (p *Parser) parseTryCatch(task *runtime.Task, ctx parser.ITry_catchContext, parent *runtime.TaskPipeline) *runtime.TryCatch {
	tc := &runtime.TryCatch{
		TryPipelines:   make([]*runtime.TaskPipeline, 0),
		CatchPipelines: make([]*runtime.TaskPipeline, 0),
		Metadata: &metadata.Metadata{
			Content: ctx.GetText(),
			Line:    ctx.GetStop().GetLine(),
			Column:  ctx.GetStop().GetColumn(),
		},
	}

	// parse try block - pass the try-catch pipeline as parent
	for _, tryBlock := range ctx.Try_block().AllPipeline_block() {
		pipeline := p.parsePipeline(task, tryBlock, parent)
		tc.TryPipelines = append(tc.TryPipelines, pipeline)
	}

	// parse optional catch block - pass the try-catch pipeline as parent
	if catchCtxs := ctx.Catch_block(); catchCtxs != nil {
		for _, catchBlock := range catchCtxs.AllPipeline_block() {
			pipeline := p.parsePipeline(task, catchBlock, parent)
			// Fixed: should append to CatchPipelines, not TryPipelines
			tc.CatchPipelines = append(tc.CatchPipelines, pipeline)
		}
	}

	return tc
}
