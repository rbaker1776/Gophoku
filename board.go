package main

import (
	"errors"
	"fmt"
	"strings"
)

// Special cell values
const (
	EmptyCell   = 0
	InvalidCell = -1
	CellCount   = 81
)

// Bitmask values
const (
    AllNine = 511
)

var (
	ErrIllegalMove = errors.New("move violates Sudoku constraints")
)

// Precomputed lookup tables for position mapping
var (
	posToRow [CellCount]int
	posToCol [CellCount]int
	posToBox [CellCount]int
)

// init initializes lookup tables for position-to-unit mappings.
// Should be called once at program start.
func init() {
	for pos := 0; pos < CellCount; pos++ {
		posToRow[pos] = int(pos / 9)
		posToCol[pos] = int(pos % 9)
		posToBox[pos] = 3*int(pos/27) + int((pos%9)/3)
	}
}

// Board represents a 9x9 Sudoku board and its constraint masks.
type Board struct {
	// cells stores values 1-9 or EmptyCell (0) in row-major order.
	// Index calculation: position = row*9 + col
	cells [CellCount]int

	// Bitmasks track placed digits in each unit (row/col/box).
	// Bit i represents digit i+1 (bit 0 = digit 1, bit 8 = digit 9).
	// A set bit means the digit is already placed in that unit.
	rowMasks [9]uint
	colMasks [9]uint
	boxMasks [9]uint

	// emptyCount tracks unfilled cells for quick completion checks.
	emptyCount int
}

// NewBoard creates and returns a new empty Sudoku board.
// NOTE: NewBoard assumes EmptyCell == 0
func NewBoard() *Board {
	b := &Board{
		emptyCount: CellCount,
	}
	return b
}

// NewBoardFromString creates a board from an 81-character string.
// Use '.' or '0' for empty cells, '1'-'9' for filled cells.
// NOTE: NewBoardFromString assumes EmptyCell == 0
func NewBoardFromString(s string) (*Board, error) {
	if len(s) != CellCount {
		return nil, fmt.Errorf("string must be exactly %d characters, got %d", CellCount, len(s))
	}

	b := NewBoard()
	for pos := 0; pos < CellCount; pos++ {
		ch := s[pos]
		switch ch {
		case '.', '0':
			// Empty cell, already initialized
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val := int(ch - '0')
			if err := b.Set(pos, val); err != nil {
				return nil, fmt.Errorf("invalid board at position %d: %w", pos, err)
			}
		default:
			return nil, fmt.Errorf("invalid character '%c' at position %d", ch, pos)
		}
	}
	return b, nil
}

// Clone creates and returns a deep copy of the board.
// The returned board is independent of the original.
func (b *Board) Clone() *Board {
	if b == nil {
		return nil
	}
	clone := *b
	return &clone
}

// Set places a value (1-9) at the given position.
// Returns an error if the move is invalid or violates Sudoku rules.
// Use board.Set(pos, EmptyCell) to clear a cell.
func (b *Board) Set(pos, val int) error {
	if err := b.validatePosition(pos); err != nil {
		return err
	}
	if err := b.validateValue(val); err != nil {
		return err
	}
	if val == EmptyCell {
		return b.Clear(pos)
	}
	if b.cells[pos] != EmptyCell {
		b.Clear(pos)
	}

	row, col, box := posToRow[pos], posToCol[pos], posToBox[pos]
	mask := uint(1 << (val - 1))

	// Check if value already exists in row, column, or box
	if b.rowMasks[row]&mask != 0 {
		return fmt.Errorf("%w: value %d already in row %d", ErrIllegalMove, val, row)
	}
	if b.colMasks[col]&mask != 0 {
		return fmt.Errorf("%w: value %d already in column %d", ErrIllegalMove, val, col)
	}
	if b.boxMasks[box]&mask != 0 {
		return fmt.Errorf("%w: value %d already in box %d", ErrIllegalMove, val, box)
	}

	// Set the value only once we know it's legal
	b.cells[pos] = val
	b.emptyCount--

	// Update candidates of affected cells
	b.rowMasks[row] |= mask
	b.colMasks[col] |= mask
	b.boxMasks[box] |= mask

	return nil
}

// SetForce places a value without validation checks.
// Use only when certain the move is valid.
func (b *Board) SetForce(pos, val int) {
	row, col, box := posToRow[pos], posToCol[pos], posToBox[pos]
	mask := uint(1 << (val - 1))

	b.cells[pos] = val
	b.rowMasks[row] |= mask
	b.colMasks[col] |= mask
	b.boxMasks[box] |= mask
	b.emptyCount--
}

// Get returns the value at the given position.
// Returns EmptyCell for empty cells or InvalidCell for invalid positions.
func (b *Board) Get(pos int) int {
	if !isValidPosition(pos) {
		return InvalidCell
	}
	return b.cells[pos]
}

// GetCell returns the value at the given row and column.
// Returns EmptyCell for empty cells or InvalidCell for invalid positions.
func (b *Board) GetCell(row, col int) int {
	if row < 0 || row >= 9 || col < 0 || col >= 9 {
		return InvalidCell
	}
	return b.cells[row*9+col]
}

// GetCandidatesMask returns the bitmask of candidates for a given position.
// A returned 0 indicates an unsolvable board or an invalid position.
func (b *Board) GetCandidatesMask(pos int) uint {
    if !isValidPosition(pos) {
        return 0
    }
	row, col, box := posToRow[pos], posToCol[pos], posToBox[pos]
    return AllNine &^ b.rowMasks[row] &^ b.colMasks[col] &^ b.boxMasks[box]
}

// GetCandidates returns a slice of candidates for a given position.
// An empty slice indicates an unsolvable board or an invalid position.
func (b *Board) GetCandidates(pos int) []int {
    mask := b.GetCandidatesMask(pos)
    candidates := make([]int, 0, 9)
    for num := 1; num <= 9; num++ {
        if mask & uint(1 << (num - 1)) != 0 {
            candidates = append(candidates, num)
        }
    }
    return candidates
}

// GetCandidatesMaskCell returns the bitmask of candidates at the given row and column.
// A returned 0 indicates an unsolvable board or an invalid position.
func (b *Board) GetCandidatesMaskCell(row, col int) uint {
	if row < 0 || row >= 9 || col < 0 || col >= 9 {
		return 0
	}
    return b.GetCandidatesMask(row*9+col)
}

// GetCandidatesCell returns a slice of candidates for the given row and column.
// An empty slice indicates an unsolvable board or an invalid position.
func (b *Board) GetCandidatesCell(row, col int) []int {
	if row < 0 || row >= 9 || col < 0 || col >= 9 {
		return []int{}
	}
    return b.GetCandidates(row*9+col)
}

// Clear removes the value at the given position.
// Returns an error if the position is invalid.
// No harm is done calling Clear on an already empty cell.
func (b *Board) Clear(pos int) error {
	if err := b.validatePosition(pos); err != nil {
		return err
	}

	// Check if the cell is already empty
	val := b.cells[pos]
	if val == EmptyCell {
		return nil
	}

	row, col, box := posToRow[pos], posToCol[pos], posToBox[pos]
	mask := uint(1 << (val - 1))

	// Clear the cell
	b.cells[pos] = EmptyCell
	b.emptyCount++

	// Update the candidates of affected cells
	b.rowMasks[row] &^= mask
	b.colMasks[col] &^= mask
	b.boxMasks[box] &^= mask

	return nil
}

// EmptyCount returns the number of empty cells on the board.
func (b *Board) EmptyCount() int {
    return b.emptyCount
}

// String returns the board as an 81-character string.
// Empty cells are represented as '.', filled cells as '1'-'9'.
func (b *Board) String() string {
	var sb strings.Builder
	sb.Grow(CellCount)

	for _, cell := range b.cells {
		if cell == EmptyCell {
			sb.WriteByte('.')
		} else {
			sb.WriteByte('0' + byte(cell))
		}
	}
	return sb.String()
}

// Format returns a human-readable board representation with grid lines.
func (b *Board) Format() string {
	var sb strings.Builder
	line := "+-------+-------+-------+\n"
	sb.WriteString(line)

	for row := 0; row < 9; row++ {
		sb.WriteString("| ")
		for col := 0; col < 9; col++ {
			val := b.GetCell(row, col)
			if val == EmptyCell {
				sb.WriteByte('.')
			} else {
				sb.WriteByte('0' + byte(val))
			}
			sb.WriteByte(' ')

			if (col+1)%3 == 0 {
				sb.WriteString("| ")
			}
		}
		sb.WriteString("\n")

		if (row+1)%3 == 0 {
			sb.WriteString(line)
		}
	}

	return sb.String()
}
