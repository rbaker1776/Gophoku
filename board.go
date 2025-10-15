package main

const (
    EmptyCell   = 0
    InvalidCell = -1
)

// Board represents a 9x9 Sudoku board.
// Each cell contains a value 1-9 | EmptyCell.
type Board [9][9]int

// New creates and returns a new empty Sudoku board.
func New() *Board {
    // NOTE: This assumes EmptyCell == 0
	return &Board{}
}

// Copy creates a deep copy of the board.
// The returned board is independent of the original.
func (b *Board) Copy() *Board {
	newBoard := New()
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			newBoard[row][col] = b[row][col]
		}
	}
	return newBoard
}

// EmptyCount returns the number of empty cells on the board.
// A completely empty board will return 81, a completely full board will return 0.
func (b *Board) EmptyCount() int {
	count := 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if b[row][col] == EmptyCell {
				count++
			}
		}
	}
	return count
}

// HintCount returns the number of filled cells on the board.
// HintCount() + EmptyCount() == 81.
func (b *Board) HintCount() int {
	return 81 - b.EmptyCount()
}
