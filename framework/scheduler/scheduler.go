package scheduler

import (
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

func (s *Scheduler) GetComponentsToDraw() []*ebiten.Image {
	spriteCount := len(s.registry.Npcs) + len(s.registry.Players)
	var items = make([]*ebiten.Image, 0, spriteCount)

	for _, s := range s.registry.Players {
		items = append(items, s.GetSprite())
	}

	for _, s := range s.registry.Npcs {
		items = append(items, s)
	}

	return items
}

func (s *Scheduler) GetComponents() []registry.Component {
	spriteCount := len(s.registry.Npcs) + len(s.registry.Players)
	var items = make([]registry.Component, 0, spriteCount)

	for _, s := range s.registry.Players {
		items = append(items, s)
	}
	return items
}

func (s *Scheduler) ScheduleUpdates() {
	for _, c := range s.GetComponents() {
		c.Update()
	}
}

func (s *Scheduler) ScheduleDrawings(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	for _, image := range s.GetComponentsToDraw() {
		screen.DrawImage(image, op)
	}
}
