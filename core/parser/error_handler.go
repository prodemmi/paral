package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"

	"os"
	"strings"
	"unicode/utf8"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	hasErrors bool
	content   string
	filename  string
}

func NewErrorListener(content, filename string) *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		content:              content,
		filename:             filename,
	}
}

func (el *ErrorListener) HasErrors() bool {
	return el.hasErrors
}

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.hasErrors = true

	// Enhanced error formatting
	el.formatError(line, column, msg, offendingSymbol)
	os.Exit(1)
}

func (el *ErrorListener) formatError(line, column int, msg string, offendingSymbol interface{}) {
	// Color codes for better visibility
	const (
		Red    = "\033[31m"
		Yellow = "\033[33m"
		Blue   = "\033[34m"
		Bold   = "\033[1m"
		Reset  = "\033[0m"
	)

	// Header with file location
	filename := el.filename
	if filename == "" {
		filename = "<input>"
	}

	fmt.Fprintf(os.Stderr, "%serror%s: %s\n", Red+Bold, Reset, msg)
	fmt.Fprintf(os.Stderr, "  %s--> %s%s:%d:%d%s\n", Blue, Bold, filename, line, column+1, Reset)

	if el.content == "" {
		return
	}

	lines := strings.Split(el.content, "\n")
	if line < 1 || line > len(lines) {
		return
	}

	// Calculate line number width for proper alignment
	maxLineNum := line + 2
	if maxLineNum > len(lines) {
		maxLineNum = len(lines)
	}
	lineNumWidth := len(fmt.Sprintf("%d", maxLineNum))

	// Show context: previous line (if exists)
	if line > 1 {
		prevLine := lines[line-2]
		fmt.Fprintf(os.Stderr, "%s%*d%s │ %s\n",
			Blue, lineNumWidth, line-1, Reset, prevLine)
	}

	// Show the problematic line
	errorLine := lines[line-1]
	fmt.Fprintf(os.Stderr, "%s%*d%s │ %s\n",
		Blue, lineNumWidth, line, Reset, errorLine)

	// Create error pointer with context
	el.createErrorPointer(errorLine, column, lineNumWidth, offendingSymbol)

	// Show context: next line (if exists)
	if line < len(lines) {
		nextLine := lines[line]
		fmt.Fprintf(os.Stderr, "%s%*d%s │ %s\n",
			Blue, lineNumWidth, line+1, Reset, nextLine)
	}

	fmt.Fprintf(os.Stderr, "\n")
}

func (el *ErrorListener) createErrorPointer(line string, column, lineNumWidth int, offendingSymbol interface{}) {
	const (
		Red    = "\033[31m"
		Yellow = "\033[33m"
		Blue   = "\033[34m"
		Bold   = "\033[1m"
		Reset  = "\033[0m"
	)

	// Calculate visual column position (handles tabs and unicode)
	visualColumn := el.calculateVisualColumn(line, column)

	// Create the pointer line
	spaces := strings.Repeat(" ", lineNumWidth)
	fmt.Fprintf(os.Stderr, "%s%s%s │ ", Blue, spaces, Reset)

	// Add spacing to reach the error position
	if visualColumn > 0 {
		// Replace tabs with spaces for proper alignment
		prefix := el.replaceTabsWithSpaces(line[:min(column, len(line))])
		prefixSpaces := strings.Repeat(" ", utf8.RuneCountInString(prefix))
		fmt.Fprintf(os.Stderr, "%s", prefixSpaces)
	}

	// Determine the length of the error highlight
	errorLength := el.getErrorLength(line, column, offendingSymbol)

	// Create the error pointer
	if errorLength <= 1 {
		fmt.Fprintf(os.Stderr, "%s%s^%s", Red+Bold, "", Reset)
	} else {
		pointer := "^" + strings.Repeat("~", errorLength-1)
		fmt.Fprintf(os.Stderr, "%s%s%s%s", Red+Bold, pointer, Reset, "")
	}

	// Add helpful hint if available
	hint := el.getErrorHint(offendingSymbol)
	if hint != "" {
		fmt.Fprintf(os.Stderr, " %s%s%s", Yellow, hint, Reset)
	}

	fmt.Fprintf(os.Stderr, "\n")
}

func (el *ErrorListener) calculateVisualColumn(line string, column int) int {
	if column >= len(line) {
		return utf8.RuneCountInString(line)
	}

	prefix := line[:column]
	// Convert tabs to 4 spaces for calculation
	expanded := strings.ReplaceAll(prefix, "\t", "    ")
	return utf8.RuneCountInString(expanded)
}

func (el *ErrorListener) replaceTabsWithSpaces(s string) string {
	return strings.ReplaceAll(s, "\t", "    ")
}

func (el *ErrorListener) getErrorLength(line string, column int, offendingSymbol interface{}) int {
	if offendingSymbol == nil {
		return 1
	}

	// Try to get token information
	if token, ok := offendingSymbol.(antlr.Token); ok {
		tokenText := token.GetText()
		if tokenText != "" && tokenText != "<EOF>" {
			return utf8.RuneCountInString(tokenText)
		}
	}

	// Try to find the next whitespace or special character
	if column < len(line) {
		remaining := line[column:]
		for i, r := range remaining {
			if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
				if i > 0 {
					return i
				}
				break
			}
		}
		// If no whitespace found, highlight the rest of the identifier
		words := strings.FieldsFunc(remaining, func(r rune) bool {
			return r == ' ' || r == '\t' || r == '(' || r == ')' || r == '{' || r == '}' ||
				r == '[' || r == ']' || r == ';' || r == ',' || r == '.'
		})
		if len(words) > 0 && len(words[0]) > 1 {
			return utf8.RuneCountInString(words[0])
		}
	}

	return 1
}

func (el *ErrorListener) getErrorHint(offendingSymbol interface{}) string {
	if offendingSymbol == nil {
		return ""
	}

	if token, ok := offendingSymbol.(antlr.Token); ok {
		tokenText := token.GetText()
		tokenType := token.GetTokenType()

		switch tokenText {
		case "<EOF>":
			return "unexpected end of file"
		case ";":
			return "unexpected semicolon"
		case "{":
			return "unmatched opening brace"
		case "}":
			return "unmatched closing brace"
		case "(":
			return "unmatched opening parenthesis"
		case ")":
			return "unmatched closing parenthesis"
		default:
			if tokenType == antlr.TokenEOF {
				return "unexpected end of input"
			}
			if tokenText != "" {
				return fmt.Sprintf("unexpected token '%s'", tokenText)
			}
		}
	}

	return ""
}
