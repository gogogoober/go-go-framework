package scheduler

import (
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/registry"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scheduler struct {
	registry *registry.Registry
}

func NewScheduler(registry *registry.Registry) *Scheduler {
	return &Scheduler{
		registry: registry,
	}
}

func (s *Scheduler) GetComponentsToDraw() []entity.Component {
	spriteCount := len(s.registry.Npcs) + len(s.registry.Players)
	var items = make([]entity.Component, 0, spriteCount)

	for _, s := range s.registry.Players {
		items = append(items, s)
	}

	for _, s := range s.registry.Npcs {
		items = append(items, s)
	}

	return items
}

func (s *Scheduler) GetPlayersComponents() []entity.Component {
	spriteCount := len(s.registry.Players)
	var items = make([]entity.Component, 0, spriteCount)

	for _, s := range s.registry.Players {
		items = append(items, s)
	}
	return items
}

func (s *Scheduler) GetNpcsComponents() []entity.Component {
	spriteCount := len(s.registry.Npcs)
	var items = make([]entity.Component, 0, spriteCount)

	for _, s := range s.registry.Npcs {
		items = append(items, s)
	}
	return items
}

func (s *Scheduler) ScheduleUpdates() {

	for _, c := range s.GetPlayersComponents() {
		c.SetNpcs(s.GetNpcsComponents())
		c.Update()
	}

	for _, c := range s.GetNpcsComponents() {
		c.Update()
	}
}

func (s *Scheduler) ScheduleDrawings(screen *ebiten.Image) {

	for _, comp := range s.GetComponentsToDraw() {
		screen.DrawImage(comp.GetSprite(), comp.GetOptions())
	}
}
