package parser

import (
	"fmt"
	parser "paral/antlr/antlr"
	"paral/core/metadata"
	"paral/core/runtime"
	"strings"
)

func (p *Parser) parseTask(ctx parser.ITask_definitionContext) *runtime.Task {
	filename := ctx.GetStop().GetInputStream().GetSourceName()
	line := ctx.GetStop().GetLine()
	column := ctx.GetStop().GetColumn()

	// Get task name from string expression
	name := ctx.IDENTIFIER().GetText()

	description := ""

	// Process directives to override name with @id if present
	for _, directiveExpr := range ctx.AllTask_directive() {
		if dcr := directiveExpr.Directive(); dcr != nil && dcr.IDENTIFIER().GetText() == "description" {
			if len(dcr.AllExpression()) > 0 {
				descCandidate := dcr.AllExpression()[0].GetText()
				if strings.HasPrefix(descCandidate, "\"") && strings.HasSuffix(descCandidate, "\"") {
					descCandidate = strings.Trim(descCandidate, "\"")
				}
				description = descCandidate
			}
		}
	}

	mt := &metadata.Metadata{
		Filename: filename,
		Content:  ctx.GetText(),
		Line:     line,
		Column:   column,
	}

	// Check for duplicate task
	existJob := p.Runtime.GetTaskByID(name)
	if existJob != nil {
		p.Runtime.Reporter.ThrowRuntimeError(
			fmt.Sprintf("task %q has already been defined", name), mt,
		)
	}

	// Create new task
	task := runtime.NewTask(p.Runtime, name, description, filename, *mt)

	// Process directives
	for _, directiveExpr := range ctx.AllTask_directive() {
		if directive := p.parseTaskDirective(directiveExpr); directive != nil {
			if err := task.AddTaskDirective(directive); err != nil {
				p.Runtime.Reporter.ThrowSyntaxError(fmt.Sprintf("\nInvalid directive: %v", err), &directive.Metadata)
			}
		}
	}

	// Process pipeline blocks
	for _, pipelineBlock := range ctx.AllPipeline_block() {
		if pipeline := p.parsePipeline(task, pipelineBlock); pipeline != nil {
			task.AddTaskPipeline(pipeline)
		}
	}

	return task
}

func (p *Parser) parseTaskDirective(ctx parser.ITask_directiveContext) *runtime.Directive {
	directive := runtime.NewTaskDirective()
	if dcr := ctx.Directive(); dcr != nil {
		directive.Type = dcr.IDENTIFIER().GetText()
		for _, dcrParam := range dcr.AllExpression() {
			text := dcrParam.GetText()
			directive.Params = append(directive.Params, text)
		}
	}
	directive.Metadata = metadata.Metadata{
		Content: ctx.GetText(),
		Line:    ctx.GetStart().GetLine(),
		Column:  ctx.GetStart().GetColumn(),
	}
	return directive
}
