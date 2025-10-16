package main

import (
	"fmt"

	"sudoku/internal/generator"
)

func main() {
	board, solution, err := generator.GenerateWithClueCount(21)
	if err == nil {
		fmt.Println(board.Format())
		fmt.Println(solution.Format())
	} else {
		fmt.Println(err)
	}
}
