package main

import (
	"github.com/codemicro/cs-cyberpet/internal/display"
	"github.com/codemicro/cs-cyberpet/internal/game"
	"github.com/codemicro/cs-cyberpet/internal/pet"
)

func main() {

	defer func() {
		display.BeforeShutdown()
	}()

	display.Scaffold()

	pet.CurrentPet = pet.NewPet("Tux")
	display.StartStatLoop(pet.CurrentPet)

	game.Play()

}
