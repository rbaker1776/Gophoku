package main

import (
    "fmt"
    "testing"
)

func BenchmarkNewPuzzleWithHints25(b *testing.B) {
    for i := 0; i < b.N; i++ {
        puzzle := NewPuzzleWithHints(25)
        if !puzzle.Solution.IsValid() {
            fmt.Errorf("solution is invalid")
        }
    }
}
