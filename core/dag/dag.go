package dag

//
//import (
//	"encoding/json"
//	"fmt"
//	"paral/core/pipelines"
//	"paral/core/runtime"
//	"paral/core/variable"
//	"strings"
//)
//
//// DAG represents the Directed Acyclic Graph
//type DAG struct {
//	Nodes   []DAGNode `json:"nodes"`
//	Edges   []DAGEdge `json:"edges"`
//	Runtime *runtime.Runtime
//}
//
//// DAGNode represents a node in the DAG (task or pipeline)
//type DAGNode struct {
//	ID       string                 `json:"id"`
//	Label    string                 `json:"label"`
//	Type     string                 `json:"type"` // "task" or "pipeline"
//	Metadata map[string]interface{} `json:"metadata"`
//}
//
//// DAGEdge represents an edge in the DAG
//type DAGEdge struct {
//	Source string `json:"source"`
//	Target string `json:"target"`
//	Label  string `json:"label,omitempty"`
//}
//
//// NewDAG creates a new DAG instance
//func NewDAG(r *runtime.Runtime) *DAG {
//	return &DAG{
//		Nodes:   []DAGNode{},
//		Edges:   []DAGEdge{},
//		Runtime: r,
//	}
//}
//
//// convertToStringMatrix converts various input types to [][]string for DAG processing
//func (d *DAG) convertToStringMatrix(input interface{}) ([][]string, error) {
//	switch v := input.(type) {
//	case int:
//		// Number: generate range [0, n-1]
//		result := make([][]string, v)
//		for i := 0; i < v; i++ {
//			result[i] = []string{fmt.Sprintf("%d", i)}
//		}
//		return result, nil
//	case []interface{}:
//		// List: treat as a single-column matrix
//		result := make([][]string, len(v))
//		for i, val := range v {
//			if str, ok := val.(string); ok {
//				result[i] = []string{str}
//			} else {
//				return nil, fmt.Errorf("list value %v is not a string", val)
//			}
//		}
//		return result, nil
//	case [][]interface{}:
//		// Matrix: convert to [][]string
//		result := make([][]string, 0, len(v))
//		for _, row := range v {
//			stringRow := make([]string, 0, len(row))
//			for _, val := range row {
//				if str, ok := val.(string); ok {
//					stringRow = append(stringRow, str)
//				} else {
//					return nil, fmt.Errorf("matrix value %v is not a string", val)
//				}
//			}
//			result = append(result, stringRow)
//		}
//		return result, nil
//	default:
//		return nil, fmt.Errorf("unsupported @for value type: %T", v)
//	}
//}
//
//// evaluateCondition evaluates @if conditions
//func (d *DAG) evaluateCondition(condition string, stash map[string]interface{}, matrix []string) (bool, error) {
//	// Resolve @value(n) placeholders
//	resolved := d.resolveCommand(condition, matrix)
//	// Simple evaluation for == and !=
//	if strings.Contains(resolved, "==") {
//		parts := strings.SplitN(resolved, "==", 2)
//		if len(parts) != 2 {
//			return false, fmt.Errorf("invalid condition: %s", condition)
//		}
//		left, right := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
//		return left == right, nil
//	}
//	if strings.Contains(resolved, "!=") {
//		parts := strings.SplitN(resolved, "!=", 2)
//		if len(parts) != 2 {
//			return false, fmt.Errorf("invalid condition: %s", condition)
//		}
//		left, right := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
//		return left != right, nil
//	}
//	if strings.HasPrefix(resolved, "@catch") {
//		return stash["error"] != nil, nil
//	}
//	if strings.HasPrefix(resolved, "@iflen") {
//		// Extract the value to check length
//		start := strings.Index(resolved, "(")
//		end := strings.LastIndex(resolved, ")")
//		if start == -1 || end == -1 {
//			return false, fmt.Errorf("invalid @iflen: %s", condition)
//		}
//		value := strings.TrimSpace(resolved[start+1 : end])
//		if v, ok := stash[value]; ok {
//			switch val := v.(type) {
//			case string:
//				return len(val) > 0, nil
//			case []interface{}:
//				return len(val) > 0, nil
//			}
//		}
//		return false, nil
//	}
//	return false, fmt.Errorf("unsupported condition: %s", condition)
//}
//
//func (d *DAG) findMatrixByName(name string) *variable.Variable {
//	for _, _var := range d.Runtime.Vars {
//		if _var.Name == name {
//			_, ok := _var.AsMatrix()
//			if ok {
//				return _var
//			}
//		}
//	}
//	return nil
//}
//
//// findDirective finds a directive by type in a list of JobDirectives
//func (d *DAG) findDirective(directives []*runtime.Directive, directiveType string) *runtime.Directive {
//	for _, d := range directives {
//		if d.Type == directiveType {
//			return d
//		}
//	}
//	return nil
//}
//
//// formatJobDirectives formats task directives for metadata
//func (d *DAG) formatJobDirectives(directives []*runtime.Directive) []map[string]interface{} {
//	var result []map[string]interface{}
//	for _, d := range directives {
//		result = append(result, map[string]interface{}{
//			"type":   d.Type,
//			"params": d.Params,
//		})
//	}
//	return result
//}
//
//// Fixed formatFunctions to handle the new Command structure with nested Functions and Conditions
//func (d *DAG) formatFunctions(pipeline *pipelines.Command) []map[string]interface{} {
//	var result []map[string]interface{}
//
//	// Handle Functions
//	for _, f := range pipeline.Functions {
//		result = append(result, map[string]interface{}{
//			"type": f.Type,
//			"args": f.Args,
//		})
//	}
//
//	return result
//}
//
//// Helper function to format Command for metadata
//func (d *DAG) formatCommandForMetadata(cmd *pipelines.Command) map[string]interface{} {
//	cmdMap := make(map[string]interface{})
//
//	if cmd.RawText != "" {
//		cmdMap["raw_text"] = cmd.RawText
//	}
//
//	if len(cmd.Functions) > 0 {
//		var functions []map[string]interface{}
//		for _, f := range cmd.Functions {
//			functions = append(functions, map[string]interface{}{
//				"type": f.Type,
//				"args": f.Args,
//			})
//		}
//		cmdMap["functions"] = functions
//	}
//
//	return cmdMap
//}
//
//// resolveCommand replaces @value(n) placeholders with matrix values
//func (d *DAG) resolveCommand(cmd string, matrix []string) string {
//	for i, val := range matrix {
//		cmd = strings.ReplaceAll(cmd, fmt.Sprintf("@value(%d)", i+1), val)
//	}
//	return cmd
//}
//
//// formatMatrixLabel formats matrix values for node labels
//func (d *DAG) formatMatrixLabel(matrix []string) string {
//	if len(matrix) == 0 {
//		return ""
//	}
//	return fmt.Sprintf(" (%s)", strings.Join(matrix, ", "))
//}
//
//// Fixedd.extractEnvironments to work with new Command structure
//func (d *DAG) extractEnvironments(pipeline *pipelines.Command, matrix []string) map[string]string {
//	envs := make(map[string]string)
//
//	// Extract from Functions
//	for _, f := range pipeline.Functions {
//		if f.Type == "setenv" && len(f.Args) == 2 {
//			if key, ok := f.Args[0].(string); ok {
//				if val, ok := f.Args[1].(string); ok {
//					envs[key] = d.resolveCommand(val, matrix)
//				}
//			}
//		}
//	}
//
//	return envs
//}
//
//// Fixed extractCondition to work with new Command structure
//func (d *DAG) extractCondition(pipeline *pipelines.Command) string {
//	// Extract from Functions
//	for _, f := range pipeline.Functions {
//		if f.Type == "if" || f.Type == "iflen" {
//			if len(f.Args) > 0 {
//				if cond, ok := f.Args[0].(string); ok {
//					return cond
//				}
//			}
//		}
//	}
//
//	return ""
//}
//
//// Fixed pipeline processing section in GenerateDAG method
//func (d *DAG) processCommandForDAG(pipeline *pipelines.Command, jobName string, matrixIdx, cmdIdx int, matrix []string, stash map[string]interface{}, dag *DAG, jobInstanceID string, seenNodes map[string]bool) (string, error) {
//	// Determine pipeline text to display
//	var displayText string
//	if pipeline.RawText != "" {
//		displayText = d.resolveCommand(pipeline.RawText, matrix)
//	} else if len(pipeline.Functions) > 0 {
//		displayText = fmt.Sprintf("@%s", pipeline.Functions[0].Type)
//	} else {
//		displayText = "<empty pipeline>"
//	}
//
//	// Create pipeline node
//	pipelineNodeID := fmt.Sprintf("cmd_%s_%d_%d", jobName, matrixIdx, cmdIdx)
//	if seenNodes[pipelineNodeID] {
//		return "", fmt.Errorf("duplicate pipeline ID %s detected", pipelineNodeID)
//	}
//	seenNodes[pipelineNodeID] = true
//
//	pipelineMetadata := map[string]interface{}{
//		"pipeline":     displayText,
//		"raw_pipeline": pipeline.RawText,
//		"functions":    d.formatFunctions(pipeline),
//	}
//
//	if len(matrix) > 0 {
//		pipelineMetadata["matrix"] = matrix
//	}
//
//	if envs := d.extractEnvironments(pipeline, matrix); len(envs) > 0 {
//		pipelineMetadata["environment"] = envs
//	}
//
//	//if condition := extractCondition(pipeline); condition != "" {
//	//	if ok, err := evaluateCondition(condition, stash, matrix); err == nil && ok {
//	//		pipelineMetadata["condition"] = condition
//	//	} else if err != nil {
//	//		r.Reporter.Warn(fmt.Sprintf("Invalid condition %s: %v", condition, err), condition.)
//	//	}
//	//}
//
//	// Process Functions for metadata
//	for _, f := range pipeline.Functions {
//		if f.Type == "try" && len(f.Args) > 0 {
//			pipelineMetadata["retries"] = f.Args[0]
//		}
//		if f.Type == "stash" && len(f.Args) == 2 {
//			if key, ok := f.Args[0].(string); ok {
//				stash[key] = f.Args[1]
//			}
//		}
//	}
//
//	dag.Nodes = append(dag.Nodes, DAGNode{
//		ID:       pipelineNodeID,
//		Label:    displayText,
//		Type:     "pipeline",
//		Metadata: pipelineMetadata,
//	})
//
//	return pipelineNodeID, nil
//}
//
//// GenerateDAG creates a detailed DAG from the Runtime struct
//func (d *DAG) GenerateDAG() *DAG {
//	// Map to track task nodes for trigger/dependency edges
//	jobNodeIDs := make(map[string][]string) // Maps task name to list of instance IDs
//	seenNodes := make(map[string]bool)      // Tracks nodes to detect cycles
//	stash := make(map[string]interface{})   // Task-level stash for condition evaluation
//
//	// Process each task
//	for _, task := range d.Runtime.Tasks {
//		// Resolve matrix values for @for directive
//		matrixValues := [][]string{{}} // Default: single instance if no @for
//		if forDirective := d.findDirective(task.Directives, "for"); forDirective != nil && len(forDirective.Params) > 0 {
//			if varName, ok := forDirective.Params[0].(string); ok {
//				// Check if varName is a variable or matrix
//				if varValue := d.Runtime.GetVariable(varName); varValue != nil {
//					// Variable: number or list
//					var err error
//					matrixValues, err = d.convertToStringMatrix(varValue.Value)
//					if err != nil {
//						d.Runtime.Reporter.Warn(fmt.Sprintf("Invalid @for variable %s: %v", varName, err), &task.Metadata)
//						continue
//					}
//				} else if matrixValue := d.findMatrixByName(varName); matrixValue != nil {
//					// Matrix
//					var err error
//					matrixValues, err = d.convertToStringMatrix(matrixValue.Value)
//					if err != nil {
//						d.Runtime.Reporter.Warn(fmt.Sprintf("Invalid @for matrix %s: %v", varName, err), &task.Metadata)
//						continue
//					}
//				} else {
//					d.Runtime.Reporter.Warn(fmt.Sprintf("@for target %s not found", varName), &task.Metadata)
//					continue
//				}
//			} else {
//				d.Runtime.Reporter.Warn("@for requires a string variable or matrix name", &task.Metadata)
//				continue
//			}
//		}
//
//		// Create a task node for each matrix combination
//		for matrixIdx, matrix := range matrixValues {
//			jobInstanceID := fmt.Sprintf("job_%s_%d", task.Name, matrixIdx)
//			if len(matrixValues) == 1 && matrixValues[0] == nil {
//				jobInstanceID = fmt.Sprintf("job_%s", task.Name)
//			}
//			if seenNodes[jobInstanceID] {
//				d.Runtime.Reporter.Warn(fmt.Sprintf("Duplicate task instance ID %s detected", jobInstanceID), &task.Metadata)
//				continue
//			}
//			seenNodes[jobInstanceID] = true
//			jobNodeIDs[task.Name] = append(jobNodeIDs[task.Name], jobInstanceID)
//
//			// Task metadata
//			jobMetadata := map[string]interface{}{
//				"filename":   task.Filename,
//				"directives": d.formatJobDirectives(task.Directives),
//			}
//
//			if len(matrix) > 0 {
//				jobMetadata["matrix"] = matrix
//			}
//
//			envMap := make(map[int]map[string]string)
//			condMap := make(map[int]string)
//			jobMetadata["environment"] = envMap
//			jobMetadata["condition"] = condMap
//
//			for i, pipeline := range task.Commands {
//				if envs := d.extractEnvironments(pipeline, matrix); len(envs) > 0 {
//					envMap[i] = envs
//				}
//
//				if condition := d.extractCondition(pipeline); condition != "" {
//					if ok, err := d.evaluateCondition(condition, stash, matrix); err != nil {
//						d.Runtime.Reporter.Warn(fmt.Sprintf("Invalid condition %s: %v", condition, err), &task.Metadata)
//					} else if ok {
//						condMap[i] = condition
//					}
//				}
//			}
//
//			if schedule := d.findDirective(task.Directives, "schedule"); schedule != nil && len(schedule.Params) > 0 {
//				if cron, ok := schedule.Params[0].(string); ok {
//					jobMetadata["schedule"] = cron
//				}
//			}
//			if manual := d.findDirective(task.Directives, "manual"); manual != nil {
//				manualValue := true
//				if len(manual.Params) > 0 {
//					if val, ok := manual.Params[0].(bool); ok {
//						manualValue = val
//					}
//				}
//				jobMetadata["manual"] = manualValue
//			}
//
//			// Add task node
//			d.Nodes = append(d.Nodes, DAGNode{
//				ID:       jobInstanceID,
//				Label:    fmt.Sprintf("%s%s", task.Name, d.formatMatrixLabel(matrix)),
//				Type:     "task",
//				Metadata: jobMetadata,
//			})
//
//			// Process pipelines for this task instance
//			var prevCommandID string
//			for cmdIdx, pipeline := range task.Commands {
//				// Resolve pipeline text with matrix values
//				resolvedCMD := d.resolveCommand(pipeline.RawText, matrix)
//
//				// Create pipeline node
//				pipelineNodeID := fmt.Sprintf("cmd_%s_%d_%d", task.Name, matrixIdx, cmdIdx)
//				if seenNodes[pipelineNodeID] {
//					d.Runtime.Reporter.Warn(fmt.Sprintf("Duplicate pipeline ID %s detected", pipelineNodeID), &task.Metadata)
//					continue
//				}
//				seenNodes[pipelineNodeID] = true
//
//				pipelineMetadata := map[string]interface{}{
//					"pipeline":     resolvedCMD,
//					"raw_pipeline": pipeline.RawText,
//					"functions":    d.formatFunctions(pipeline),
//				}
//				if len(matrix) > 0 {
//					pipelineMetadata["matrix"] = matrix
//				}
//				if envs := d.extractEnvironments(pipeline, matrix); len(envs) > 0 {
//					pipelineMetadata["environment"] = envs
//				}
//				if condition := d.extractCondition(pipeline); condition != "" {
//					if ok, err := d.evaluateCondition(condition, stash, matrix); err == nil && ok {
//						pipelineMetadata["condition"] = condition
//					} else if err != nil {
//						d.Runtime.Reporter.Warn(fmt.Sprintf("Invalid condition %s: %v", condition, err), &task.Metadata)
//					}
//				}
//				for _, dir := range pipeline.Functions {
//					if dir.Type == "try" && len(dir.Args) > 0 {
//						pipelineMetadata["retries"] = dir.Args[0]
//					}
//					if dir.Type == "stash" && len(dir.Args) == 2 {
//						if key, ok := dir.Args[0].(string); ok {
//							stash[key] = dir.Args[1]
//						}
//					}
//				}
//
//				d.Nodes = append(d.Nodes, DAGNode{
//					ID:       pipelineNodeID,
//					Label:    resolvedCMD,
//					Type:     "pipeline",
//					Metadata: pipelineMetadata,
//				})
//
//				// Add edge from task to its first pipeline
//				if cmdIdx == 0 {
//					d.Edges = append(d.Edges, DAGEdge{
//						Source: jobInstanceID,
//						Target: pipelineNodeID,
//						Label:  "",
//					})
//				}
//
//				// Add edge between consecutive pipelines
//				if prevCommandID != "" {
//					d.Edges = append(d.Edges, DAGEdge{
//						Source: prevCommandID,
//						Target: pipelineNodeID,
//						Label:  "next",
//					})
//				}
//				prevCommandID = pipelineNodeID
//
//				// Add edges for trigger and depend directives
//				for _, directive := range pipeline.Functions {
//					if directive.Type == "trigger" || directive.Type == "depend" {
//						if len(directive.Args) > 0 {
//							if targetJob, ok := directive.Args[0].(string); ok {
//								if targetJobIDs, exists := jobNodeIDs[targetJob]; exists {
//									for _, targetJobID := range targetJobIDs {
//										d.Edges = append(d.Edges, DAGEdge{
//											Source: jobInstanceID,
//											Target: targetJobID,
//											Label:  directive.Type,
//										})
//									}
//								} else {
//									d.Runtime.Reporter.Warn(fmt.Sprintf("%s target %s not found", directive.Type, targetJob), &task.Metadata)
//								}
//							}
//						}
//					}
//				}
//			}
//
//			// Add edges for task-level trigger and depend directives
//			for _, directive := range task.Directives {
//				if directive.Type == "trigger" || directive.Type == "depend" {
//					if len(directive.Params) > 0 {
//						if targetJob, ok := directive.Params[0].(string); ok {
//							if targetJobIDs, exists := jobNodeIDs[targetJob]; exists {
//								for _, targetJobID := range targetJobIDs {
//									d.Edges = append(d.Edges, DAGEdge{
//										Source: jobInstanceID,
//										Target: targetJobID,
//										Label:  directive.Type,
//									})
//								}
//							} else {
//								d.Runtime.Reporter.Warn(fmt.Sprintf("%s target %s not found", directive.Type, targetJob), &task.Metadata)
//							}
//						}
//					}
//				}
//				if directive.Type == "defer" {
//					if len(directive.Params) > 0 {
//						if targetJob, ok := directive.Params[0].(string); ok {
//							if targetJobIDs, exists := jobNodeIDs[targetJob]; exists {
//								for _, targetJobID := range targetJobIDs {
//									d.Edges = append(d.Edges, DAGEdge{
//										Source: jobInstanceID,
//										Target: targetJobID,
//										Label:  "defer",
//									})
//								}
//							} else {
//								d.Runtime.Reporter.Warn(fmt.Sprintf("defer target %s not found", targetJob), &task.Metadata)
//							}
//						}
//					}
//				}
//			}
//		}
//	}
//
//	// Validate for cycles
//	if hasCycle(d) {
//		//c.Reporter.Warn("Potential cycle detected in DAG", "", 0, 0)
//	}
//
//	return d
//}
//
//// ToJSON serializes the DAG to JSON
//func (d *DAG) ToJSON() (string, error) {
//	data, err := json.MarshalIndent(d, "", "  ")
//	if err != nil {
//		return "", err
//	}
//	return string(data), nil
//}
//
//// hasCycle checks for cycles in the DAG (placeholder implementation)
//func hasCycle(dag *DAG) bool {
//	// Implement cycle detection if needed
//	return false
//}
