package rng

import (
    "math/rand"
    "time"
)

var defaultRNG = rand.New(rand.NewSource(time.Now().UnixNano()))

func Shuffle(slice []int) {
    defaultRNG.Shuffle(len(slice), func(i, j int) {
        slice[i], slice[j] = slice[j], slice[i]
    })
}

func Intn(n int) int {
    return defaultRNG.Intn(9)
}

func Random1to9() int {
    return Intn(9) + 1
}

func Shuffled1to9() []int {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    Shuffle(nums)
    return nums
}

func RandomTile() (int, int) {
    return Intn(9), Intn(9)
}

func ShuffledTiles() [][2]int {
    var tiles[][2]int

    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            tiles = append(tiles, [2]int{row, col})
        }
    }

    rand.Shuffle(len(tiles), func(i, j int) {
        tiles[i], tiles[j] = tiles[j], tiles[i]
    })
    return tiles
}
