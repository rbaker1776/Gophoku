package cmd

import (
	"fmt"
	"github.com/rybkr/sudoku/internal/generator"
	"github.com/spf13/cobra"
	"time"
)

var (
	numPuzzles int
	clueCount  int
	timeout    time.Duration
)

func init() {
	genCmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate Sudoku puzzles",
		Long: `Generate one or more Sudoku puzzles with a specified difficulty level.

Examples:
  sudoku gen --clueCount 40
  sudoku gen -n 5 --clueCount 30
  sudoku gen --clueCount 20 --timeout 15s`,
		RunE: runGen,
	}

	genCmd.Flags().IntVarP(&numPuzzles, "number", "n", 1, "Number of puzzles to generate")
	genCmd.Flags().IntVarP(&clueCount, "clueCount", "c", generator.DefaultClueCount, "Number of clues 17-80")
	genCmd.Flags().DurationVar(&timeout, "timeout", 10*time.Second, "Generation timeout per puzzle")

	rootCmd.AddCommand(genCmd)
}

func runGen(cmd *cobra.Command, args []string) error {
	for i := 0; i < numPuzzles; i++ {
		opts := generator.DefaultOptions(clueCount)
		opts.Timeout = timeout
		gen := generator.New(opts)

		puzzle, solution, err := gen.Generate()
		if err != nil {
			return fmt.Errorf("generation failed: %w", err)
		}

		fmt.Println("Puzzle:")
		fmt.Println(puzzle.Format())
		fmt.Println("\nSolution:")
		fmt.Println(solution.Format())
		fmt.Println()
	}

	return nil
}
