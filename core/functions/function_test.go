package functions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSprintf(t *testing.T) {
	nestedSprintf := NewTestFunction("sprintf", "My Name is %s", "Emad")
	def := NewTestFunction("sprintf", "Hello, %s", nestedSprintf)

	result, _ := def.Call()
	assert.Equal(t, "Hello, My Name is Emad", result)
}

func TestPrint(t *testing.T) {
	nestedSprintf := NewTestFunction("sprintf", "My Name is %s", "Emad")
	def := NewTestFunction("print", "Hello, ", nestedSprintf)

	result, _ := def.Call()
	expected := "Hello, My Name is Emad"
	assert.Equal(t, expected, result)
}

func TestPrintf(t *testing.T) {
	def := NewTestFunction("printf", "Value: %d and name: %s", 42, "Layra")

	result, _ := def.Call()
	expected := "Value: 42 and name: Layra"
	assert.Equal(t, expected, result)
}

func TestNestedFunctions(t *testing.T) {
	inner := NewTestFunction("sprintf", "Inner %d", 100)
	middle := NewTestFunction("sprintf", "Middle %s", inner)
	outer := NewTestFunction("sprintf", "Outer %s", middle)

	result, _ := outer.Call()
	expected := "Outer Middle Inner 100"
	assert.Equal(t, expected, result)
}

func TestEmptyArgs(t *testing.T) {
	def := NewTestFunction("sprintf", 1, 1)
	result, _ := def.Call()
	assert.Equal(t, "", result)

	defPrint := NewTestFunction("print")
	result2, _ := defPrint.Call()
	assert.Equal(t, "", result2)

	defPrintf := NewTestFunction("printf")
	result3, _ := defPrintf.Call()
	assert.Equal(t, "", result3)
}

func TestGetenvAndSetenv(t *testing.T) {
	_ = os.Unsetenv("APP_MODE")

	// getenv with default
	getEnvFunc := NewTestFunction("getenv", "APP_MODE", "development")
	val, _ := getEnvFunc.Call()
	assert.Equal(t, "development", val)

	// setenv
	setEnvFunc := NewTestFunction("setenv", "APP_MODE", "production")
	_, _ = setEnvFunc.Call()

	// getenv again, now it should return "production"
	getEnvFunc2 := NewTestFunction("getenv", "APP_MODE")
	val2, _ := getEnvFunc2.Call()
	assert.Equal(t, "production", val2)
}

func TestJoin(t *testing.T) {
	f := NewTestFunction("join", "a", []interface{}{"b", "c"}, [][]interface{}{{"d", "e"}})
	result, _ := f.Call()
	assert.Equal(t, "a b c d e", result)
}

func TestUpperLowerTrimLen(t *testing.T) {
	f1 := NewTestFunction("upper", "hello")
	r1, _ := f1.Call()
	assert.Equal(t, "HELLO", r1)

	f2 := NewTestFunction("lower", "HELLO")
	r2, _ := f2.Call()
	assert.Equal(t, "hello", r2)

	f3 := NewTestFunction("trim", "  space  ")
	r3, _ := f3.Call()
	assert.Equal(t, "space", r3)

	f4 := NewTestFunction("len", []interface{}{1, 2, 3})
	r4, _ := f4.Call()
	assert.Equal(t, 3, r4)
}

func TestReplaceContains(t *testing.T) {
	f1 := NewTestFunction("replace", "a", "b", "a cat and a hat")
	r1, _ := f1.Call()
	assert.Equal(t, "b cbt bnd b hbt", r1)

	f2 := NewTestFunction("contains", "cat", "a cat and a hat")
	r2, _ := f2.Call()
	assert.Equal(t, true, r2)
}

func TestFirstLastUniqueReverse(t *testing.T) {
	f1 := NewTestFunction("first", []interface{}{1, 2, 3})
	r1, _ := f1.Call()
	assert.Equal(t, 1, r1)

	f2 := NewTestFunction("last", []interface{}{1, 2, 3})
	r2, _ := f2.Call()
	assert.Equal(t, 3, r2)

	f3 := NewTestFunction("unique", []interface{}{1, 2, 2, 1, 3})
	r3, _ := f3.Call()
	assert.ElementsMatch(t, []interface{}{1, 2, 3}, r3)

	f4 := NewTestFunction("reverse", []interface{}{1, 2, 3})
	r4, _ := f4.Call()
	assert.Equal(t, []interface{}{3, 2, 1}, r4)
}

func TestWritefile(t *testing.T) {
	filename := "test_output.txt"
	defer os.Remove(filename)

	f := NewTestFunction("writefile", filename, "hello world")
	r, _ := f.Call()
	assert.Equal(t, "ok", r)

	data, err := os.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, "hello world", string(data))
}

func TestSplitAppend(t *testing.T) {
	f1 := NewTestFunction("split", ",", "a,b,c")
	r1, _ := f1.Call()
	assert.ElementsMatch(t, []interface{}{"a", "b", "c"}, r1)

	f2 := NewTestFunction("append", []interface{}{1, 2}, 3)
	r2, _ := f2.Call()
	assert.Equal(t, []interface{}{1, 2, 3}, r2)
}

func TestRegex(t *testing.T) {
	f := NewTestFunction("regex", "\\d+", "abc123def456")
	r, _ := f.Call()
	assert.ElementsMatch(t, []interface{}{"123", "456"}, r)
}

func TestSlice(t *testing.T) {
	f := NewTestFunction("slice", []interface{}{"a", "b", "c"}, 0, 2)
	r, _ := f.Call()
	assert.Equal(t, []interface{}{"a", "b"}, r)
}

func TestTime(t *testing.T) {
	f := NewTestFunction("time")
	r, _ := f.Call()
	assert.NotEmpty(t, r)
}

func TestMiddleMinMaxTransposeDetSum(t *testing.T) {
	f1 := NewTestFunction("middle", "abcde")
	r1, _ := f1.Call()
	assert.Equal(t, "c", r1)

	f2 := NewTestFunction("min", []interface{}{3, 1, 4})
	r2, _ := f2.Call()
	assert.Equal(t, 1, r2)

	f3 := NewTestFunction("max", []interface{}{3, 1, 4})
	r3, _ := f3.Call()
	assert.Equal(t, 4, r3)

	matrix := [][]interface{}{{1.0, 2.0}, {3.0, 4.0}}
	f4 := NewTestFunction("transpose", matrix)
	r4, _ := f4.Call()
	assert.Equal(t, [][]interface{}{{1.0, 3.0}, {2.0, 4.0}}, r4)

	f5 := NewTestFunction("det", matrix)
	r5, _ := f5.Call()
	assert.Equal(t, int(-2), r5)

	f6 := NewTestFunction("sum", []interface{}{1, 2, 3})
	r6, _ := f6.Call()
	assert.Equal(t, 6, r6)
}
