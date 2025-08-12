package parser

import (
	parser "paral/antlr/antlr"
	"paral/core/runtime"
)

func (p *Parser) parsePipeline(task *runtime.Task, ctx parser.IPipeline_blockContext) *runtime.TaskPipeline {
	return p.parsePipelineContent(task, ctx.Pipeline_content())
}

func (p *Parser) parsePipelineContent(task *runtime.Task, ctx parser.IPipeline_contentContext) *runtime.TaskPipeline {
	var pipeline *runtime.TaskPipeline
	for blockIndex, item := range ctx.AllPipeline_item() {
		if condCtx := item.Condition(); condCtx != nil {
			if ifCtx := condCtx.If_condition(); ifCtx != nil {
				ifCondition := p.parseIfCondition(task, ifCtx)
				return &runtime.TaskPipeline{
					Condition: &runtime.Condition{
						IfCondition: ifCondition,
						RawText:     ifCtx.GetText(),
					},
					Block: blockIndex,
				}
			}
		}
		if bufCtx := item.Buf(); bufCtx != nil {
			buf := p.parseBuf(task, bufCtx)
			return &runtime.TaskPipeline{
				Buf:   buf,
				Block: blockIndex,
			}
		}
		if stashCtx := item.Stash(); stashCtx != nil {
			stash := p.parseStash(task, stashCtx)
			return &runtime.TaskPipeline{
				Stash: stash,
				Block: blockIndex,
			}
		}
		if fnCtx := item.Function(); fnCtx != nil {
			function := p.parseFunction(task, fnCtx)
			return &runtime.TaskPipeline{
				Function: function,
				Block:    blockIndex,
			}
		}
	}
	return pipeline
}
