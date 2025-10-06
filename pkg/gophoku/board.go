package gophoku

import (
    "fmt"
    "gophoku/internal/rng"
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

    return newBoard
}

func (b *Board) EmptyCount() int {
    count := 0
    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            if b[row][col] == 0 {
                count++
            }
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
        if b[row][c] == num {
            return false
        }
    }

    for r := 0; r < 9; r++ {
        if b[r][col] == num {
            return false
        }
    }

    startRow := int(row / 3) * 3
    startCol := int(col / 3) * 3
    for r := startRow; r < startRow + 3; r++ {
        for c := startCol; c < startCol + 3; c++ {
            if b[r][c] == num {
                return false
            }
        }
    }

    return true
}

func (b *Board) IsSolved() bool {
    return b.EmptyCount() == 0 && b.IsValid()
}

func (b *Board) MinCandidatesTile() (int, int, []int) {
    bestRow, bestCol := -1, -1
    var bestCandidates []int

    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            if b[row][col] == 0 {
                candidates := b.Candidates(row, col)
                if bestRow == -1 || len(candidates) < len(bestCandidates) {
                    bestRow, bestCol = row, col
                    bestCandidates = candidates
                }
            }
        }
    }

    return bestRow, bestCol, bestCandidates
}

func (b *Board) Candidates(row, col int) []int {
    var candidates []int 
    if b[row][col] != 0 {
        return candidates
    }

    for num := 1; num <= 9; num++ {
        if b.CanPlace(row, col, num) {
            candidates = append(candidates, num)
        }
    }

    return candidates
}

func (b *Board) Solve() bool {
    if b.EmptyCount() == 0 {
        return b.IsSolved()
    } else if b.HintCount() == 0 {
        b.fillDiagonalBoxes()
    }

    row, col, candidates := b.MinCandidatesTile()
    if len(candidates) == 0 {
        return false
    }
    rng.Shuffle(candidates)

    for _, candidate := range(candidates) {
        b[row][col] = candidate
        if b.Solve() {
            return true
        }
        b[row][col] = 0
    }

    return false
}

func (b *Board) fillDiagonalBoxes() {
    for box := 0; box < 3; box++ {
        startRow, startCol := box * 3, box * 3
        i, nums := 0, rng.Shuffled1to9()
        for row := startRow; row < startRow + 3; row++ {
            for col := startCol; col < startCol + 3; col++ {
                b[row][col] = nums[i]
                i++
            }
        }
    }
}

func (b *Board) HasUniqueSolution() bool {
    solutionCount := 0
    b.countSolutions(&solutionCount, 2)
    return solutionCount == 1
}

func (b *Board) countSolutions(count *int, maxCount int) {
    if *count >= maxCount {
        return
    }

    if b.EmptyCount() == 0 {
        if b.IsValid() {
            *count++
        }
    }

    row, col, candidates := b.MinCandidatesTile()
    if len(candidates) == 0 {
        return
    }
    rng.Shuffle(candidates)
    
    for _, candidate := range candidates {
        b[row][col] = candidate
        b.countSolutions(count, maxCount)
        b[row][col] = 0
        
        if *count >= maxCount {
            return
        }
    }
}

func (b *Board) String() string {
    s, l := "", "+-------+-------+-------+\n"
	for i := 0; i < 9; i++ {
		if i % 3 == 0 {
            s += l
        }
		s += "| "
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
                s += ". "
            } else {
                s += fmt.Sprintf("%d ", b[i][j])
            }
			if (j + 1) % 3 == 0 {
                s += "| "
            }
		}
		s += "\n"
	}
	return s + l
}
