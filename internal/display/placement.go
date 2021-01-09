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

	CharacterXPos           int
	CharacterYPos           int
	LongestCharacterSection int

	ClearCurrentCharacter func()
)

func init() {
	infoBoxSizeX = tools.FindLongestStringLen(pet.StatNames) + 4 + statTickerLen // plus four compensates for weird spacing
	infoBoxSizeY = len(pet.DefaultPetStats) + 1                                  // plus one compensating for the top bottom border
}

func Scaffold() {
	_, screenY := Screen.Size()

	ShowCharacterInCenter(textart.Tux)

	Box(infoBoxPosX, infoBoxPosY, infoBoxPosX+infoBoxSizeX, infoBoxPosY+infoBoxSizeY, " STATS ")
	PrintLine(screenY-BottomLineHeight, '─', false)
	PrintString(">", 0, InputLineNumber)
}

func ShowCharacterInCenter(character []string) {
	LongestCharacterSection = tools.FindLongestStringLen(character)
	screenX, screenY := Screen.Size()
	CharacterXPos = (screenX - LongestCharacterSection) / 2
	CharacterYPos = (screenY - len(character)) / 2

	if ClearCurrentCharacter != nil {
		ClearCurrentCharacter()
	}

	PrintMultiString(character, CharacterXPos, CharacterYPos)

	ClearCurrentCharacter = func() {
		blankString := string(tools.MakeRuneSlice(' ', LongestCharacterSection))
		for i := 0; i < len(character); i += 1 {
			rawPrintString(blankString, CharacterXPos, CharacterYPos+i)
		}
		Screen.Show()
	}
}
