package core

import (
	"fmt"
	"strings"
)

type Command struct {
	Functions  []Function
	Conditions []Condition
	Stashes    []Stash
	Bufes      []Buf
	RawText    string
	Block      int
}

func (c *Command) GetResult(forValue interface{}, forIndex int) string {
	cmd := c.RawText

	for _, function := range c.Functions {
		function.forValues = forValue
		function.forIndex = forIndex
		function.call()

		funcRaw := function.raw
		result := function.GetResult()

		switch function.Type {
		case "print", "println", "printf":
			output := strings.TrimRight(fmt.Sprint(result), "\n")
			if strings.HasPrefix(strings.TrimSpace(cmd), "@") {
				cmd = output
			} else {
				cmd = strings.ReplaceAll(cmd, funcRaw, output)
			}

		default:
			switch res := result.(type) {
			case string:
				cmd = strings.ReplaceAll(cmd, funcRaw, res)
			case []interface{}:
				var strValues []string
				for _, v := range res {
					strValues = append(strValues, fmt.Sprint(v))
				}
				cmd = strings.ReplaceAll(cmd, funcRaw, strings.Join(strValues, " "))
			default:
				cmd = strings.ReplaceAll(cmd, funcRaw, fmt.Sprintf("%v", res))
			}
		}
	}

	return cmd
}
