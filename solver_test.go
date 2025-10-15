package main

import (
    "testing"
)

func BenchmarkSolveEmpty(b *testing.B) {
    for i := 0; i < b.N; i++ {
        board := NewBoard()
        board.Solve()
    }
}
