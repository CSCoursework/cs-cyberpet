package ui

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

// BeforeShutdown is called before the program exits. It clears the console and terminates the tcell.Screen instance,
// which means that any error messages or stacktraces that may have arisen can be printed normally and automatically by
// Go without being messed up.
func BeforeShutdown() {
	displayLock.Lock()
	Screen.Clear()
	Screen.Fini()
	displayLock.Unlock()
}

// PrintString prints a string in a specified position. This function acquires and releases the display lock before it
// does anything and after is has updated the display accordingly.
func PrintString(in string, posX, posY int) {
	displayLock.Lock()
	rawPrintString(in, posX, posY)
	Screen.Show()
	displayLock.Unlock()
}

// rawPrintString prints the provided string to the screen at a specified position. This function should not be called
// without first locking displayLock
func rawPrintString(in string, posX, posY int) {
	rawPrintRunes([]rune(in), posX, posY)
}

// rawPrintRunes prints the slice of runes to the screen at a specified position. This function should not be called
// without first locking displayLock
func rawPrintRunes(in []rune, posX, posY int) {
	if len(in) == 0 {
		return
	}
	for i, char := range in {
		Screen.SetContent(posX+i, posY, char, nil, tcell.StyleDefault)
	}
}

// rawTransparentPrintString prints the provided string to the screen at a specified position without printing spaces
// (" "). This function should not be called without first locking displayLock
func rawTransparentPrintString(in string, posX, posY int) {
	rawTransparentPrintRunes([]rune(in), posX, posY)
}

// rawTransparentPrintRunes prints the provided string to the screen at a specified position without printing spaces
// (" "). This function should not be called without first locking displayLock
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

// PrintTransparentMultiString prints a slice of strings beneath each other while not printing spaces (" "), with the
// top left hand corner of the printed set being at the provided coordinate.
func PrintTransparentMultiString(in []string, posX, posY int) {
	displayLock.Lock()
	for i, x := range in {
		rawTransparentPrintString(x, posX, posY+i)
	}
	Screen.Show()
	displayLock.Unlock()
}

// PrintTransparentMultiString prints a slice of slices of runes beneath each other while not printing spaces (" "),
// with the top left hand corner of the printed set being at the provided coordinate.
func PrintTransparentMultiRuneSlice(in [][]rune, posX, posY int) {
	displayLock.Lock()
	for i, x := range in {
		rawTransparentPrintRunes(x, posX, posY+i)
	}
	Screen.Show()
	displayLock.Unlock()
}

// PrintLine prints a line of the specified rune along a certain point on the X or Y axis. If isVertical is true, then
// the line is printed along x = fixedPos. Else, if isVertical is false, the line is printed along y = fixedPos
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

// CharacterSay prints the specified string prefixed by "> " at the specified offset from the top right hand corner of
// the currently displayed character. If the string has newlines in it, it is split and each line is printed underneath
// the previous, with a prefix of "  ". A function that can be used to clear the text printed by this function is
// returned.
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
