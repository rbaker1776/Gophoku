package main

import (
	"flag"
	"fmt"
	"gophoku/pkg/gophoku"
	"os"
)

func main() {
	// Define subcommands
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	hintCount := generateCmd.Int("hintCount", 30, "Number of hints in the puzzle (17-81)")
	numPuzzles := generateCmd.Int("n", 1, "Number of puzzles to generate")
	verbose := generateCmd.Bool("v", false, "Verbose output (show full puzzle boards)")
	compressedOnly := generateCmd.Bool("c", false, "Output compressed format only")

	// Check if a subcommand was provided
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Parse the subcommand
	switch os.Args[1] {
	case "generate":
		generateCmd.Parse(os.Args[2:])
		handleGenerate(*hintCount, *numPuzzles, *verbose, *compressedOnly)
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func handleGenerate(hintCount, numPuzzles int, verbose, compressedOnly bool) {
	// Validate hint count
	if hintCount < gophoku.MinValidHints || hintCount > gophoku.MaxValidHints {
		fmt.Fprintf(os.Stderr, "Error: hintCount must be between %d and %d\n", 
			gophoku.MinValidHints, gophoku.MaxValidHints)
		os.Exit(1)
	}

	// Validate number of puzzles
	if numPuzzles < 1 {
		fmt.Fprintln(os.Stderr, "Error: number of puzzles must be at least 1")
		os.Exit(1)
	}

	// Generate puzzles
	for i := 0; i < numPuzzles; i++ {
		puzzle := gophoku.NewPuzzleWithHints(hintCount)
		
		if compressedOnly {
			// Just output the compressed string
			fmt.Println(puzzle.Board.CompressedString())
		} else if verbose {
			// Show full puzzle board
			if numPuzzles > 1 {
				fmt.Printf("\n=== Puzzle %d/%d ===\n", i+1, numPuzzles)
			}
			fmt.Println(puzzle.Board.String())
			fmt.Printf("Compressed: %s\n", puzzle.Board.CompressedString())
		} else {
			// Default: compressed string with minimal formatting
			if numPuzzles > 1 {
				fmt.Printf("Puzzle %d: ", i+1)
			}
			fmt.Println(puzzle.Board.CompressedString())
		}
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
	fmt.Printf("        Number of hints in the puzzle (default 30, range %d-%d)\n", 
		gophoku.MinValidHints, gophoku.MaxValidHints)
	fmt.Println("  -n int")
	fmt.Println("        Number of puzzles to generate (default 1)")
	fmt.Println("  -v")
	fmt.Println("        Verbose output (show full puzzle boards)")
	fmt.Println("  -c")
	fmt.Println("        Compressed output only (no labels)")
	fmt.Println("\nExamples:")
	fmt.Println("  gophoku generate --hintCount 30 -n 100")
	fmt.Println("  gophoku generate -hintCount 25 -v")
	fmt.Println("  gophoku generate -n 5 -c")
}
