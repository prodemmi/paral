package functions

import (
	"fmt"
	"time"
)

func (f *Function) time() (interface{}, error) {
	if len(f.argResults) == 0 {
		return time.Now().Format(time.RFC3339), nil
	}

	if len(f.argResults) == 1 {
		format, ok := f.argResults[0].(string)
		if !ok {
			return nil, fmt.Errorf("time: first argument must be a string (format)")
		}
		return time.Now().Format(format), nil
	}

	if len(f.argResults) == 2 {
		timeStr, ok := f.argResults[0].(string)
		if !ok {
			return nil, fmt.Errorf("time: first argument must be a string (time string)")
		}
		format, ok := f.argResults[1].(string)
		if !ok {
			return nil, fmt.Errorf("time: second argument must be a string (format)")
		}
		parsedTime, err := time.Parse(format, timeStr)
		if err != nil {
			return nil, fmt.Errorf("time: failed to parse time string '%s' with format '%s': %v", timeStr, format, err)
		}
		return parsedTime.Format(format), nil
	}

	return nil, fmt.Errorf("time: invalid number of arguments, expected 0, 1, or 2, got %d", len(f.argResults))
}
