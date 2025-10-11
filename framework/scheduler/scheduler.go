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

func (gs *Scheduler) GetComponentToDraw() []*ebiten.Image {
	var items = make([]*ebiten.Image, 0)

	items = append(items, gs.registry.Players...)
	return items
}

func (gs *Scheduler) GetItemsToDraw() []*ebiten.Image {
	var items = make([]*ebiten.Image, 0)

	items = append(items, gs.registry.Scene)

	for _, n := range gs.registry.Npcs {
		items = append(items, n)
	}

	return items
}

func (gs *Scheduler) ScheduleDrawings(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for _, image := range gs.GetItemsToDraw() {
		screen.DrawImage(image, op)
	}

	for _, image := range gs.GetComponentToDraw() {
		screen.DrawImage(image, op)
	}
}
