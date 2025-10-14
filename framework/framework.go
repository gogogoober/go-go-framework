package framework

import (
	"fmt"
	"go-go-Framework/framework/registry"
	"go-go-Framework/framework/scheduler"
	"go-go-Framework/framework/services"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GoGoFramework struct {
	GameName  string
	Window    *GoGoWindow
	Registry  *registry.Registry
	Scheduler *scheduler.Scheduler
	services.Services
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

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		fmt.Printf("TEST")

	}

	return nil
}

func (g *GoGoFramework) Draw(screen *ebiten.Image) {
	g.Scheduler.ScheduleDrawings(screen)
}

func (g *GoGoFramework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Width, g.Window.Height
}
