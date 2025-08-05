package parser

import (
	"paral/antlr/antlr"
	"paral/core/runtime"
)

func (p *Parser) parsePipeline(task *runtime.Task, ctx parser.IPipeline_blockContext) *runtime.TaskPipeline {
	return p.parsePipelineContent(task, ctx.Pipeline_content())
}

func (p *Parser) parsePipelineContent(task *runtime.Task, ctx parser.IPipeline_contentContext) *runtime.TaskPipeline {
	for blockIndex, item := range ctx.AllPipeline_item() {
		if bufCtx := item.Buf(); bufCtx != nil {
			buf := p.parseBuf(task, bufCtx)
			task.AddTaskPipeline(&runtime.TaskPipeline{
				Buf:   buf,
				Block: blockIndex,
			})
			continue
		}
		if stashCtx := item.Stash(); stashCtx != nil {
			stash := p.parseStash(task, stashCtx)
			task.AddTaskPipeline(&runtime.TaskPipeline{
				Stash: stash,
				Block: blockIndex,
			})
			continue
		}
		if fnCtx := item.Function(); fnCtx != nil {
			task.AddTaskPipeline(&runtime.TaskPipeline{
				Function: p.parseFunction(task, fnCtx),
				Block:    blockIndex,
			})
			continue
		}
		if raw := item.COMMAND_RAW_TEXT(); raw != nil {
			pipeline := &runtime.Command{
				RawText: raw.GetText(),
			}
			task.AddTaskPipeline(&runtime.TaskPipeline{
				Command: pipeline,
				Block:   blockIndex,
			})
			continue
		}
		if cond := item.Condition(); cond != nil {
			panic("condition in pipeline not supported yet")
		}
	}
	return nil
}
