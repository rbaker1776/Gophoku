package main

import (
    "fmt"
    "gophoku/pkg/gophoku"
)

func main() {
    board := gophoku.NewBoard() 
    board.Solve()
    fmt.Println(board.String())
}
