package runtime

import (
	"paral/core/metadata"
	"strings"
)

type TryCatch struct {
	TryPipelines   []*TaskPipeline
	CatchPipelines []*TaskPipeline
	Metadata       *metadata.Metadata
	LastError      string // Store the last error for @error access
}

func (tc *TryCatch) Execute(ctx *ExecutionContext, task *Task, cmdExecutor *CommandExecutor) ([]bool, []string, []bool) {
	var successes []bool
	var outputs []string
	var prints []bool

	// Clear any previous error
	tc.LastError = ""

	// Execute try block
	trySuccess := true
	var tryOutputs []string
	var tryPrints []bool
	var trySuccesses []bool

	for _, pipeline := range tc.TryPipelines {
		success, output, shouldPrint := pipeline.GetResult(ctx, task, cmdExecutor)

		// Store try results temporarily
		trySuccesses = append(trySuccesses, success)
		tryOutputs = append(tryOutputs, output)
		tryPrints = append(tryPrints, shouldPrint)

		if !success {
			trySuccess = false
			tc.LastError = strings.TrimLeft(output, " ▶")
			break // Stop executing remaining try pipelines on first error
		}
	}

	// If try block succeeded, include all try results
	if trySuccess {
		successes = append(successes, trySuccesses...)
		outputs = append(outputs, tryOutputs...)
		prints = append(prints, tryPrints...)
	} else {
		// If try block failed, execute catch block instead
		if len(tc.CatchPipelines) > 0 {
			// Set the error context for the catch block
			cmdExecutor.Runtime.SetTryCatchError(tc.LastError)

			for _, pipeline := range tc.CatchPipelines {
				success, output, shouldPrint := pipeline.GetResult(ctx, task, cmdExecutor)

				successes = append(successes, success)
				outputs = append(outputs, output)
				prints = append(prints, shouldPrint)
			}

			// Clear the error context after catch execution
			cmdExecutor.Runtime.ClearTryCatchError()
		} else {
			// No catch block, so include the failed try results
			successes = append(successes, trySuccesses...)
			outputs = append(outputs, tryOutputs...)
			prints = append(prints, tryPrints...)
		}
	}

	return successes, outputs, prints
}
