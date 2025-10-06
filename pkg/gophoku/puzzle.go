package gophoku

import (
    "gophoku/internal/rng"
)

type Puzzle struct {
    Board    *Board
    Solution *Board
    Hints    int
}

type Generator struct {
    board *Board
}

func NewGenerator(board *Board) *Generator {
    return &Generator{board: board}
}

// Generate creates a new puzzle with the specified number of hints.
// Returns a Puzzle containing both the initial state and its solution.
func (g *Generator) Generate(hintCount int) *Puzzle {
    puzzle := &Puzzle{
        Board: g.board.Copy(),
    }

    // Check that the number of hints can form a valid Sudoku puzzle
    if hintCount < 17 || hintCount > 81 || g.board.HintCount() > hintCount {
        return puzzle
    }

    for g.board.HintCount() != hintCount {
        // Generate a complete solution
        if !g.board.Solve() {
            return puzzle
        }

        // Store the solution
        puzzle.Solution = g.board.Copy()

        // Remove cells while maintaining a unique solution
        puzzle.Hints = 9 * 9 - g.removeCells(9 * 9 - hintCount)
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
        if g.board[row][col] == 0 {
            // The tile was already removed
            continue
        }

        num := g.board[row][col]
        g.board[row][col] = 0

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
