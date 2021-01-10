package game

import (
	"errors"
	"github.com/codemicro/cs-cyberpet/internal/display"
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

func Play() {

	for {
		_, chosenOption, err := display.SelectOption(options)
		if err != nil {
			if errors.Is(err, display.ErrorInputTerminated) {
				return
			} else {
				panic(err)
			}
		}

		if !pet.CurrentPet.IsDead {

			switch chosenOption {
			case "sleep":

				display.ShowCharacterInCenter(textart.Bed)

				cf := display.CharacterSay("zzzzZZZZZzzzz", 3, 3)
				time.Sleep(2 * time.Second)

				pet.CurrentPet.SetStat("Fatigue", 0)
				pet.CurrentPet.SetStatDelta("Health", 20)

				cf()

				display.ShowCharacterInCenter(textart.Tux)

			case "play":

				// TODO: Minigame. Tictactoe or something?

				cf := display.CharacterSay("wheee such fun", 3, 0)
				time.Sleep(time.Second * 2)
				pet.CurrentPet.SetStatDelta("Boredom", -20)
				cf()

			case "eat":

				// randomly pick a food
				var food []string
				if random.Intn(2) == 1 {
					food = textart.Pizza
				} else {
					food = textart.Cheese
				}

				cf := display.CharacterSay("*nom nom nom*", 3, 0)
				display.AnimateSlideIn(food)
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

				cf := display.CharacterSay("*sluuuuurp*", 3, 0)
				display.AnimateSlideIn(drink)
				cf()
				pet.CurrentPet.SetStat("Thirst", -15)

			case "quit":
				return

			}

		}

	}

}
