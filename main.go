package main

import (
	"fmt"
	"io"
	"os"

	"github.com/olgasafonova/md-to-mediawiki-go/md-to-mediawiki-plus/converter"
	flag "github.com/spf13/pflag"
)

const version = "1.0.0"

func main() {
	// Define CLI flags
	var (
		inputFile   string
		outputFile  string
		withCSS     bool
		concurrent  bool
		showVersion bool
		showHelp    bool
	)

	flag.StringVarP(&inputFile, "input", "i", "", "Input Markdown file")
	flag.StringVarP(&outputFile, "output", "o", "", "Output MediaWiki file (default: stdout)")
	flag.BoolVar(&withCSS, "with-css", false, "Include CSS styling in output")
	flag.BoolVarP(&concurrent, "concurrent", "c", false, "Use concurrent processing for large files (>50KB)")
	flag.BoolVarP(&showVersion, "version", "v", false, "Show version information")
	flag.BoolVarP(&showHelp, "help", "h", false, "Show help information")

	flag.Parse()

	// Show version
	if showVersion {
		fmt.Printf("md-to-mediawiki-go version %s\n", version)
		fmt.Println("Markdown to MediaWiki converter with Tietoevry branding")
		os.Exit(0)
	}

	// Show help
	if showHelp || inputFile == "" {
		showUsage()
		os.Exit(0)
	}

	// Read input file
	var inputData []byte
	var err error

	if inputFile == "-" {
		// Read from stdin
		inputData, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Read from file
		inputData, err = os.ReadFile(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file '%s': %v\n", inputFile, err)
			os.Exit(1)
		}
	}

	// Convert
	config := converter.Config{
		AddStyling: withCSS,
		Concurrent: concurrent,
	}

	output := converter.Convert(string(inputData), config)

	// Write output
	if outputFile == "" || outputFile == "-" {
		// Write to stdout
		fmt.Print(output)
	} else {
		// Write to file
		err = os.WriteFile(outputFile, []byte(output), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file '%s': %v\n", outputFile, err)
			os.Exit(1)
		}

		cssNote := ""
		if withCSS {
			cssNote = " (with CSS)"
		}
		concurrentNote := ""
		if concurrent {
			concurrentNote = " [concurrent mode]"
		}
		fmt.Printf("✅ Converted '%s' -> '%s'%s%s\n", inputFile, outputFile, cssNote, concurrentNote)
	}
}

func showUsage() {
	fmt.Println("Markdown to MediaWiki Converter")
	fmt.Println("Converts Obsidian-style Markdown to MediaWiki format with Tietoevry branding")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  md-to-mediawiki-go -i <input.md> [-o <output.txt>] [options]")
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  # Basic conversion")
	fmt.Println("  md-to-mediawiki-go -i example.md -o output.txt")
	fmt.Println()
	fmt.Println("  # With CSS styling")
	fmt.Println("  md-to-mediawiki-go -i example.md -o output.txt --with-css")
	fmt.Println()
	fmt.Println("  # Concurrent processing for large files")
	fmt.Println("  md-to-mediawiki-go -i large-doc.md -o output.txt -c")
	fmt.Println()
	fmt.Println("  # Read from stdin, write to stdout")
	fmt.Println("  cat example.md | md-to-mediawiki-go -i - > output.txt")
	fmt.Println()
	fmt.Println("Features:")
	fmt.Println("  ✅ Purple gradient headings (Tietoevry colors)")
	fmt.Println("  ✅ Peach + Hero Blue inline code styling")
	fmt.Println("  ✅ Reversed changelog (newest first)")
	fmt.Println("  ✅ Green emoji checkmarks")
	fmt.Println("  ✅ Clean table handling")
	fmt.Println("  ✅ WCAG AA compliant colors")
	fmt.Println("  ✅ Concurrent processing for performance")
	fmt.Println()
}
