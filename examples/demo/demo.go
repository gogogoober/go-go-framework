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
	var width = 500
	var height = 500
	var grid = 10

	var gameOptions = framework.GoGoFrameworkNewOptions{
		Window:   &window.GoGoWindow{Width: width, Height: height},
		GridSize: &window.GoGoWindow{Width: grid, Height: grid},
	}

	var framework = framework.NewGoGoFramework(gameOptions)

	demo := &Game{
		framework: framework,
	}

	demo.framework.Registry.AddPlayer(player.NewPlayerComponent(player.NewPlayerParams{IsPlayerControled: true, X: width / 2, Y: width / 2, Width: width / grid, Height: height / grid, Framework: framework}))
	demo.framework.Init()

	err := ebiten.RunGame(demo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
