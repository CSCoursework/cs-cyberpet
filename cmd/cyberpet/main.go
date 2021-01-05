package main

import (
	"github.com/codemicro/cs-cyberpet/internal/display"
	"github.com/codemicro/cs-cyberpet/internal/textart"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"os"
	"time"
)

func main() {
	display.Scaffold()
	screenX, screenY := display.Screen.Size()
	tuxX, tuxY := display.FindTopLeftCoord(tools.FindLongestString(textart.Tux), len(textart.Tux), screenX, screenY-display.BottomLineHeight)
	display.PrintMultiLine(textart.Tux, tuxX, tuxY)

	inp, _ := display.CollectInputAtPosition(os.Stdin, 2, screenY-display.BottomLineHeight + 1, true, 0)

	display.PrintString(inp, 0, 0)

	time.Sleep(time.Second * 2)
}
