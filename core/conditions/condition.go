package conditions

import (
	"paral/core/runtime"
)

type Condition struct {
	Type         string
	Expression   *runtime.Command
	ThenCommand  *runtime.Command
	ThenCommands []*runtime.Command
	ElseCommand  *runtime.Command
	ElseCommands []*runtime.Command
}
