package core

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Function struct {
	Type       string
	Args       []interface{}
	raw        string
	result     interface{}
	argResults []interface{}
	forValues  interface{}
	forIndex   int
	Report     string
	Metadata   Metadata
}

func NewFunction(name string, filename string, line, column int, args ...interface{}) *Function {
	f := &Function{
		Type: name,
		Args: args,
		Metadata: Metadata{
			Filename: filename,
			Line:     line,
			Column:   column,
		},
	}
	f.raw = f.getRaw()
	return f
}

func (f *Function) callArgs() error {
	f.argResults = make([]interface{}, len(f.Args))
	var err error
	for i := len(f.Args) - 1; i >= 0; i-- {
		f.argResults[i], err = resolveValue(f.Args[i], f.forValues, f.forIndex)
		if err != nil {
			return err
		}
	}
	return nil
}

func resolveValue(arg interface{}, forValue interface{}, forIndex int) (interface{}, error) {
	// Handle @value and @key directly
	if str, ok := arg.(string); ok {
		if str == "@value" && forValue != nil {
			return forValue, nil
		}
		if str == "@key" && forIndex >= 0 {
			return forIndex, nil
		}
	}

	switch v := arg.(type) {
	case Var:
		switch val := v.Value.(type) {
		case StringValue:
			return val.Value, nil
		case IntValue:
			return val.Value, nil
		case FloatValue:
			return val.Value, nil
		case BoolValue:
			return val.Value, nil
		case ListValue:
			return val.Value, nil
		case MatrixValue:
			return val.Value, nil
		default:
			return val, nil
		}
	case Function:
		v.forValues = forValue
		v.forIndex = forIndex
		return v.call()
	case *Function:
		v.forValues = forValue
		v.forIndex = forIndex
		return v.call()
	case []interface{}:
		resolved := make([]interface{}, len(v))
		for i, item := range v {
			val, err := resolveValue(item, forValue, forIndex)
			if err != nil {
				return nil, err
			}
			resolved[i] = val
		}
		return resolved, nil
	case [][]interface{}:
		resolved := make([][]interface{}, len(v))
		for i, inner := range v {
			resolved[i] = make([]interface{}, len(inner))
			for j, item := range inner {
				val, err := resolveValue(item, forValue, forIndex)
				if err != nil {
					return nil, err
				}
				resolved[i][j] = val
			}
		}
		return resolved, nil
	default:
		return v, nil
	}
}

func (f *Function) call() (interface{}, error) {
	var err error

	// اجرای callArgs و چک کردن خطا
	err = f.callArgs()
	if err != nil {
		return nil, err
	}

	switch f.Type {
	case "sprintf":
		f.result, err = f.sprintf()
	case "print":
		f.result, err = f.print()
	case "println":
		f.result, err = f.println()
	case "printf":
		f.result, err = f.printf()
	case "join":
		f.result, err = f.join()
	case "upper":
		f.result, err = f.upper()
	case "lower":
		f.result, err = f.lower()
	case "len":
		f.result, err = f.len()
	case "trim":
		f.result, err = f.trim()
	case "replace":
		f.result, err = f.replace()
	case "contains":
		f.result, err = f.contains()
	case "first":
		f.result, err = f.first()
	case "last":
		f.result, err = f.last()
	case "unique":
		f.result, err = f.unique()
	case "reverse":
		f.result, err = f.reverse()
	case "writefile":
		f.result, err = f.writefile()
	case "split":
		f.result, err = f.split()
	case "append":
		f.result, err = f.append()
	case "regex":
		f.result, err = f.regex()
	case "slice":
		f.result, err = f.slice()
	case "time":
		f.result, err = f.time()
	case "middle":
		f.result, err = f.middle()
	case "min":
		f.result, err = f.min()
	case "max":
		f.result, err = f.max()
	case "transpose":
		f.result, err = f.transpose()
	case "det":
		f.result, err = f.det()
	case "sum":
		f.result, err = f.sum()
	case "getenv":
		var envName string
		var defaultVal interface{} = nil

		if len(f.argResults) == 0 {
			// آرگومان envName الزامی است
			return nil, fmt.Errorf("getenv: missing environment variable name argument")
		}

		if s, ok := f.argResults[0].(string); ok {
			envName = s
		} else {
			return nil, fmt.Errorf("getenv: first argument must be a string")
		}

		if len(f.argResults) > 1 {
			defaultVal = f.argResults[1]
		}

		f.result, err = f.getenv(envName, defaultVal)
	case "setenv":
		if len(f.argResults) < 2 {
			return nil, fmt.Errorf("setenv: requires at least 2 arguments")
		}
		key, ok := f.argResults[0].(string)
		if !ok {
			return nil, fmt.Errorf("setenv: first argument must be a string")
		}
		f.result, err = f.setenv(key, f.argResults[1])
	default:
		f.result = f.raw
		err = fmt.Errorf("@%s: unknown function", f.Type)
	}

	return f.result, err
}

func (f *Function) sprintf() (interface{}, error) {
	if len(f.argResults) == 0 {
		return "", nil
	}
	format, ok := f.argResults[0].(string)
	if !ok {
		return "", fmt.Errorf("sprintf: first argument must be string")
	}
	args := []interface{}{}
	for _, a := range f.argResults[1:] {
		args = append(args, a)
	}
	return fmt.Sprintf(format, args...), nil
}

func (f *Function) print() (interface{}, error) {
	parts := []string{}
	for _, a := range f.argResults {
		str := fmt.Sprint(a)
		if unescaped, err := strconv.Unquote(`"` + str + `"`); err == nil {
			parts = append(parts, unescaped)
		} else {
			parts = append(parts, str)
		}
	}
	f.result = strings.Join(parts, "")
	return f.result, nil
}

func (f *Function) println() (interface{}, error) {
	parts := []string{}
	for _, a := range f.argResults {
		str := fmt.Sprint(a)
		if unescaped, err := strconv.Unquote(`"` + str + `"`); err == nil {
			parts = append(parts, unescaped)
		} else {
			parts = append(parts, str)
		}
	}
	f.result = strings.Join(parts, "") + "\n"
	return f.result, nil
}

func (f *Function) printf() (interface{}, error) {
	if len(f.argResults) == 0 {
		f.result = ""
		return f.result, nil
	}
	format, ok := f.argResults[0].(string)
	if !ok {
		f.result = ""
		return f.result, fmt.Errorf("printf: first argument must be string")
	}
	if unescaped, err := strconv.Unquote(`"` + format + `"`); err == nil {
		format = unescaped
	}
	args := f.argResults[1:]
	for i, arg := range args {
		if str, ok := arg.(string); ok {
			if unescaped, err := strconv.Unquote(`"` + str + `"`); err == nil {
				args[i] = unescaped
			}
		}
	}
	f.result = fmt.Sprintf(format, args...)
	return f.result, nil
}

func (f *Function) getenv(envName string, defaultValue interface{}) (interface{}, error) {
	val, exists := os.LookupEnv(envName)
	if !exists {
		if defaultValue != nil {
			return fmt.Sprint(defaultValue), nil
		}
		return "", nil
	}
	return val, nil
}

func (f *Function) setenv(key string, value interface{}) (interface{}, error) {
	strVal := fmt.Sprint(value)
	err := os.Setenv(key, strVal)
	return nil, err
}

func (f *Function) join() (interface{}, error) {
	if len(f.argResults) == 0 {
		return "", nil
	}
	var result []string
	for _, arg := range f.argResults {
		switch v := arg.(type) {
		case string:
			result = append(result, v)
		case []interface{}:
			for _, item := range v {
				result = append(result, fmt.Sprint(item))
			}
		case [][]interface{}:
			for _, list := range v {
				for _, item := range list {
					result = append(result, fmt.Sprint(item))
				}
			}
		default:
			result = append(result, fmt.Sprint(v))
		}
	}
	return strings.Join(result, " "), nil
}

func (f *Function) upper() (interface{}, error) {
	if len(f.argResults) == 0 {
		return "", nil
	}
	return strings.ToUpper(fmt.Sprint(f.argResults[0])), nil
}

func (f *Function) lower() (interface{}, error) {
	if len(f.argResults) == 0 {
		return "", nil
	}
	return strings.ToLower(fmt.Sprint(f.argResults[0])), nil
}

func (f *Function) len() (interface{}, error) {
	if len(f.argResults) == 0 {
		return 0, nil
	}
	switch v := f.argResults[0].(type) {
	case string:
		return len(v), nil
	case []interface{}:
		return len(v), nil
	case [][]interface{}:
		return len(v), nil
	default:
		return 0, nil
	}
}

func (f *Function) trim() (interface{}, error) {
	if len(f.argResults) == 0 {
		return "", nil
	}
	return strings.TrimSpace(fmt.Sprint(f.argResults[0])), nil
}

func (f *Function) replace() (interface{}, error) {
	if len(f.argResults) < 3 {
		return "", fmt.Errorf("replace: requires 3 arguments")
	}
	old := fmt.Sprint(f.argResults[0])
	newVal := fmt.Sprint(f.argResults[1])
	str := fmt.Sprint(f.argResults[2])
	return strings.ReplaceAll(str, old, newVal), nil
}

func (f *Function) contains() (interface{}, error) {
	if len(f.argResults) < 2 {
		return false, fmt.Errorf("contains: requires 2 arguments")
	}
	substr := fmt.Sprint(f.argResults[0])
	str := fmt.Sprint(f.argResults[1])
	return strings.Contains(str, substr), nil
}

func (f *Function) unique() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, nil
	}
	seen := make(map[string]bool)
	switch v := f.argResults[0].(type) {
	case string:
		var result []string
		for _, r := range v {
			s := string(r)
			if !seen[s] {
				seen[s] = true
				result = append(result, s)
			}
		}
		return strings.Join(result, ""), nil
	case []interface{}:
		var result []interface{}
		for _, item := range v {
			s := fmt.Sprint(item)
			if !seen[s] {
				seen[s] = true
				result = append(result, item)
			}
		}
		return result, nil
	}
	return nil, nil
}

func (f *Function) reverse() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, nil
	}
	switch v := f.argResults[0].(type) {
	case string:
		runes := []rune(v)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes), nil
	case []interface{}:
		n := len(v)
		reversed := make([]interface{}, n)
		for i := 0; i < n; i++ {
			reversed[i] = v[n-1-i]
		}
		return reversed, nil
	case [][]interface{}:
		n := len(v)
		reversed := make([][]interface{}, n)
		for i := 0; i < n; i++ {
			reversed[i] = v[n-1-i]
		}
		return reversed, nil
	}
	return nil, nil
}

func (f *Function) writefile() (interface{}, error) {
	if len(f.argResults) != 2 {
		return nil, fmt.Errorf("@writefile requires exactly 2 arguments (filename string, data string)")
	}

	filenameVal := f.argResults[0]
	filename, ok := filenameVal.(string)
	if !ok || strings.TrimSpace(filename) == "" {
		return nil, fmt.Errorf("@writefile: first argument must be a non-empty string (filename)")
	}

	dataVal := f.argResults[1]
	data, ok := dataVal.(string)
	if !ok {
		return nil, fmt.Errorf("@writefile: second argument must be a string (data to write)")
	}

	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return nil, err
	}

	return "ok", nil
}

func (f *Function) split() (interface{}, error) {
	if len(f.argResults) != 2 {
		return nil, fmt.Errorf("split: requires exactly 2 arguments (separator, input)")
	}
	separator, ok := f.argResults[0].(string)
	if !ok {
		return nil, fmt.Errorf("split: first argument must be a string (separator)")
	}
	input, ok := f.argResults[1].(string)
	if !ok {
		return nil, fmt.Errorf("split: second argument must be a string (input)")
	}
	result := strings.Split(input, separator)
	interfaceResult := make([]interface{}, len(result))
	for i, v := range result {
		interfaceResult[i] = v
	}
	return interfaceResult, nil
}

func (f *Function) append() (interface{}, error) {
	if len(f.argResults) != 2 {
		return nil, fmt.Errorf("append: requires exactly 2 arguments (input, appendTo)")
	}
	appendTo := f.argResults[1]
	switch input := f.argResults[0].(type) {
	case string:
		// Append to string
		appendStr, ok := appendTo.(string)
		if !ok {
			return nil, fmt.Errorf("append: second argument must be a string when input is a string")
		}
		return input + appendStr, nil
	case []interface{}:
		// Append to list
		result := make([]interface{}, len(input)+1)
		copy(result, input)
		result[len(input)] = appendTo
		return result, nil
	case [][]interface{}:
		// Append to matrix (as a new row)
		newRow, ok := appendTo.([]interface{})
		if !ok {
			return nil, fmt.Errorf("append: second argument must be a list when input is a matrix")
		}
		result := make([][]interface{}, len(input)+1)
		copy(result, input)
		result[len(input)] = newRow
		return result, nil
	default:
		return nil, fmt.Errorf("append: first argument must be a string, list, or matrix")
	}
}

func (f *Function) regex() (interface{}, error) {
	if len(f.argResults) != 2 {
		return nil, fmt.Errorf("regex: requires exactly 2 arguments (regex, input)")
	}
	pattern, ok := f.argResults[0].(string)
	if !ok {
		return nil, fmt.Errorf("regex: first argument must be a string (regex pattern)")
	}
	input, ok := f.argResults[1].(string)
	if !ok {
		return nil, fmt.Errorf("regex: second argument must be a string (input)")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("regex: invalid pattern %s: %v", pattern, err)
	}
	matches := re.FindAllString(input, -1)
	if len(matches) == 0 {
		return nil, nil
	}
	result := make([]interface{}, len(matches))
	for i, match := range matches {
		result[i] = match
	}
	return result, nil
}

func (f *Function) slice() (interface{}, error) {
	if len(f.argResults) != 3 {
		return nil, fmt.Errorf("slice: requires exactly 3 arguments (input, from, to)")
	}
	from, ok := f.argResults[1].(int)
	if !ok {
		return nil, fmt.Errorf("slice: second argument must be an integer (from index)")
	}
	to, ok := f.argResults[2].(int)
	if !ok {
		return nil, fmt.Errorf("slice: third argument must be an integer (to index)")
	}
	if from < 0 || from > to {
		return nil, fmt.Errorf("slice: invalid indices (from=%d, to=%d)", from, to)
	}
	switch input := f.argResults[0].(type) {
	case string:
		if to > len(input) {
			return nil, fmt.Errorf("slice: to index (%d) exceeds string length (%d)", to, len(input))
		}
		return input[from:to], nil
	case []interface{}:
		if to > len(input) {
			return nil, fmt.Errorf("slice: to index (%d) exceeds list length (%d)", to, len(input))
		}
		return input[from:to], nil
	case [][]interface{}:
		if to > len(input) {
			return nil, fmt.Errorf("slice: to index (%d) exceeds matrix length (%d)", to, len(input))
		}
		return input[from:to], nil
	default:
		return nil, fmt.Errorf("slice: first argument must be a string, list, or matrix")
	}
}

func (f *Function) time() (interface{}, error) {
	if len(f.argResults) == 0 {
		// Case: @time() - Current time in RFC3339 format
		return time.Now().Format(time.RFC3339), nil
	}

	if len(f.argResults) == 1 {
		// Case: @time("format") - Current time in specified format
		format, ok := f.argResults[0].(string)
		if !ok {
			return nil, fmt.Errorf("time: first argument must be a string (format)")
		}
		return time.Now().Format(format), nil
	}

	if len(f.argResults) == 2 {
		// Case: @time("time_string", "format") - Parse time string and format
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

func (f *Function) first() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, nil
	}
	switch v := f.argResults[0].(type) {
	case string:
		if len(v) > 0 {
			return string(v[0]), nil
		}
	case []interface{}:
		if len(v) > 0 {
			return v[0], nil
		}
	case [][]interface{}:
		if len(v) > 0 && len(v[0]) > 0 {
			return v[0][0], nil
		}
	}
	return nil, nil
}

func (f *Function) last() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, nil
	}
	switch v := f.argResults[0].(type) {
	case string:
		if len(v) > 0 {
			return string(v[len(v)-1]), nil
		}
	case []interface{}:
		if len(v) > 0 {
			return v[len(v)-1], nil
		}
	case [][]interface{}:
		if len(v) > 0 && len(v[len(v)-1]) > 0 {
			lastList := v[len(v)-1]
			return lastList[len(lastList)-1], nil
		}
	}
	return nil, nil
}

func (f *Function) middle() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, nil
	}
	switch v := f.argResults[0].(type) {
	case string:
		if len(v) <= 2 {
			return "", nil
		}
		midIndex := len(v) / 2
		if len(v)%2 == 0 {
			return string(v[midIndex-1 : midIndex+1]), nil
		}
		return string(v[midIndex]), nil
	case []interface{}:
		if len(v) <= 2 {
			return nil, nil
		}
		midIndex := len(v) / 2
		if len(v)%2 == 0 {
			return v[midIndex-1 : midIndex+1], nil
		}
		return v[midIndex], nil
	case [][]interface{}:
		if len(v) <= 2 {
			return nil, nil
		}
		midIndex := len(v) / 2
		if len(v)%2 == 0 {
			return v[midIndex-1 : midIndex+1], nil
		}
		return v[midIndex], nil
	}
	return nil, fmt.Errorf("middle: first argument must be a string, list, or matrix")
}

func (f *Function) min() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, fmt.Errorf("min: at least one argument required")
	}

	// Case 1: Single argument (list or matrix)
	if len(f.argResults) == 1 {
		switch v := f.argResults[0].(type) {
		case []interface{}:
			if len(v) == 0 {
				return nil, fmt.Errorf("min: empty list")
			}
			minVal, err := toFloat64(v[0])
			if err != nil {
				return nil, fmt.Errorf("min: list elements must be numbers, got %v", v[0])
			}
			for _, item := range v[1:] {
				val, err := toFloat64(item)
				if err != nil {
					return nil, fmt.Errorf("min: list elements must be numbers, got %v", item)
				}
				if val < minVal {
					minVal = val
				}
			}
			if isInt(minVal) {
				return int(minVal), nil
			}
			return minVal, nil
		case [][]interface{}:
			if len(v) == 0 || len(v[0]) == 0 {
				return nil, fmt.Errorf("min: empty matrix")
			}
			minVal, err := toFloat64(v[0][0])
			if err != nil {
				return nil, fmt.Errorf("min: matrix elements must be numbers, got %v", v[0][0])
			}
			for _, row := range v {
				for _, item := range row {
					val, err := toFloat64(item)
					if err != nil {
						return nil, fmt.Errorf("min: matrix elements must be numbers, got %v", item)
					}
					if val < minVal {
						minVal = val
					}
				}
			}
			if isInt(minVal) {
				return int(minVal), nil
			}
			return minVal, nil
		default:
			return nil, fmt.Errorf("min: first argument must be a list or matrix")
		}
	}

	// Case 2: Multiple number arguments
	minVal, err := toFloat64(f.argResults[0])
	if err != nil {
		return nil, fmt.Errorf("min: arguments must be numbers, got %v", f.argResults[0])
	}
	for _, arg := range f.argResults[1:] {
		val, err := toFloat64(arg)
		if err != nil {
			return nil, fmt.Errorf("min: arguments must be numbers, got %v", arg)
		}
		if val < minVal {
			minVal = val
		}
	}
	if isInt(minVal) {
		return int(minVal), nil
	}
	return minVal, nil
}

func (f *Function) max() (interface{}, error) {
	if len(f.argResults) == 0 {
		return nil, fmt.Errorf("max: at least one argument required")
	}

	// Case 1: Single argument (list or matrix)
	if len(f.argResults) == 1 {
		switch v := f.argResults[0].(type) {
		case []interface{}:
			if len(v) == 0 {
				return nil, fmt.Errorf("max: empty list")
			}
			maxVal, err := toFloat64(v[0])
			if err != nil {
				return nil, fmt.Errorf("max: list elements must be numbers, got %v", v[0])
			}
			for _, item := range v[1:] {
				val, err := toFloat64(item)
				if err != nil {
					return nil, fmt.Errorf("max: list elements must be numbers, got %v", item)
				}
				if val > maxVal {
					maxVal = val
				}
			}
			if isInt(maxVal) {
				return int(maxVal), nil
			}
			return maxVal, nil
		case [][]interface{}:
			if len(v) == 0 || len(v[0]) == 0 {
				return nil, fmt.Errorf("max: empty matrix")
			}
			maxVal, err := toFloat64(v[0][0])
			if err != nil {
				return nil, fmt.Errorf("max: matrix elements must be numbers, got %v", v[0][0])
			}
			for _, row := range v {
				for _, item := range row {
					val, err := toFloat64(item)
					if err != nil {
						return nil, fmt.Errorf("max: matrix elements must be numbers, got %v", item)
					}
					if val > maxVal {
						maxVal = val
					}
				}
			}
			if isInt(maxVal) {
				return int(maxVal), nil
			}
			return maxVal, nil
		default:
			return nil, fmt.Errorf("max: first argument must be a list or matrix")
		}
	}

	// Case 2: Multiple number arguments
	maxVal, err := toFloat64(f.argResults[0])
	if err != nil {
		return nil, fmt.Errorf("max: arguments must be numbers, got %v", f.argResults[0])
	}
	for _, arg := range f.argResults[1:] {
		val, err := toFloat64(arg)
		if err != nil {
			return nil, fmt.Errorf("max: arguments must be numbers, got %v", arg)
		}
		if val > maxVal {
			maxVal = val
		}
	}
	if isInt(maxVal) {
		return int(maxVal), nil
	}
	return maxVal, nil
}

func (f *Function) transpose() (interface{}, error) {
	if len(f.argResults) != 1 {
		return nil, fmt.Errorf("transpose: requires exactly 1 argument (matrix)")
	}
	matrix, ok := f.argResults[0].([][]interface{})
	if !ok {
		return nil, fmt.Errorf("transpose: argument must be a matrix")
	}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]interface{}{}, nil
	}
	// Check if matrix is valid (all rows have the same length)
	cols := len(matrix[0])
	for _, row := range matrix[1:] {
		if len(row) != cols {
			return nil, fmt.Errorf("transpose: matrix rows must have equal length")
		}
	}
	// Create transposed matrix
	rows := len(matrix)
	result := make([][]interface{}, cols)
	for i := range result {
		result[i] = make([]interface{}, rows)
		for j := range matrix {
			result[i][j] = matrix[j][i]
		}
	}
	return result, nil
}

func (f *Function) det() (interface{}, error) {
	if len(f.argResults) != 1 {
		return nil, fmt.Errorf("det: requires exactly 1 argument (matrix)")
	}
	matrix, ok := f.argResults[0].([][]interface{})
	if !ok {
		return nil, fmt.Errorf("det: argument must be a matrix")
	}
	if len(matrix) == 0 {
		return nil, fmt.Errorf("det: empty matrix")
	}
	if len(matrix) != len(matrix[0]) {
		return nil, fmt.Errorf("det: matrix must be square")
	}
	n := len(matrix)
	// Convert to float64 matrix
	m := make([][]float64, n)
	for i, row := range matrix {
		m[i] = make([]float64, n)
		for j, val := range row {
			fVal, err := toFloat64(val)
			if err != nil {
				return nil, fmt.Errorf("det: matrix elements must be numbers, got %v", val)
			}
			m[i][j] = fVal
		}
	}
	// Calculate determinant
	detVal := determinant(m)
	if isInt(detVal) {
		return int(detVal), nil
	}
	return detVal, nil
}

func (f *Function) sum() (interface{}, error) {
	if len(f.argResults) != 1 {
		return nil, fmt.Errorf("sum: requires exactly 1 argument (list or matrix)")
	}
	switch v := f.argResults[0].(type) {
	case []interface{}:
		if len(v) == 0 {
			return 0, nil
		}
		sumVal, err := toFloat64(v[0])
		if err != nil {
			return nil, fmt.Errorf("sum: list elements must be numbers, got %v", v[0])
		}
		for _, item := range v[1:] {
			val, err := toFloat64(item)
			if err != nil {
				return nil, fmt.Errorf("sum: list elements must be numbers, got %v", item)
			}
			sumVal += val
		}
		if isInt(sumVal) {
			return int(sumVal), nil
		}
		return sumVal, nil
	case [][]interface{}:
		if len(v) == 0 || len(v[0]) == 0 {
			return 0, nil
		}
		sumVal, err := toFloat64(v[0][0])
		if err != nil {
			return nil, fmt.Errorf("sum: matrix elements must be numbers, got %v", v[0][0])
		}
		for _, row := range v {
			for _, item := range row {
				val, err := toFloat64(item)
				if err != nil {
					return nil, fmt.Errorf("sum: matrix elements must be numbers, got %v", item)
				}
				sumVal += val
			}
		}
		if isInt(sumVal) {
			return int(sumVal), nil
		}
		return sumVal, nil
	default:
		return nil, fmt.Errorf("sum: argument must be a list or matrix")
	}
}

// Helper function to convert interface{} to float64
func toFloat64(val interface{}) (float64, error) {
	switch v := val.(type) {
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		return 0, fmt.Errorf("cannot convert string to number: %v", v)
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

// Helper function to check if a float64 is effectively an integer
func isInt(val float64) bool {
	return math.Mod(val, 1) == 0
}

// Helper function to calculate determinant of a square matrix
func determinant(m [][]float64) float64 {
	n := len(m)
	if n == 1 {
		return m[0][0]
	}
	if n == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}
	det := 0.0
	for j := 0; j < n; j++ {
		det += m[0][j] * float64(cofactor(m, 0, j))
	}
	return det
}

// Helper function to calculate cofactor for determinant
func cofactor(m [][]float64, row, col int) float64 {
	subMatrix := make([][]float64, len(m)-1)
	for i := 0; i < len(m)-1; i++ {
		subMatrix[i] = make([]float64, len(m)-1)
		for j := 0; j < len(m)-1; j++ {
			r := i
			if i >= row {
				r++
			}
			c := j
			if j >= col {
				c++
			}
			subMatrix[i][j] = m[r][c]
		}
	}
	sign := 1.0
	if (row+col)%2 != 0 {
		sign = -1.0
	}
	return sign * determinant(subMatrix)
}

func (f *Function) getRaw() string {
	parts := []string{}
	for _, a := range f.Args {
		switch v := a.(type) {
		case *Function:
			parts = append(parts, v.getRaw())
		case string:
			if v == "@value" && f.forValues != nil {
				parts = append(parts, fmt.Sprintf("%v", f.forValues))
			} else if v == "@key" && f.forIndex >= 0 {
				parts = append(parts, fmt.Sprintf("%d", f.forIndex))
			} else {
				parts = append(parts, fmt.Sprintf(`"%s"`, v))
			}
		default:
			parts = append(parts, fmt.Sprintf("%v", v))
		}
	}
	return fmt.Sprintf("@%s(%s)", f.Type, strings.Join(parts, ","))
}

func (f *Function) GetResult() interface{} {
	return f.result
}

func (f *Function) String() string {
	if f.result != nil {
		return fmt.Sprint(f.result)
	}
	return f.getRaw()
}
