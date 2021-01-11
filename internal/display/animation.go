package display

import (
	"github.com/codemicro/cs-cyberpet/internal/textart"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"time"
)

func AnimateSlideIn(image []string) {
	screenX, _ := Screen.Size()

	// find center point in screen Y
	_, ypos := FindTopLeftCoord(image, tools.FindLongestStringLen(image))

	// find center point in screen X for the barrier
	barrierX := screenX / 2

	cf := func(){}

	modimg := make([]string, len(image))
	copy(modimg, image)

	for xpos := 0; xpos < barrierX; xpos += 4 {

		diff := barrierX - xpos
		for i, ln := range modimg {
			if len(ln) > diff {
				modimg[i] = ln[:diff]
			}
		}

		cf()
		PrintTransparentMultiString(modimg, xpos, ypos)
		cf = MakeClearFunction(modimg, xpos, ypos)
		ShowCharacterInCenter(textart.Tux)
		time.Sleep(time.Second / 8)
	}
}