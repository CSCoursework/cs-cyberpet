package ui

import (
	"errors"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"github.com/gdamore/tcell/v2"
	"io"
	"os"
)

// ErrorInputTerminated is the error returned when a user presses CTRL+C during a text input
var ErrorInputTerminated = errors.New("input interrupted")

// CollectInputAtPosition collects input from the provided io.Reader from the specified position on the screen. If
// clearAfter is true, the input echoed to the display is cleared after a newline is entered. limit specifies the
// maximum number of bytes that will be accepted as input.
func CollectInputAtPosition(reader io.Reader, posX, posY int, clearAfter bool, limit int) (string, error) {
	var inp []byte
	buf := make([]byte, 1)
	startX := posX

	Screen.ShowCursor(posX, posY)
	defer Screen.HideCursor() // run when this function exists

	for {
		// read a singly byte
		_, err := reader.Read(buf)
		if err != nil {
			return "", err
		}

		if buf[0] == 13 || buf[0] == 10 { // Carriage return (first part of the Windows line delimiter) and line feed
			// values (OSX/Unix line delimiter) for specifying when input is completed
			break
		}

		if buf[0] == 3 { // ctrl+c
			return "", ErrorInputTerminated
		}

		displayLock.Lock()
		if buf[0] == 8 { // backspace
			if len(inp) != 0 {
				inp = inp[:len(inp)-1]
				Screen.SetContent(posX+len(inp), posY, ' ', nil, tcell.StyleDefault)
			}
		} else {
			if limit == 0 || len(inp) < limit {
				inp = append(inp, buf[0])
				Screen.SetContent(posX+len(inp)-1, posY, rune(buf[0]), nil, tcell.StyleDefault)
			}
		}

		Screen.ShowCursor(posX+len(inp), posY)
		Screen.Show()
		displayLock.Unlock()
	}

	if clearAfter {
		displayLock.Lock()
		rawPrintRunes(tools.MakeRuneSlice(' ', len(inp)), startX, posY)
		Screen.Show()
		displayLock.Unlock()
	}

	return string(inp), nil
}

// ShowOptions shows a set of options with a corresponding letter based off the position in the slice in the options
// line of the bottom bar of the display.
func ShowOptions(opts []string) {
	var outputString string
	for i, val := range opts {
		outputString += tools.GetAlphabetChar(i) + ") " + val + "  "
	}
	PrintString(outputString, 0, OptionsLineNumber)
}

// SelectOption wraps ShowOptions, but also adds a prompt for the user to input one of the options. The selected option
// index and value are returned, and the printed text is cleared afterwards
func SelectOption(opts []string) (int, string, error) {

	defer func() {
		// This function is called automagically whenever SelectOption returns, which can be multiple places depending
		// on if an error was returned by CollectInputAtPosition
		PrintLine(OptionsLineNumber, ' ', false)
		PrintLine(StatusLineNumber, ' ', false)
	}()

	ShowOptions(opts)
	PrintString("Select an option", 0, StatusLineNumber)

	// this is some slightly fancy loop control that says "while the selected option number is larger than the number of
	// options or smaller than zero, keep prompting for an input". The loop starting value for the option number is one
	//greater than the number of options to ensure that this loop is run at least once.
	optNum := len(opts) + 1
	for optNum > len(opts) || optNum < 0 {
		inp, err := CollectInputAtPosition(os.Stdin, 2, InputLineNumber, true, 1)
		if err != nil {
			return 0, "", err
		}

		optNum = tools.GetCharNumber(inp)
		PrintString("Invalid option. Please select another", 0, StatusLineNumber)
	}

	return optNum, opts[optNum], nil
}
