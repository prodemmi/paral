package parser

import (
	"fmt"
	parser "paral/antlr/antlr"
	"paral/core/runtime"
	"strings"
)

func (p *Parser) parsePipeline(task *runtime.Task, ctx parser.IPipeline_blockContext, parent *runtime.TaskPipeline) *runtime.TaskPipeline {
	return p.parsePipelineContent(task, ctx.Pipeline_content(), parent)
}

func (p *Parser) parsePipelineContent(task *runtime.Task, ctx parser.IPipeline_contentContext, parent *runtime.TaskPipeline) *runtime.TaskPipeline {
	for blockIndex, item := range ctx.AllPipeline_item() {
		if condCtx := item.Condition(); condCtx != nil {
			if ifCtx := condCtx.If_condition(); ifCtx != nil {
				// Create the condition pipeline first
				conditionPipeline := &runtime.TaskPipeline{
					Parent: parent,
					Block:  blockIndex,
				}

				// Parse the if condition with the condition pipeline as parent for nested pipelines
				ifCondition := p.parseIfCondition(task, ifCtx, conditionPipeline)

				// Set the condition on the pipeline
				conditionPipeline.Condition = &runtime.Condition{
					IfCondition: ifCondition,
					RawText:     ifCtx.GetText(),
				}

				return conditionPipeline
			}
		}

		if tryCatchCtx := item.Try_catch(); tryCatchCtx != nil {
			// Create the try-catch pipeline first
			tryCatchPipeline := &runtime.TaskPipeline{
				Parent: parent,
				Block:  blockIndex,
			}

			// Parse the try-catch with the try-catch pipeline as parent for nested pipelines
			tryCatch := p.parseTryCatch(task, tryCatchCtx, tryCatchPipeline)

			// Set the try-catch on the pipeline
			tryCatchPipeline.TryCatch = tryCatch

			return tryCatchPipeline
		}

		if bufCtx := item.Buf(); bufCtx != nil {
			buf := p.parseBuf(task, bufCtx, parent)
			return &runtime.TaskPipeline{
				Parent: parent,
				Buf:    buf,
				Block:  blockIndex,
			}
		}

		if stashCtx := item.Stash(); stashCtx != nil {
			stash := p.parseStash(task, stashCtx, parent)
			return &runtime.TaskPipeline{
				Parent: parent,
				Stash:  stash,
				Block:  blockIndex,
			}
		}

		if fnCtx := item.Function(); fnCtx != nil {
			function := p.parseFunction(task, fnCtx)
			return &runtime.TaskPipeline{
				Parent:   parent,
				Function: function,
				Block:    blockIndex,
			}
		}

		if unknownCtx := item.Unknown_command(); unknownCtx != nil {
			command := p.parseUnknownCommand(task, unknownCtx)
			return &runtime.TaskPipeline{
				Parent:  parent,
				Command: command,
				Block:   blockIndex,
			}
		}
	}

	// If we reach here, no specific pipeline item was found
	if ctx.GetText() != "" {
		fmt.Printf("Warning: Unhandled pipeline content: %s\n", ctx.GetText())
		return &runtime.TaskPipeline{
			Parent: parent,
			Command: &runtime.Command{
				RawText:   strings.TrimSpace(ctx.GetText()),
				Functions: nil,
			},
		}
	}

	return &runtime.TaskPipeline{
		Parent: parent,
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
