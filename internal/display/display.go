package display

import (
	"fmt"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"github.com/gdamore/tcell/v2"
	"os"
	"strings"
	"sync"
	"time"
)

const BottomLineHeight = 4

var (
	Screen tcell.Screen

	StatusLineNumber  int
	InputLineNumber   int
	OptionsLineNumber int

	displayLock sync.Mutex
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

func BeforeShutdown() {
	displayLock.Lock()
	Screen.Clear()
	Screen.Fini()
	displayLock.Unlock()
}

func PrintString(in string, posX, posY int) {
	displayLock.Lock()
	rawPrintString(in, posX, posY)
	Screen.Show()
	displayLock.Unlock()
}

func rawPrintString(in string, posX, posY int) {
	rawPrintRunes([]rune(in), posX, posY)
}

func rawPrintRunes(in []rune, posX, posY int) {
	if len(in) == 0 {
		return
	}
	for i, char := range in {
		Screen.SetContent(posX+i, posY, char, nil, tcell.StyleDefault)
	}
}

func rawTransparentPrintString(in string, posX, posY int) {
	rawTransparentPrintRunes([]rune(in), posX, posY)
}

func rawTransparentPrintRunes(in []rune, posX, posY int) {
	if len(in) == 0 {
		return
	}
	for i, char := range in {
		if char != ' ' {
			Screen.SetContent(posX+i, posY, char, nil, tcell.StyleDefault)
		}
	}
}

func PrintMultiString(in []string, posX, posY int) {
	displayLock.Lock()
	for i, x := range in {
		rawPrintString(x, posX, posY+i)
	}
	Screen.Show()
	displayLock.Unlock()
}

func PrintTransparentMultiString(in []string, posX, posY int) {
	displayLock.Lock()
	for i, x := range in {
		rawTransparentPrintString(x, posX, posY+i)
	}
	Screen.Show()
	displayLock.Unlock()
}

func PrintLine(fixedPos int, char rune, isVertical bool) {
	displayLock.Lock()
	var x, y, totalLen int
	xs, ys := Screen.Size()
	if isVertical {
		totalLen = ys
		x = fixedPos
	} else {
		totalLen = xs
		y = fixedPos
	}
	rawPrintRunes(tools.MakeRuneSlice(char, totalLen), x, y)
	Screen.Show()
	displayLock.Unlock()
}

func CharacterSay(in string, yOffset, xOffset int) (clearFunc func()) {
	displayLock.Lock()
	splitLines := strings.Split(in, "\n")
	for i, x := range splitLines {
		starter := "  "
		if i == 0 {
			starter = "< "
		}
		for ii, char := range starter + x {
			Screen.ShowCursor(CharacterXPos+LongestCharacterSection+xOffset+ii+1, CharacterYPos+yOffset+i)
			rawPrintString(string(char), CharacterXPos+LongestCharacterSection+xOffset+ii, CharacterYPos+yOffset+i)
			Screen.Show()
			time.Sleep(time.Millisecond * 50)
		}
	}
	Screen.HideCursor()
	displayLock.Unlock()
	return func() {
		displayLock.Lock()
		blankString := tools.MakeRuneSlice(' ', tools.FindLongestStringLen(splitLines)+2)
		for i := 0; i < len(splitLines); i += 1 {
			rawPrintRunes(blankString, CharacterXPos+LongestCharacterSection+xOffset, CharacterYPos+yOffset+i)
		}
		Screen.Show()
		displayLock.Unlock()
	}
}
