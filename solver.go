package main

import (
    "math/rand"
)

func (b *Board) Solve() bool {
    if b.emptyCount == 0 {
        return b.IsValid()
    }

    // If starting with an empty board, fill diagonal boxes for efficiency
	if b.emptyCount == CellCount {
		b.fillDiagonalBoxes()
	}

    pos, candidates := b.minCandidatesCell()
	if len(candidates) == 0 {
		return false
	}
	// Shuffle candidates for randomness in solution generation
	// We want randomness so that we can use Solve to generate random puzzles
	rand.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})

	// Try each candidate using backtracking
	for _, candidate := range candidates {
	    b.Set(pos, candidate)	
		if b.Solve() {
			return true
		} else {
			b.Clear(pos)
		}
	}

	return false
}

// minCandidatesCell finds the empty cell with the fewest valid candidates.
// Returns the cell position and the list of valid candidates.
// An empty candidates list indicates an unsolvable board.
func (b *Board) minCandidatesCell() (int, []int) {
    bestPos := CellCount
    var bestCandidates []int

    for pos := 0; pos < CellCount; pos++ {
        if b.cells[pos] == EmptyCell {
            candidates := b.candidates(pos)

            // Choose the tile with the fewest candidates
            if !isValidPosition(bestPos) || len(candidates) < len(bestCandidates) {
                bestPos = pos
                bestCandidates = candidates
            }

            // Early exit if we find a cell with zero candidates
            if len(candidates) == 0 {
                break
            }
        }
    }

    return bestPos, bestCandidates
}

// candidates finds the numbers that can be placed in a cell without creating an immediate violation.
// Returns a list of numbers [1-9] that can be placed in the provided position.
// If the position is already occupied, Candidates returns an empty list.
func (b *Board) candidates(pos int) []int {
    row, col, box := posToUnits(pos)
    candidateBits := b.rowCandidates[row] & b.colCandidates[col] & b.boxCandidates[box]
    var candidates []int

    for i := 0; i < 9; i++ {
        if candidateBits & (1 << i) != 0 {
            candidates = append(candidates, i + 1)
        }
    }

	return candidates
}

// fillDiagonalBoxes fills three 3x3 boxes on a sudoku board (27 squares total) that are independent.
// This is used by Solve for efficiency when the Sudoku board is empty.
func (b *Board) fillDiagonalBoxes() {
	for box := 0; box < 3; box++ {
		nums := rand.Perm(9)
		idx := 0
		
		startRow := box * 3
		startCol := box * 3
		
		for row := startRow; row < startRow+3; row++ {
			for col := startCol; col < startCol+3; col++ {
				pos := row*9 + col
				b.Set(pos, nums[idx]+1)
				idx++
			}
		}
	}
}
