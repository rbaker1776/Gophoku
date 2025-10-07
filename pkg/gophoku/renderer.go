package gophoku

import (
    "fmt"
    "os"
)

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

        for row := 0; row < 9; row++ {
            for col := 0; col < 9; col++ {
                num := puzzle.Board[row][col]
                if num == 0 {
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
