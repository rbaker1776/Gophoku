package gophoku

const (
	// Board dimensions
	BoardSize = 9
	BoxSize   = 3
	CellCount = BoardSize * BoardSize

	// Cell Values
	EmptyCell    = 0
	MinValue     = 1
	MaxValue     = 9
    InvalidValue = -1

	// Puzzle constraints
	MinValidHints = 17 // Mathematical minimum to produce a unique Sudoku solution
	MaxValidHints = CellCount

	// Generation parameters
	MaxGenerationAttempts = 99
)
