package display

import (
	"strings"
)

func Scaffold() {
	_, dispY := Screen.Size()
	PrintLine(dispY - BottomLineHeight, '─', false)

	PrintString(">", 0, dispY - BottomLineHeight + 2)

}

func FindTopLeftCoord(imgX, imgY, sizeX, sizeY int) (int, int) {
	remX := sizeX - imgX
	remY := sizeY - imgY
	return remX / 2, remY / 2
}

func CharacterSay(in string, tuxX, tuxY, longestPart, yOffset int) {
	for i, x := range strings.Split(in, "\n") {
		starter := "  "
		if i == 0 {
			starter = "< "
		}
		PrintString(starter + x, tuxX + longestPart, tuxY + yOffset + i)
	}
}