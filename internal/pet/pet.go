package pet

import (
	"errors"
	"strings"
	"sync"
	"time"
)

var (
	DefaultPetStats = []*Stat{
		{Name: "Health", Value: 100},
		{Name: "Boredom", Value: 0, Delta: 5},
		{Name: "Thirst", Value: 0, Delta: 3},
		{Name: "Hunger", Value: 0, Delta: 2},
		{Name: "Fatigue", Value: 0, Delta: 2},
	}
	StatNames          []string
	StatUpdateInterval = time.Second * 5

	CurrentPet *Pet
)

func init() {
	for _, stat := range DefaultPetStats {
		StatNames = append(StatNames, stat.Name)
	}
}

type Pet struct {
	Name   string
	IsDead bool

	StatLock           *sync.RWMutex
	Stats              []*Stat
	StatUpdateNotifier chan bool
}

type Stat struct {
	Name  string
	Value int
	Delta int
}

// FixStats ensures all stat values are within the range 0<=n<=100
func (p *Pet) FixStats() {
	for _, stat := range p.Stats {
		nv := stat.Value
		if nv > 100 {
			nv = 100
		} else if nv < 0 {
			nv = 0
		}
		stat.Value = nv
	}
}

// modval modifies the value of a set statistic by name, using the provided function
func (p *Pet) modval(sname string, modfunc func(int) int) {
	p.StatLock.Lock()
	defer func() {
		p.FixStats()
		p.StatLock.Unlock()
		p.StatUpdateNotifier <- true
	}()
	for _, stat := range p.Stats {
		if strings.EqualFold(stat.Name, sname) {
			stat.Value = modfunc(stat.Value)
			return
		}
	}
	panic(errors.New("modval: specified value not found"))
}

// SetStat sets the statistic by name to a certain value
func (p *Pet) SetStat(sname string, val int) {
	p.modval(sname, func(_ int) int {
		return val
	})
}

// SetStatDelta applies a delta to a statistic by name
func (p *Pet) SetStatDelta(sname string, delta int) {
	p.modval(sname, func(x int) int {
		return x + delta
	})
}

// NewPet creates a new Pet instance and returns a reference to it
func NewPet(name string) *Pet {

	p := &Pet{
		Name:               name,
		StatLock:           new(sync.RWMutex),
		Stats:              DefaultPetStats,
		StatUpdateNotifier: make(chan bool),
	}

	// background stat update worker
	// Further explanation about how this works can be found in the README.md file
	go func() {
		for {

			if p.IsDead {
				p.StatUpdateNotifier <- false // stop updating stats, trigger gravestone
				return
			}

			time.Sleep(StatUpdateInterval)

			var healthStatIndex int
			for i, stat := range p.Stats {
				if stat.Name == "Health" {
					healthStatIndex = i
				}
			}

			p.StatLock.Lock()
			for _, stat := range p.Stats {
				nv := stat.Value + stat.Delta
				if nv > 100 {
					nv = 100
				} else if nv < 0 {
					nv = 0
				}
				stat.Value = nv

				switch stat.Name {
				case "Thirst":
					if stat.Value == 100 {
						p.Stats[healthStatIndex].Value -= 5
					}
				case "Hunger":
					if stat.Value == 100 {
						p.Stats[healthStatIndex].Value -= 3
					}
				case "Fatigue":
					if stat.Value == 100 {
						p.Stats[healthStatIndex].Value -= 1
					}
				}

			}

			if p.Stats[healthStatIndex].Value <= 0 {
				p.IsDead = true
				p.Stats[healthStatIndex].Value = 0
			}

			p.StatLock.Unlock()

			p.StatUpdateNotifier <- true // keep updating stats
		}
	}()

	return p
}
