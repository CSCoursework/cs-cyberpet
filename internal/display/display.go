package display

import (
	"fmt"
	"github.com/gdamore/tcell"
	"os"
)

const BottomLineHeight = 3

var Screen tcell.Screen

func init() {
	var err error
	Screen, err = tcell.NewScreen()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create new tcell Screen: %s", err.Error())
		os.Exit(1)
	}
	err = Screen.Init()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to initialise tcell Screen: %s", err.Error())
		os.Exit(1)
	}
}

func PrintString(in string, posX, posY int) {
	if len(in) == 0 {
		return
	}
	Screen.SetContent(posX, posY, rune(in[0]), []rune(in[1:]), 0)
	Screen.Show()
}

func PrintMultiLine(in []string, posX, posY int) {
	for i, x := range in {
		PrintString(x, posX, posY + i)
	}
}

func PrintLine(fixedPos int, char rune, isVertical bool) {
	var x, y, totalLen int
	xs, ys := Screen.Size()
	if isVertical {
		totalLen = ys
		x = fixedPos
	} else {
		totalLen = xs
		y = fixedPos
	}
	modPos := func(i int) int {

		if isVertical {
			y += 1
		} else {
			x += 1
		}

		return i + 1
	}
	for i := 0; i < totalLen; i = modPos(i) {
		Screen.SetContent(x, y, char, nil, 0)
	}
	Screen.Show()
}