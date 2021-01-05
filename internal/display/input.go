package display

import (
	"errors"
	"fmt"
	"io"
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
				Screen.SetContent(posX+len(inp), posY, ' ', nil, 0)
			}
		} else {
			if limit == 0 || len(inp) < limit {
				inp = append(inp, buf[0])
				Screen.SetContent(posX+len(inp) - 1, posY, rune(buf[0]), nil, 0)
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
		Screen.SetContent(startX, posY, ' ', runeBuf, 0)
		Screen.Show()
	}

	return fmt.Sprint(inp), nil
}
