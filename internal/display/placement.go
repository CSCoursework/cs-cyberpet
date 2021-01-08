package display

import (
	"github.com/codemicro/cs-cyberpet/internal/pet"
	"github.com/codemicro/cs-cyberpet/internal/textart"
	"github.com/codemicro/cs-cyberpet/internal/tools"
)

const (
	infoBoxPosX = 2
	infoBoxPosY = 1
)

var (
	infoBoxSizeX int
	infoBoxSizeY int

	CharacterXPos  int
	CharacterYPos  int
	LongestCharacterSection int
)

func init() {
	infoBoxSizeX = tools.FindLongestStringLen(pet.StatNames) + 4 + statTickerLen // plus four compensates for weird spacing
	infoBoxSizeY = len(pet.DefaultPetStats) + 1 // plus one compensating for the top bottom border
}

func Scaffold(character []string) {
	screenX, screenY := Screen.Size()

	LongestCharacterSection = tools.FindLongestStringLen(textart.Tux)
	CharacterXPos, CharacterYPos = FindTopLeftCoord(LongestCharacterSection, len(textart.Tux), screenX, screenY-BottomLineHeight)
	PrintMultiString(character, CharacterXPos, CharacterYPos)

	Box(infoBoxPosX, infoBoxPosY, infoBoxPosX+infoBoxSizeX, infoBoxPosY+infoBoxSizeY, " STATS ")
	PrintLine(screenY- BottomLineHeight, '─', false)
	PrintString(">", 0, screenY- BottomLineHeight + 2)

}

func FindTopLeftCoord(imgX, imgY, sizeX, sizeY int) (int, int) {
	remX := sizeX - imgX
	remY := sizeY - imgY
	return remX / 2, remY / 2
}