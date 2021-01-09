package display

import (
	"github.com/codemicro/cs-cyberpet/internal/pet"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"github.com/gdamore/tcell/v2"
)

const (
	fullBlock     = '█'
	emptyBlock    = ' '
	statTickerLen = 20
)

func UpdateStats(petInfo *pet.Pet) {

	longestStatNameLen := tools.FindLongestStringLen(pet.StatNames) + 2 // adding two adds a gap between the label and the ticker bar

	petInfo.StatLock.RLock()

	for i, stat := range petInfo.Stats {
		lineNum := infoBoxPosY + i + 1 // plus one to allow for the border, plus one for nice spacing

		// make label
		asRunes := []rune(tools.RightPadString(stat.Name, longestStatNameLen, ' '))

		// make indicator tape
		numBlocks := stat.Value / (100 / statTickerLen)
		numSpaces := statTickerLen - numBlocks
		asRunes = append(asRunes, tools.MakeRuneSlice(fullBlock, numBlocks)...)
		asRunes = append(asRunes, tools.MakeRuneSlice(emptyBlock, numSpaces)...)

		Screen.SetContent(infoBoxPosX + 2, lineNum, asRunes[0], asRunes[1:], tcell.StyleDefault)
	}

	petInfo.StatLock.RUnlock()

	Screen.Show()
}

func StartStatLoop(pt *pet.Pet) {
	UpdateStats(pt)
	go func(pt *pet.Pet) {
		for {
			UpdateStats(pt)
			time.Sleep(pet.StatUpdateInterval)
		}
	}(pt)
}