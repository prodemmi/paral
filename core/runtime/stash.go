package runtime

import (
	"paral/core/metadata"
)

type Stash struct {
	Name     string
	TaskID   string
	RawText  string
	Pipeline *TaskPipeline
	Metadata metadata.Metadata
}

func NewStash(name, taskID, rawText string, metadata metadata.Metadata, pipeline *TaskPipeline) *Stash {
	return &Stash{
		Name:     name,
		TaskID:   taskID,
		RawText:  rawText,
		Pipeline: pipeline,
		Metadata: metadata,
	}
}

func (s *Stash) GetValue(ctx *ExecutionContext, task *Task, cmdExecutor *CommandExecutor) (string, bool) {
	if s.Pipeline.Command != nil {
		result, ok, _ := s.Pipeline.Command.GetResult(ctx, task, cmdExecutor)
		return result, ok
	}

	if s.Pipeline.Function != nil {
		result, err := s.Pipeline.Function.Call()
		return result.(string), err == nil
	}

	return "", false
}

func (s *Stash) GetRawValue() (string, bool) {
	if s.Pipeline.Command != nil {
		result := s.Pipeline.Command.GetRawResult()
		return result, true
	}

	if s.Pipeline.Command != nil {
		result, ok := s.Pipeline.Function.Call()
		return result.(string), ok != nil
	}

	return "", false
}
