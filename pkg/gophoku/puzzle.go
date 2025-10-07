package gophoku

type Puzzle struct {
	Board    *Board
	Solution *Board
	Hints    int
}

func NewPuzzle(board *Board) *Puzzle {
	solution := board.Copy()
	solution.Solve()
	return &Puzzle{
		Board:    board,
		Solution: solution,
		Hints:    board.HintCount(),
	}
}
