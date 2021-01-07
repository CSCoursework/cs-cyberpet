package display

import (
	"github.com/codemicro/cs-cyberpet/internal/pet"
	"github.com/codemicro/cs-cyberpet/internal/tools"
)

const (
	infoBoxPosX = 2
	infoBoxPosY = 1
)

var (
	infoBoxSizeX int
	infoBoxSizeY int
)

func init() {
	infoBoxSizeX = tools.FindLongestStringLen(pet.StatNames) + 4 + statTickerLen // plus four compensates for weird spacing
	infoBoxSizeY = len(pet.DefaultPetStats) + 1 // plus one compensating for the top bottom border
}

func Scaffold() {
	_, dispY := Screen.Size()
	Box(infoBoxPosX, infoBoxPosY, infoBoxPosX+infoBoxSizeX, infoBoxPosY+infoBoxSizeY, " STATS ")
	PrintLine(dispY - BottomLineHeight, '─', false)
	PrintString(">", 0, dispY - BottomLineHeight + 2)

}

func FindTopLeftCoord(imgX, imgY, sizeX, sizeY int) (int, int) {
	remX := sizeX - imgX
	remY := sizeY - imgY
	return remX / 2, remY / 2
}