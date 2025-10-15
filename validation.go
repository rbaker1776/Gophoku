package main

// isValidPosition reports whether a given position is in bounds of a Sudoku board.
func isValidPosition(pos int) bool {
	return pos >= 0 && pos < CellCount
}

// isValidNumber reports whether a given number is valid on a Sudoku board.
func isValidNumber(num int) bool {
    return num >= 1 && num <= 9
}
