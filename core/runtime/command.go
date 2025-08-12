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

	// If the resolved command doesn't start with '@', execute it as a shell command
	if !strings.HasPrefix(resolvedCmd, "@") {
		if loopContext == nil {
			loopContext = &TaskLoopContext{}
		}
		output, cmdSuccess := executor.ExecuteShellCommand(resolvedCmd, loopContext.Value, ctx)
		if !cmdSuccess {
			return string(output), false, fmt.Errorf("\n\t\t\t%s", output)
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
