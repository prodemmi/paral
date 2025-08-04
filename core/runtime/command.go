package runtime

import (
	"fmt"
	"strings"
)

type Command struct {
	Functions []*Function
	RawText   string
}

func (c *Command) GetRawResult(loopContext *TaskLoopContext, ctx *ExecutionContext) string {
	cmd := c.RawText
	for _, function := range c.Functions {
		if loopContext != nil {
			function.SetLoopContext(loopContext)
		}
		result, err := function.Call()
		if err != nil {
			continue
		}
		funcRaw := function.GetRaw()
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
	return cmd
}

func (c *Command) GetResult(ctx *ExecutionContext, task *Task, executor *CommandExecutor) (string, bool, error) {
	rawResult := c.RawText
	loopContext := task.GetActiveLoopContext() // Type assertion for LoopContext
	success := true
	for _, fn := range c.Functions {
		if loopContext != nil {
			fn.SetLoopContext(loopContext)
		}
		fnResult, err := fn.Call()
		if err != nil {
			return rawResult, false, err
		}
		if fnResult != nil {
			rawResult = fmt.Sprint(fnResult)
		}
	}

	if !strings.HasPrefix(strings.TrimSpace(c.RawText), "@") {
		if loopContext == nil {
			loopContext = &TaskLoopContext{}
		}
		output, cmdSuccess := executor.ExecuteShellCommand(c.RawText, loopContext.Value, ctx)
		if !cmdSuccess {
			return string(output), false, nil
		}
		return string(output), true, nil
	}

	return rawResult, success, nil
}

func (c *Command) getIndentation(task interface{}) string {
	return "  "
}
