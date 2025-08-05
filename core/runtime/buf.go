package runtime

import "paral/core/metadata"

type Buf struct {
	Name         string
	SourceTaskID string
	RawText      string
	Command      *Command
	Metadata     metadata.Metadata
}

func NewBuf(name, sourceTaskID, rawText string, metadata metadata.Metadata, pipeline *Command) *Buf {
	return &Buf{
		Name:         name,
		SourceTaskID: sourceTaskID,
		RawText:      rawText,
		Command:      pipeline,
		Metadata:     metadata,
	}
}

func (s *Buf) GetValue(ctx *ExecutionContext, task *Task, cmdExecutor *CommandExecutor) (string, bool) {
	result, ok, _ := s.Command.GetResult(ctx, task, cmdExecutor)
	return result, ok
}
