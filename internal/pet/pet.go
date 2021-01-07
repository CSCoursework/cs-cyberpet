package pet

type Pet struct {
	Name string

	Stats []PetStat
}

func NewPet(name string) *Pet {

	// TODO: gradual decrease goroutine

	return &Pet{
		Name:    name,
		Stats: []PetStat{
			{Name: "Health", Value: 100},
			{Name: "Boredom", Value: 0},
			{Name: "Thirst", Value: 0},
			{Name: "Hunger", Value: 0},
			{Name: "Fatigue", Value: 0},
		},
	}
}

type PetStat struct {
	Name string
	Value int
	Delta int
}