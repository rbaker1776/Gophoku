package gophoku

import (
	"gophoku/internal/rng"
)

// Solve attmpts to solve the Sudoku board using backtracking.
// Returns true if ~a~ valid solution is found, false otherwise.
// If multiple solutions are possible, Solve will choose a random one.
func (b *Board) Solve() bool {
	if b.EmptyCount() == 0 {
		return b.IsSolved()
	}

	// If starting with an empty board, fill diagonal boxes for efficiency
	if b.HintCount() == 0 {
		b.fillDiagonalBoxes()
	}

	row, col, candidates := b.MinCandidatesTile()
	if len(candidates) == 0 {
		return false
	}
	// Shuffle candidates for randomness in solution generation
	// We want randomness so that we can use Solve to generate random puzzles
	rng.Shuffle(candidates)

	// Try each candidate using backtracking
	for _, candidate := range candidates {
		b[row][col] = candidate
		if b.Solve() {
			return true
		} else {
			b[row][col] = EmptyCell
		}
	}

	return false
}

// MinCandidatesTile finds the empty tile with the fewest valid candidates.
// This implements the "minimum remaining values" heuristic for efficient solving.
// Returns the tile position and the list of valid candidates.
// An empty candidates list indicates an unsolvable board.
func (b *Board) MinCandidatesTile() (int, int, []int) {
	bestRow, bestCol := -1, -1
	var bestCandidates []int

	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if b[row][col] == EmptyCell {
				candidates := b.Candidates(row, col)

				// Choose the tile with minimum candidates
				if bestRow == -1 || len(candidates) < len(bestCandidates) {
					bestRow, bestCol = row, col
					bestCandidates = candidates
				}

				// Early exit if we find a tile with no valid candidates
				if len(candidates) == 0 {
					return bestRow, bestCol, bestCandidates
				}
			}
		}
	}

	return bestRow, bestCol, bestCandidates
}

// Candidates finds the numbers that can be placed in a tile without creating an immediate violation.
// Returns a list of numbers 1-9 that can be placed in the provided position.
// If the position is already occupied, Candidates returns an empty list.
func (b *Board) Candidates(row, col int) []int {
	var candidates []int

	// Early exit if the square is already occupied
	if b[row][col] != EmptyCell {
		return candidates
	}

	// Test each candidate
	for num := MinValue; num <= MaxValue; num++ {
		if b.CanPlace(row, col, num) {
			candidates = append(candidates, num)
		}
	}

	return candidates
}

// HasUniqueSolution checks if the current board has exactly one solution.
// This is useful for puzzle generation to ensure a valid, solvable puzzle.
func (b *Board) HasUniqueSolution() bool {
	solutionCount := 0
	b.countSolutions(&solutionCount, 2)
	return solutionCount == 1
}

// countSolutions counts the number of solutions up to maxCount.
// This is used internally by HasUniqueSolution to efficiently check uniqueness without finding all possible solutions.
func (b *Board) countSolutions(count *int, maxCount int) {
	// Early exit if we've already found enough solutions
	if *count >= maxCount {
		return
	}

	// If the board is complete, no more solutions can be found
	if b.EmptyCount() == 0 {
		if b.IsValid() {
			*count++
		}
		return
	}

	row, col, candidates := b.MinCandidatesTile()
	if len(candidates) == 0 {
		return
	}

	// Try each candidate using backtracking
	for _, candidate := range candidates {
		b[row][col] = candidate
		b.countSolutions(count, maxCount)
		b[row][col] = EmptyCell

		// Early exit if we've found enough solutions
		if *count >= maxCount {
			return
		}
	}
}

// FillDiagonalBoxes fills three 3x3 boxes on a sudoku board (27 squares total) that are independent.
// This is used by Solve for efficiency purposes when the Sudoku board is empty.
func (b *Board) fillDiagonalBoxes() {
	for box := 0; box < BoxSize; box++ {
		startRow, startCol := box*BoxSize, box*BoxSize
		i, nums := 0, rng.Shuffled1to9()
		for row := startRow; row < startRow+BoxSize; row++ {
			for col := startCol; col < startCol+BoxSize; col++ {
				b[row][col] = nums[i]
				i++
			}
		}
	}
}
