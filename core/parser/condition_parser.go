package parser

import (
	parser "paral/antlr/antlr"
	"paral/core/runtime"
)

// parseIfCondition converts an if/elif/else AST into runtime.IfCondition.
func (p *Parser) parseIfCondition(task *runtime.Task, ctx parser.IIf_conditionContext, parent *runtime.TaskPipeline) *runtime.IfCondition {
	var branches []runtime.ConditionalBranch

	// --- IF branch ---
	mainExpr := p.parseExpression(task, ctx.Expression())
	// Pass the condition pipeline as parent to nested pipelines
	mainPipelines := p.parsePipelineBlocks(task, ctx.AllPipeline_block(), parent)
	branches = append(branches, runtime.ConditionalBranch{
		Expression: mainExpr,
		Pipelines:  mainPipelines,
		RawText:    ctx.GetText(),
	})

	// --- ELIF branches ---
	for _, elifCtx := range ctx.AllElseif_condition() {
		elifExpr := p.parseExpression(task, elifCtx.Expression())
		// Pass the condition pipeline as parent to nested pipelines
		elifPipelines := p.parsePipelineBlocks(task, elifCtx.AllPipeline_block(), parent)
		branches = append(branches, runtime.ConditionalBranch{
			Expression: elifExpr,
			Pipelines:  elifPipelines,
			RawText:    ctx.GetText(),
		})
	}

	// --- ELSE branch ---
	if elseCtx := ctx.Else_condition(); elseCtx != nil {
		// Pass the condition pipeline as parent to nested pipelines
		elsePipelines := p.parsePipelineBlocks(task, elseCtx.AllPipeline_block(), parent)
		branches = append(branches, runtime.ConditionalBranch{
			Expression: nil, // else has no expression
			Pipelines:  elsePipelines,
			RawText:    ctx.GetText(),
		})
	}

	return &runtime.IfCondition{
		Branches: branches,
		RawText:  ctx.GetText(),
	}
}

// parsePipelineBlocks maps a list of pipeline_block nodes to []*TaskPipeline.
func (p *Parser) parsePipelineBlocks(task *runtime.Task, blocks []parser.IPipeline_blockContext, parent *runtime.TaskPipeline) []*runtime.TaskPipeline {
	var pipelines []*runtime.TaskPipeline
	for _, block := range blocks {
		if pipeline := p.parsePipeline(task, block, parent); pipeline != nil {
			pipelines = append(pipelines, pipeline)
		}
	}
	return pipelines
}
