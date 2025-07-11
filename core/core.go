package core

import (
	"fmt"
	"strings"
	"sync"
)

type Core struct {
	mu       sync.Mutex
	Vars     []Var
	Matrixes []Matrix
	Jobs     []Job

	Filename string
}

// NewCore creates and returns a new Core instance
func NewCore(filename string) *Core {
	return &Core{
		Vars:     []Var{},
		Matrixes: []Matrix{},
		Jobs:     []Job{},

		Filename: filename,
	}
}

// AddVar safely adds a variable
func (c *Core) AddVar(v Var) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := v.Value.([]interface{}); ok {
		v.Type = "list"
	}
	c.Vars = append(c.Vars, v)
}

// AddMatrix safely adds a matrix
func (c *Core) AddMatrix(m Matrix) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Matrixes = append(c.Matrixes, m)
}

// AddJob safely adds a job
func (c *Core) AddJob(j Job) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Jobs = append(c.Jobs, j)
}

// ResolveMatrix returns matrix values for a given matrix name
func (c *Core) ResolveMatrix(name string) ([][]string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, matrix := range c.Matrixes {
		if matrix.Name == name {
			// Convert [][]interface{} to [][]string
			result := make([][]string, len(matrix.Values))
			for i, row := range matrix.Values {
				result[i] = make([]string, len(row))
				for j, val := range row {
					switch v := val.(type) {
					case string:
						result[i][j] = v
					case int:
						result[i][j] = fmt.Sprintf("%d", v)
					case bool:
						result[i][j] = fmt.Sprintf("%t", v)
					case float64:
						result[i][j] = fmt.Sprintf("%f", v)
					default:
						return nil, false // Unsupported type
					}
				}
			}
			return result, true
		}
	}
	return nil, false
}

// GetValues prints all variables
func (c *Core) PrintValues() {
	c.mu.Lock()
	defer c.mu.Unlock()

	var b strings.Builder

	// Print Vars
	b.WriteString("🔸 Vars:\n")
	for _, v := range c.Vars {
		b.WriteString(fmt.Sprintf("  - %s (%s) = %v\n", v.Name, v.Type, v.Value))
	}

	// Print Matrixes
	b.WriteString("\n🔸 Matrixes:\n")
	for _, m := range c.Matrixes {
		b.WriteString(fmt.Sprintf("  - %s:\n", m.Name))
		for _, combo := range m.GetMatrixCombinations() {
			b.WriteString(fmt.Sprintf("      • %v\n", combo))
		}
	}

	// Print Jobs
	b.WriteString("\n🔸 Jobs:\n")
	for _, j := range c.Jobs {
		b.WriteString(fmt.Sprintf("  %s:\n", j.Name))
		for _, command := range j.Commands {
			b.WriteString(fmt.Sprintf("    → %s\n", command.CMD))
		}
	}

	fmt.Println(b.String())
}
