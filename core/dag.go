package core

import (
	"encoding/json"
	"fmt"
	"strings"
)

// DAGNode represents a node in the DAG (job or command)
type DAGNode struct {
	ID       string                 `json:"id"`
	Label    string                 `json:"label"`
	Type     string                 `json:"type"` // "job" or "command"
	Metadata map[string]interface{} `json:"metadata"`
}

// DAGEdge represents an edge in the DAG
type DAGEdge struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Label  string `json:"label,omitempty"`
}

// DAG represents the Directed Acyclic Graph
type DAG struct {
	Nodes []DAGNode `json:"nodes"`
	Edges []DAGEdge `json:"edges"`
}

// NewDAG creates a new DAG instance
func NewDAG() *DAG {
	return &DAG{
		Nodes: []DAGNode{},
		Edges: []DAGEdge{},
	}
}

// convertToStringMatrix converts various input types to [][]string for DAG processing
func convertToStringMatrix(input interface{}) ([][]string, error) {
	switch v := input.(type) {
	case int:
		// Number: generate range [0, n-1]
		result := make([][]string, v)
		for i := 0; i < v; i++ {
			result[i] = []string{fmt.Sprintf("%d", i)}
		}
		return result, nil
	case []interface{}:
		// List: treat as a single-column matrix
		result := make([][]string, len(v))
		for i, val := range v {
			if str, ok := val.(string); ok {
				result[i] = []string{str}
			} else {
				return nil, fmt.Errorf("list value %v is not a string", val)
			}
		}
		return result, nil
	case [][]interface{}:
		// Matrix: convert to [][]string
		result := make([][]string, 0, len(v))
		for _, row := range v {
			stringRow := make([]string, 0, len(row))
			for _, val := range row {
				if str, ok := val.(string); ok {
					stringRow = append(stringRow, str)
				} else {
					return nil, fmt.Errorf("matrix value %v is not a string", val)
				}
			}
			result = append(result, stringRow)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported @for value type: %T", v)
	}
}

// evaluateCondition evaluates @if conditions
func evaluateCondition(condition string, stash map[string]interface{}, matrix []string) (bool, error) {
	// Resolve @value(n) placeholders
	resolved := resolveCommand(condition, matrix)
	// Simple evaluation for == and !=
	if strings.Contains(resolved, "==") {
		parts := strings.SplitN(resolved, "==", 2)
		if len(parts) != 2 {
			return false, fmt.Errorf("invalid condition: %s", condition)
		}
		left, right := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		return left == right, nil
	}
	if strings.Contains(resolved, "!=") {
		parts := strings.SplitN(resolved, "!=", 2)
		if len(parts) != 2 {
			return false, fmt.Errorf("invalid condition: %s", condition)
		}
		left, right := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		return left != right, nil
	}
	if strings.HasPrefix(resolved, "@catch") {
		return stash["error"] != nil, nil
	}
	if strings.HasPrefix(resolved, "@iflen") {
		// Extract the value to check length
		start := strings.Index(resolved, "(")
		end := strings.LastIndex(resolved, ")")
		if start == -1 || end == -1 {
			return false, fmt.Errorf("invalid @iflen: %s", condition)
		}
		value := strings.TrimSpace(resolved[start+1 : end])
		if v, ok := stash[value]; ok {
			switch val := v.(type) {
			case string:
				return len(val) > 0, nil
			case []interface{}:
				return len(val) > 0, nil
			}
		}
		return false, nil
	}
	return false, fmt.Errorf("unsupported condition: %s", condition)
}

func (c *Core) findVarByName(name string) *Var {
	for _, _var := range c.Vars {
		if _var.Name == name {
			return &_var
		}
	}
	return nil
}

func (c *Core) findMatrixByName(name string) *Var {
	for _, _var := range c.Vars {
		if _var.Name == name {
			_, ok := _var.AsMatrix()
			if ok {
				return &_var
			}
		}
	}
	return nil
}

// findDirective finds a directive by type in a list of JobDirectives
func findDirective(directives []Directive, directiveType string) *Directive {
	for _, d := range directives {
		if d.Type == directiveType {
			return &d
		}
	}
	return nil
}

// formatJobDirectives formats job directives for metadata
func formatJobDirectives(directives []Directive) []map[string]interface{} {
	var result []map[string]interface{}
	for _, d := range directives {
		result = append(result, map[string]interface{}{
			"type":   d.Type,
			"params": d.Params,
		})
	}
	return result
}

// Fixed formatFunctions to handle the new Command structure with nested Functions and Conditions
func formatFunctions(command Command) []map[string]interface{} {
	var result []map[string]interface{}

	// Handle Functions
	for _, f := range command.Functions {
		result = append(result, map[string]interface{}{
			"type": f.Type,
			"args": f.Args,
		})
	}

	// Handle Conditions recursively
	for _, c := range command.Conditions {
		conditionMap := map[string]interface{}{
			"type": c.Type,
		}

		// Process Expression as nested Command
		if c.Expression.RawText != "" || len(c.Expression.Functions) > 0 || len(c.Expression.Conditions) > 0 {
			conditionMap["expression"] = formatCommandForMetadata(c.Expression)
		}

		// Process ThenCommand as nested Command
		if c.ThenCommand.RawText != "" || len(c.ThenCommand.Functions) > 0 || len(c.ThenCommand.Conditions) > 0 {
			conditionMap["then_command"] = formatCommandForMetadata(c.ThenCommand)
		}

		// Process ThenCommands as array of nested Commands
		if len(c.ThenCommands) > 0 {
			var thenCommands []map[string]interface{}
			for _, thenCmd := range c.ThenCommands {
				thenCommands = append(thenCommands, formatCommandForMetadata(thenCmd))
			}
			conditionMap["then_commands"] = thenCommands
		}

		// Process ElseCommand as nested Command
		if c.ElseCommand.RawText != "" || len(c.ElseCommand.Functions) > 0 || len(c.ElseCommand.Conditions) > 0 {
			conditionMap["else_command"] = formatCommandForMetadata(c.ElseCommand)
		}

		// Process ElseCommands as array of nested Commands
		if len(c.ElseCommands) > 0 {
			var elseCommands []map[string]interface{}
			for _, elseCmd := range c.ElseCommands {
				elseCommands = append(elseCommands, formatCommandForMetadata(elseCmd))
			}
			conditionMap["else_commands"] = elseCommands
		}

		result = append(result, conditionMap)
	}

	return result
}

// Helper function to format Command for metadata
func formatCommandForMetadata(cmd Command) map[string]interface{} {
	cmdMap := make(map[string]interface{})

	if cmd.RawText != "" {
		cmdMap["raw_text"] = cmd.RawText
	}

	if len(cmd.Functions) > 0 {
		var functions []map[string]interface{}
		for _, f := range cmd.Functions {
			functions = append(functions, map[string]interface{}{
				"type": f.Type,
				"args": f.Args,
			})
		}
		cmdMap["functions"] = functions
	}

	if len(cmd.Conditions) > 0 {
		var conditions []map[string]interface{}
		for _, c := range cmd.Conditions {
			conditionMap := map[string]interface{}{
				"type": c.Type,
			}

			if c.Expression.RawText != "" || len(c.Expression.Functions) > 0 || len(c.Expression.Conditions) > 0 {
				conditionMap["expression"] = formatCommandForMetadata(c.Expression)
			}

			if c.ThenCommand.RawText != "" || len(c.ThenCommand.Functions) > 0 || len(c.ThenCommand.Conditions) > 0 {
				conditionMap["then_command"] = formatCommandForMetadata(c.ThenCommand)
			}

			if len(c.ThenCommands) > 0 {
				var thenCommands []map[string]interface{}
				for _, thenCmd := range c.ThenCommands {
					thenCommands = append(thenCommands, formatCommandForMetadata(thenCmd))
				}
				conditionMap["then_commands"] = thenCommands
			}

			conditions = append(conditions, conditionMap)
		}
		cmdMap["conditions"] = conditions
	}

	return cmdMap
}

// resolveCommand replaces @value(n) placeholders with matrix values
func resolveCommand(cmd string, matrix []string) string {
	for i, val := range matrix {
		cmd = strings.ReplaceAll(cmd, fmt.Sprintf("@value(%d)", i+1), val)
	}
	return cmd
}

// formatMatrixLabel formats matrix values for node labels
func formatMatrixLabel(matrix []string) string {
	if len(matrix) == 0 {
		return ""
	}
	return fmt.Sprintf(" (%s)", strings.Join(matrix, ", "))
}

// Fixed extractEnvironments to work with new Command structure
func extractEnvironments(command Command, matrix []string) map[string]string {
	envs := make(map[string]string)

	// Extract from Functions
	for _, f := range command.Functions {
		if f.Type == "setenv" && len(f.Args) == 2 {
			if key, ok := f.Args[0].(string); ok {
				if val, ok := f.Args[1].(string); ok {
					envs[key] = resolveCommand(val, matrix)
				}
			}
		}
	}

	// Extract from Conditions recursively
	for _, c := range command.Conditions {
		// Check expression
		if exprEnvs := extractEnvironments(c.Expression, matrix); len(exprEnvs) > 0 {
			for k, v := range exprEnvs {
				envs[k] = v
			}
		}

		// Check then command
		if thenEnvs := extractEnvironments(c.ThenCommand, matrix); len(thenEnvs) > 0 {
			for k, v := range thenEnvs {
				envs[k] = v
			}
		}

		// Check then commands
		for _, thenCmd := range c.ThenCommands {
			if thenEnvs := extractEnvironments(thenCmd, matrix); len(thenEnvs) > 0 {
				for k, v := range thenEnvs {
					envs[k] = v
				}
			}
		}

		// Check else command
		if elseEnvs := extractEnvironments(c.ElseCommand, matrix); len(elseEnvs) > 0 {
			for k, v := range elseEnvs {
				envs[k] = v
			}
		}

		// Check else commands
		for _, elseCmd := range c.ElseCommands {
			if elseEnvs := extractEnvironments(elseCmd, matrix); len(elseEnvs) > 0 {
				for k, v := range elseEnvs {
					envs[k] = v
				}
			}
		}
	}

	return envs
}

// Fixed extractCondition to work with new Command structure
func extractCondition(command Command) string {
	// Extract from Functions
	for _, f := range command.Functions {
		if f.Type == "if" || f.Type == "iflen" {
			if len(f.Args) > 0 {
				if cond, ok := f.Args[0].(string); ok {
					return cond
				}
			}
		}
	}

	// Extract from Conditions
	for _, c := range command.Conditions {
		if c.Type == "if" || c.Type == "iflen" {
			// Return the raw text of the expression if available
			if c.Expression.RawText != "" {
				return c.Expression.RawText
			}
			// Or try to extract condition from nested structure
			if nestedCond := extractCondition(c.Expression); nestedCond != "" {
				return nestedCond
			}
		}
	}

	return ""
}

// Fixed command processing section in GenerateDAG method
func (c *Core) processCommandForDAG(command Command, jobName string, matrixIdx, cmdIdx int, matrix []string, stash map[string]interface{}, dag *DAG, jobInstanceID string, seenNodes map[string]bool) (string, error) {
	// Determine command text to display
	var displayText string
	if command.RawText != "" {
		displayText = resolveCommand(command.RawText, matrix)
	} else if len(command.Functions) > 0 {
		displayText = fmt.Sprintf("@%s", command.Functions[0].Type)
	} else if len(command.Conditions) > 0 {
		displayText = fmt.Sprintf("@%s", command.Conditions[0].Type)
	} else {
		displayText = "<empty command>"
	}

	// Create command node
	commandNodeID := fmt.Sprintf("cmd_%s_%d_%d", jobName, matrixIdx, cmdIdx)
	if seenNodes[commandNodeID] {
		return "", fmt.Errorf("duplicate command ID %s detected", commandNodeID)
	}
	seenNodes[commandNodeID] = true

	commandMetadata := map[string]interface{}{
		"command":     displayText,
		"raw_command": command.RawText,
		"functions":   formatFunctions(command),
	}

	if len(matrix) > 0 {
		commandMetadata["matrix"] = matrix
	}

	if envs := extractEnvironments(command, matrix); len(envs) > 0 {
		commandMetadata["environment"] = envs
	}

	if condition := extractCondition(command); condition != "" {
		if ok, err := evaluateCondition(condition, stash, matrix); err == nil && ok {
			commandMetadata["condition"] = condition
		} else if err != nil {
			Warn(fmt.Sprintf("Invalid condition %s: %v", condition, err), "", 0, 0)
		}
	}

	// Process Functions for metadata
	for _, f := range command.Functions {
		if f.Type == "try" && len(f.Args) > 0 {
			commandMetadata["retries"] = f.Args[0]
		}
		if f.Type == "stash" && len(f.Args) == 2 {
			if key, ok := f.Args[0].(string); ok {
				stash[key] = f.Args[1]
			}
		}
	}

	dag.Nodes = append(dag.Nodes, DAGNode{
		ID:       commandNodeID,
		Label:    displayText,
		Type:     "command",
		Metadata: commandMetadata,
	})

	return commandNodeID, nil
}

// GenerateDAG creates a detailed DAG from the Core struct
func (c *Core) GenerateDAG() *DAG {
	dag := NewDAG()

	// Map to track job nodes for trigger/dependency edges
	jobNodeIDs := make(map[string][]string) // Maps job name to list of instance IDs
	seenNodes := make(map[string]bool)      // Tracks nodes to detect cycles
	stash := make(map[string]interface{})   // Job-level stash for condition evaluation

	// Process each job
	for _, job := range c.Jobs {
		// Resolve matrix values for @for directive
		matrixValues := [][]string{{}} // Default: single instance if no @for
		if forDirective := findDirective(job.Directives, "for"); forDirective != nil && len(forDirective.Params) > 0 {
			if varName, ok := forDirective.Params[0].(string); ok {
				// Check if varName is a variable or matrix
				if varValue := c.findVarByName(varName); varValue != nil {
					// Variable: number or list
					var err error
					matrixValues, err = convertToStringMatrix(varValue.Value)
					if err != nil {
						Warn(fmt.Sprintf("Invalid @for variable %s: %v", varName, err), job.Filename, 0, 0)
						continue
					}
				} else if matrixValue := c.findMatrixByName(varName); matrixValue != nil {
					// Matrix
					var err error
					matrixValues, err = convertToStringMatrix(matrixValue.Value)
					if err != nil {
						Warn(fmt.Sprintf("Invalid @for matrix %s: %v", varName, err), job.Filename, 0, 0)
						continue
					}
				} else {
					Warn(fmt.Sprintf("@for target %s not found", varName), job.Filename, 0, 0)
					continue
				}
			} else {
				Warn("@for requires a string variable or matrix name", job.Filename, 0, 0)
				continue
			}
		}

		// Create a job node for each matrix combination
		for matrixIdx, matrix := range matrixValues {
			jobInstanceID := fmt.Sprintf("job_%s_%d", job.Name, matrixIdx)
			if len(matrixValues) == 1 && matrixValues[0] == nil {
				jobInstanceID = fmt.Sprintf("job_%s", job.Name)
			}
			if seenNodes[jobInstanceID] {
				Warn(fmt.Sprintf("Duplicate job instance ID %s detected", jobInstanceID), job.Filename, 0, 0)
				continue
			}
			seenNodes[jobInstanceID] = true
			jobNodeIDs[job.Name] = append(jobNodeIDs[job.Name], jobInstanceID)

			// Job metadata
			jobMetadata := map[string]interface{}{
				"filename":   job.Filename,
				"directives": formatJobDirectives(job.Directives),
			}

			if len(matrix) > 0 {
				jobMetadata["matrix"] = matrix
			}

			envMap := make(map[int]map[string]string)
			condMap := make(map[int]string)
			jobMetadata["environment"] = envMap
			jobMetadata["condition"] = condMap

			for i, command := range job.Commands {
				if envs := extractEnvironments(command, matrix); len(envs) > 0 {
					envMap[i] = envs
				}

				if condition := extractCondition(command); condition != "" {
					if ok, err := evaluateCondition(condition, stash, matrix); err != nil {
						Warn(fmt.Sprintf("Invalid condition %s: %v", condition, err), job.Filename, 0, 0)
					} else if ok {
						condMap[i] = condition
					}
				}
			}

			if schedule := findDirective(job.Directives, "schedule"); schedule != nil && len(schedule.Params) > 0 {
				if cron, ok := schedule.Params[0].(string); ok {
					jobMetadata["schedule"] = cron
				}
			}
			if manual := findDirective(job.Directives, "manual"); manual != nil {
				manualValue := true
				if len(manual.Params) > 0 {
					if val, ok := manual.Params[0].(bool); ok {
						manualValue = val
					}
				}
				jobMetadata["manual"] = manualValue
			}

			// Add job node
			dag.Nodes = append(dag.Nodes, DAGNode{
				ID:       jobInstanceID,
				Label:    fmt.Sprintf("%s%s", job.Name, formatMatrixLabel(matrix)),
				Type:     "job",
				Metadata: jobMetadata,
			})

			// Process commands for this job instance
			var prevCommandID string
			for cmdIdx, command := range job.Commands {
				// Resolve command text with matrix values
				resolvedCMD := resolveCommand(command.RawText, matrix)

				// Create command node
				commandNodeID := fmt.Sprintf("cmd_%s_%d_%d", job.Name, matrixIdx, cmdIdx)
				if seenNodes[commandNodeID] {
					Warn(fmt.Sprintf("Duplicate command ID %s detected", commandNodeID), job.Filename, 0, 0)
					continue
				}
				seenNodes[commandNodeID] = true

				commandMetadata := map[string]interface{}{
					"command":     resolvedCMD,
					"raw_command": command.RawText,
					"functions":   formatFunctions(command),
				}
				if len(matrix) > 0 {
					commandMetadata["matrix"] = matrix
				}
				if envs := extractEnvironments(command, matrix); len(envs) > 0 {
					commandMetadata["environment"] = envs
				}
				if condition := extractCondition(command); condition != "" {
					if ok, err := evaluateCondition(condition, stash, matrix); err == nil && ok {
						commandMetadata["condition"] = condition
					} else if err != nil {
						Warn(fmt.Sprintf("Invalid condition %s: %v", condition, err), job.Filename, 0, 0)
					}
				}
				for _, dir := range command.Functions {
					if dir.Type == "try" && len(dir.Args) > 0 {
						commandMetadata["retries"] = dir.Args[0]
					}
					if dir.Type == "stash" && len(dir.Args) == 2 {
						if key, ok := dir.Args[0].(string); ok {
							stash[key] = dir.Args[1]
						}
					}
				}

				dag.Nodes = append(dag.Nodes, DAGNode{
					ID:       commandNodeID,
					Label:    resolvedCMD,
					Type:     "command",
					Metadata: commandMetadata,
				})

				// Add edge from job to its first command
				if cmdIdx == 0 {
					dag.Edges = append(dag.Edges, DAGEdge{
						Source: jobInstanceID,
						Target: commandNodeID,
						Label:  "",
					})
				}

				// Add edge between consecutive commands
				if prevCommandID != "" {
					dag.Edges = append(dag.Edges, DAGEdge{
						Source: prevCommandID,
						Target: commandNodeID,
						Label:  "next",
					})
				}
				prevCommandID = commandNodeID
				fmt.Println("command.Directive", command.Functions)

				// Add edges for trigger and depend directives
				for _, directive := range command.Functions {
					if directive.Type == "trigger" || directive.Type == "depend" {
						if len(directive.Args) > 0 {
							if targetJob, ok := directive.Args[0].(string); ok {
								if targetJobIDs, exists := jobNodeIDs[targetJob]; exists {
									for _, targetJobID := range targetJobIDs {
										dag.Edges = append(dag.Edges, DAGEdge{
											Source: jobInstanceID,
											Target: targetJobID,
											Label:  directive.Type,
										})
									}
								} else {
									Warn(fmt.Sprintf("%s target %s not found", directive.Type, targetJob), job.Filename, 0, 0)
								}
							}
						}
					}
				}
			}

			// Add edges for job-level trigger and depend directives
			for _, directive := range job.Directives {
				if directive.Type == "trigger" || directive.Type == "depend" {
					if len(directive.Params) > 0 {
						if targetJob, ok := directive.Params[0].(string); ok {
							if targetJobIDs, exists := jobNodeIDs[targetJob]; exists {
								for _, targetJobID := range targetJobIDs {
									dag.Edges = append(dag.Edges, DAGEdge{
										Source: jobInstanceID,
										Target: targetJobID,
										Label:  directive.Type,
									})
								}
							} else {
								Warn(fmt.Sprintf("%s target %s not found", directive.Type, targetJob), job.Filename, 0, 0)
							}
						}
					}
				}
				if directive.Type == "dependall" {
					if len(directive.Params) > 0 {
						if targetJob, ok := directive.Params[0].(string); ok {
							if targetJobIDs, exists := jobNodeIDs[targetJob]; exists {
								for _, targetJobID := range targetJobIDs {
									dag.Edges = append(dag.Edges, DAGEdge{
										Source: jobInstanceID,
										Target: targetJobID,
										Label:  "dependall",
									})
								}
							} else {
								Warn(fmt.Sprintf("dependall target %s not found", targetJob), job.Filename, 0, 0)
							}
						}
					}
				}
			}
		}
	}

	// Validate for cycles
	if hasCycle(dag) {
		Warn("Potential cycle detected in DAG", "", 0, 0)
	}

	return dag
}

// ToJSON serializes the DAG to JSON
func (d *DAG) ToJSON() (string, error) {
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// hasCycle checks for cycles in the DAG (placeholder implementation)
func hasCycle(dag *DAG) bool {
	// Implement cycle detection if needed
	return false
}
