package minigame

import (
	"errors"
	"github.com/codemicro/cs-cyberpet/internal/ui"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var tttBoard = [][]rune{
	[]rune("    A   B   C"),
	[]rune("  ┌───┬───┬───┐"),
	[]rune("1 │ Q │ Q │ Q │"),
	[]rune("  ├───┼───┼───┤"),
	[]rune("2 │ Q │ Q │ Q │"),
	[]rune("  ├───┼───┼───┤"),
	[]rune("3 │ Q │ Q │ Q │"),
	[]rune("  └───┴───┴───┘"),
}

var (
	inputValidationRegex = regexp.MustCompile(`(?m)[a-cA-C][1-3]`)
	modPositions [3][3][2]int
)

func init() {
	modPositions[0][0] = [2]int{2, 4}
	modPositions[0][1] = [2]int{2, 8}
	modPositions[0][2] = [2]int{2, 12}
	modPositions[1][0] = [2]int{4, 4}
	modPositions[1][1] = [2]int{4, 8}
	modPositions[1][2] = [2]int{4, 12}
	modPositions[2][0] = [2]int{6, 4}
	modPositions[2][1] = [2]int{6, 8}
	modPositions[2][2] = [2]int{6, 12}
}

func createBoardFromState(state [3][3]rune) [][]rune {

	outputBoard := make([][]rune, len(tttBoard))
	copy(outputBoard, tttBoard)
	for i := range outputBoard {
		copy(outputBoard[i], tttBoard[i])
	}

	for colNum, col := range state {
		for rowNum := range col {

			irPos := modPositions[colNum][rowNum]

			outputBoard[irPos[0]][irPos[1]] = state[colNum][rowNum]

		}
	}

	return outputBoard
}

// Tictactoe runs a tictactoe game... or would, if I finished it.
// A return value of true would mean the player wins the game - conversely, a return value of false means the player
// lost.
func Tictactoe() bool {

	var gameState [3][3]rune

	xPos, screenY := ui.Screen.Size()
	yPos := (screenY - len(tttBoard)) / 2

	xPos -= 30

	defer func() {
		var s []string
		for _, l := range tttBoard {
			s = append(s, string(l))
		}
		ui.MakeClearFunction(s, xPos, yPos)()
	}()

	ui.PrintTransparentMultiRuneSlice(createBoardFromState(gameState), xPos, yPos)

	for i := 0; i < 6; i += 1 {
		ui.PrintLine(ui.OptionsLineNumber, ' ', false)
		ui.PrintString("Input a coordinate (eg A1)", 0, ui.OptionsLineNumber)

		for {
			var inp string
			for !inputValidationRegex.MatchString(inp) {

				if inp != "" {
					ui.PrintLine(ui.OptionsLineNumber, ' ', false)
					ui.PrintString("That coordinate was invalid. Input another coordinate (eg A1)", 0, ui.OptionsLineNumber)
				}

				var err error
				inp, err = ui.CollectInputAtPosition(os.Stdin, 2, ui.InputLineNumber, true, 2)
				if err != nil {
					if errors.Is(err, ui.ErrorInputTerminated) {
						return false
					}
					panic(err)
				}
			}
			colLtr := inp[0]
			rowStr := inp[1]

			row := tools.GetCharNumber(strings.ToLower(string(colLtr)))
			col, _ := strconv.Atoi(string(rowStr))

			// arrays start at zero!
			col -= 1

			if gameState[col][row] == 0 {
				gameState[col][row] = 'x'
				break
			} else {
				ui.PrintLine(ui.OptionsLineNumber, ' ', false)
				ui.PrintString("That position is already occupied. Input another coordinate (eg A1)", 0, ui.OptionsLineNumber)
			}
		}

		ui.PrintTransparentMultiRuneSlice(createBoardFromState(gameState), xPos, yPos)
	}

	return true

}