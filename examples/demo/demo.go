package demo

import (
	"fmt"
	"go-go-Framework/examples/demo/npc"
	"go-go-Framework/examples/demo/snake"
	"go-go-Framework/framework"
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/window"
	"log"
	"math/rand"

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
	var grid = 10

	var gameOptions = framework.GoGoFrameworkNewOptions{
		Window:   &window.GoGoWindow{Width: width, Height: height},
		GridSize: &window.GoGoWindow{Width: width / grid, Height: height / grid},
	}

	var framework = framework.NewGoGoFramework(gameOptions)

	demo := &Game{
		framework: framework,
	}

	snake := snake.NewSnake(snake.NewSnakeOptions{Height: width / grid, Framework: framework})
	fmt.Println(snake)
	demo.framework.Registry.AddComponent(snake, "snake")

	// var pp = getPositions(framework)
	// var newPos = getNewPosition(pp, make([]entity.Component, 0))

	// player := player.NewPlayerComponent(player.NewPlayerParams{IsPlayerControled: true, X: newPos.x, Y: newPos.y, Width: width / grid, Height: height / grid, Framework: framework})
	apple := npc.NewNpcComponent(npc.NewNPCOptions{Width: width / grid, Height: height / grid, Framework: framework})

	// demo.framework.Registry.AddPlayer(player)
	demo.framework.Registry.AddNpc(apple)

	demo.framework.Init()

	err := ebiten.RunGame(demo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func getPositions(Framework *framework.GoGoFramework) []GridPosisition {
	gh := Framework.Options.Window.Height
	gw := Framework.Options.Window.Width
	ggh := Framework.Options.GridSize.Height
	ggw := Framework.Options.GridSize.Width

	var hg = (gh / ggh)
	// var wg = (gw / ggw)

	var positions = make([]GridPosisition, hg)
	var count = 0

	for i := 0; i < gh; i = i + ggh {
		positions[count].y = i
		count++
	}
	count = 0
	for i := 0; i < gw; i = i + ggw {
		positions[count].x = i
		count++

	}
	return positions
}

func getNewPosition(pp []GridPosisition, players []entity.Component) GridPosisition {
	var newPosition []GridPosisition

	if len(players) == 0 {
		newPosition = pp
	} else {
		for i := range players {
			for j := range pp {
				var pPos = players[i].GetPosition()
				if pp[j].x != pPos.X && pp[j].y != pPos.Y {
					newPosition = append(newPosition, pp[j])

				}
			}
		}

	}

	var posPick = rand.Intn(len(newPosition))
	var _np = newPosition[posPick]
	return _np
}
