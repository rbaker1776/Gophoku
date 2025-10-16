package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	board := NewBoard()
	fmt.Println(board.Format())
}
