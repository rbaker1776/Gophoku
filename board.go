package main

import (
    "strconv"
    "strings"
    "fmt"
)

const (
    // Special cell values
	EmptyCell   = 0
	InvalidCell = -1
    CellCount   = 81

    // Bitmask values
    AllNine = 511
)

// Board represents a 9x9 Sudoku board.
type Board struct {
    // Each cell contains a value 1-9 | EmptyCell.
    cells [CellCount]int

    // Candidate bitmasks store the candidates that have yet to be placed in an area.
    // bit i is set = value i+1 can be placed in the area.
    rowCandidates [9]uint
    colCandidates [9]uint
    boxCandidates [9]uint

    // Track the number of empty cells on the board.
    emptyCount int
}

// NewBoard creates and returns a new empty Sudoku board.
// NOTE: NewBoard assumes EmptyCell == 0
func NewBoard() *Board {
    b := &Board{
        emptyCount: CellCount,
    }
    for i := 0; i < 9; i++ {
        b.rowCandidates[i] = AllNine
        b.colCandidates[i] = AllNine
        b.boxCandidates[i] = AllNine
    }
    return b
}

// Clone creates and returns a deep copy of the board.
// The returned board is independent of the original.
func (b *Board) Clone() *Board {
	newBoard := NewBoard()
    for pos:= 0; pos < CellCount; pos++ {
		newBoard.Set(pos, b.Get(pos))
    }
	return newBoard
}

// Set places the value at the position and reports whether the placement is valid.
// Set checks for Sudoku legality.
// NOTE: Set assumes that val != EmptyCell, use Clear to set a position to EmptyCell
func (b *Board) Set(pos, val int) bool {
    if !isValidPosition(pos) || !isValidNumber(val) {
        return false
    }
    if b.cells[pos] != EmptyCell {
        b.Clear(pos)
    }

    // Check Sudoku legality
    row, col, box := posToUnits(pos)
    mask := uint(1 << (val - 1))
    if b.rowCandidates[row] & b.colCandidates[col] & b.boxCandidates[box] & mask == 0 {
        return false
    }

    // Set the value only once we know it's legal
    b.cells[pos] = val
    b.emptyCount--

    // Update candidates of affected cells
    b.rowCandidates[row] &^= mask
    b.colCandidates[col] &^= mask
    b.boxCandidates[box] &^= mask

    return true
}

// Get retrieves the value at the position.
// Returns InvalidCell if the position is invalid.
func (b *Board) Get(pos int) int {
    if !isValidPosition(pos) {
        return InvalidCell
    }
    return b.cells[pos]
}

// Clear removes the value from the board at position and reports whether the clearing is valid.
// NOTE: Clear assumes the there is a value present at the position.
func (b *Board) Clear(pos int) bool {
    if !isValidPosition(pos) || !isValidNumber(b.cells[pos]) {
        return false
    }

    val := b.cells[pos]
    b.cells[pos] = EmptyCell
    b.emptyCount++

    // Update the candidates of affected cells
    row, col, box := posToUnits(pos)
    mask := uint(1 << (val - 1))
    b.rowCandidates[row] |= mask
    b.colCandidates[col] |= mask
    b.boxCandidates[box] |= mask

    return true
}

// posToUnits decomposes a position into a row [0-9), column [0-9), and box number [0-9).
// posToUnits makes no attempt to validate pos.
func posToUnits(pos int) (int, int, int) {
    row := int(pos / 9)
    col := pos % 9
    box := 3 * int(row / 3) + int(col / 3)
    return row, col, box
}

// String returns a string representation of a Board.
func (b *Board) String() string {
    s := ""
    for pos := 0; pos < CellCount; pos++ {
        if b.cells[pos] == EmptyCell {
            s += "."
        } else {
            s += strconv.Itoa(b.cells[pos])
        }
    }
    return s
}

// PrettyString returns a human-readable representation of the board.
func (b *Board) PrettyString() string {
	var sb strings.Builder
	line := "+-------+-------+-------+\n"
	
	for row := 0; row < 9; row++ {
		if row%3 == 0 {
			sb.WriteString(line)
		}
		sb.WriteString("| ")
		for col := 0; col < 9; col++ {
			pos := row*9 + col
			if b.cells[pos] == 0 {
				sb.WriteString(". ")
			} else {
				fmt.Fprintf(&sb, "%d ", b.cells[pos])
			}
			if (col+1)%3 == 0 {
				sb.WriteString("| ")
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString(line)
	return sb.String()
}
