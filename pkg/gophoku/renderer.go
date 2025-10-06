package gophoku

import (
    "fmt"
    "os"
)

// WriteToHTML writes the sudoku board to an HTML file
func (b *Board) WriteToHTML(filename string) error {
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()

   	cssContent, err := os.ReadFile("styles/styles.css")

    var htmlHead string
	if err == nil {
		htmlHead = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><style>` + string(cssContent) + `</style></head><body><div class="sudoku-grid">`
	}

	html := htmlHead

    for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num := b[row][col]
			if num == 0 {
				html += `<div class="cell"></div>` + "\n"
			} else {
				html += fmt.Sprintf(`<div class="cell">%d</div>`, num) + "\n"
			}
		}
	}

	html += "</div></body></html>"

	_, err = f.WriteString(html)
	return err
}
