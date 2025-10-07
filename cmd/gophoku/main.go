package main

import (
	"flag"
	"fmt"
	"gophoku/pkg/gophoku"
	"os"
)

func main() {
    // Subcommand: generate
	generateCmd    := flag.NewFlagSet("generate", flag.ExitOnError)
	hintCount      := generateCmd.Int("hintCount", 30, "Number of hints in the puzzle")
	numPuzzles     := generateCmd.Int("n", 1, "Number of puzzles to generate")

	// Check if a subcommand was provided
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		generateCmd.Parse(os.Args[2:])
		handleGenerate(*hintCount, *numPuzzles)
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func handleGenerate(hintCount, numPuzzles int) {
	// Validate hint count
	if hintCount < gophoku.MinValidHints || hintCount > gophoku.MaxValidHints {
		fmt.Fprintf(os.Stderr, "Error: hintCount must be between %d and %d\n", gophoku.MinValidHints, gophoku.MaxValidHints)
		os.Exit(1)
	}
	if numPuzzles < 1 {
		fmt.Fprintln(os.Stderr, "Error: number of puzzles must be at least 1")
		os.Exit(1)
	}

	for i := 0; i < numPuzzles; i++ {
		puzzle := gophoku.NewPuzzleWithHints(hintCount)
	    fmt.Println(puzzle.Board.CompressedString())
	}
}

func printUsage() {
	fmt.Println("Gophoku - A Sudoku puzzle generator")
	fmt.Println("\nUsage:")
	fmt.Println("  gophoku <command> [options]")
	fmt.Println("\nCommands:")
	fmt.Println("  generate    Generate Sudoku puzzles")
	fmt.Println("  help        Show this help message")
	fmt.Println("\nGenerate Options:")
	fmt.Println("  -hintCount int")
	fmt.Printf("        Number of hints in the puzzle (default 30, range %d-%d)\n", gophoku.MinValidHints, gophoku.MaxValidHints)
	fmt.Println("  -n int")
	fmt.Println("        Number of puzzles to generate (default 1)")
}
