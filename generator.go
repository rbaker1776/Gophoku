package main

import (
    "errors"
    "math/rand"
    "time"
)

const (
    MinValidClueCount = 17
    MaxValidClueCount = 80
    DefaultClueCount  = 32
)

var (
	ErrGenerationFailed = errors.New("failed to generate valid puzzle")
    ErrInvalidClueCount = errors.New("clue count must be between 17 and 80")
    ErrDiggingFailed    = errors.New("failed to remove proper number of clues")
)

// GeneratorOptions configures puzzle generation behavior.
type GeneratorOptions struct {
    ClueCount    int           // Number of clues to add to the puzzle
    Tolerance    int           // Tolerance allows clues within +/- tolerance of ClueCount
	Timeout      time.Duration // Timeout limits generation time
    Seed         int64         // Seed for reproducible puzzles (0 = random)
    EnsureUnique bool          // EnsureUnique verifies single solution
    MaxAttempts  int           // MaxAttempts limits generation retries
}

// DefaultGeneratorOptions returns standard generator options.
func DefaultGeneratorOptions(clueCount int) *GeneratorOptions {
    clueCount = min(max(clueCount, MinValidClueCount), MaxValidClueCount)
	return &GeneratorOptions{
        ClueCount:    clueCount,
        Tolerance:    0,
		Timeout:      30 * time.Second,
		Seed:         0,
		MaxAttempts:  100,
		EnsureUnique: true,
	}
}

// Generator creates Sudoku puzzles.
type Generator struct {
	options *GeneratorOptions
	rng     *rand.Rand
}

// NewGenerator creates a puzzle generator with the given options.
func NewGenerator(options *GeneratorOptions) *Generator {
	if options == nil {
		options = DefaultGeneratorOptions(DefaultClueCount)
	}

	seed := options.Seed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	return &Generator{
		options: options,
		rng:     rand.New(rand.NewSource(seed)),
	}
}

// Generate creates a new Sudoku puzzle.
// Returns the puzzle and its solution, or an error if generation fails.
func (g *Generator) Generate() (puzzle *Board, solution *Board, err error) {
    if g.options.ClueCount < MinValidClueCount || g.options.ClueCount > MaxValidClueCount {
		return nil, nil, ErrInvalidClueCount
	}

	for attempt := 0; attempt < g.options.MaxAttempts; attempt++ {
		// Generate a complete valid board
		solution, err = g.generateSolution()
		if err != nil {
			continue
		}

		// Remove clues to create the puzzle
		puzzle, err = g.removeCells(solution)
		if err != nil {
			continue
		}

		// Verify uniqueness if required
		if g.options.EnsureUnique {
			if !g.hasUniqueSolution(puzzle) {
				continue
			}
		}

		return puzzle, solution, nil
	}

	return nil, nil, ErrGenerationFailed
}

// generateSolution creates a complete valid Sudoku board.
func (g *Generator) generateSolution() (*Board, error) {
	board := NewBoard()

	// Use solver with randomization to generate a complete board
	solver := NewSolver(board, &SolverOptions{
		MaxSolutions: 1,
		Randomize:    true,
		Timeout:      g.options.Timeout,
	})

	return solver.Solve()
}

// removeCells removes clues from a complete board to create a puzzle.
func (g *Generator) removeCells(solution *Board) (*Board, error) {
	puzzle := solution.Clone()

	// Calculate how many cells to remove
	targetClues := g.options.ClueCount
	cellsToRemove := CellCount - targetClues

	// Create shuffled list of all positions
	positions := g.rng.Perm(CellCount)

	// Remove cells until we reach target clues
    cellsRemoved := 0
	for _, pos := range positions {
        if cellsRemoved >= cellsToRemove {
			break
		}

		// Try removing this cell
		val := puzzle.Get(pos)
		if val == EmptyCell {
			continue
		}

		puzzle.Clear(pos)
		cellsRemoved++

		// Verify the puzzle still has a unique solution
		if g.options.EnsureUnique {
			if !g.hasUniqueSolution(puzzle) {
				// Restore the cells
				puzzle.SetForce(pos, val)
				cellsRemoved--
			}
		}
	}

    if cellsRemoved == cellsToRemove {
	    return puzzle, nil
    } else {
        return puzzle, ErrDiggingFailed
    }
}

// hasUniqueSolution checks if the puzzle has exactly one solution.
func (g *Generator) hasUniqueSolution(puzzle *Board) bool {
	solver := NewSolver(puzzle, &SolverOptions{
		MaxSolutions: 2,
		Randomize:    false,
		Timeout:      g.options.Timeout,
	})

	solutions := g.countSolutions(solver)
	return solutions == 1
}

// countSolutions counts the number of solutions for a puzzle.
func (g *Generator) countSolutions(solver *Solver) int {
	count := 0

	// Use backtracking to count solutions
	var backtrack func(*Board) bool
	backtrack = func(b *Board) bool {
		// Apply constraint propagation
		tempSolver := NewSolver(b, &SolverOptions{
			MaxSolutions: 1,
			Randomize:    false,
		})

		if err := tempSolver.propagateConstraints(); err != nil {
			return false
		}

		// Check if solved
		if tempSolver.board.EmptyCount() == 0 {
			count++
			return count < 2 // Stop after finding 2 solutions
		}

		// Find MRV cell
		pos, candidates := tempSolver.findMRVCell()
		if len(candidates) == 0 {
			return false
		}

		for _, val := range candidates {
			if count >= 2 {
				return false
			}

			clone := tempSolver.board.Clone()
			clone.SetForce(pos, val)

			backtrack(clone)
		}

		return count < 2
	}

	backtrack(solver.board.Clone())
	return count
}

// GenerateWithDifficulty is a convenience function to generate a puzzle with a specific difficulty level.
func GenerateWithClueCount(clueCount int) (*Board, *Board, error) {
	generator := NewGenerator(DefaultGeneratorOptions(clueCount))
	return generator.Generate()
}
