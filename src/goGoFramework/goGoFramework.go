package goGoFramework

import "github.com/hajimehoshi/ebiten/v2"

type GoGoFramework struct {
	GameName string
	Scene    *ebiten.Image
}

func (g *GoGoFramework) Update() error {
	return nil
}

func (g *GoGoFramework) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Scene, nil)
}

func (g *GoGoFramework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
