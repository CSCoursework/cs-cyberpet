package pet

var DefaultPetStats = []Stat{
	{Name: "Health", Value: 100},
	{Name: "Boredom", Value: 0},
	{Name: "Thirst", Value: 0},
	{Name: "Hunger", Value: 0},
	{Name: "Fatigue", Value: 0},
}

type Pet struct {
	Name string

	Stats []Stat
}

func NewPet(name string) *Pet {

	// TODO: gradual decrease goroutine

	return &Pet{
		Name:    name,
		Stats: DefaultPetStats,
	}
}

type Stat struct {
	Name string
	Value int
	Delta int
}