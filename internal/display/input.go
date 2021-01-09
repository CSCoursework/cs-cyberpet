package display

import (
	"errors"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"github.com/gdamore/tcell/v2"
	"io"
	"os"
)

var ErrorInputTerminated = errors.New("input interrupted")

func CollectInputAtPosition(reader io.Reader, posX, posY int, clearAfter bool, limit int) (string, error) {
	var inp []byte
	buf := make([]byte, 1)
	startX := posX

	Screen.ShowCursor(posX, posY)
	defer Screen.HideCursor()

	for {
		// read a char
		_, err := reader.Read(buf)
		if err != nil {
			return "", err
		}

		if buf[0] == 13 || buf[0] == 10 { // line endings
			break
		}

		if buf[0] == 3 {
			return "", ErrorInputTerminated
		}

		if buf[0] == 8 { // backspace
			if len(inp) != 0 {
				inp = inp[:len(inp) - 1]
				Screen.SetContent(posX+len(inp), posY, ' ', nil, tcell.StyleDefault)
			}
		} else {
			if limit == 0 || len(inp) < limit {
				inp = append(inp, buf[0])
				Screen.SetContent(posX+len(inp) - 1, posY, rune(buf[0]), nil, tcell.StyleDefault)
			}
		}

		Screen.ShowCursor(posX+len(inp), posY)
		Screen.Show()
	}

	if clearAfter {
		var runeBuf []rune
		for i := 1; i < len(inp); i += 1 {
			runeBuf = append(runeBuf, ' ')
		}
		Screen.SetContent(startX, posY, ' ', runeBuf, tcell.StyleDefault)
		Screen.Show()
	}

	return string(inp), nil
}

func ShowOptions(opts []string) {
	var outputString string
	for i, val := range opts {
		outputString += tools.GetAlphabetChar(i) + ") " + val + "  "
	}
	PrintString(outputString, 0, OptionsLineNumber)
}

func SelectOption(opts []string) (int, string, error) {

	defer func() {
		PrintLine(OptionsLineNumber, ' ', false)
		PrintLine(StatusLineNumber, ' ', false)
	}()

	ShowOptions(opts)
	PrintString("Select an option", 0, StatusLineNumber)

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