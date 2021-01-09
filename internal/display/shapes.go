package display

import (
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"github.com/gdamore/tcell/v2"
)

const (
	tlChar = '┌'
	trChar = '┐'
	brChar = '┘'
	blChar = '└'
	vertChar = '│'
	horiChar = '─'
)

func Box(topLeftX, topLeftY, bottomRightX, bottomRightY int, title string) {
	rawBox(topLeftX, topLeftY, bottomRightX, bottomRightY, title)
	Screen.Show()
}

func rawBox(topLeftX, topLeftY, bottomRightX, bottomRightY int, title string) {
	width := bottomRightX - topLeftX
	height := bottomRightY - topLeftY

	var (
		topLine []rune
		bottomLine []rune

		leftLine []rune
		rightLine []rune
	)

	// generate some lines for the box

	{
		rs := tools.MakeRuneSlice(horiChar, width + 1)

		topLine = make([]rune, len(rs))
		bottomLine = make([]rune, len(rs))
		copy(topLine, rs)
		copy(bottomLine, rs)

		// add corner sections

		topLine[0] = tlChar
		topLine[len(topLine) - 1] = trChar
		for i, v := range title {
			topLine[i + 1] = v
		}

		bottomLine[0] = blChar
		bottomLine[len(bottomLine) - 1] = brChar
	}

	{
		rs := tools.MakeRuneSlice(vertChar, height - 1) // -1 to allow for the top and bottom line

		leftLine = make([]rune, len(rs))
		rightLine = make([]rune, len(rs))
		copy(leftLine, rs)
		copy(rightLine, rs)
	}

	// place these lines

	Screen.SetContent(topLeftX, topLeftY, topLine[0], topLine[1:], tcell.StyleDefault)
	Screen.SetContent(topLeftX, bottomRightY, bottomLine[0], bottomLine[1:], tcell.StyleDefault)

	for i, v := range leftLine {
		Screen.SetContent(topLeftX, topLeftY + 1 + i, v, nil, tcell.StyleDefault)
	}

	for i, v := range rightLine {
		Screen.SetContent(bottomRightX, topLeftY + 1 + i, v, nil, tcell.StyleDefault)
	}
}
