package gophoku

func (p *Puzzle) Difficulty() int {
    return p.Board.Copy().traceDifficulty()    
}

func (b *Board) traceDifficulty() int {
    if b.EmptyCount() == 0 {
        // If the board is full, the difficulty is zero
        return 0
    }

    row, col, candidates := b.MinCandidatesTile()
    if len(candidates) == 0 {
        // If no tiles can be placed, the difficulty is zero and nothing more can be done
        return 0
    }

    score := 0
    for _, candidate := range candidates {
        b[row][col] = candidate
        score += 1 + b.traceDifficulty()
        b[row][col] = 0
    }
    return score
}
