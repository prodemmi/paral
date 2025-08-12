package runtime

// Condition is a wrapper for different condition types in your DSL.
type Condition struct {
	IfCondition    *IfCondition
	MatchCondition *MatchCondition
	RawText        string
}

// ConditionalBranch represents one branch in an if/elif/else chain.
// Expression == nil means it's the `else` branch.
type ConditionalBranch struct {
	Expression *Expression
	Pipelines  []*TaskPipeline
	RawText    string
}

// IfCondition represents an if-elif-else chain.
type IfCondition struct {
	Branches []ConditionalBranch
	RawText  string
}

// MatchCondition represents a match/case structure.
type MatchCondition struct {
	Expression *Expression
	Pipelines  map[*Expression][]*TaskPipeline
}

func (i *IfCondition) Execute(ctx *ExecutionContext, task *Task, cmdExecutor *CommandExecutor) (success bool, displayResult string, shouldPrint bool) {
	for _, branch := range i.Branches {
		if branch.Expression == nil || branch.Expression.IsTrue() {
			for _, pipeline := range branch.Pipelines {
				return pipeline.GetResult(ctx, task, cmdExecutor)
			}
			break
		}
	}
	return false, "", true
}
