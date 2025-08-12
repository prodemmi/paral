package parser

import (
	"fmt"
	parser "paral/antlr/antlr"
	"paral/core/runtime"
)

// parseCondition decides which type of condition to parse (if/match).
func (p *Parser) parseCondition(task *runtime.Task, ctx parser.IConditionContext) *runtime.Condition {
	if ifCondCtx := ctx.If_condition(); ifCondCtx != nil {
		ifCondition := p.parseIfCondition(task, ifCondCtx)
		return &runtime.Condition{
			IfCondition: ifCondition,
		}
	}

	// TODO: add MatchCondition parsing here when ready
	return nil
}

// parseIfCondition converts an if/elif/else AST into runtime.IfCondition.
func (p *Parser) parseIfCondition(task *runtime.Task, ctx parser.IIf_conditionContext) *runtime.IfCondition {
	var branches []runtime.ConditionalBranch

	// --- IF branch ---
	mainExpr := p.parseExpression(task, ctx.Expression())
	mainPipelines := p.parsePipelineBlocks(task, ctx.AllPipeline_block())
	branches = append(branches, runtime.ConditionalBranch{
		Expression: mainExpr,
		Pipelines:  mainPipelines,
		RawText:    ctx.GetText(),
	})

	// --- ELIF branches ---
	for _, elifCtx := range ctx.AllElseif_condition() {
		elifExpr := p.parseExpression(task, elifCtx.Expression())
		elifPipelines := p.parsePipelineBlocks(task, elifCtx.AllPipeline_block())
		branches = append(branches, runtime.ConditionalBranch{
			Expression: elifExpr,
			Pipelines:  elifPipelines,
			RawText:    ctx.GetText(),
		})
	}

	// --- ELSE branch ---
	if elseCtx := ctx.Else_condition(); elseCtx != nil {
		elsePipelines := p.parsePipelineBlocks(task, elseCtx.AllPipeline_block())
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
func (p *Parser) parsePipelineBlocks(task *runtime.Task, blocks []parser.IPipeline_blockContext) []*runtime.TaskPipeline {

	var pipelines []*runtime.TaskPipeline
	for _, block := range blocks {
		if pipeline := p.parsePipeline(task, block); pipeline != nil {
			if pipeline.Command != nil {
				fmt.Println("pipeline", pipeline.Command.RawText)
			}
			pipelines = append(pipelines, pipeline)
		}
	}
	return pipelines
}
