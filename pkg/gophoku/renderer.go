package gophoku

import (
	"fmt"
	"os"
)

// String returns a human-readable string representation of the board.
// Example output:
//
//	+-------+-------+-------+
//	| . . . | . 4 7 | 5 . . |
//	| . . 3 | . . . | . . 4 |
//	| 1 . . | . . . | . . . |
//	+-------+-------+-------+
//	| . . . | . 9 . | 3 1 . |
//	| 5 . . | 3 6 . | . . . |
//	| . 9 1 | . 5 . | . . 6 |
//	+-------+-------+-------+
//	| . . . | . 7 . | 8 . . |
//	| 6 . . | 1 . . | . . 2 |
//	| . . . | . . 8 | . 4 . |
//	+-------+-------+-------+
func (b *Board) String() string {
	s, l := "", "+-------+-------+-------+\n"
	for row := 0; row < BoardSize; row++ {
		if row % BoxSize == 0 {
			s += l
		}
		s += "| "
		for col := 0; col < BoardSize; col++ {
			if b[row][col] == EmptyCell {
				s += ". "
			} else {
				s += fmt.Sprintf("%d ", b[row][col])
			}
			if (col+1)%BoxSize == 0 {
				s += "| "
			}
		}
		s += "\n"
	}
	return s + l
}

// Compressed string returns a compressed string representaiton of the board.
// The returned string can be passed into NewBoardFromString() to reproduce the board.
func (b *Board) CompressedString() string {
    s := ""
    for row := 0; row < BoardSize; row++ {
        for col := 0; col < BoardSize; col++ {
            if b[row][col] == EmptyCell {
                s += "."
            } else {
                s += string(b[row][col] + '0')
            }
        }
    }
    return s
}

// WriteToHTML writes the sudoku board to an HTML file
func WriteToHTML(filename string, puzzles []Puzzle) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	cssContent, err := os.ReadFile("styles/styles.css")

	var htmlHead string
	if err == nil {
		// Embed the CSS
		htmlHead = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Sudoku Puzzles</title>
    <style>
` + string(cssContent) + `
    </style>
</head>
<body>
`
	} else {
		// Link to external CSS
		htmlHead = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Sudoku Puzzles</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
`
	}

	html := htmlHead

	for i, puzzle := range puzzles {
		if i > 0 {
			html += `    <div class="page-break"></div>` + "\n"
		}
		html += fmt.Sprintf(`    <h1>Sudoku Puzzle %d</h1>`, i+1) + "\n"
		html += `    <div class="sudoku-grid">` + "\n"

		for row := 0; row < BoardSize; row++ {
			for col := 0; col < BoardSize; col++ {
				num := puzzle.Board[row][col]
				if num == EmptyCell {
					html += `<div class="cell"></div>` + "\n"
				} else {
					html += fmt.Sprintf(`<div class="cell">%d</div>`, num) + "\n"
				}
			}
		}

		html += `    </div>` + "\n"
	}

	html += "</body></html>"

	_, err = f.WriteString(html)
	return err
}
