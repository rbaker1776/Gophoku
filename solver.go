package main

import (
    "math/rand"
    "math/bits"
)

// Solve attempts to solve a Sudoku board and reports whether ~a~ solution was found.
// If multiple solutions are possible, a random one will be chosen.
// NOTE: Solve modifies the board directly. If this is undesirable, use Solution.
func (b *Board) Solve() bool {
    if b.emptyCount == 0 {
        return b.IsValid()
    }

    // If starting with an empty board, fill diagonal boxes for efficiency
    if b.emptyCount == CellCount {
        b.fillDiagonalBoxes()
    }

    // Apply constraint propagation before backtracking
    if !b.propagateConstraints() {
        return false
    }

    // If propagation solved the puzzle, we're done
    if b.emptyCount == 0 {
        return b.IsValid()
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

// propagateConstraints applies constraint propagation techniques.
// Returns false if the board is unsolvable.
func (b *Board) propagateConstraints() bool {
    changed := true
    for changed {
        changed = false
        
        // Naked singles: cells with only one candidate
        for pos := 0; pos < CellCount; pos++ {
            if b.cells[pos] == EmptyCell {
                row, col, box := posToUnits(pos)
                candidateBits := b.rowCandidates[row] & b.colCandidates[col] & b.boxCandidates[box]
                
                if candidateBits == 0 {
                    return false // No valid candidates, unsolvable
                }
                
                // If only one bit is set, we have a naked single
                if bits.OnesCount(candidateBits) == 1 {
                    value := bits.TrailingZeros(candidateBits) + 1
                    b.Set(pos, value)
                    changed = true
                }
            }
        }
        
        // Hidden singles: values that can only go in one place in a unit
        if !changed {
            changed = b.findHiddenSingles()
        }
    }
    return true
}

// findHiddenSingles finds values that can only be placed in one position within a unit
// Reports whether any cells were solved.
func (b *Board) findHiddenSingles() bool {
    changed := false
    
    // Check rows
    for row := 0; row < 9; row++ {
        for val := 1; val <= 9; val++ {
            mask := uint(1 << (val - 1))
            if b.rowCandidates[row]&mask == 0 {
                continue // Value already placed
            }
            
            count := 0
            lastPos := -1
            for col := 0; col < 9; col++ {
                pos := row*9 + col
                if b.cells[pos] == EmptyCell {
                    _, _, box := posToUnits(pos)
                    if b.colCandidates[col]&mask != 0 && b.boxCandidates[box]&mask != 0 {
                        count++
                        lastPos = pos
                    }
                }
            }
            if count == 1 && lastPos != -1 {
                b.Set(lastPos, val)
                changed = true
            }
        }
    }
    
    // Check columns
    for col := 0; col < 9; col++ {
        for val := 1; val <= 9; val++ {
            mask := uint(1 << (val - 1))
            if b.colCandidates[col]&mask == 0 {
                continue
            }
            
            count := 0
            lastPos := -1
            for row := 0; row < 9; row++ {
                pos := row*9 + col
                if b.cells[pos] == EmptyCell {
                    _, _, box := posToUnits(pos)
                    if b.rowCandidates[row]&mask != 0 && b.boxCandidates[box]&mask != 0 {
                        count++
                        lastPos = pos
                    }
                }
            }
            if count == 1 && lastPos != -1 {
                b.Set(lastPos, val)
                changed = true
            }
        }
    }
    
    return changed
}

// Solution attempts to solve a Sudoku board and returns the solution as a new board.
func (b *Board) Solution() *Board {
    newBoard := b.Clone()
    newBoard.Solve()
    return newBoard
}

// HasUniqueSolution checks if the current board has exactly one solution.
// This is useful for puzzle generation to ensure a valid, solvable puzzle.
func (b *Board) HasUniqueSolution() bool {
	solutionCount := 0
	b.Clone().countSolutions(&solutionCount, 2)
	return solutionCount == 1
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

// countSolutions counts the number of solutions up to maxCount.
// This is used internally by HasUniqueSolution to check uniqueness without finding all possible solutions.
func (b *Board) countSolutions(count *int, maxCount int) {
	// Early exit if we've already found enough solutions
	if *count >= maxCount {
		return
	}

	// If the board is complete, no more solutions can be found
	if b.emptyCount == 0 {
		if b.IsValid() {
			*count++
		}
		return
	}

	pos, candidates := b.minCandidatesCell()
	if len(candidates) == 0 {
		return
	}

	// Try each candidate using backtracking
	for _, candidate := range candidates {
        b.Set(pos, candidate)
		b.countSolutions(count, maxCount)
		b.Clear(pos)

		// Early exit if we've found enough solutions
		if *count >= maxCount {
			return
		}
	}
}
