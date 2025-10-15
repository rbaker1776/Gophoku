package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

    puzzle := NewPuzzleWithHints(25)

    fmt.Println(puzzle.Board.PrettyString())
}
