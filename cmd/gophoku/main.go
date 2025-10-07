package main

import (
    "fmt"
    "gophoku/pkg/gophoku"
    "gophoku/internal/rng"
)

func main() {
    hardest := 0
    for true {
        board := gophoku.NewBoard() 
        generator := gophoku.NewGenerator(board)
        puzzle := generator.Generate(rng.Intn(10) + 17)
        difficulty := puzzle.Difficulty()
        if difficulty > hardest {
            hardest = difficulty
            fmt.Println("\n\n==============================================\n\n")
            fmt.Println(puzzle.Board.String())
            fmt.Println(puzzle.Difficulty())
        }
    }
}
