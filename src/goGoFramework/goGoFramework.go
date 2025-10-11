package gogoFramework

import (
	"go-go-framework/src/gogoFramework/component"
	"go-go-framework/src/gogoFramework/registry"
	"go-go-framework/src/gogoFramework/scheduler"

	"github.com/hajimehoshi/ebiten/v2"
)

type GoGoFramework struct {
	GameName string
	Scene    *ebiten.Image
	npc      []*ebiten.Image
	player   []*component.Component
	Window   *GoGoWindow
	registry.Registry
}

type GoGoWindow struct {
	Width  int
	Height int
}

func NewGoGoFrameworkWithDefaults(gameName string) *GoGoFramework {
	return &GoGoFramework{
		GameName: gameName,
		Window: &GoGoWindow{
			Width:  620,
			Height: 620,
		},
	}
}

func (g *GoGoFramework) Update() error {
	return nil
}

func (g *GoGoFramework) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	var images = scheduler.NewScheduler(g.Registry.Scene, g.Registry.Npcs, g.Registry.Players)
	for _, image := range images.GetItemsToDraw() {
		screen.DrawImage(image, op)
	}

	for _, image := range images.GetComponentToDraw() {
		screen.DrawImage(image, op)
	}
}

func (g *GoGoFramework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Width, g.Window.Height
}
