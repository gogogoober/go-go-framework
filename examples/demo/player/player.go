package player

import (
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/inputservice"
	"go-go-Framework/framework/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerComponent struct {
	entity.Component
	X       int
	Y       int
	sprites []*ebiten.Image
	op      ebiten.DrawImageOptions
}

func NewPlayerComponent() *PlayerComponent {
	rect := utils.NewRect(120, 80, color.RGBA{60, 160, 255, 255})
	op := ebiten.DrawImageOptions{}

	return &PlayerComponent{
		X:       0,
		Y:       0,
		sprites: []*ebiten.Image{rect},
		op:      op,
	}
}

func (pc *PlayerComponent) GetSprite() *ebiten.Image {
	return pc.sprites[0]
}

func (pc *PlayerComponent) GetOptions() *ebiten.DrawImageOptions {
	return &pc.op
}

func (pc *PlayerComponent) Update() {
	pc.handleMovement()
}

func (pc *PlayerComponent) handleMovement() {
	if inputservice.IsKeyStringPressed("w") {
		pc.op.GeoM.Translate(moveUp())
	}

	if inputservice.IsKeyStringPressed("a") {
		pc.op.GeoM.Translate(moveLeft())
	}

	if inputservice.IsKeyStringPressed("s") {
		pc.op.GeoM.Translate(moveDown())
	}

	if inputservice.IsKeyStringPressed("d") {
		pc.op.GeoM.Translate(moveRight())
	}
}

func moveUp() (dx float64, dy float64) {
	return 0, -5
}

func moveDown() (dx float64, dy float64) {
	return 0, 5
}

func moveRight() (dx float64, dy float64) {
	return 5, 0
}

func moveLeft() (dx float64, dy float64) {
	return -5, 0
}
