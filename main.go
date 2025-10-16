package main

import (
	"fmt"
)

func main() {
    board, solution, err := GenerateWithClueCount(21)
    if err == nil {
	    fmt.Println(board.Format())
	    fmt.Println(solution.Format())
    } else {
        fmt.Println(err)
    }
}
