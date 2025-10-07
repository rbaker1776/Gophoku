package gophoku

import (
	"gophoku/internal/rng"
)

type Generator struct {
	board *Board
}

func NewGenerator() *Generator {
	return &Generator{board: NewBoard()}
}

// Generate creates a new puzzle with the specified number of hints.
// Returns a Puzzle containing both the initial state and its solution.
func (g *Generator) Generate(hintCount int) *Puzzle {
	puzzle := NewPuzzle(NewBoard())

	// Check that the number of hints can form a valid Sudoku puzzle
	if hintCount < MinValidHints || hintCount > MaxValidHints {
		return puzzle
	}

	for i := 0; g.board.HintCount() != hintCount && i < MaxGenerationAttempts; i++ {
		g.board = NewBoard()

		// Generate a complete solution
		if !g.board.Solve() {
			return puzzle
		}

		// Store the solution
		puzzle.Solution = g.board.Copy()

		// Remove cells while maintaining a unique solution
		puzzle.Hints = CellCount - g.removeCells(CellCount-hintCount)
	}

	puzzle.Board = g.board
	return puzzle
}

// removeCells attempts to remove removeCount cells from the board while maintaining a unique solution.
// Returns the number of cells actually removed.
// The final board state is guaranteed to have one unique solution.
func (g *Generator) removeCells(removeCount int) int {
	// Try removing tiles in a random order
	removeQueue := rng.ShuffledTiles()

	removed := 0
	for i := 0; i < len(removeQueue) && removed < removeCount; i++ {
		row, col := removeQueue[i][0], removeQueue[i][1]
		if g.board[row][col] == EmptyCell {
			// The tile was already removed
			continue
		}

		num := g.board[row][col]
		g.board[row][col] = EmptyCell

		testBoard := g.board.Copy()
		if testBoard.HasUniqueSolution() {
			removed++
		} else {
			// This tile is integral to having a unique solution
			g.board[row][col] = num
		}
	}

	return removed
}
