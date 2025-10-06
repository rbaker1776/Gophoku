package gophoku

import (
    "fmt"
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
func (g *Generator) Generate(hintCount int) (*Puzzle, error) {
    // Generate a complete solution
    if !g.board.Solve() {
        return nil, fmt.Errorf("failed to generate complete solution")
    }

    // Store the solution
    solution := g.board.Copy()

    // Remove cells while maintaining a unique solution
    removedCount := g.removeCells(9 * 9 - hintCount)

    puzzle := &Puzzle{
        Board:    g.board,
        Solution: solution,
        Hints:    9 * 9 - removedCount,
    }
    return puzzle, nil
}

func (g *Generator) removeCells(removeCount int) int {
    removeQueue := rng.ShuffledTiles()

    removed := 0
    for i := 0; i < len(removeQueue) && removed < removeCount; i++ {
        row, col := removeQueue[i][0], removeQueue[i][1]
        if g.board[row][col] == 0 {
            continue
        }

        num := g.board[row][col]
        g.board[row][col] = 0

        testBoard := g.board.Copy()
        if testBoard.HasUniqueSolution() {
            removed++
        } else {
            g.board[row][col] = num
        }
    }

    return removed
}
