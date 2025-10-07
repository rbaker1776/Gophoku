package gophoku

// IsValid checks if the current board state satisfies all Sudoku constraints.
// Empty cells are ignored in validation.
func (b *Board) IsValid() bool {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if !(b[row][col] >= MinValue && b[row][col] <= MaxValue) {
				return false
			}
			if b[row][col] != EmptyCell {
				num := b[row][col]
				if !b.CanPlace(row, col, num) {
					return false
				}
			}
		}
	}
	return true
}

// CanPlace determines if a number can be legally placed at the specified position.
// Returns true if the placement is legal, false otherwise.
func (b *Board) CanPlace(row, col, num int) bool {
	// Check input bounds
	if row < 0 || row >= BoardSize || col < 0 || col >= BoardSize {
		return false
	}

	// Check row for duplicate
	for c := 0; c < BoardSize; c++ {
		if c != col && b[row][c] == num {
			return false
		}
	}

	// Check column for duplicate
	for r := 0; r < BoardSize; r++ {
		if r != row && b[r][col] == num {
			return false
		}
	}

	// Check 3x3 box for duplicate
	startRow := int(row/BoxSize) * BoxSize
	startCol := int(col/BoxSize) * BoxSize
	for r := startRow; r < startRow+BoxSize; r++ {
		for c := startCol; c < startCol+BoxSize; c++ {
			if (r != row || c != col) && b[r][c] == num {
				return false
			}
		}
	}

	return true
}
