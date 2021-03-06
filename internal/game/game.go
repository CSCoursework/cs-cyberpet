package game

import (
	"errors"
	"github.com/codemicro/cs-cyberpet/internal/ui"
	"github.com/codemicro/cs-cyberpet/internal/game/minigame"
	"github.com/codemicro/cs-cyberpet/internal/pet"
	"github.com/codemicro/cs-cyberpet/internal/textart"
	"math/rand"
	"time"
)

var options = []string{
	"sleep",
	"play",
	"eat",
	"drink",
	"quit",
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Play runs the main loop for the game.
func Play() {

	for {
		_, chosenOption, err := ui.SelectOption(options)
		if err != nil {
			if errors.Is(err, ui.ErrorInputTerminated) {
				return
			} else {
				panic(err)
			}
		}

		if !pet.CurrentPet.IsDead {

			switch chosenOption {
			case "sleep":

				ui.ShowCharacterInCenter(textart.Bed)

				cf := ui.CharacterSay("zzzzZZZZZzzzz", 3, 3)
				time.Sleep(2 * time.Second)

				pet.CurrentPet.SetStat("Fatigue", 0)
				pet.CurrentPet.SetStatDelta("Health", 20)

				cf()

				ui.ShowCharacterInCenter(textart.Tux)

			case "play":

				cf := ui.CharacterSay("wheee such fun", 3, 0)
				minigame.Tictactoe()

				// clear anything left over in the options block
				ui.PrintLine(ui.StatusLineNumber, ' ', false)
				ui.PrintLine(ui.OptionsLineNumber, ' ', false)

				time.Sleep(time.Second * 2)
				pet.CurrentPet.SetStatDelta("Boredom", -20)
				cf()

			case "eat":

				// randomly pick a food
				var food []string
				if random.Intn(2) == 1 { // this is picking a number from 0<=n<2, *not* 0<=n<=2
					food = textart.Pizza
				} else {
					food = textart.Cheese
				}

				cf := ui.CharacterSay("*nom nom nom*", 3, 0)
				ui.AnimateSlideIn(food)
				cf()
				pet.CurrentPet.SetStatDelta("Hunger", -15)

			case "drink":

				// randomly pick a drink
				var drink []string
				if random.Intn(2) == 1 {
					drink = textart.Tea
				} else {
					drink = textart.Wine
				}

				cf := ui.CharacterSay("*sluuuuurp*", 3, 0)
				ui.AnimateSlideIn(drink)
				cf()
				pet.CurrentPet.SetStat("Thirst", -15)

			case "quit":
				return

			}

		}

	}

}
