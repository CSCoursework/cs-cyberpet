package minigame

import (
	"errors"
	"fmt"
	"github.com/codemicro/cs-cyberpet/internal/display"
	"github.com/codemicro/cs-cyberpet/internal/tools"
	"os"
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

var modPositions [3][3][2]int

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

func Tictactoe() bool {

	var gameState [3][3]rune

	xPos, screenY := display.Screen.Size()
	yPos := (screenY - len(tttBoard)) / 2

	xPos -= 30

	display.PrintTransparentMultiRuneSlice(tttBoard, xPos, yPos)

	for i := 0; i < 6; i += 1 {
		display.PrintString(fmt.Sprintf("loop %d input a coordinate, eg A1", i), 0, display.OptionsLineNumber)

		inp, err := display.CollectInputAtPosition(os.Stdin, 2, display.InputLineNumber, false, 2)
		if err != nil {
			if errors.Is(err, display.ErrorInputTerminated) {
				return false
			}
			panic(err)
		}
		colLtr := inp[0]
		rowStr := inp[1]

		row := tools.GetCharNumber(strings.ToLower(string(colLtr)))
		col, _ := strconv.Atoi(string(rowStr))

		// arrays start at zero!
		col -= 1

		gameState[col][row] = 'x'

		display.PrintTransparentMultiRuneSlice(createBoardFromState(gameState), xPos, yPos)
	}

	// Yes, this is the beginnings of the game. No, it never got finished because I ran out of time.

	return true

}