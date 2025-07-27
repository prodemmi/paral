package parser

import (
	parser "paral/antlr/antlr"
	"paral/core/runtime"
)

type Parser struct {
	*parser.BaseParalParserListener
	Runtime  *runtime.Runtime
	Reporter *runtime.Reporter
}

func (p *Parser) ExitVariable_assignment(ctx *parser.Variable_assignmentContext) {
	if variable := p.parseVariable(ctx); variable != nil {
		p.Runtime.AddVar(variable)
	}
}

func (p *Parser) ExitTask_definition(ctx *parser.Task_definitionContext) {
	if task := p.parseJob(ctx); task != nil {
		p.Runtime.AddJob(task)
	}
}
