package runtime

import (
	"fmt"
	"paral/core/metadata"
	"paral/core/variable"
)

type Runtime struct {
	Vars     []*variable.Variable
	Jobs     []*Task
	Metadata *metadata.Metadata
	Reporter *Reporter
}

func NewRuntime(metadata *metadata.Metadata, reporter *Reporter) *Runtime {
	return &Runtime{
		Vars:     make([]*variable.Variable, 0),
		Jobs:     make([]*Task, 0),
		Metadata: metadata,
		Reporter: reporter,
	}
}

func NewTestRuntime() *Runtime {
	return &Runtime{
		Vars: make([]*variable.Variable, 0),
		Jobs: make([]*Task, 0),
		Metadata: &metadata.Metadata{
			Filename: "test.paral",
			Content:  "",
		},
	}
}

func (r *Runtime) AddVar(v *variable.Variable) {
	r.Vars = append(r.Vars, v)
}

func (r *Runtime) AddJob(j *Task) {
	r.Jobs = append(r.Jobs, j)
}

func (r *Runtime) GetVariable(name string) *variable.Variable {
	for _, v := range r.Vars {
		if v.Name == name {
			return v
		}
	}
	return nil
}

func (r *Runtime) GetTaskByID(id string) *Task {
	var targetJob *Task
	for i := range r.Jobs {
		task := r.Jobs[i]
		if task.Name == id {
			targetJob = task
			break
		}
		// Also check if task has an @id directive matching the jobName
		for _, directive := range task.Directives {
			if directive.Type == "id" && len(directive.Params) > 0 {
				if fmt.Sprint(directive.Params[0]) == id {
					targetJob = task
					break
				}
			}
		}
		if targetJob != nil {
			break
		}
	}
	return targetJob
}
