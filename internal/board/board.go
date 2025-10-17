package board

import (
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
	allNine = 511
)

// Board represents a 9x9 Sudoku board.
type Board struct {
	cells [CellCount]int

	// Bitmasks track placed digits in each unit (row/col/box).
	// Bit i represents digit i+1 (bit 0 = digit 1, bit 8 = digit 9).
	// This allows for O(1) validation.
	rowMasks [9]uint
	colMasks [9]uint
	boxMasks [9]uint

	// emptyCount tracks unfilled cells for quick completion checks.
	// Once initialized, emptyCount should only be touched inside Set and Clear.
	emptyCount int
}

// New creates an empty Board.
func New() *Board {
	b := &Board{
		emptyCount: CellCount,
	}
	return b
}

// NewFromString creates a Board from an 81-character string.
// Use '.' or '0' for empty cells, '1'-'9' for filled cells.
func NewFromString(s string) (*Board, error) {
	if len(s) != CellCount {
		return nil, fmt.Errorf("string must be exactly %d characters, got %d", CellCount, len(s))
	}

	b := New()
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

// Clone creates an independent copy of the Board.
func (b *Board) Clone() *Board {
	if b == nil {
		return nil
	}
	clone := *b
	return &clone
}

// Set attempts to place a value 1-9 at the given position.
// Returns an error if the placement violates Sudoku rules or parameters are invalid.
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

	// Check if value already exists in row, column, or box for Sudoku rules
	if b.rowMasks[row]&mask != 0 {
		return fmt.Errorf("%w: value %d already in row %d", ErrIllegalMove, val, row)
	}
	if b.colMasks[col]&mask != 0 {
		return fmt.Errorf("%w: value %d already in column %d", ErrIllegalMove, val, col)
	}
	if b.boxMasks[box]&mask != 0 {
		return fmt.Errorf("%w: value %d already in box %d", ErrIllegalMove, val, box)
	}

	// Modify the board only once we know it's legal to do so
	b.cells[pos] = val
	b.rowMasks[row] |= mask
	b.colMasks[col] |= mask
	b.boxMasks[box] |= mask
	b.emptyCount--

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

// Clear removes the value at the given position.
// Returns an error if the position is invalid.
// No harm is done calling Clear on an already empty cell.
func (b *Board) Clear(pos int) error {
	if err := b.validatePosition(pos); err != nil {
		return err
	}

	// Exit early if the cell is already empty, no harm no foul
	val := b.cells[pos]
	if val == EmptyCell {
		return nil
	}

	row, col, box := posToRow[pos], posToCol[pos], posToBox[pos]
	mask := uint(1 << (val - 1))

	b.cells[pos] = EmptyCell
	b.rowMasks[row] &^= mask
	b.colMasks[col] &^= mask
	b.boxMasks[box] &^= mask
	b.emptyCount++

	return nil
}

// Get returns the value at the given position.
// Returns InvalidCell for invalid positions.
func (b *Board) Get(pos int) int {
	if !isValidPosition(pos) {
		return InvalidCell
	}
	return b.cells[pos]
}

// GetCandidatesMask returns the bitmask of candidates for a given position.
// A returned 0 indicates an unsolvable board or an invalid position.
func (b *Board) GetCandidatesMask(pos int) uint {
	if !isValidPosition(pos) {
		return 0
	}
	row, col, box := posToRow[pos], posToCol[pos], posToBox[pos]
	return allNine &^ b.rowMasks[row] &^ b.colMasks[col] &^ b.boxMasks[box]
}

// GetCandidates returns a slice of candidates 1-9 for a given position.
// An empty slice indicates an unsolvable board or an invalid position.
func (b *Board) GetCandidates(pos int) []int {
	mask := b.GetCandidatesMask(pos)
	candidates := make([]int, 0, 9)
	for num := 1; num <= 9; num++ {
		if mask&uint(1<<(num-1)) != 0 {
			candidates = append(candidates, num)
		}
	}
	return candidates
}

// EmptyCount returns the number of empty cells on the board.
func (b *Board) EmptyCount() int {
	return b.emptyCount
}

// ClueCount returns the number of filled cells on the board.
func (b *Board) ClueCount() int {
	return CellCount - b.emptyCount
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
			val := b.Get(MakePos(row, col))
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

// Precomputed lookup tables for position mapping
var (
	posToRow [CellCount]int
	posToCol [CellCount]int
	posToBox [CellCount]int
)

// MakePos transforms a row and column into a linear position.
// Returns InvalidCell if row and/or col are invalid.
func MakePos(row, col int) int {
	if row < 0 || row >= 9 || col < 0 || col >= 9 {
		return InvalidCell
	}
	return 9*row + col
}

// init initializes lookup tables for position-to-unit mappings.
// Should be called once at program start.
func init() {
	for pos := 0; pos < CellCount; pos++ {
		posToRow[pos] = int(pos / 9)
		posToCol[pos] = int(pos % 9)
		posToBox[pos] = 3*int(pos/27) + int((pos%9)/3)
	}
}
