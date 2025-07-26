package main

import (
	"flag"
	"fmt"
	"os"

	parser "paral/antlr/antlr"
	"paral/core"
	"paral/server"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	hasErrors bool
}

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.hasErrors = true
	_, _ = fmt.Fprintf(os.Stderr, "line %d:%d %s\n", line, column, msg)
}

type ExecutionOptions struct {
	Graph      bool
	Describe   bool
	Summary    bool
	Verbose    bool
	Silent     bool
	Short      bool
	DryRun     bool
	NoColor    bool
	Profile    bool
	Version    bool
	Help       bool
	OutputFile string
	Timeout    int
}

func parseFlags() *ExecutionOptions {
	opts := &ExecutionOptions{}

	// Existing flags
	flag.BoolVar(&opts.Graph, "graph", false, "show file graph")
	flag.BoolVar(&opts.Describe, "describe", false, "describe the paral file")

	// New output control flags
	flag.BoolVar(&opts.Summary, "s", false, "only show execution summary")
	flag.BoolVar(&opts.Summary, "summary", false, "only show execution summary")

	flag.BoolVar(&opts.Verbose, "v", false, "verbose output with detailed information")
	flag.BoolVar(&opts.Verbose, "verbose", false, "verbose output with detailed information")

	flag.BoolVar(&opts.Silent, "q", false, "silent mode - suppress all output except errors")
	flag.BoolVar(&opts.Silent, "silent", false, "silent mode - suppress all output except errors")

	flag.BoolVar(&opts.Short, "short", false, "short output format")

	// Execution control flags
	flag.BoolVar(&opts.DryRun, "dry-run", false, "show what would be executed without running commands")
	flag.BoolVar(&opts.DryRun, "n", false, "show what would be executed without running commands")

	// Output formatting flags
	flag.BoolVar(&opts.NoColor, "no-color", false, "disable colored output")
	flag.StringVar(&opts.OutputFile, "o", "", "write output to file instead of stdout")
	flag.StringVar(&opts.OutputFile, "output", "", "write output to file instead of stdout")

	// Performance and debugging flags
	flag.BoolVar(&opts.Profile, "profile", false, "show execution time profiling")

	// Parallel execution flags
	flag.IntVar(&opts.Timeout, "timeout", 0, "command timeout in seconds (0 = no timeout)")

	// Utility flags
	flag.BoolVar(&opts.Version, "version", false, "show version information")
	flag.BoolVar(&opts.Help, "h", false, "show help information")
	flag.BoolVar(&opts.Help, "help", false, "show help information")

	flag.Parse()

	return opts
}

func showVersion() {
	fmt.Println("Paral v1.0.0")
	fmt.Println("A parallel task execution tool")
}

func showHelp() {
	fmt.Println("Paral - Parallel Task Execution Tool")
	fmt.Println()
	fmt.Println("Usage: paral [options] <file>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  File Operations:")
	fmt.Println("    --graph              Show file dependency graph")
	fmt.Println("    --describe           Describe the paral file structure")
	fmt.Println()
	fmt.Println("  Output Control:")
	fmt.Println("    -s, --summary        Only show execution summary")
	fmt.Println("    -v, --verbose        Verbose output with detailed information")
	fmt.Println("    -q, --silent         Silent mode - suppress all output except errors")
	fmt.Println("    --short              Short output format")
	fmt.Println("    --no-color           Disable colored output")
	fmt.Println("    -o, --output FILE    Write output to file instead of stdout")
	fmt.Println()
	fmt.Println("  Execution Control:")
	fmt.Println("    -n, --dry-run        Show what would be executed without running")
	fmt.Println("    --timeout SEC        Command timeout in seconds (0 = no timeout)")
	fmt.Println()
	fmt.Println("  Debugging & Performance:")
	fmt.Println("    --profile            Show execution time profiling")
	fmt.Println("    --log-level LEVEL    Set log level (debug, info, warn, error)")
	fmt.Println()
	fmt.Println("  Information:")
	fmt.Println("    --version            Show version information")
	fmt.Println("    -h, --help           Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  paral build.paral                    # Execute the file")
	fmt.Println("  paral -v build.paral                 # Execute with verbose output")
	fmt.Println("  paral --dry-run build.paral          # Show what would be executed")
	fmt.Println("  paral --describe build.paral         # Show file structure")
}

func validateOptions(opts *ExecutionOptions) error {
	// Check for conflicting options
	if opts.Silent && opts.Verbose {
		return fmt.Errorf("cannot use both --silent and --verbose flags")
	}

	if opts.Silent && opts.Summary {
		return fmt.Errorf("cannot use both --silent and --summary flags")
	}

	if opts.Timeout < 0 {
		return fmt.Errorf("timeout cannot be negative")
	}

	return nil
}

func main() {
	opts := parseFlags()

	// Handle utility flags first
	if opts.Help {
		showHelp()
		return
	}

	if opts.Version {
		showVersion()
		return
	}

	// Validate options
	if err := validateOptions(opts); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Check for input file
	if flag.NArg() < 1 {
		core.ThrowSyntaxErrorSimple("No input file provided")
	}

	inputFile := flag.Arg(0)

	// Parse the input file
	input, err := antlr.NewFileStream(inputFile)
	if err != nil {
		core.ThrowSyntaxErrorSimple("Failed to read input file: " + err.Error())
	}

	lexer := parser.NewParalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	paralParser := parser.NewParalParser(stream)

	// Add custom error listener
	errorListener := &ErrorListener{}
	paralParser.RemoveErrorListeners()
	paralParser.AddErrorListener(errorListener)

	// Create core instance and parse
	coreInstance := core.NewCore(inputFile, input.String())
	visitor := &core.CoreVisitor{Core: coreInstance}
	antlr.ParseTreeWalkerDefault.Walk(visitor, paralParser.Program())

	if errorListener.hasErrors {
		_, _ = fmt.Fprintf(os.Stderr, "Parsing failed, exiting.\n")
		os.Exit(1)
	}

	// Execute based on options
	if opts.Graph {
		server.NewGraphServer(coreInstance)
	} else if opts.Describe {
		coreInstance.Describe()
	} else {
		// Create execution configuration
		execConfig := &core.ExecutionConfig{
			SummaryOnly: opts.Summary,
			Verbose:     opts.Verbose,
			Silent:      opts.Silent,
			Short:       opts.Short,
			DryRun:      opts.DryRun,
			NoColor:     opts.NoColor,
			Profile:     opts.Profile,
			Timeout:     opts.Timeout,
			OutputFile:  opts.OutputFile,
		}

		coreInstance.ExecWithConfig(execConfig)
	}
}
