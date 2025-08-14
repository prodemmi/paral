package parser

import (
	parser "paral/antlr/antlr"
	"paral/core/metadata"
	"paral/core/runtime"
)

func (p *Parser) parseMatch(task *runtime.Task, ctx parser.IMatch_statementContext, parent *runtime.TaskPipeline) *runtime.Match {
	m := &runtime.Match{
		Expression:     p.parseExpression(task, ctx.Expression()), // the main @match(...) expression
		MatchPipelines: make([]*runtime.MatchPipeline, 0),
	}

	// loop through match_block rules
	for _, mb := range ctx.AllMatch_block() {
		// match_expression can be `_` or normal expression
		var expr *runtime.Expression
		if underscore := mb.Match_expression().UNDERSCORE(); underscore != nil {
			// create an expression representing the default `_`
			expr = runtime.NewLiteralExpression("_", &metadata.Metadata{
				Content: underscore.GetText(),
				Line:    underscore.GetSymbol().GetLine(),
				Column:  underscore.GetSymbol().GetColumn(),
			})
		} else {
			expr = p.parseExpression(task, mb.Match_expression().Expression())
		}

		// pipeline_block* for this match arm
		pipeline := p.parsePipeline(task, mb.Pipeline_block(), parent)
		m.MatchPipelines = append(m.MatchPipelines, &runtime.MatchPipeline{
			Expression: expr,
			Pipeline:   pipeline,
		})
	}

	return m
}
