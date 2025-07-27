package runtime

import (
	"fmt"
	"paral/core/functions"
	"strings"
)

type Command struct {
	Functions []*functions.Function
	Stashes   []Stash
	Buffs     []Buf
	RawText   string
	Block     int
}

func (c *Command) GetResult(forValue interface{}, forIndex int) string {
	cmd := c.RawText

	for _, function := range c.Functions {
		function.SetLoopData(forValue, forIndex)
		result, err := function.Call()
		if err != nil {
			continue // Skip failed functions
		}

		funcRaw := function.GetRaw()

		switch function.Type {
		case "print", "println", "printf":
			output := strings.TrimRight(fmt.Sprint(result), "\n")
			if strings.HasPrefix(strings.TrimSpace(cmd), "@") {
				return output // Replace entire command for print functions
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
