package gophoku

import (
    "fmt"
)

type Board [9][9]int

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

func (b *Board) IsSolved() bool {
    return b.EmptyCount() == 0 && b.IsValid()
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
