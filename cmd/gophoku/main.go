package main

import (
	"flag"
	"fmt"
	"gophoku/pkg/gophoku"
	"os"
    "time"
)

var elapsed string

func main() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	hintCount   := generateCmd.Int("hintCount", gophoku.DefaultHintCount, "Number of hints in the puzzle")
	numPuzzles  := generateCmd.Int("n", 1, "Number of puzzles to generate")

    prettyPrint := generateCmd.Bool("pretty", false, "Pretty print the puzzles")
    showStats   := generateCmd.Bool("runtimestats", false, "Show runtime stats")

	// Check if a subcommand was provided
	if len(os.Args) < 2 {
		printUsage()
        exit(1)
	}

	switch os.Args[1] {
	case "generate":
		generateCmd.Parse(os.Args[2:])
        gophoku.RuntimeStatsEnabled = *showStats
        start := time.Now()
		handleGenerate(*hintCount, *numPuzzles, *prettyPrint)
        elapsed = time.Since(start).String()
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printUsage()
        exit(1)
	}

    exit(0) 
}

func exit(exitCode int) {
    if gophoku.RuntimeStatsEnabled {
        printStats()
    }
    os.Exit(exitCode)
}

func handleGenerate(hintCount, numPuzzles int, prettyPrint bool) {
	// Validate hint count
	if hintCount < gophoku.MinValidHints || hintCount > gophoku.MaxValidHints {
		fmt.Fprintf(os.Stderr, "Error: hintCount must be between %d and %d\n", gophoku.MinValidHints, gophoku.MaxValidHints)
        exit(1)
	}
	if numPuzzles < 1 {
		fmt.Fprintln(os.Stderr, "Error: number of puzzles must be at least 1")
        exit(1)
	}

	for i := 0; i < numPuzzles; i++ {
		puzzle := gophoku.NewPuzzleWithHints(hintCount)
        if prettyPrint {
            fmt.Println(puzzle.Board.String())
        } else {
            fmt.Println(puzzle.Board.CompressedString())
        }
	}
}

func printUsage() {
	fmt.Println("Gophoku - A Sudoku puzzle generator")
	fmt.Println("\nUsage:")
	fmt.Println("  gophoku <command> [options]")
	fmt.Println("\nCommands:")
	fmt.Println("  generate     Generate Sudoku puzzles")
	fmt.Println("  help         Show this help message")
	fmt.Println("\nGenerate Options:")
	fmt.Println("  -hintCount int")
	fmt.Printf("       Number of hints in the puzzle (default %d, range %d-%d)\n", gophoku.DefaultHintCount, gophoku.MinValidHints, gophoku.MaxValidHints)
	fmt.Println("  -n int")
	fmt.Println("       Number of puzzles to generate (default 1)")
    fmt.Println("\nGeneric Options:")
    fmt.Println("  -pretty bool")
    fmt.Println("       Pretty print the puzzles")
    fmt.Println("  -runtimestats bool")
    fmt.Println("       Print runtime stats")
}

func printStats() {
    fmt.Println("\nGophoku runtime stats")
    fmt.Printf("\nTotal board instances created: %d\n", gophoku.TotalBoardCount)
    fmt.Printf("Maximum board instances concurrent: %d\n", gophoku.MaxBoardCopies)
    fmt.Printf("\nTook %s\n", elapsed)
}
