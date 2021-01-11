package ui

import (
	"github.com/codemicro/cs-cyberpet/internal/pet"
	"github.com/codemicro/cs-cyberpet/internal/textart"
	"github.com/codemicro/cs-cyberpet/internal/tools"
)

const (

	// These constants represent the position of the top left corner of the information box.
	infoBoxPosX = 2
	infoBoxPosY = 1
)

var (
	infoBoxSizeX int
	infoBoxSizeY int

	// The following four global variables are updated by the ShowCharacterInCenter function
	CharacterXPos           int
	CharacterYPos           int
	LongestCharacterSection int
	ClearCurrentCharacter func()
)

func init() {
	// The size of the info box is determined based on the number of pet statistics and the statistic name lengths at
	//startup

	infoBoxSizeX = tools.FindLongestStringLen(pet.StatNames) + 4 + statTickerLen // plus four compensates for weird
	// spacing
	infoBoxSizeY = len(pet.DefaultPetStats) + 1                                  // plus one compensating for the top
	// bottom border
}

// Scaffold draws the inital screen layout, with the Tux character, the info box, a line across the bottom of the screen
// for the status bar and the ">" symbol where text is inputted from.
func Scaffold() {
	_, screenY := Screen.Size()

	ShowCharacterInCenter(textart.Tux)

	Box(infoBoxPosX, infoBoxPosY, infoBoxPosX+infoBoxSizeX, infoBoxPosY+infoBoxSizeY, " STATS ")
	PrintLine(screenY-BottomLineHeight, '─', false)
	PrintString(">", 0, InputLineNumber)
}

// FindTopLeftCoord is used to find the top left coordinate of an image based on its length and the length of its
// longest part if that image is to be placed in the center of the console
func FindTopLeftCoord(character []string, longestStringLen int) (int, int) {
	screenX, screenY := Screen.Size()
	xpos := (screenX - longestStringLen) / 2
	ypos := (screenY - len(character)) / 2
	return xpos, ypos
}

// MakeClearFunction returns a function the clear any images that are printed to the console based on the image itself
// and position it was printed in.
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

// ShowCharacterInCenter prints the current character in the center of the console and updates the following global variables:
// * ClearCurrentCharacter
// * LongestCharacterSection
// * CharacterXPos
// * CharacterYPos
func ShowCharacterInCenter(character []string) {

	if ClearCurrentCharacter != nil {
		ClearCurrentCharacter()
	}

	LongestCharacterSection = tools.FindLongestStringLen(character)
	CharacterXPos, CharacterYPos = FindTopLeftCoord(character, LongestCharacterSection)

	PrintTransparentMultiString(character, CharacterXPos, CharacterYPos)

	ClearCurrentCharacter = MakeClearFunction(character, CharacterXPos, CharacterYPos)
}
