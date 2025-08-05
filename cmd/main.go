package main

import (
	"flag"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/gosuri/uilive"
	"os"
	ant "paral/antlr/antlr"
	"paral/core/metadata"
	parser "paral/core/parser"
	"paral/core/runtime"
	"time"
)

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

func ParseFlags() *ExecutionOptions {
	opts := &ExecutionOptions{}

	// Existing flags
	flag.BoolVar(&opts.Graph, "graph", false, "show file graph")
	flag.BoolVar(&opts.Describe, "describe", false, "describe the paral file")

	// Output control flags
	flag.BoolVar(&opts.Summary, "s", false, "only show execution summary")
	flag.BoolVar(&opts.Summary, "summary", false, "only show execution summary")

	flag.BoolVar(&opts.Verbose, "v", false, "verbose output with detailed information")
	flag.BoolVar(&opts.Verbose, "verbose", false, "verbose output with detailed information")

	flag.BoolVar(&opts.Silent, "q", false, "silent mode - suppress all output except errors")
	flag.BoolVar(&opts.Silent, "silent", false, "silent mode - suppress all output except errors")

	flag.BoolVar(&opts.Short, "short", false, "short output format")

	// Execution control flags
	flag.BoolVar(&opts.DryRun, "dry-run", false, "show what would be executed without running pipelines")
	flag.BoolVar(&opts.DryRun, "n", false, "show what would be executed without running pipelines")

	// Output formatting flags
	flag.BoolVar(&opts.NoColor, "no-color", false, "disable colored output")
	flag.StringVar(&opts.OutputFile, "o", "", "write output to file instead of stdout")
	flag.StringVar(&opts.OutputFile, "output", "", "write output to file instead of stdout")

	// Performance and debugging flags
	flag.BoolVar(&opts.Profile, "profile", false, "show execution time profiling")

	flag.IntVar(&opts.Timeout, "timeout", 0, "pipeline timeout in seconds (0 = no timeout)")

	// Utility flags
	flag.BoolVar(&opts.Version, "version", false, "show version information")
	flag.BoolVar(&opts.Help, "h", false, "show help information")
	flag.BoolVar(&opts.Help, "help", false, "show help information")

	flag.Parse()

	return opts
}

func ShowVersion() {
	fmt.Println("Paral v1.0.0")
	fmt.Println("A task execution tool")
}

func ShowHelp() {
	fmt.Println("Paral - Task Execution Tool")
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

func ValidateOptions(opts *ExecutionOptions) error {
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

func ShowLoading(inputFile string, opts *ExecutionOptions) chan struct{} {
	if !opts.Silent && !opts.Short {
		done := make(chan struct{})
		go func() {
			writer := uilive.New()
			writer.Start()

			dots := []string{"..", "...", "."}
			ticker := time.NewTicker(500 * time.Millisecond)
			defer ticker.Stop()

			fmt.Fprintf(writer, "\rLoading %s%s\n", inputFile, dots[len(dots)-1])

			i := 0
			for {
				select {
				case <-done:
					fmt.Fprintf(writer, "\r%-100s\r", "") // Clear the line with enough padding
					writer.Flush()
					writer.Stop()
					return
				case <-ticker.C:
					fmt.Fprintf(writer, "\rLoading %s%s\n", inputFile, dots[i%len(dots)])
					writer.Flush()
					i++
				}
			}
		}()
		time.Sleep(1 * time.Second)
		return done
	}
	return nil
}

func clearLoading(loader chan struct{}) {
	close(loader)
	time.Sleep(200 * time.Millisecond) // Ensure cleanup completes
}

func main() {
	opts := ParseFlags()

	// Handle utility flags first
	if opts.Help {
		ShowHelp()
		return
	}

	if opts.Version {
		ShowVersion()
		return
	}

	// Validate options
	if err := ValidateOptions(opts); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Check for input file
	if flag.NArg() < 1 {
		fmt.Println("* No input file provided")
		ShowHelp()
		os.Exit(1)
	}

	inputFile := flag.Arg(0)

	// Start loading animation
	loader := ShowLoading(inputFile, opts)

	// Parse the input file
	input, err := antlr.NewFileStream(inputFile)
	globalMetadata := metadata.Metadata{
		Filename: inputFile,
		Content:  input.String(),
	}
	reporter := runtime.NewReporter(&globalMetadata)
	paralRuntime := runtime.NewRuntime(&globalMetadata, reporter)

	if err != nil {
		paralRuntime.Reporter.ThrowErrorMessage("Failed to read input file: " + err.Error())
	}

	lexer := ant.NewParalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	paralParser := ant.NewParalParser(stream)

	// Add custom error listener
	errorListener := parser.NewErrorListener(input.String(), inputFile)
	paralParser.RemoveErrorListeners()
	paralParser.AddErrorListener(errorListener)

	// Create core instance and parse
	visitor := &parser.Parser{Runtime: paralRuntime}
	antlr.ParseTreeWalkerDefault.Walk(visitor, paralParser.Program())

	// Stop loader immediately after parsing, before any execution output
	if !opts.Silent && !opts.Short {
		clearLoading(loader)
	}

	// Check for parsing errors
	if errorListener.HasErrors() {
		clearLoading(loader)
		_, _ = fmt.Fprintf(os.Stderr, "Parsing failed, exiting.\n")
		os.Exit(1)
	}

	// Execute based on options
	if opts.Graph {
		//server.NewGraphServer(paralRuntime)
	} else if opts.Describe {
		// TODO: implement new describer.go
		//coreInstance.Describe()
	} else {
		// Create execution configuration
		execConfig := &runtime.ExecutionConfig{
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
		executor := runtime.NewExecutor(paralRuntime)
		executor.ExecWithConfig(execConfig)
	}
}
