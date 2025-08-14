package runtime

import (
	"fmt"
	"paral/core/metadata"
)

type Match struct {
	Expression     *Expression
	MatchPipelines []*MatchPipeline
	Metadata       *metadata.Metadata
}

type MatchPipeline struct {
	Expression *Expression // _ or an expression
	Pipeline   *TaskPipeline
}

func (m *Match) Execute(ctx *ExecutionContext, task *Task, cmdExecutor *CommandExecutor) (bool, string, bool) {
	// evaluate the main match expression
	matchVal := m.Expression.Result

	for _, arm := range m.MatchPipelines {
		// check if this arm matches
		armMatches := false

		if arm.Expression != nil {
			if arm.Expression.IsLiteral("_") { // default case
				armMatches = true
			} else {
				armVal := arm.Expression.Result
				if fmt.Sprint(armVal) == fmt.Sprint(matchVal) {
					armMatches = true
				}
			}
		}

		if armMatches {
			// execute and return the result of the matched pipeline
			success, output, shouldPrint := arm.Pipeline.GetResult(ctx, task, cmdExecutor)
			return success, output, shouldPrint
		}
	}

	// no match found
	return false, "", false
}
