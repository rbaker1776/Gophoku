package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

    board := NewBoard()
    board.Solve()
    fmt.Println(board.PrettyString())
}
