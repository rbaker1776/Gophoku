package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

    board := NewBoard()
    board.Set(1, 2)
    fmt.Println(board.String())
}
