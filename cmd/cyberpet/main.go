package main

import (
	"github.com/codemicro/cs-cyberpet/internal/ui"
	"github.com/codemicro/cs-cyberpet/internal/game"
	"github.com/codemicro/cs-cyberpet/internal/pet"
)

func main() {

	defer func() {
		ui.BeforeShutdown()
	}()

	ui.Scaffold()

	pet.CurrentPet = pet.NewPet("Tux")
	ui.StartStatLoop(pet.CurrentPet)

	game.Play()

}
