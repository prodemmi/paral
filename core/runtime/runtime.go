package runtime

import (
	"fmt"
	"paral/core/metadata"
	"paral/core/variable"
)

type Runtime struct {
	Vars     []*variable.Variable
	Tasks    []*Task
	Metadata *metadata.Metadata
	Reporter *Reporter
	Executor *TaskExecutor
}

func NewRuntime(metadata *metadata.Metadata, reporter *Reporter) *Runtime {
	runtime := &Runtime{
		Vars:     make([]*variable.Variable, 0),
		Tasks:    make([]*Task, 0),
		Metadata: metadata,
		Reporter: reporter,
	}
	runtime.Executor = NewExecutor(runtime)
	return runtime
}

func NewTestRuntime() *Runtime {
	return &Runtime{
		Vars:  make([]*variable.Variable, 0),
		Tasks: make([]*Task, 0),
		Metadata: &metadata.Metadata{
			Filename: "test.paral",
			Content:  "",
		},
	}
}

func (r *Runtime) AddVar(v *variable.Variable) {
	r.Vars = append(r.Vars, v)
}

func (r *Runtime) AddTask(j *Task) {
	r.Tasks = append(r.Tasks, j)
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
	for i := range r.Tasks {
		task := r.Tasks[i]
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

// GetTaskDependencies returns the list of task IDs that this task depends on
func (r *Runtime) GetTaskDependencies(task *Task) []string {
	var dependencies []string
	for _, directive := range task.Directives {
		if directive.Type == "depend" {
			for _, param := range directive.Params {
				dependencies = append(dependencies, fmt.Sprint(param))
			}
		}
	}
	return dependencies
}

// HasDeferDirective checks if a task has @defer directive
func (r *Runtime) HasDeferDirective(task *Task) bool {
	for _, directive := range task.Directives {
		if directive.Type == "defer" {
			return true
		}
	}
	return false
}

// GetExecutionOrder returns tasks sorted by dependency order
func (r *Runtime) GetExecutionOrder() ([]*Task, []*Task, error) {
	var independentTasks []*Task
	var deferTasks []*Task

	// Separate tasks into independent and defer categories
	for _, task := range r.Tasks {
		if task.IsManual() {
			continue
		}

		if r.HasDeferDirective(task) {
			deferTasks = append(deferTasks, task)
		} else if len(r.GetTaskDependencies(task)) == 0 {
			independentTasks = append(independentTasks, task)
		} else {
			// Tasks with dependencies are handled in topological sort
			independentTasks = append(independentTasks, task)
		}
	}

	// Sort independent tasks by dependencies using topological sort
	sortedTasks, err := r.topologicalSort(independentTasks)
	if err != nil {
		return nil, nil, err
	}

	return sortedTasks, deferTasks, nil
}

// topologicalSort sorts tasks based on their dependencies
func (r *Runtime) topologicalSort(tasks []*Task) ([]*Task, error) {
	// Build adjacency list and in-degree count
	graph := make(map[string][]*Task)
	inDegree := make(map[string]int)
	taskMap := make(map[string]*Task)

	// Initialize
	for _, task := range tasks {
		taskId := task.GetTaskId()
		taskMap[taskId] = task
		inDegree[taskId] = 0
	}

	// Build graph
	for _, task := range tasks {
		taskId := task.GetTaskId()
		dependencies := r.GetTaskDependencies(task)

		for _, depId := range dependencies {
			if _, exists := taskMap[depId]; !exists {
				return nil, fmt.Errorf("task '%s' depends on undefined task '%s'", taskId, depId)
			}
			graph[depId] = append(graph[depId], task)
			inDegree[taskId]++
		}
	}

	// Kahn's algorithm
	var queue []*Task
	var result []*Task

	// Find all tasks with no dependencies
	for _, task := range tasks {
		taskId := task.GetTaskId()
		if inDegree[taskId] == 0 {
			queue = append(queue, task)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		currentId := current.GetTaskId()
		for _, dependent := range graph[currentId] {
			dependentId := dependent.GetTaskId()
			inDegree[dependentId]--
			if inDegree[dependentId] == 0 {
				queue = append(queue, dependent)
			}
		}
	}

	if len(result) != len(tasks) {
		return nil, fmt.Errorf("circular dependency detected in tasks")
	}

	return result, nil
}
