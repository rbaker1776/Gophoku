package main

import (
    "fmt"
    "gophoku/pkg/gophoku"
)

func main() {
    board := gophoku.NewBoard() 
    generator := gophoku.NewGenerator(board)
    puzzle, err := generator.Generate(17)
    if err != nil {
        fmt.Println(puzzle.Board.String())
    } else {
        fmt.Println(puzzle.Board.String())
    }
}
