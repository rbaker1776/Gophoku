package gophoku

import (
)

type Board [9][9]int

type Tile struct {
    Row, Col int
}

func NewBoard() *Board {
    return &Board{}
}

func (b *Board) Copy() *Board {
    newBoard := NewBoard()

    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            newBoard[row][col] = b[row][col]
        }
    }

    return &newBoard
}

func (b *Board) EmptyCount() int {
    count := 0
    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            count += int(b[row][col] == 0)
        }
    }
    return count
}

func (b *Board) HintCount() int {
    return 9 * 9 - b.EmptyCount()
}

func (b *Board) IsValid() bool {
    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            if b[row][col] != 0 {
                num := b[row][col]
                b[row][col] = 0
                if !b.CanPlace(row, col, num) {
                    return false
                }
                b[row][col] = num
            }
        }
    }
    return true
}

func (b *Board) CanPlace(row, col, num int) bool {
    if b[row][col] != 0 {
        return false
    }

    for c := 0; c < 9; c++ {
        if board[row][c] == num {
            return false
        }
    }

    for r := 0; r < 9; r++ {
        if board[r][col] == num {
            return false
        }
    }

    startRow := int(row / 3) * 3
    startCol := int(col / 3) * 3
    for r := startRow; r < startRow + 3; r++ {
        for c := startCol; c < startCol + 3; c++ {
            if board[r][c] == num {
                return false
            }
        }
    }

    return true
}

func (b *Board) IsSolved() bool {
    return b.EmptyCount == 0 && b.IsValid()
}

func (b *Board) MinCandidatesTile() (int, int, []int) {
    bestRow, bestCol := -1, -1
    var candidates []int = {1, 2, 3, 4, 5, 6, 7, 8, 9}

    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            if b[row][col] == 0 {
                if len(b.Candidates(row, col)) < len(candidates) {
                    candidates = b.Candidates(row, col)
                }
                bestRow, bestCol = row, col
            }
        }
    }

    return bestRow, bestCol, candidates
}

func (b *Board) Candidates(row, col) []int {
    var candidates []int 
    if b[row][col] != 0 {
        return candidtates
    }

    for num := 0; num < 9; num++ {
        if b.CanPlace(row, col, num) {
            candidates = append(candidates, num)
        }
    }

    return candidates
}

func (b *Board) Solve() {
    if b.EmptyCount() == 0 {
        return
    }

    row, col, candidates := b.MinCandidatesTile()
    if len(candidates) == 0 {
        return
    }

    for _, candidate := range(candidates) {
        b[row][col] = candidate
        b.Solve()
        if b.IsSolved() {
            return
        }
    }
}
