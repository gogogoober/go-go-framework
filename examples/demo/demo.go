package demo

import (
	"go-go-framework/src/gogoFramework"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	framework *gogoFramework.GoGoFramework
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
		framework: gogoFramework.NewGoGoFrameworkWithDefaults("DEMO"),
	}

	newScene := ebiten.NewImage(100, 100)
	newScene.Fill(color.RGBA{0xee, 10, 50, 0xff})

	demo.framework.RegisterScene(newScene)

	ebiten.SetWindowSize(620, 620)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")
	err := ebiten.RunGame(demo)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
