package main

// IsValid reports whether a board satisfies Sudoku constraints.
// Empty cells are ignored for validation.
func (b *Board) IsValid() bool {
    var rowConstraints, colConstraints, boxConstraints [9]uint

    for pos := 0; pos < CellCount; pos++ {
        val := b.Get(pos) 
        if val == EmptyCell {
            continue
        }

        // Check Sudoku constraints
        row, col, box := posToUnits(pos)
        mask := uint(1 << (val - 1))
        if (rowConstraints[row] & mask != 0) || (colConstraints[col] & mask != 0) || (boxConstraints[box] & mask != 0) {
            return false
        }

        rowConstraints[row] |= mask
        colConstraints[col] |= mask
        boxConstraints[box] |= mask
    }

    return true
}

// isValidPosition reports whether a given position is in bounds of a Sudoku board.
func isValidPosition(pos int) bool {
	return pos >= 0 && pos < CellCount
}

// isValidNumber reports whether a given number is valid on a Sudoku board.
func isValidNumber(num int) bool {
    return num >= 1 && num <= 9
}
