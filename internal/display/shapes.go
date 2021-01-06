package display

import (
	"github.com/codemicro/cs-cyberpet/internal/tools"
)

const (
	tlChar = '┌'
	trChar = '┐'
	brChar = '┘'
	blChar = '└'
	vertChar = '│'
	horiChar = '─'
)

func Box(topLeftX, topLeftY, bottomRightX, bottomRightY int) {
	rawBox(topLeftX, topLeftY, bottomRightX, bottomRightY)
	Screen.Show()
}

func rawBox(topLeftX, topLeftY, bottomRightX, bottomRightY int) {
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

	Screen.SetContent(topLeftX, topLeftY, topLine[0], topLine[1:], 0)
	Screen.SetContent(topLeftX, bottomRightY, bottomLine[0], bottomLine[1:], 0)

	for i, v := range leftLine {
		Screen.SetContent(topLeftX, topLeftY + 1 + i, v, nil, 0)
	}

	for i, v := range rightLine {
		Screen.SetContent(bottomRightX, topLeftY + 1 + i, v, nil, 0)
	}
}
