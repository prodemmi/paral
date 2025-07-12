package core

import (
	"fmt"
	"strings"
	"sync"
)

type Core struct {
	Vars     []Var
	Matrixes []Matrix
	Jobs     []Job
	mu       sync.Mutex
}

func NewCore() *Core {
	return &Core{
		Vars:     make([]Var, 0),
		Matrixes: make([]Matrix, 0),
		Jobs:     make([]Job, 0),
	}
}

func (c *Core) AddVar(v Var) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Vars = append(c.Vars, v)
}

func (c *Core) AddMatrix(m Matrix) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Matrixes = append(c.Matrixes, m)
}

func (c *Core) AddJob(j Job) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Jobs = append(c.Jobs, j)
}

func (c *Core) ResolveMatrix(name string) ([][]interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, m := range c.Matrixes {
		if m.Name == name {
			return m.GetMatrixCombinations(), true
		}
	}
	return nil, false
}

func (c *Core) ResolveVariable(name string) (Var, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, v := range c.Vars {
		if v.Name == name {
			return v, true
		}
	}
	return Var{}, false
}

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
		// Print Directives
		b.WriteString("    ~Job Directives:\n")
		for _, dir := range j.Directives {
			b.WriteString(fmt.Sprintf("      @%s = %v\n", dir.Type, dir.Value))
		}
		// Print Commands
		b.WriteString("    ~Commands:\n")
		for i, command := range j.Commands {
			if command.CMD != "" {
				b.WriteString(fmt.Sprintf("      CMD: %s\n", command.CMD))
				if len(command.Directive) > 0 {
					for _, d := range command.Directive {
						if len(d.Params) > 0 {
							args := make([]string, len(d.Params))
							for i, v := range d.Params {
								args[i] = fmt.Sprintf("%v", v)
							}
							b.WriteString(fmt.Sprintf("      CMD Directives: %s\n", fmt.Sprintf("@%s(%s)", d.Type, strings.Join(args, ", "))))
						}
					}
				}
			}
			if i < len(j.Commands)-1 {
				b.WriteString("    --------------\n")
			}
		}
	}

	fmt.Println(b.String())
}
