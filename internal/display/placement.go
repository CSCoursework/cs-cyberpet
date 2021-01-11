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

func FindTopLeftCoord(character []string, longestStringLen int) (int, int) {
	screenX, screenY := Screen.Size()
	xpos := (screenX - longestStringLen) / 2
	ypos := (screenY - len(character)) / 2
	return xpos, ypos
}

func MakeClearFunction(character []string, printedXPos, printedYPos int) func() {
	return func() {
		displayLock.Lock()
		for i, line := range character {

			// count number of spaces before first character
			var alignmentSpaces int
			for ii, char := range line {
				if char != ' ' {
					alignmentSpaces = ii
					break
				}
			}
			// make rune array of spaces that is representative of the amount of non-alignment spaces in the text art
			blankRunes := tools.MakeRuneSlice(' ', len(line) - alignmentSpaces)
			// print this rune array at the specified offset
			rawPrintRunes(blankRunes, printedXPos+alignmentSpaces, printedYPos+i)
		}
		Screen.Show()
		displayLock.Unlock()
	}
}

func ShowCharacterInCenter(character []string) {

	if ClearCurrentCharacter != nil {
		ClearCurrentCharacter()
	}

	LongestCharacterSection = tools.FindLongestStringLen(character)
	CharacterXPos, CharacterYPos = FindTopLeftCoord(character, LongestCharacterSection)

	PrintTransparentMultiString(character, CharacterXPos, CharacterYPos)

	ClearCurrentCharacter = MakeClearFunction(character, CharacterXPos, CharacterYPos)
}
