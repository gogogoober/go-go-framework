package framework

import (
	"go-go-Framework/framework/registry"
	"go-go-Framework/framework/scheduler"
	"go-go-Framework/framework/window"

	"github.com/hajimehoshi/ebiten/v2"
)

type GoGoFramework struct {
	Registry  *registry.Registry
	Scheduler *scheduler.Scheduler
	options   GoGoFrameworkNewOptions
}

type GoGoFrameworkNewOptions struct {
	GameName string
	Window   *window.GoGoWindow
	GridSize *window.GoGoWindow
}

func NewGoGoFramework(options GoGoFrameworkNewOptions) *GoGoFramework {
	reg := registry.NewRegistry()
	if options.GameName == "" {
		options.GameName = "New Game"
	}
	if options.Window == nil {
		options.Window = &window.GoGoWindow{Width: 500, Height: 500}
	}
	if options.GridSize == nil {
		options.GridSize = &window.GoGoWindow{Width: 1, Height: 1}
	}

	return &GoGoFramework{
		Registry:  reg,
		Scheduler: scheduler.NewScheduler(reg),
		options:   options,
	}
}

func (g *GoGoFramework) Init() {
	ebiten.SetWindowSize(g.options.Window.Width, g.options.Window.Height)
	ebiten.SetWindowTitle(g.options.GameName)
}

func (g *GoGoFramework) Update() error {
	g.Scheduler.ScheduleUpdates()
	return nil
}

func (g *GoGoFramework) Draw(screen *ebiten.Image) {
	g.Scheduler.ScheduleDrawings(screen)
}

func (g *GoGoFramework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.options.Window.Width, g.options.Window.Height
}
