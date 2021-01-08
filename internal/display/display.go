package display

import (
	"fmt"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"github.com/gdamore/tcell"
	"os"
	"strings"
	"time"
)

const BottomLineHeight = 4

var (
	Screen tcell.Screen

	StatusLineNumber int
	InputLineNumber int
	OptionsLineNumber int
)

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

	_, dimY := Screen.Size()
	StatusLineNumber = dimY - BottomLineHeight + 1
	OptionsLineNumber = dimY - BottomLineHeight + 2
	InputLineNumber = dimY - BottomLineHeight + 3

}

func PrintString(in string, posX, posY int) {
	rawPrintString(in, posX, posY)
	Screen.Show()
}

func rawPrintString(in string, posX, posY int) {
	if len(in) == 0 {
		return
	}
	Screen.SetContent(posX, posY, rune(in[0]), []rune(in[1:]), 0)
}

func PrintMultiString(in []string, posX, posY int) {
	for i, x := range in {
		rawPrintString(x, posX, posY + i)
	}
	Screen.Show()
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
	Screen.SetContent(x, y, char, tools.MakeRuneSlice(char, totalLen - 1), 0)
	Screen.Show()
}

func CharacterSay(in string, characterX, characterY, longestPart, yOffset int) (clearFunc func()) {
	splitLines := strings.Split(in, "\n")
	for i, x := range splitLines {
		starter := "  "
		if i == 0 {
			starter = "< "
		}
		for ii, char := range starter + x {
			Screen.ShowCursor(characterX + longestPart + ii + 1, characterY + yOffset + i)
			rawPrintString(string(char), characterX + longestPart + ii, characterY + yOffset + i)
			Screen.Show()
			time.Sleep(time.Millisecond * 50)
		}
	}
	Screen.HideCursor()
	return func() {
		blankString := string(tools.MakeRuneSlice(' ', tools.FindLongestStringLen(splitLines) + 2))
		for i := 0; i < len(splitLines); i += 1 {
			rawPrintString(blankString, characterX+ longestPart, characterY+ yOffset + i)
		}
		Screen.Show()
	}
}
