package gophoku

import (
	"gophoku/internal/rng"
)

// Puzzle represents a single Sudoku puzzle, with an initial state and a solution.
type Puzzle struct {
	Board    *Board
	Solution *Board
}

// NewPuzzle creates and returns a new empty Puzzle.
func NewPuzzle() *Puzzle {
    return &Puzzle{
        Board:    NewBoard(),
        Solution: NewBoard(),
    }
}

// NewPuzzleWithHints creates and returns a new Puzzle with hintCount hints.
// hintCount should be between MinValidHints and MaxValidHints.
func NewPuzzleWithHints(hintCount int) *Puzzle {
    puzzle := NewPuzzle()

	// Check that the number of hints can form a valid Sudoku puzzle
	if hintCount < MinValidHints || hintCount > MaxValidHints {
		return puzzle
	}

	for i := 0; i < MaxGenerationAttempts; i++ {
		puzzle.Board = NewBoard()

		// Generate a complete solution
		if !puzzle.Board.Solve() {
            continue
		}

		// Store the solution
		puzzle.Solution = puzzle.Board.Copy()

		// Remove cells while maintaining a unique solution
		if puzzle.Board.removeCells(CellCount-hintCount) {
            break
        }
	}

    return puzzle
}

// removeCells attempts to remove removeCount cells from the board while maintaining a unique solution.
// The final board state is guaranteed to be valid have one unique solution, even if it means removing fewer than removeCount cells.
func (b *Board) removeCells(removeCount int) bool {
    if removeCount == 0 {
        return true
    }

    hintTiles := b.HintTiles()
    // Try removing tiles in a random ordefr
    rng.ShuffleTiles(hintTiles)

    for _, tile := range(hintTiles) {
        row, col := tile[0], tile[1]
        num := b[row][col]
        b[row][col] = EmptyCell

        if b.HasUniqueSolution() && b.removeCells(removeCount - 1) {
            return true
        } else {
            b[row][col] = num
        }
    }

    return false
}
