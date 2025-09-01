package functions

import (
	"fmt"
)

func IsNot(args ...interface{}) (interface{}, error) {
	if len(args) == 2 {
		// equality check (negated)
		return args[0] != args[1], nil
	}

	if len(args) == 3 {
		left := args[0]
		op, ok := args[1].(string)
		if !ok {
			return false, fmt.Errorf("operator must be a string")
		}
		right := args[2]

		// try numeric comparison
		leftF, lOk := toFloat(left)
		rightF, rOk := toFloat(right)
		if lOk && rOk {
			switch op {
			case ">", "gt":
				return !(leftF > rightF), nil
			case "<", "lt":
				return !(leftF < rightF), nil
			case ">=", "gte":
				return !(leftF >= rightF), nil
			case "<=", "lte":
				return !(leftF <= rightF), nil
			case "==", "=", "eq":
				return !(leftF == rightF), nil
			case "!=", "ne":
				return !(leftF != rightF), nil
			}
		}

		// fallback: string comparison
		ls := fmt.Sprint(left)
		rs := fmt.Sprint(right)
		switch op {
		case "==", "=", "eq":
			return !(ls == rs), nil
		case "!=", "ne":
			return !(ls != rs), nil
		}
		return false, fmt.Errorf("unsupported operator %q for types %T and %T", op, left, right)
	}

	return false, fmt.Errorf("invalid number of arguments")
}
