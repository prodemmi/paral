package parser

import (
	"paral/antlr/antlr"
	"paral/core/runtime"
)

func (p *Parser) parseIfCondition(task *runtime.Task, ctx parser.IIf_conditionContext) *runtime.IfCondition {
	thenPipelines := make([]*runtime.TaskPipeline, len(ctx.AllPipeline_block()))
	for i, pipelineBlock := range ctx.AllPipeline_block() {
		if pipeline := p.parsePipeline(task, pipelineBlock); pipeline != nil {
			thenPipelines[i] = pipeline
		}
	}

	var expression *runtime.Expression
	if expr := ctx.Expression(); expr != nil {
		expression = p.parseExpression(expr)
	}

	return &runtime.IfCondition{
		Type:          "if",
		Expression:    expression,
		ThenPipelines: thenPipelines,
		ElsePipelines: nil,
	}
}
