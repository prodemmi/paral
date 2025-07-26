package core

type Condition struct {
	Type         string
	Expression   Command
	ThenCommand  Command
	ThenCommands []Command
	ElseCommand  Command
	ElseCommands []Command
}
