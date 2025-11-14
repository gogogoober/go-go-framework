package scheduler

import (
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/registry"
	"sync"

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

	var wg sync.WaitGroup

	for _, g := range s.registry.GetLiveComponents() {
		for _, comp := range g {
			wg.Add(1)
			go func(c entity.Component, t int) {
				defer wg.Done()
				c.Update(t)
			}(comp, tick)
		}
	}
	wg.Wait()
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
