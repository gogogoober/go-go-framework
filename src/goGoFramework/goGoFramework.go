package goGoFramework

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

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

	demoScene := ebiten.NewImage(50, 50)
	demoScene.Fill(color.RGBA{0xff, 0, 0, 0xff})

	return &GoGoFramework{
		GameName: gameName,
		Scene:    demoScene,
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
	screen.DrawImage(g.Scene, op)

	// if screen == nil {
	// 	screen.DrawImage(g.Scene, op)
	// } else {
	// 	screen.DrawImage(screen, op)

	// }

}

func (g *GoGoFramework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Width, g.Window.Height
}
