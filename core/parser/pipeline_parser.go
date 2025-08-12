package parser

import (
	"fmt"
	parser "paral/antlr/antlr"
	"paral/core/runtime"
	"strings"
)

func (p *Parser) parsePipeline(task *runtime.Task, ctx parser.IPipeline_blockContext) *runtime.TaskPipeline {
	return p.parsePipelineContent(task, ctx.Pipeline_content())
}

func (p *Parser) parsePipelineContent(task *runtime.Task, ctx parser.IPipeline_contentContext) *runtime.TaskPipeline {

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

		if unknownCtx := item.Unknown_command(); unknownCtx != nil {
			command := p.parseUnknownCommand(task, unknownCtx)
			return &runtime.TaskPipeline{
				Command: command,
				Block:   blockIndex,
			}
		}
	}

	// If we reach here, no specific pipeline item was found
	if ctx.GetText() != "" {
		fmt.Printf("Warning: Unhandled pipeline content: %s\n", ctx.GetText())
		return &runtime.TaskPipeline{
			Command: &runtime.Command{
				RawText:   strings.TrimSpace(ctx.GetText()),
				Functions: nil,
			},
		}
	}

	return &runtime.TaskPipeline{
		Command: &runtime.Command{
			RawText:   strings.TrimSpace(ctx.GetText()),
			Functions: nil,
		},
	}
}

// parseUnknownCommand extracts only functions from the parsed expressions
func (p *Parser) parseUnknownCommand(task *runtime.Task, ctx parser.IUnknown_commandContext) *runtime.Command {
	command := &runtime.Command{
		RawText:   ctx.GetText(),
		Functions: make([]*runtime.Function, 0),
	}

	// Parse all expressions and extract only the functions
	for _, exprCtx := range ctx.AllExpression() {
		p.extractFunctionsFromExpression(task, command, exprCtx)
	}

	return command
}

// extractFunctionsFromExpression recursively extracts functions from expressions
func (p *Parser) extractFunctionsFromExpression(task *runtime.Task, command *runtime.Command, ctx parser.IExpressionContext) {
	// Check if this expression is a function call
	if fnCtx := ctx.Function(); fnCtx != nil {
		function := p.parseFunction(task, fnCtx)
		if function != nil {
			command.Functions = append(command.Functions, function)
		}
	}
}
