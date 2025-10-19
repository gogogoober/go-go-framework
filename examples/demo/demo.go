package demo

import (
	"go-go-Framework/examples/demo/player"
	"go-go-Framework/framework"
	"go-go-Framework/framework/window"
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
		framework: framework.NewGoGoFramework(framework.GoGoFrameworkNewParams{"Game Name", window.GoGoWindow{620, 620}}),
	}

	demo.framework.Registry.AddPlayer(player.NewPlayerComponent(true))
	demo.framework.Registry.AddPlayer(player.NewPlayerComponent(false))
	demo.framework.Init()

	err := ebiten.RunGame(demo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
