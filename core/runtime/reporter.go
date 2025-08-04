package runtime

import (
	"fmt"
	"os"
	"paral/core/metadata"
)

type Reporter struct {
	metadata *metadata.Metadata
}

func NewReporter(metadata *metadata.Metadata) *Reporter {
	return &Reporter{
		metadata: metadata,
	}
}

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
func (c *Reporter) formatMessage(msgType, msg string, metadata *metadata.Metadata) string {
	var color, header string
	switch msgType {
	case "Error":
		color = Red
		header = "[Error]"
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
	file := c.metadata.Filename
	line := 0
	column := 0
	if metadata != nil {
		line = metadata.Line
		column = metadata.Column
	}

	if file != "" {
		if line > 0 || column > 0 {
			location = fmt.Sprintf("%s%s:%d:%d\n\t%s ", Blue, file, line, column, Reset)
		} else {
			location = fmt.Sprintf("%s%s\n\t%s ", Blue, file, Reset)
		}
	}

	return fmt.Sprintf(
		"%s%s%s%s %s%s\n",
		color, Bold, header, Reset, location, msg,
	)
}

// ThrowSyntaxError prints a formatted syntax error and exits.
func (c *Reporter) ThrowSyntaxError(msg string, metadata *metadata.Metadata) {
	_, _ = fmt.Fprintln(os.Stderr, c.formatMessage("Syntax Error", msg, metadata))
	os.Exit(1)
}

// ThrowRuntimeError prints a formatted runtime error and exits.
func (c *Reporter) ThrowRuntimeError(msg string, metadata *metadata.Metadata) {
	_, _ = fmt.Fprintln(os.Stderr, c.formatMessage("Runtime Error", msg, metadata))
	os.Exit(1)
}

// ThrowErrorMessage prints a formatted error message.
func (c *Reporter) ThrowErrorMessage(msg string) {
	_, _ = fmt.Fprintln(os.Stderr, c.formatMessage("Error", msg, nil))
	os.Exit(1)
}

// Warn prints a formatted warning message without exiting.
func (c *Reporter) Warn(msg string, metadata *metadata.Metadata) {
	_, _ = fmt.Fprintln(os.Stderr, c.formatMessage("Warning", msg, metadata))
}

// Info prints a formatted informational message.
func (c *Reporter) Info(msg string, metadata *metadata.Metadata) {
	_, _ = fmt.Fprintln(os.Stdout, c.formatMessage("Info", msg, metadata))
}

// ThrowSyntaxErrorSimple prints a syntax error without file or line info and exits.
func (c *Reporter) ThrowSyntaxErrorSimple(msg string, metadata *metadata.Metadata) {
	c.ThrowSyntaxError(msg, metadata)
}

// ThrowRuntimeErrorSimple prints a runtime error without file or line info and exits.
func (c *Reporter) ThrowRuntimeErrorSimple(msg string, metadata *metadata.Metadata) {
	c.ThrowRuntimeError(msg, metadata)
}

// WarnSimple prints a warning without file or line info.
func (c *Reporter) WarnSimple(msg string, metadata *metadata.Metadata) {
	c.Warn(msg, metadata)
}

// InfoSimple prints an info message without file or line info.
func (c *Reporter) InfoSimple(msg string, metadata *metadata.Metadata) {
	c.Info(msg, metadata)
}
