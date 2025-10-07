package gophoku

// Board represents a BoardSize*BoardSize Sudoku board.
// Each cell contains a value MinValue-MaxValue | EmptyCell.
// Valid Sudoku numbers are 1-9.
type Board [BoardSize][BoardSize]int

// NewBoard creates and returns a new empty Sudoku board.
func NewBoard() *Board {
	// NOTE: If EmptyCell != 0, then this function will be invalid because the default value is zero
	return &Board{}
}

// Copy creates a deep copy of the board.
// The returned board is independent of the original and can be modified without affecting the source board.
func (b *Board) Copy() *Board {
	newBoard := NewBoard()

	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			newBoard[row][col] = b[row][col]
		}
	}

	return newBoard
}

// EmptyCount returns the number of empty cells on the board.
// A completely empty board will return CellCount, a completely full board will return 0.
func (b *Board) EmptyCount() int {
	count := 0
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if b[row][col] == EmptyCell {
				count++
			}
		}
	}
	return count
}

// HintCount returns the number of filled cells on the board.
// This is the complement of EmptyCount and always satisfies HintCount() + EmptyCount() == CellCount.
func (b *Board) HintCount() int {
	return CellCount - b.EmptyCount()
}

// IsSolved returns true if the board is completely filled and valid according to Sudoku rules.
// A solved board has no empty cells and satisfies all Sudoku constraints.
func (b *Board) IsSolved() bool {
	return b.EmptyCount() == 0 && b.IsValid()
}
