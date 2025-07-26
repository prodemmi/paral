package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSprintf(t *testing.T) {
	nestedSprintf := NewFunction("sprintf", "My Name is %s", "Emad")
	def := NewFunction("sprintf", "Hello, %s", nestedSprintf)

	result := def.call()
	assert.Equal(t, "Hello, My Name is Emad", result)
}

func TestPrint(t *testing.T) {
	nestedSprintf := NewFunction("sprintf", "My Name is %s", "Emad")
	def := NewFunction("print", "Hello,", nestedSprintf, 123)

	result := def.call()
	expected := "Hello, My Name is Emad 123"
	assert.Equal(t, expected, result)
}

func TestPrintf(t *testing.T) {
	def := NewFunction("printf", "Value: %d and name: %s", 42, "Layra")

	result := def.call()
	expected := "Value: 42 and name: Layra"
	assert.Equal(t, expected, result)
}

func TestNestedFunctions(t *testing.T) {
	inner := NewFunction("sprintf", "Inner %d", 100)
	middle := NewFunction("sprintf", "Middle %s", inner)
	outer := NewFunction("sprintf", "Outer %s", middle)

	result := outer.call()
	expected := "Outer Middle Inner 100"
	assert.Equal(t, expected, result)
}

func TestEmptyArgs(t *testing.T) {
	def := NewFunction("sprintf")
	assert.Equal(t, "", def.call())

	defPrint := NewFunction("print")
	assert.Equal(t, "", defPrint.call())

	defPrintf := NewFunction("printf")
	assert.Equal(t, "", defPrintf.call())
}

func TestGetenvAndSetenv(t *testing.T) {
	// پاک کردن متغیر محیطی برای اطمینان
	_ = os.Unsetenv("APP_MODE")

	// تست getEnv با مقدار پیش‌فرض
	getEnvFunc := NewFunction("getenv", "APP_MODE", "development")
	val := getEnvFunc.call()
	assert.Equal(t, "development", val)

	// تست setEnv برای تنظیم مقدار
	setEnvFunc := NewFunction("setenv", "APP_MODE", "production")
	setEnvFunc.call()

	// دوباره getEnv بدون مقدار پیش‌فرض چون الان تنظیم شده
	getEnvFunc2 := NewFunction("getenv", "APP_MODE")
	val2 := getEnvFunc2.call()
	assert.Equal(t, "production", val2)
}
