package demo

import (
	"go-go-Framework/examples/demo/player"
	"go-go-Framework/framework"
	"go-go-Framework/framework/utils"

	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	framework *framework.GoGoFramework
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
		framework: framework.NewGoGoFrameworkWithDefaults("DEMO"),
	}

	newScene := ebiten.NewImage(100, 100)
	newScene.Fill(color.RGBA{0xee, 10, 50, 0xff})

	circ := utils.NewCircle(40, color.RGBA{0xee, 0x10, 0x32, 0xff})
	rect := utils.NewRect(120, 80, color.RGBA{60, 160, 255, 255})

	// player := utils.NewCircle(40, color.RGBA{0xee, 0x10, 0x32, 0xff})

	demo.framework.Registry.SetScene(newScene)
	demo.framework.Registry.AddNpc(rect)
	demo.framework.Registry.AddNpc(circ)

	playerC := player.NewPlayerComponent()
	demo.framework.Registry.AddPlayer(playerC)

	ebiten.SetWindowSize(620, 620)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")
	err := ebiten.RunGame(demo)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
