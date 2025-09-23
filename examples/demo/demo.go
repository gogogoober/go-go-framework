package demo

import (
	"go-go-framework/src/goGoFramework"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	framework *goGoFramework.GoGoFramework
}

func (g *Game) Update() error {
	return g.framework.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.framework.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.framework.Layout(outsideWidth, outsideHeight)
}

func RunDemo() {
	demo := &Game{
		framework: goGoFramework.NewGoGoFrameworkWithDefaults("DEMO"),
	}
	ebiten.SetWindowSize(620*2, 620*2)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")
	err := ebiten.RunGame(demo)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
