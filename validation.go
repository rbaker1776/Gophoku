package main

// IsValid reports whether the board state satisfies all Sudoku constraints.
// Empty cells are ignored in validation.
func (b *Board) IsValid() bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
            num := b.Get(row, col)
            if !(isValidNumber(num) || num == EmptyCell) {
				return false
			}
            if num != EmptyCell {
				if !b.canPlace(row, col, num) {
					return false
				}
			}
		}
	}
	return true
}

// canPlace reports whether a number can be legally placed at the specified position.
func (b *Board) canPlace(row, col, num int) bool {
	// Check input bounds
	if !isValidPosition(row, col) || !isValidNumber(num) {
		return false
	}

	// Check row for duplicate
	for c := 0; c < 9; c++ {
		if c != col && b.Get(row, c) == num {
			return false
		}
	}

	// Check column for duplicate
	for r := 0; r < 9; r++ {
		if r != row && b.Get(r, col) == num {
			return false
		}
	}

	// Check 3x3 box for duplicate
	startRow := int(row/3) * 3
	startCol := int(col/3) * 3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if (r != row || c != col) && b.Get(r, c) == num {
				return false
			}
		}
	}

	return true
}

// isValidCell reports whether a given position is in bounds of a Sudoku board.
func isValidCell(row, col int) bool {
	return row >= 0 && row < 9 && col >= 0 && col < 9
}

// isValidNumber reports whether a given number is valid on a Sudoku board.
func isValidNumber(num int) bool {
    return num >= 1 && num <= 9
}
