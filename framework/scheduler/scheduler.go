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

func (s *Scheduler) GetComponentToDraw() []*ebiten.Image {
	var items = make([]*ebiten.Image, 0)

	items = append(items, s.registry.Players...)
	return items
}

func (s *Scheduler) GetItemsToDraw() []*ebiten.Image {
	var items = make([]*ebiten.Image, 0)

	items = append(items, s.registry.Scene)

	for _, n := range s.registry.Npcs {
		items = append(items, n)
	}

	return items
}

func (s *Scheduler) ScheduleDrawings(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for _, image := range s.GetItemsToDraw() {
		screen.DrawImage(image, op)
	}

	for _, image := range s.GetComponentToDraw() {
		screen.DrawImage(image, op)
	}
}
