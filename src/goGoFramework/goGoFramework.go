package goGoFramework

import "github.com/hajimehoshi/ebiten/v2"

type GoGoFramework struct {
	GameName string
	Scene    *ebiten.Image
	Window   *GoGoWindow
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
	// op := &ebiten.DrawImageOptions{}

	// screen.DrawImage(screen, op)
}

func (g *GoGoFramework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Width, g.Window.Height
}
