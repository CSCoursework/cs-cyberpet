package ui

import (
	"github.com/codemicro/cs-cyberpet/internal/pet"
	"github.com/codemicro/cs-cyberpet/internal/textart"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"os"
	"time"
)

const (
	fullBlock     = '█'
	emptyBlock    = ' '
	// the amount of characters available for a single stats progress bar
	statTickerLen = 20
)

// UpdateStats takes a reference to a pet and shows its statistics in the info box in the console
func UpdateStats(petInfo *pet.Pet) {

	longestStatNameLen := tools.FindLongestStringLen(pet.StatNames) + 2 // adding two adds a gap between the label and
	// the ticker bar

	displayLock.Lock()
	petInfo.StatLock.RLock()

	for i, stat := range petInfo.Stats {
		lineNum := infoBoxPosY + i + 1 // plus one to allow for the top border

		// make label
		asRunes := []rune(tools.RightPadString(stat.Name, longestStatNameLen, ' '))

		// make indicator tape
		numBlocks := stat.Value / (100 / statTickerLen)
		numSpaces := statTickerLen - numBlocks
		asRunes = append(asRunes, tools.MakeRuneSlice(fullBlock, numBlocks)...)
		asRunes = append(asRunes, tools.MakeRuneSlice(emptyBlock, numSpaces)...)

		rawPrintRunes(asRunes, infoBoxPosX+2, lineNum)
	}

	petInfo.StatLock.RUnlock()

	Screen.Show()
	displayLock.Unlock()
}

// StartStatLoop starts a background worker that updates that statistics in the UI whenever they are modified in the
// worker. Further explanation about how this function works can be found in the README.md file for this project.
func StartStatLoop(pt *pet.Pet) {
	UpdateStats(pt)
	go func(pt *pet.Pet) {
		for <-pt.StatUpdateNotifier {
			UpdateStats(pt)
		}

		ShowCharacterInCenter(textart.Gravestone)

		CharacterSay("oops, I died", 3, 3)

		time.Sleep(time.Second * 2)

		BeforeShutdown()
		os.Exit(0)

	}(pt)
}
