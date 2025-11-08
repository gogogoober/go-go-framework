package demo

import (
	"go-go-Framework/examples/demo/hud"
	"go-go-Framework/examples/demo/npc"
	"go-go-Framework/examples/demo/snake"
	"go-go-Framework/framework"
	"go-go-Framework/framework/window"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	framework *framework.GoGoFramework
}

type GridPosisition struct {
	x int
	y int
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
	var grid = 20

	var gameOptions = framework.GoGoFrameworkNewOptions{
		Window:   &window.GoGoWindow{Width: width, Height: height},
		GridSize: &window.GoGoWindow{Width: width / grid, Height: height / grid},
	}

	var framework = framework.NewGoGoFramework(gameOptions)

	demo := &Game{
		framework: framework,
	}

	hud1 := hud.NewHud(framework)
	snake1 := snake.NewSnake(snake.NewSnakeOptions{Height: width / grid, Framework: framework})
	apple := npc.NewNpcComponent(npc.NewNPCOptions{Width: width / grid, Height: height / grid, Framework: framework})

	demo.framework.Registry.AddComponent(hud1, "hud")
	demo.framework.Registry.AddComponent(snake1, "snake")
	demo.framework.Registry.AddComponent(apple, "apple")

	demo.framework.Init()

	err := ebiten.RunGame(demo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
