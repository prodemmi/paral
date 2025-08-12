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

func (i *IfCondition) Execute(ctx *ExecutionContext, task *Task, cmdExecutor *CommandExecutor) ([]bool, []string, []bool) {
	var successResults []bool
	var displayResults []string
	var shouldPrintResults []bool

	for _, branch := range i.Branches {
		if branch.Expression == nil || branch.Expression.IsTrue() {
			successResults = make([]bool, len(branch.Pipelines))
			displayResults = make([]string, len(branch.Pipelines))
			shouldPrintResults = make([]bool, len(branch.Pipelines))

			for idx, pipeline := range branch.Pipelines {
				successResults[idx], displayResults[idx], shouldPrintResults[idx] =
					pipeline.GetResult(ctx, task, cmdExecutor)
			}
			break
		}
	}

	return successResults, displayResults, shouldPrintResults
}
