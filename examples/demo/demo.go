package demo

import (
	"log"

	"github.com/gogogoober/go-go-framework/src/goGoFramework"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {

	return goGoFramework.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func RunDemo() {
	err := ebiten.RunGame(&Game{})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
