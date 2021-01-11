package main

import (
	"github.com/codemicro/cs-cyberpet/internal/ui"
	"github.com/codemicro/cs-cyberpet/internal/game"
	"github.com/codemicro/cs-cyberpet/internal/pet"
)

func main() {

	defer func() { // runs before this function exits - meaning even if something in the main thread panics and hard
		// quits, this still gets run
		ui.BeforeShutdown()
	}()

	ui.Scaffold()

	pet.CurrentPet = pet.NewPet("Tux")
	ui.StartStatLoop(pet.CurrentPet)

	game.Play()

}
