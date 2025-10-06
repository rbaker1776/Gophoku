package main

import (
    "fmt"
    "gophoku/pkg/gophoku"
)

func main() {
    board := gophoku.NewBoard() 
    generator := gophoku.NewGenerator(board)
    puzzle := generator.Generate(30)
    fmt.Println(puzzle.Board.String())
    fmt.Println(puzzle.Difficulty())
}
