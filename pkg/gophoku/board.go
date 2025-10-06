package gophoku

import (
    "fmt"
)

// Board represents a 9x9 Sudoku board.
// Each cell contains a value 0-9, where 0 represents an empty cell.
// Valid Sudoku numbers are 1-9.
type Board [9][9]int

// NewBoard creates and returns a new empty Sudoku board.
func NewBoard() *Board {
    return &Board{}
}

// Copy creates a deep copy of the board.
// The returned board is independent of the original and can be modified without affecting the source board.
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
            if b[row][col] == 0 {
                count++
            }
        }
    }
    return count
}

// HintCount returns the number of filled cells on the board.
// This is the complement of EmptyCount and always satisfies HintCount() + EmptyCount() == 81.
func (b *Board) HintCount() int {
    return 9 * 9 - b.EmptyCount()
}

// IsSolved returns true if the board is completely filled and valid according to Sudoku rules.
// A solved board has no empty cells and satisfies all Sudoku constraints.
func (b *Board) IsSolved() bool {
    return b.EmptyCount() == 0 && b.IsValid()
}

// IsValid checks if the current board state satisfies all Sudoku constraints.
// Empty cells are ignored in validation.
// NOTE: This method temporarily modifies and restores cells during validation.
func (b *Board) IsValid() bool {
    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            if b[row][col] < 0 || b[row][col] > 9 {
                return false
            }
            if b[row][col] != 0 {
                num := b[row][col]
                // Temporarily modify the board and see if placement is valid
                b[row][col] = 0
                if !b.CanPlace(row, col, num) {
                    return false
                }
                // Restore the board
                b[row][col] = num
            }
        }
    }
    return true
}

// CanPlace determines if a number can be legally placed at the specified position.
// Returns true if the placement is legal, false otherwise.
func (b *Board) CanPlace(row, col, num int) bool {
    // Check input bounds
    if row < 0 || row >= 9 || col < 0 || col >= 9 {
        return false
    }

    // Check if the square is already occupied
    if b[row][col] != 0 {
        return false
    }

    // Check row for duplicate
    for c := 0; c < 9; c++ {
        if b[row][c] == num {
            return false
        }
    }

    // Check column for duplicate
    for r := 0; r < 9; r++ {
        if b[r][col] == num {
            return false
        }
    }

    // Check 3x3 box for duplicate
    startRow := int(row / 3) * 3
    startCol := int(col / 3) * 3
    for r := startRow; r < startRow + 3; r++ {
        for c := startCol; c < startCol + 3; c++ {
            if b[r][c] == num {
                return false
            }
        }
    }

    return true
}

// String returns a human-readable string representation of the board.
// Example output:
//  +-------+-------+-------+
//  | . . . | . 4 7 | 5 . . |
//  | . . 3 | . . . | . . 4 |
//  | 1 . . | . . . | . . . |
//  +-------+-------+-------+
//  | . . . | . 9 . | 3 1 . |
//  | 5 . . | 3 6 . | . . . |
//  | . 9 1 | . 5 . | . . 6 |
//  +-------+-------+-------+
//  | . . . | . 7 . | 8 . . |
//  | 6 . . | 1 . . | . . 2 |
//  | . . . | . . 8 | . 4 . |
//  +-------+-------+-------+
func (b *Board) String() string {
    s, l := "", "+-------+-------+-------+\n"
	for i := 0; i < 9; i++ {
		if i % 3 == 0 {
            s += l
        }
		s += "| "
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
                s += ". "
            } else {
                s += fmt.Sprintf("%d ", b[i][j])
            }
			if (j + 1) % 3 == 0 {
                s += "| "
            }
		}
		s += "\n"
	}
	return s + l
}
