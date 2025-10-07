package gophoku

// Board represents a BoardSize*BoardSize Sudoku board.
// Each cell contains a value MinValue-MaxValue | EmptyCell.
// Valid Sudoku numbers are 1-9.
type Board [BoardSize][BoardSize]int

// NewBoard creates and returns a new empty Sudoku board.
// NOTE: NewBoard assumed EmptyCell == 0 because the default value is zero.
func NewBoard() *Board {
	return &Board{}
}

// NewBoardFromString creates and returns a new Sudoku board based on an input string.
// A cell will only get filled by a valid value 1-9, other characters will be skipped.
// An empty cell can be noted by '.' or '0', assuming EmptyCell == 0.
func NewBoardFromString(s string) *Board {
    board := NewBoard()
    row, col := 0, 0

    for _, ch := range(s) {
        switch {
        case ch == '.' || int(ch - '0') == EmptyCell:
        case isValidNumber(int(ch - '0')):
            board[row][col] = int(ch - '0')
        default:
            // Opt to skip whitespace and other characters
            continue
        }

        col++
        if col % BoardSize == 0 {
            row++
            col = 0
        }

        if !isValidPosition(row, col) {
            break
        }
    }

    return board
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

// Get returns the value at the specified position.
// Returns InvalidValue if the position is out of bounds.
func (b *Board) Get(row, col int) int {
	if !isValidPosition(row, col) {
		return InvalidValue
	}
	return b[row][col]
}

// Set sets the value at the specified position.
// Returns false if the position is invalid or the value is out of range.
// Does not check Sudoku validity, use CanPlace for validation.
func (b *Board) Set(row, col, value int) bool {
	if !isValidPosition(row, col) {
		return false
	}
	if !isValidNumber(value) {
		return false
	}
	b[row][col] = value
	return true
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

// HintTiles returns a list of board tiles that are not empty.
func (b *Board) HintTiles() [][2]int {
    var hints [][2]int
    for row := 0; row < BoardSize; row++ {
        for col := 0; col < BoardSize; col++ {
            if b[row][col] != EmptyCell {
                hints = append(hints, [2]int{row, col})
            }
        }
    }
    return hints
}

// IsSolved returns true if the board is completely filled and valid according to Sudoku rules.
// A solved board has no empty cells and satisfies all Sudoku constraints.
func (b *Board) IsSolved() bool {
	return b.EmptyCount() == 0 && b.IsValid()
}

// isValidPosition reports whether a given position is in bounds of a Sudoku board.
func isValidPosition(row, col int) bool {
	return row >= 0 && row < BoardSize && col >= 0 && col < BoardSize
}

// isValidNumber reports whether a given number is valid on a Sudoku board.
func isValidNumber(num int) bool {
    return num >= MinValue && num <= MaxValue
}
