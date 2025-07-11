package core

import (
	"fmt"
	"os"
)

// ANSI escape codes for color
const (
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Green  = "\033[32m"
	Reset  = "\033[0m"
	Bold   = "\033[1m"
)

// formatMessage formats a message with type, content, file, line, and optional hint.
func formatMessage(msgType, msg, file string, line int, column int) string {
	var color, header string
	switch msgType {
	case "Syntax Error":
		color = Red
		header = "[Syntax Error]"
	case "Runtime Error":
		color = Red
		header = "[Runtime Error]"
	case "Warning":
		color = Yellow
		header = "[Warning]"
	case "Info":
		color = Green
		header = "[Info]"
	default:
		color = Reset
		header = "[Message]"
	}

	var location string
	if file != "" {
		if line > 0 || column > 0 {
			location = fmt.Sprintf("%s%s:%d:%d %s ", Blue, file, line, column, Reset)
		} else {
			location = fmt.Sprintf("%s%s %s ", Blue, file, Reset)
		}
	}

	return fmt.Sprintf(
		"%s%s%s%s %s%s\n",
		color, Bold, header, Reset, location, msg,
	)
}

// ThrowSyntaxError prints a formatted syntax error and exits.
func ThrowSyntaxError(msg, file string, line int, column int) {
	fmt.Fprintln(os.Stderr, formatMessage("Syntax Error", msg, file, line, column))
	os.Exit(1)
}

// ThrowRuntimeError prints a formatted runtime error and exits.
func ThrowRuntimeError(msg, file string, line int, column int) {
	fmt.Fprintln(os.Stderr, formatMessage("Runtime Error", msg, file, line, column))
	os.Exit(1)
}

// Warn prints a formatted warning message without exiting.
func Warn(msg, file string, line int, column int) {
	fmt.Fprintln(os.Stderr, formatMessage("Warning", msg, file, line, column))
}

// Info prints a formatted informational message.
func Info(msg, file string, line int, column int) {
	fmt.Fprintln(os.Stdout, formatMessage("Info", msg, file, line, column))
}

// ThrowSyntaxErrorSimple prints a syntax error without file or line info and exits.
func ThrowSyntaxErrorSimple(msg string) {
	ThrowSyntaxError(msg, "", 0, 0)
}

// ThrowRuntimeErrorSimple prints a runtime error without file or line info and exits.
func ThrowRuntimeErrorSimple(msg string) {
	ThrowRuntimeError(msg, "", 0, 0)
}

// WarnSimple prints a warning without file or line info.
func WarnSimple(msg string) {
	Warn(msg, "", 0, 0)
}

// InfoSimple prints an info message without file or line info.
func InfoSimple(msg string) {
	Info(msg, "", 0, 0)
}
