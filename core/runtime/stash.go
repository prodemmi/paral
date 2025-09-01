package runtime

import (
	"fmt"
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

func (s *Stash) GetValue(ctx *ExecutionContext, task *Task, cmdExecutor *CommandExecutor) (interface{}, error) {
	if s.Pipeline.Command != nil {
		result, _, err := s.Pipeline.Command.GetResult(ctx, task, cmdExecutor)
		return result, err
	}

	if s.Pipeline.Function != nil {
		result, err := s.Pipeline.Function.Call()
		return result, err
	}

	if s.Pipeline.Expression != nil {
		result, err := s.Pipeline.Expression.EvaluateValue(task.GetActiveLoopContext())
		return result, err
	}

	return "", fmt.Errorf("failed to evaluate stash")
}

func (s *Stash) GetRawValue() (string, bool, error) {
	if s.Pipeline.Command != nil {
		result, err := s.Pipeline.Command.GetRawResult()
		if err != nil {
			return "", false, err
		}
		return result, true, nil
	}

	if s.Pipeline.Function != nil {
		result, ok := s.Pipeline.Function.Call()
		return result.(string), ok != nil, nil
	}

	if s.Pipeline.Expression != nil {
		result := s.Pipeline.Expression.RawText
		return result, true, nil
	}

	return "", false, nil
}
