package main

import (
    "fmt"
    "gophoku/pkg/gophoku"
)

func main() {
    puzzle := gophoku.NewPuzzleWithHints(25)
    fmt.Println(puzzle.Board.String())
    fmt.Println(puzzle.Difficulty())
}
