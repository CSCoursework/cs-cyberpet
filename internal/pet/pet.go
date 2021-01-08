package pet

import (
	"sync"
	"time"
)

var (
	DefaultPetStats = []*Stat{
		{Name: "Health", Value: 100},
		{Name: "Boredom", Value: 0, Delta: 5},
		{Name: "Thirst", Value: 0},
		{Name: "Hunger", Value: 0},
		{Name: "Fatigue", Value: 0},
	}
	StatNames []string
	StatUpdateInterval = time.Second * 5

	CurrentPet *Pet
)

func init() {
	for _, stat := range DefaultPetStats {
		StatNames = append(StatNames, stat.Name)
	}
}

type Pet struct {
	Name string

	StatLock *sync.RWMutex
	Stats []*Stat
}

func NewPet(name string) *Pet {

	p := &Pet{
		Name:    name,
		StatLock: new(sync.RWMutex),
		Stats: DefaultPetStats,
	}

	go func() {
		for {
			p.StatLock.Lock()
			time.Sleep(StatUpdateInterval)
			for _, stat := range p.Stats {
				nv := stat.Value + stat.Delta
				if nv > 100 {
					nv = 100
				} else if nv < 0 {
					nv = 0
				}
				stat.Value = nv
			}
			p.StatLock.Unlock()
		}
	}()

	return p
}

type Stat struct {
	Name string
	Value int
	Delta int
}