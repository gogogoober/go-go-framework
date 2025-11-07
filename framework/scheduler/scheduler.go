package scheduler

import (
	"fmt"
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

func (s *Scheduler) ScheduleUpdates(tick int) {
	if tick%60 == 0 {
		fmt.Println("---------")
		fmt.Println("Global Tick")
	}

	for _, g := range s.registry.GetComponents() {
		for _, c := range g {
			c.Update(tick)
		}
	}
}

func (s *Scheduler) ScheduleDrawings(screen *ebiten.Image) {
	for _, g := range s.registry.GetComponents() {
		for _, c := range g {
			screen.DrawImage(c.GetSprite(), c.GetOptions())
		}
	}
}
