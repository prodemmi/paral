package runtime

import (
	"fmt"
	"strings"
)

type Command struct {
	Functions []*Function
	RawText   string
}

func (c *Command) GetRawResult() string {
	cmd := c.RawText
	for _, function := range c.Functions {
		result, err := function.Call()
		if err != nil {
			// Log error but continue to allow partial resolution
			continue
		}
		funcRaw := function.RawText
		switch res := result.(type) {
		case string:
			cmd = strings.ReplaceAll(cmd, funcRaw, res)
		case []interface{}:
			var strValues []string
			for _, v := range res {
				strValues = append(strValues, fmt.Sprint(v))
			}
			cmd = strings.ReplaceAll(cmd, funcRaw, strings.Join(strValues, " "))
		default:
			cmd = strings.ReplaceAll(cmd, funcRaw, fmt.Sprintf("%v", res))
		}
	}
	return strings.TrimSpace(cmd)
}

func (c *Command) GetResult(ctx *ExecutionContext, task *Task, executor *CommandExecutor) (string, bool, error) {
	// Resolve all function calls in RawText
	loopContext := task.GetActiveLoopContext()
	resolvedCmd := c.GetRawResult()

	// Check if there are unresolved functions (indicated by @ symbols that weren't replaced)
	if strings.Contains(resolvedCmd, "@") {
		// Check for common unresolved function patterns
		if strings.Contains(resolvedCmd, "@value") && task.GetActiveLoopContext() == nil {
			return "", false, fmt.Errorf("@value can only be used within a @for directive")
		}
		if strings.Contains(resolvedCmd, "@key") && task.GetActiveLoopContext() == nil {
			return "", false, fmt.Errorf("@key can only be used within a @for directive")
		}
		// Add more specific error checks for other context-dependent functions as needed

		// Generic check for any remaining @ symbols (potential unresolved functions)
		for _, part := range strings.Fields(resolvedCmd) {
			if strings.HasPrefix(part, "@") && !strings.HasPrefix(part, "@@") { // @@ might be escaped
				return "", false, fmt.Errorf("unresolved function or variable: %s", part)
			}
		}
	}

	// If the resolved command doesn't start with '@', execute it as a shell command
	if !strings.HasPrefix(resolvedCmd, "@") {
		if loopContext == nil {
			loopContext = &TaskLoopContext{}
		}
		output, cmdSuccess := executor.ExecuteShellCommand(resolvedCmd, loopContext.Value, ctx)
		if !cmdSuccess {
			errorMsg := string(output)
			// Clean up the error message - remove leading/trailing whitespace and newlines
			errorMsg = strings.TrimSpace(errorMsg)
			return errorMsg, false, fmt.Errorf("\n\t\t\t%s", errorMsg)
		}
		return strings.TrimRight(string(output), "\n"), true, nil
	}

	// Handle function-only commands
	result := resolvedCmd
	for _, fn := range c.Functions {
		fnResult, err := fn.Call()
		if err != nil {
			return result, false, err
		}
		if fnResult != nil {
			result = fmt.Sprint(fnResult)
		}
	}

	return strings.TrimRight(result, "\n"), true, nil
}

func (c *Command) getIndentation(task interface{}) string {
	return "  "
}
