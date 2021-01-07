package pet

type Pet struct {
	Name string

	Health int
	Boredom int
	Thirst int
	Hunger int
	Fatigue int
}

func NewPet(name string) *Pet {

	// TODO: gradual decrease goroutine

	return &Pet{
		Name:    name,
		Health:  100,
		Boredom: 0,
		Thirst:  0,
		Hunger:  0,
		Fatigue: 0,
	}
}