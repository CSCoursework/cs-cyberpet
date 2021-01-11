package ui

import (
	"github.com/codemicro/cs-cyberpet/internal/textart"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"time"
)

// AnimateSlideIn shows an animation of the specified image scrolling in from the left hand side of the screen and
// vanishing as is passes the center point of the display on the X axis.
func AnimateSlideIn(image []string) {
	screenX, _ := Screen.Size()

	// find center point in screen Y
	_, ypos := FindTopLeftCoord(image, tools.FindLongestStringLen(image))

	// find center point in screen X for the barrier
	barrierX := screenX / 2

	// the clear function is set as an empty function so it can be called before any clear function is generated on the
	// first iteration of the loop below
	cf := func(){}

	// This has to be copied because slices in Go are only sets of pointers to memory addresses, so even if one is
	// passed by value, only the header of the slice is copied and the pointers to the values remain the same
	modimg := make([]string, len(image))
	copy(modimg, image)

	// while the x position of the image is less than the barrier, ie to the left of it
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