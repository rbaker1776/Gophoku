package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "sudoku",
	Short: "A modern Sudoku CLI for generating and solving puzzles",
	Long:  `Sudoku is a modern command-line tool for generating and solving Sudoku puzzles with customizable difficulty and reproducible results.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
