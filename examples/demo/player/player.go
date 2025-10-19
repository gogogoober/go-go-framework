package player

import (
	"fmt"
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/inputservice"
	"go-go-Framework/framework/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerComponent struct {
	isPlayer bool
	sprites  []*ebiten.Image
	op       ebiten.DrawImageOptions
	entity.Component
	entity.Position
}

func NewPlayerComponent(isPlayer bool) *PlayerComponent {
	rect := utils.NewRect(120, 80, color.RGBA{60, 160, 255, 255})
	op := ebiten.DrawImageOptions{}
	var w = rect.Bounds()

	startP := entity.Position{0, 0, w.Size().X, w.Size().Y}

	if !isPlayer {
		op.GeoM.Translate(100, 100)
		startP.X = 100
		startP.Y = 100
	}
	fmt.Print(w.Size())

	return &PlayerComponent{
		sprites:  []*ebiten.Image{rect},
		op:       op,
		Position: startP,
		isPlayer: isPlayer,
	}
}

func (pc *PlayerComponent) GetSprite() *ebiten.Image {
	return pc.sprites[0]
}

func (pc *PlayerComponent) GetOptions() *ebiten.DrawImageOptions {
	return &pc.op
}

func (pc *PlayerComponent) GetPosition() *entity.Position {
	return &pc.Position
}

func (pc *PlayerComponent) Update() {
	if pc.isPlayer {
		pc.handleMovement()
	}
}

func (pc *PlayerComponent) handleMovement() {
	var dx = int(0)
	var dy = int(0)
	if inputservice.IsKeyStringPressed("w") {
		dx, dy = moveUp()
	}

	if inputservice.IsKeyStringPressed("a") {
		dx, dy = moveLeft()
	}

	if inputservice.IsKeyStringPressed("s") {
		dx, dy = moveDown()
	}

	if inputservice.IsKeyStringPressed("d") {
		dx, dy = moveRight()
	}
	pc.Position.X += dx
	pc.Position.Y += dy
	pc.op.GeoM.Translate(float64(dx), float64(dy))
}

func moveUp() (dx, dy int) {
	return 0, -5
}

func moveDown() (dx, dy int) {
	return 0, 5
}

func moveRight() (dx, dy int) {
	return 5, 0
}

func moveLeft() (dx, dy int) {
	return -5, 0
}
