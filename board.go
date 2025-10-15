package main

import "fmt"

const (
    EmptyCell   = 0
    InvalidCell = -1
)

// Board represents a 9x9 Sudoku board.
// Each cell contains a value 1-9 | EmptyCell.
type Board [9][9]int

// New creates and returns a new empty Sudoku board.
func NewBoard() *Board {
    // NOTE: This assumes EmptyCell == 0
	return &Board{}
}

// Copy creates a deep copy of the board.
// The returned board is independent of the original.
func (b *Board) Copy() *Board {
	newBoard := NewBoard()
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
// String returns a human-readable string representation of the board.
// Example output:
//
//	+-------+-------+-------+
//	| . . . | . 4 7 | 5 . . |
//	| . . 3 | . . . | . . 4 |
//	| 1 . . | . . . | . . . |
//	+-------+-------+-------+
//	| . . . | . 9 . | 3 1 . |
//	| 5 . . | 3 6 . | . . . |
//	| . 9 1 | . 5 . | . . 6 |
//	+-------+-------+-------+
//	| . . . | . 7 . | 8 . . |
//	| 6 . . | 1 . . | . . 2 |
//	| . . . | . . 8 | . 4 . |
//	+-------+-------+-------+
func (b *Board) String() string {
	s, l := "", "+-------+-------+-------+\n"
	for row := 0; row < 9; row++ {
		if row % 3 == 0 {
			s += l
		}
		s += "| "
		for col := 0; col < 9; col++ {
			if b[row][col] == EmptyCell {
				s += ". "
			} else {
				s += fmt.Sprintf("%d ", b[row][col])
			}
			if (col+1)%3 == 0 {
				s += "| "
			}
		}
		s += "\n"
	}
	return s + l
}
