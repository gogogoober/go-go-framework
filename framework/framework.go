package framework

import (
	"go-go-Framework/framework/registry"
	"go-go-Framework/framework/scheduler"

	"github.com/hajimehoshi/ebiten/v2"
)

type GoGoFramework struct {
	GameName  string
	Window    *GoGoWindow
	Registry  *registry.Registry
	Scheduler *scheduler.Scheduler
}

type GoGoWindow struct {
	Width  int
	Height int
}

func NewGoGoFrameworkWithDefaults(gameName string) *GoGoFramework {
	reg := registry.NewRegistry()

	return &GoGoFramework{
		GameName:  gameName,
		Registry:  reg,
		Scheduler: scheduler.NewScheduler(reg),
		Window: &GoGoWindow{
			Width:  620,
			Height: 620,
		},
	}
}

func (g *GoGoFramework) Update() error {

	// if inputservice.IsKeyStringPressed("a") {
	// 	fmt.Printf("Escape")
	// }
	g.Scheduler.ScheduleUpdates()

	return nil
}

func (g *GoGoFramework) Draw(screen *ebiten.Image) {
	g.Scheduler.ScheduleDrawings(screen)
}

func (g *GoGoFramework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Width, g.Window.Height
}
