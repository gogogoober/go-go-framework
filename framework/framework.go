package framework

import (
	"go-go-Framework/framework/interactionsystem"
	"go-go-Framework/framework/registry"
	"go-go-Framework/framework/scheduler"
	"go-go-Framework/framework/services/timeservice"
	"go-go-Framework/framework/window"

	"github.com/hajimehoshi/ebiten/v2"
)

type GoGoFramework struct {
	tickCount         int
	Registry          *registry.Registry
	Scheduler         *scheduler.Scheduler
	Options           GoGoFrameworkNewOptions
	Services          *GoGoServices
	InteractionSystem *interactionsystem.InteractionSystem[interactionsystem.InteractionItem]
}

type GoGoServices struct {
	*timeservice.TimeService
}

type GoGoFrameworkNewOptions struct {
	GlobalTick        int
	GameName          string
	Window            *window.GoGoWindow
	GridSize          *window.GoGoWindow
	InteractionSystem *interactionsystem.InteractionSystem[interactionsystem.InteractionItem]
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
	if options.GlobalTick == 0 {
		options.GlobalTick = 10
	}

	var GoGoServices = &GoGoServices{
		timeservice.NewTimeSerice(),
	}

	return &GoGoFramework{
		Registry:          reg,
		Scheduler:         scheduler.NewScheduler(reg, options.GlobalTick),
		Options:           options,
		Services:          GoGoServices,
		InteractionSystem: options.InteractionSystem,
	}
}

func (g *GoGoFramework) Init() {
	ebiten.SetWindowSize(g.Options.Window.Width, g.Options.Window.Height)
	ebiten.SetWindowTitle(g.Options.GameName)
}

func (g *GoGoFramework) Update() error {
	g.tickCount++
	g.Services.TimeService.TotalTicks = g.tickCount
	g.Registry.UpdateComponents()
	g.Registry.UpdateSnapshot()
	g.Registry.InitAddedComponents()
	g.Registry.ResetQueues()
	g.Scheduler.ScheduleUpdates(g.tickCount)
	g.InteractionSystem.Reset()
	return nil
}

func (g *GoGoFramework) Draw(screen *ebiten.Image) {
	g.Scheduler.ScheduleDrawings(screen)
}

func (g *GoGoFramework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Options.Window.Width, g.Options.Window.Height
}
