package main

import (
	"fmt"
)

func main() {
	board := NewBoard()
    solver := NewSolver(board, GenerateOptions())
    solved, err := solver.Solve()
    if err == nil {
	    fmt.Println(solved.Format())
    } else {
        fmt.Println(err)
    }
}
