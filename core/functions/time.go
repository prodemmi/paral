package functions

import (
	"fmt"
	"time"
)

func Time(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return time.Now().Format(time.RFC3339), nil
	}

	if len(args) == 1 {
		format, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("first argument must be a string (format)")
		}
		return time.Now().Format(format), nil
	}

	if len(args) == 2 {
		timeStr, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("first argument must be a string (time string)")
		}
		format, ok := args[1].(string)
		if !ok {
			return nil, fmt.Errorf("second argument must be a string (format)")
		}
		parsedTime, err := time.Parse(format, timeStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse time string '%s' with format '%s': %v", timeStr, format, err)
		}
		return parsedTime.Format(format), nil
	}

	return nil, fmt.Errorf("invalid number of arguments, expected 0, 1, or 2, got %d", len(args))
}
