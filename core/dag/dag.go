package dag

import (
	"encoding/json"
	"fmt"
	"paral/core/runtime"
	"paral/core/variable"
	"strings"
)

// DAG represents the Directed Acyclic Graph
type DAG struct {
	Nodes   []DAGNode        `json:"nodes"`
	Edges   []DAGEdge        `json:"edges"`
	Runtime *runtime.Runtime `json:"-"`
}

// DAGNode represents a node in the DAG (task, pipeline, or control structure)
type DAGNode struct {
	ID       string                 `json:"id"`
	Label    string                 `json:"label"`
	Type     string                 `json:"type"` // "task", "pipeline", "condition", "trycatch", "match", "stash", "buf"
	Metadata map[string]interface{} `json:"metadata"`
}

// DAGEdge represents an edge in the DAG
type DAGEdge struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Label  string `json:"label,omitempty"`
}

// MarshalJSON customizes JSON serialization to avoid cycles
func (d *DAG) MarshalJSON() ([]byte, error) {
	// Create a sanitized copy of the DAG for JSON serialization
	sanitized := struct {
		Nodes []DAGNode `json:"nodes"`
		Edges []DAGEdge `json:"edges"`
	}{
		Nodes: make([]DAGNode, len(d.Nodes)),
		Edges: d.Edges, // Edges are safe as they contain only strings
	}

	// Deep-copy nodes and sanitize metadata
	for i, node := range d.Nodes {
		sanitized.Nodes[i] = DAGNode{
			ID:       node.ID,
			Label:    node.Label,
			Type:     node.Type,
			Metadata: sanitizeMetadataForJSON(node.Metadata),
		}
	}

	return json.Marshal(&sanitized)
}

// sanitizeMetadataForJSON ensures metadata is JSON-serializable
func sanitizeMetadataForJSON(metadata map[string]interface{}) map[string]interface{} {
	sanitized := make(map[string]interface{})
	for key, value := range metadata {
		switch v := value.(type) {
		case string, int, float64, bool:
			sanitized[key] = v
		case []string:
			sanitized[key] = v
		case []interface{}:
			// Convert to []string if possible
			strSlice := make([]string, 0, len(v))
			for _, item := range v {
				if s, ok := item.(string); ok {
					strSlice = append(strSlice, s)
				} else {
					strSlice = append(strSlice, fmt.Sprint(item))
				}
			}
			sanitized[key] = strSlice
		case map[string]interface{}:
			sanitized[key] = sanitizeMetadataForJSON(v)
		default:
			// Convert complex types to string to avoid serialization issues
			sanitized[key] = fmt.Sprint(v)
		}
	}
	return sanitized
}

// NewDAG creates a new DAG instance
func NewDAG(r *runtime.Runtime) *DAG {
	return &DAG{
		Nodes:   []DAGNode{},
		Edges:   []DAGEdge{},
		Runtime: r,
	}
}

// GenerateDAG creates a detailed DAG from the Runtime struct
func (d *DAG) GenerateDAG() *DAG {
	// Map to track task nodes for dependency edges
	jobNodeIDs := make(map[string][]string) // Maps task name to list of instance IDs
	seenNodes := make(map[string]bool)      // Tracks nodes to detect cycles
	stash := make(map[string]interface{})   // Global stash for condition evaluation
	buf := make(map[string]interface{})     // Global buf for condition evaluation

	// First pass: create all task nodes and collect instance IDs
	for _, task := range d.Runtime.Tasks {
		// Resolve matrix values for @for directive
		matrixValues := [][]string{{}} // Default: single instance if no @for
		if forDirective := d.findDirective(task.Directives, "for"); forDirective != nil && len(forDirective.Params) > 0 {
			if varName, ok := forDirective.Params[0].(*runtime.Expression); ok {
				// Check if varName is a variable or matrix
				if varValue, err := varName.EvaluateValue(task.GetActiveLoopContext()); err == nil {
					var err error
					matrixValues, err = d.convertToStringMatrix(varValue)
					if err != nil {
						d.Runtime.Reporter.Warn(fmt.Sprintf("Invalid @for variable %s: %v", varName, err), &task.Metadata)
						continue
					}
				} else {
					d.Runtime.Reporter.Warn(fmt.Sprintf("@for target %s not found", varName), &task.Metadata)
					continue
				}
			} else {
				d.Runtime.Reporter.Warn("@for requires a string variable or matrix name", &task.Metadata)
				continue
			}
		}

		// Create a task node for each matrix combination
		for matrixIdx, matrix := range matrixValues {
			taskInstanceID := fmt.Sprintf("task_%s_%d", task.Name, matrixIdx)
			if len(matrixValues) == 1 && len(matrixValues[0]) == 0 {
				taskInstanceID = fmt.Sprintf("task_%s", task.Name)
			}
			if seenNodes[taskInstanceID] {
				d.Runtime.Reporter.Warn(fmt.Sprintf("Duplicate task instance ID %s detected", taskInstanceID), &task.Metadata)
				continue
			}
			seenNodes[taskInstanceID] = true
			jobNodeIDs[task.Name] = append(jobNodeIDs[task.Name], taskInstanceID)

			// Task metadata
			taskMetadata := map[string]interface{}{
				"filename":    task.Filename,
				"description": task.Description,
				"directives":  d.formatTaskDirectives(task.Directives),
				"pipelines":   len(task.Pipelines),
			}

			if len(matrix) > 0 {
				taskMetadata["matrix"] = matrix
			}

			// Extract task-level properties
			if schedule := d.findDirective(task.Directives, "schedule"); schedule != nil && len(schedule.Params) > 0 {
				if cron, ok := schedule.Params[0].(string); ok {
					taskMetadata["schedule"] = cron
				}
			}
			if manual := d.findDirective(task.Directives, "manual"); manual != nil {
				manualValue := true
				if len(manual.Params) > 0 {
					if val, ok := manual.Params[0].(bool); ok {
						manualValue = val
					}
				}
				taskMetadata["manual"] = manualValue
			}

			// Add task node
			d.Nodes = append(d.Nodes, DAGNode{
				ID:       taskInstanceID,
				Label:    fmt.Sprintf("%s%s", task.Name, d.formatMatrixLabel(matrix)),
				Type:     "task",
				Metadata: taskMetadata,
			})

			// Process pipelines for this task instance
			var prevPipelineIDs []string
			for pipelineIdx, pipeline := range task.Pipelines {
				pipelineNodeIDs, err := d.processPipelineForDAG(pipeline, task.Name, matrixIdx, pipelineIdx, matrix, stash, buf, seenNodes)
				if err != nil {
					d.Runtime.Reporter.Warn(fmt.Sprintf("Error processing pipeline %d for task %s: %v", pipelineIdx, task.Name, err), &task.Metadata)
					continue
				}
				// Connect task to its first pipeline(s) or previous pipeline(s)
				if pipelineIdx == 0 {
					for _, pipelineNodeID := range pipelineNodeIDs {
						d.Edges = append(d.Edges, DAGEdge{
							Source: taskInstanceID,
							Target: pipelineNodeID,
							Label:  "pipeline",
						})
					}
				} else {
					// Connect previous pipeline(s) to current pipeline(s)
					for _, prevID := range prevPipelineIDs {
						for _, currID := range pipelineNodeIDs {
							d.Edges = append(d.Edges, DAGEdge{
								Source: prevID,
								Target: currID,
								Label:  "pipeline",
							})
						}
					}
				}
				prevPipelineIDs = pipelineNodeIDs
			}
		}

		// Handle dependencies (@depend directive)
		if dependDirective := d.findDirective(task.Directives, "depend"); dependDirective != nil {
			for _, dep := range dependDirective.Params {
				depName, ok := dep.(string)
				if !ok {
					d.Runtime.Reporter.Warn(fmt.Sprintf("Invalid dependency %v for task %s", dep, task.Name), &task.Metadata)
					continue
				}
				if depNodeIDs, exists := jobNodeIDs[depName]; exists {
					for _, depNodeID := range depNodeIDs {
						for _, taskNodeID := range jobNodeIDs[task.Name] {
							d.Edges = append(d.Edges, DAGEdge{
								Source: depNodeID,
								Target: taskNodeID,
								Label:  "dependency",
							})
						}
					}
				} else {
					d.Runtime.Reporter.Warn(fmt.Sprintf("Dependency %s not found for task %s", depName, task.Name), &task.Metadata)
				}
			}
		}

		// Handle defer dependencies (@defer directive)
		if deferDirective := d.findDirective(task.Directives, "defer"); deferDirective != nil {
			for depTaskName, depNodeIDs := range jobNodeIDs {
				if depTaskName != task.Name {
					for _, depNodeID := range depNodeIDs {
						for _, taskNodeID := range jobNodeIDs[task.Name] {
							d.Edges = append(d.Edges, DAGEdge{
								Source: depNodeID,
								Target: taskNodeID,
								Label:  "defer",
							})
						}
					}
				}
			}
		}
	}

	return d
}

// processPipelineForDAG processes a single pipeline and returns node IDs
func (d *DAG) processPipelineForDAG(pipeline *runtime.TaskPipeline, taskName string, matrixIdx, pipelineIdx int, matrix []string, stash map[string]interface{}, buf map[string]interface{}, seenNodes map[string]bool) ([]string, error) {
	var nodeIDs []string

	// Generate a unique ID for the pipeline
	pipelineID := fmt.Sprintf("pipeline_%s_%d_%d", taskName, matrixIdx, pipelineIdx)
	if seenNodes[pipelineID] {
		return nil, fmt.Errorf("duplicate pipeline ID %s detected", pipelineID)
	}
	seenNodes[pipelineID] = true

	// Pipeline metadata
	pipelineMetadata := map[string]interface{}{
		"task":   taskName,
		"matrix": matrix,
		"index":  pipelineIdx,
	}

	// Process pipeline components
	if pipeline.Command != nil {
		// Sanitize command raw text for metadata
		commandText, err := pipeline.Command.GetRawResult()
		if err != nil {
			commandText = fmt.Sprintf("error: %v", err)
		}
		pipelineMetadata["command"] = commandText
		d.Nodes = append(d.Nodes, DAGNode{
			ID:       pipelineID,
			Label:    fmt.Sprintf("Command_%s_%d", taskName, pipelineIdx),
			Type:     "command",
			Metadata: pipelineMetadata,
		})
		nodeIDs = append(nodeIDs, pipelineID)
	} else if pipeline.Function != nil {
		// Sanitize function raw text for metadata
		functionText := pipeline.Function.GetRaw()
		pipelineMetadata["function"] = functionText
		d.Nodes = append(d.Nodes, DAGNode{
			ID:       pipelineID,
			Label:    fmt.Sprintf("%s_%s_%d", pipeline.Function.Type, taskName, pipelineIdx),
			Type:     "function",
			Metadata: pipelineMetadata,
		})
		nodeIDs = append(nodeIDs, pipelineID)
	} else if pipeline.Stash != nil {
		stashName := pipeline.Stash.Name
		stashValue, ok, err := pipeline.Stash.GetRawValue()
		if err != nil {
			pipelineMetadata["stash_error"] = fmt.Sprint(err)
		} else if ok {
			pipelineMetadata["stash_value"] = stashValue
		}
		stash[stashName] = stashValue
		d.Nodes = append(d.Nodes, DAGNode{
			ID:       pipelineID,
			Label:    fmt.Sprintf("Stash_%s_%d", taskName, pipelineIdx),
			Type:     "stash",
			Metadata: pipelineMetadata,
		})
		nodeIDs = append(nodeIDs, pipelineID)
	} else if pipeline.Buf != nil {
		bufName := pipeline.Buf.Name
		bufValue, ok, err := pipeline.Buf.GetRawValue()
		if err != nil {
			pipelineMetadata["buf_error"] = fmt.Sprint(err)
		} else if ok {
			pipelineMetadata["buf_value"] = bufValue
		}
		buf[bufName] = bufValue
		d.Nodes = append(d.Nodes, DAGNode{
			ID:       pipelineID,
			Label:    fmt.Sprintf("Buf_%s_%d", taskName, pipelineIdx),
			Type:     "buf",
			Metadata: pipelineMetadata,
		})
		nodeIDs = append(nodeIDs, pipelineID)
	} else if pipeline.Condition != nil && pipeline.Condition.IfCondition != nil {
		// Process if condition branches
		for branchIdx, branch := range pipeline.Condition.IfCondition.Branches {
			branchID := fmt.Sprintf("%s_branch_%d", pipelineID, branchIdx)
			if seenNodes[branchID] {
				return nil, fmt.Errorf("duplicate branch ID %s detected", branchID)
			}
			seenNodes[branchID] = true

			branchMetadata := map[string]interface{}{
				"task":     taskName,
				"matrix":   matrix,
				"index":    pipelineIdx,
				"branch":   branchIdx,
				"raw_text": branch.RawText,
			}

			d.Nodes = append(d.Nodes, DAGNode{
				ID:       branchID,
				Label:    fmt.Sprintf("Condition_%s_%d_%d", taskName, pipelineIdx, branchIdx),
				Type:     "condition",
				Metadata: branchMetadata,
			})
			nodeIDs = append(nodeIDs, branchID)

			// Process pipelines in the branch
			var prevBranchPipelineIDs []string
			for subPipelineIdx, subPipeline := range branch.Pipelines {
				// Avoid recursive processing by limiting depth and sanitizing metadata
				if subPipelineIdx > 100 { // Arbitrary limit to prevent infinite recursion
					return nil, fmt.Errorf("excessive subpipeline depth in condition for task %s", taskName)
				}
				subPipelineID := fmt.Sprintf("%s_subpipeline_%d", branchID, subPipelineIdx)
				if seenNodes[subPipelineID] {
					return nil, fmt.Errorf("duplicate subpipeline ID %s detected", subPipelineID)
				}
				seenNodes[subPipelineID] = true

				// Process sub-pipeline without allowing further nesting of conditions/trycatch/match
				if subPipeline.Command != nil || subPipeline.Function != nil || subPipeline.Stash != nil || subPipeline.Buf != nil {
					subNodeIDs, err := d.processPipelineForDAG(subPipeline, taskName, matrixIdx, subPipelineIdx, matrix, stash, buf, seenNodes)
					if err != nil {
						return nil, err
					}
					for _, subNodeID := range subNodeIDs {
						d.Edges = append(d.Edges, DAGEdge{
							Source: branchID,
							Target: subNodeID,
							Label:  "branch",
						})
					}
					if subPipelineIdx > 0 {
						for _, prevID := range prevBranchPipelineIDs {
							for _, currID := range subNodeIDs {
								d.Edges = append(d.Edges, DAGEdge{
									Source: prevID,
									Target: currID,
									Label:  "pipeline",
								})
							}
						}
					}
					prevBranchPipelineIDs = subNodeIDs
				} else {
					d.Runtime.Reporter.Warn(fmt.Sprintf("Nested condition/trycatch/match in task %s pipeline %d branch %d not supported in DAG", taskName, pipelineIdx, branchIdx), nil)
				}
			}
		}
	} else if pipeline.TryCatch != nil {
		tryCatchID := fmt.Sprintf("%s_trycatch", pipelineID)
		if seenNodes[tryCatchID] {
			return nil, fmt.Errorf("duplicate try-catch ID %s detected", tryCatchID)
		}
		seenNodes[tryCatchID] = true

		pipelineMetadata["trycatch"] = "try-catch block"
		d.Nodes = append(d.Nodes, DAGNode{
			ID:       tryCatchID,
			Label:    fmt.Sprintf("TryCatch_%s_%d", taskName, pipelineIdx),
			Type:     "trycatch",
			Metadata: pipelineMetadata,
		})
		nodeIDs = append(nodeIDs, tryCatchID)

		// Process try pipelines
		var prevTryPipelineIDs []string
		for tryIdx, tryPipeline := range pipeline.TryCatch.TryPipelines {
			tryPipelineID := fmt.Sprintf("%s_try_%d", tryCatchID, tryIdx)
			if seenNodes[tryPipelineID] {
				return nil, fmt.Errorf("duplicate try pipeline ID %s detected", tryPipelineID)
			}
			seenNodes[tryPipelineID] = true

			// Process try pipeline without allowing further nesting
			if tryPipeline.Command != nil || tryPipeline.Function != nil || tryPipeline.Stash != nil || tryPipeline.Buf != nil {
				tryNodeIDs, err := d.processPipelineForDAG(tryPipeline, taskName, matrixIdx, tryIdx, matrix, stash, buf, seenNodes)
				if err != nil {
					return nil, err
				}
				for _, tryNodeID := range tryNodeIDs {
					d.Edges = append(d.Edges, DAGEdge{
						Source: tryCatchID,
						Target: tryNodeID,
						Label:  "try",
					})
				}
				if tryIdx > 0 {
					for _, prevID := range prevTryPipelineIDs {
						for _, currID := range tryNodeIDs {
							d.Edges = append(d.Edges, DAGEdge{
								Source: prevID,
								Target: currID,
								Label:  "pipeline",
							})
						}
					}
				}
				prevTryPipelineIDs = tryNodeIDs
			} else {
				d.Runtime.Reporter.Warn(fmt.Sprintf("Nested condition/trycatch/match in task %s try pipeline %d not supported in DAG", taskName, tryIdx), nil)
			}
		}

		// Process catch pipelines
		var prevCatchPipelineIDs []string
		for catchIdx, catchPipeline := range pipeline.TryCatch.CatchPipelines {
			catchPipelineID := fmt.Sprintf("%s_catch_%d", tryCatchID, catchIdx)
			if seenNodes[catchPipelineID] {
				return nil, fmt.Errorf("duplicate catch pipeline ID %s detected", catchPipelineID)
			}
			seenNodes[catchPipelineID] = true

			// Process catch pipeline without allowing further nesting
			if catchPipeline.Command != nil || catchPipeline.Function != nil || catchPipeline.Stash != nil || catchPipeline.Buf != nil {
				catchNodeIDs, err := d.processPipelineForDAG(catchPipeline, taskName, matrixIdx, catchIdx, matrix, stash, buf, seenNodes)
				if err != nil {
					return nil, err
				}
				for _, catchNodeID := range catchNodeIDs {
					d.Edges = append(d.Edges, DAGEdge{
						Source: tryCatchID,
						Target: catchNodeID,
						Label:  "catch",
					})
				}
				if catchIdx > 0 {
					for _, prevID := range prevCatchPipelineIDs {
						for _, currID := range catchNodeIDs {
							d.Edges = append(d.Edges, DAGEdge{
								Source: prevID,
								Target: currID,
								Label:  "pipeline",
							})
						}
					}
				}
				prevCatchPipelineIDs = catchNodeIDs
			} else {
				d.Runtime.Reporter.Warn(fmt.Sprintf("Nested condition/trycatch/match in task %s catch pipeline %d not supported in DAG", taskName, catchIdx), nil)
			}
		}
	} else if pipeline.Match != nil {
		matchID := fmt.Sprintf("%s_match", pipelineID)
		if seenNodes[matchID] {
			return nil, fmt.Errorf("duplicate match ID %s detected", matchID)
		}
		seenNodes[matchID] = true

		pipelineMetadata["match"] = "match block"
		d.Nodes = append(d.Nodes, DAGNode{
			ID:       matchID,
			Label:    fmt.Sprintf("Match_%s_%d", taskName, pipelineIdx),
			Type:     "match",
			Metadata: pipelineMetadata,
		})
		nodeIDs = append(nodeIDs, matchID)

		// Process match pipelines
		var prevMatchPipelineIDs []string
		for matchIdx, matchPipeline := range pipeline.Match.MatchPipelines {
			matchPipelineID := fmt.Sprintf("%s_matchpipeline_%d", matchID, matchIdx)
			if seenNodes[matchPipelineID] {
				return nil, fmt.Errorf("duplicate match pipeline ID %s detected", matchPipelineID)
			}
			seenNodes[matchPipelineID] = true

			// Process match pipeline without allowing further nesting
			if matchPipeline.Pipeline.Command != nil || matchPipeline.Pipeline.Function != nil || matchPipeline.Pipeline.Stash != nil || matchPipeline.Pipeline.Buf != nil {
				matchNodeIDs, err := d.processPipelineForDAG(matchPipeline.Pipeline, taskName, matrixIdx, matchIdx, matrix, stash, buf, seenNodes)
				if err != nil {
					return nil, err
				}
				for _, matchNodeID := range matchNodeIDs {
					d.Edges = append(d.Edges, DAGEdge{
						Source: matchID,
						Target: matchNodeID,
						Label:  "match",
					})
				}
				if matchIdx > 0 {
					for _, prevID := range prevMatchPipelineIDs {
						for _, currID := range matchNodeIDs {
							d.Edges = append(d.Edges, DAGEdge{
								Source: prevID,
								Target: currID,
								Label:  "pipeline",
							})
						}
					}
				}
				prevMatchPipelineIDs = matchNodeIDs
			} else {
				d.Runtime.Reporter.Warn(fmt.Sprintf("Nested condition/trycatch/match in task %s match pipeline %d not supported in DAG", taskName, matchIdx), nil)
			}
		}
	}

	return nodeIDs, nil
}

// findDirective finds a directive by type
func (d *DAG) findDirective(directives []*runtime.Directive, directiveType string) *runtime.Directive {
	for _, directive := range directives {
		if directive.Type == directiveType {
			return directive
		}
	}
	return nil
}

// formatTaskDirectives formats directives for metadata
func (d *DAG) formatTaskDirectives(directives []*runtime.Directive) []map[string]interface{} {
	var formatted []map[string]interface{}
	for _, directive := range directives {
		sanitizedParams := make([]interface{}, len(directive.Params))
		for i, param := range directive.Params {
			switch v := param.(type) {
			case *runtime.Expression:
				sanitizedParams[i] = v.RawText // Use raw text to avoid cycles
			default:
				sanitizedParams[i] = v
			}
		}
		formatted = append(formatted, map[string]interface{}{
			"type":   directive.Type,
			"params": sanitizedParams,
		})
	}
	return formatted
}

// formatMatrixLabel formats the matrix label for display
func (d *DAG) formatMatrixLabel(matrix []string) string {
	if len(matrix) == 0 {
		return ""
	}
	return fmt.Sprintf("[%s]", strings.Join(matrix, ","))
}

// convertToStringMatrix converts a variable value to a string matrix
func (d *DAG) convertToStringMatrix(value interface{}) ([][]string, error) {
	switch v := value.(type) {
	case *variable.Variable:
		if mv, ok := v.AsMatrix(); ok {
			result := make([][]string, len(mv.Value))
			for i, list := range mv.Value {
				result[i] = make([]string, len(list))
				for j, val := range list {
					result[i][j] = fmt.Sprint(val)
				}
			}
			return result, nil
		}
		if lv, ok := v.AsList(); ok {
			result := make([][]string, len(lv.Value))
			for i, val := range lv.Value {
				result[i] = []string{fmt.Sprint(val)}
			}
			return result, nil
		}
		return [][]string{{fmt.Sprint(v.Value)}}, nil
	default:
		return [][]string{{fmt.Sprint(v)}}, nil
	}
}

// hasCycle checks for cycles in the DAG
func (d *DAG) hasCycle() bool {
	graph := make(map[string][]string)
	for _, edge := range d.Edges {
		graph[edge.Source] = append(graph[edge.Source], edge.Target)
	}

	state := make(map[string]int) // 0: unvisited, 1: visiting, 2: visited

	var dfs func(node string) bool
	dfs = func(node string) bool {
		if state[node] == 1 {
			return true // Back edge found, cycle detected
		}
		if state[node] == 2 {
			return false // Already fully explored
		}

		state[node] = 1 // Mark as visiting
		for _, neighbor := range graph[node] {
			if dfs(neighbor) {
				return true
			}
		}
		state[node] = 2 // Mark as visited
		return false
	}

	// Check each node
	for _, node := range d.Nodes {
		if state[node.ID] == 0 {
			if dfs(node.ID) {
				return true
			}
		}
	}

	return false
}

// GetTaskExecutionOrder returns tasks in dependency order for execution
func (d *DAG) GetTaskExecutionOrder() ([]string, error) {
	// Build graph of task dependencies only
	taskGraph := make(map[string][]string)
	taskInDegree := make(map[string]int)
	allTasks := make(map[string]bool)

	// Initialize all task nodes
	for _, node := range d.Nodes {
		if node.Type == "task" {
			taskID := node.ID
			allTasks[taskID] = true
			taskInDegree[taskID] = 0
		}
	}

	// Build dependency graph from edges
	for _, edge := range d.Edges {
		// Only consider edges between tasks for execution order
		if allTasks[edge.Source] && allTasks[edge.Target] {
			switch edge.Label {
			case "dependency", "defer":
				taskGraph[edge.Source] = append(taskGraph[edge.Source], edge.Target)
				taskInDegree[edge.Target]++
			}
		}
	}

	// Topological sort using Kahn's algorithm
	var queue []string
	var result []string

	// Find tasks with no dependencies
	for taskID := range allTasks {
		if taskInDegree[taskID] == 0 {
			queue = append(queue, taskID)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, dependent := range taskGraph[current] {
			taskInDegree[dependent]--
			if taskInDegree[dependent] == 0 {
				queue = append(queue, dependent)
			}
		}
	}

	if len(result) != len(allTasks) {
		return nil, fmt.Errorf("circular dependency detected in task dependencies")
	}

	return result, nil
}

// GetTaskStatistics returns statistics about the DAG
func (d *DAG) GetTaskStatistics() map[string]interface{} {
	stats := make(map[string]interface{})

	nodeTypes := make(map[string]int)
	edgeTypes := make(map[string]int)

	for _, node := range d.Nodes {
		nodeTypes[node.Type]++
	}

	for _, edge := range d.Edges {
		if edge.Label != "" {
			edgeTypes[edge.Label]++
		} else {
			edgeTypes["unlabeled"]++
		}
	}

	stats["total_nodes"] = len(d.Nodes)
	stats["total_edges"] = len(d.Edges)
	stats["node_types"] = nodeTypes
	stats["edge_types"] = edgeTypes

	// Count tasks with various properties
	scheduledTasks := 0
	manualTasks := 0
	matrixTasks := 0

	for _, node := range d.Nodes {
		if node.Type == "task" {
			if schedule, exists := node.Metadata["schedule"]; exists && schedule != nil {
				scheduledTasks++
			}
			if manual, exists := node.Metadata["manual"]; exists && manual.(bool) {
				manualTasks++
			}
			if matrix, exists := node.Metadata["matrix"]; exists && matrix != nil {
				matrixTasks++
			}
		}
	}

	stats["scheduled_tasks"] = scheduledTasks
	stats["manual_tasks"] = manualTasks
	stats["matrix_tasks"] = matrixTasks

	return stats
}

// ToJSON serializes the DAG to JSON
func (d *DAG) ToJSON() ([]byte, error) {
	return json.Marshal(&d)
}

// ValidateDAG performs comprehensive validation of the DAG
func (d *DAG) ValidateDAG() []string {
	var warnings []string

	// Check for cycles
	if d.hasCycle() {
		warnings = append(warnings, "Circular dependency detected in DAG")
	}

	// Check for orphaned nodes
	connectedNodes := make(map[string]bool)
	for _, edge := range d.Edges {
		connectedNodes[edge.Source] = true
		connectedNodes[edge.Target] = true
	}

	for _, node := range d.Nodes {
		if !connectedNodes[node.ID] && node.Type != "task" {
			warnings = append(warnings, fmt.Sprintf("Orphaned node detected: %s", node.ID))
		}
	}

	// Check for broken references
	nodeExists := make(map[string]bool)
	for _, node := range d.Nodes {
		nodeExists[node.ID] = true
	}

	for _, edge := range d.Edges {
		if !nodeExists[edge.Source] {
			warnings = append(warnings, fmt.Sprintf("Edge references non-existent source node: %s", edge.Source))
		}
		if !nodeExists[edge.Target] {
			warnings = append(warnings, fmt.Sprintf("Edge references non-existent target node: %s", edge.Target))
		}
	}

	return warnings
}
