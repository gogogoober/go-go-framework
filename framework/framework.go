package framework

import (
	"go-go-Framework/framework/registry"
	"go-go-Framework/framework/scheduler"
	"go-go-Framework/framework/window"

	"github.com/hajimehoshi/ebiten/v2"
)

type GoGoFramework struct {
	GameName  string
	Window    *window.GoGoWindow
	Registry  *registry.Registry
	Scheduler *scheduler.Scheduler
}

type GoGoFrameworkNewParams struct {
	GameName string
	window.GoGoWindow
}

func NewGoGoFramework(params GoGoFrameworkNewParams) *GoGoFramework {
	reg := registry.NewRegistry()

	return &GoGoFramework{
		GameName:  params.GameName,
		Registry:  reg,
		Scheduler: scheduler.NewScheduler(reg),
		Window:    &params.GoGoWindow,
	}
}

func (g *GoGoFramework) Init() {
	ebiten.SetWindowSize(g.Window.Width, g.Window.Height)
	ebiten.SetWindowTitle(g.GameName)
}

func (g *GoGoFramework) Update() error {
	g.Scheduler.ScheduleUpdates()
	return nil
}

func (g *GoGoFramework) Draw(screen *ebiten.Image) {
	g.Scheduler.ScheduleDrawings(screen)
}

func (g *GoGoFramework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Width, g.Window.Height
}
