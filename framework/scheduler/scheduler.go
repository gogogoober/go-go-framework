package scheduler

import (
	"fmt"
	"go-go-Framework/framework/registry"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scheduler struct {
	registry   *registry.Registry
	globalTick int
}

func NewScheduler(registry *registry.Registry, globalTick int) *Scheduler {
	return &Scheduler{
		registry:   registry,
		globalTick: globalTick,
	}
}

func (s *Scheduler) ScheduleUpdates(tick int) {
	if tick%s.globalTick == 0 {
		fmt.Println("---------")
		fmt.Println("Global Tick")
	}

	for _, g := range s.registry.GetLiveComponents() {
		for _, c := range g {
			c.Update(tick)
		}
	}
}

func (s *Scheduler) ScheduleDrawings(screen *ebiten.Image) {
	for _, g := range s.registry.GetLiveComponents() {
		for _, c := range g {
			sprite := c.GetSprite(screen)
			if sprite == nil {
				continue
			}
			screen.DrawImage(sprite, c.GetOptions())
		}
	}
}
