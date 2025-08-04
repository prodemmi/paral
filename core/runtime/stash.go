package runtime

import (
	"paral/core/metadata"
)

type Stash struct {
	Name     string
	TaskID   string
	RawText  string
	Command  *Command
	Metadata metadata.Metadata
}

func NewStash(name, taskID, rawText string, metadata metadata.Metadata, pipeline *Command) *Stash {
	return &Stash{
		Name:     name,
		TaskID:   taskID,
		RawText:  rawText,
		Command:  pipeline,
		Metadata: metadata,
	}
}

func (s *Stash) GetValue(ctx *ExecutionContext, task *Task, cmdExecutor *CommandExecutor) (string, bool) {
	result, ok, _ := s.Command.GetResult(ctx, task, cmdExecutor)
	return result, ok
}
