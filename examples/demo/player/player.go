package player

import (
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/collisionservice"
	"go-go-Framework/framework/services/inputservice"
	"go-go-Framework/framework/services/positionservice"
	"go-go-Framework/framework/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerComponent struct {
	IsPlayerControled bool
	sprites           []*ebiten.Image
	op                ebiten.DrawImageOptions
	width             int
	height            int
	entity.Component
	positionservice.Position
	npcs []entity.Component
}

type NewPlayerParams struct {
	IsPlayerControled bool
	X                 int
	Y                 int
}

func NewPlayerComponent(playerParams NewPlayerParams) *PlayerComponent {
	rect := utils.NewRect(120, 80, color.RGBA{60, 160, 255, 255})
	op := ebiten.DrawImageOptions{}
	var w = rect.Bounds()
	return &PlayerComponent{
		sprites:           []*ebiten.Image{rect},
		op:                op,
		width:             w.Size().X,
		height:            w.Size().Y,
		Position:          positionservice.GetRectanglePosition(0, 0, w.Size().X, w.Size().Y),
		IsPlayerControled: playerParams.IsPlayerControled,
	}
}

func (pc *PlayerComponent) GetSprite() *ebiten.Image {
	return pc.sprites[0]
}

func (pc *PlayerComponent) GetOptions() *ebiten.DrawImageOptions {
	return &pc.op
}

func (pc *PlayerComponent) GetPosition() *positionservice.Position {
	return &pc.Position
}

func (pc *PlayerComponent) Update() {

	if pc.IsPlayerControled {
		pc.handleMovement()
	}
	pc.npcs = make([]entity.Component, 0)

}

func (pc *PlayerComponent) SetNpcs(npcs []entity.Component) {
	pc.npcs = npcs

}

func (pc *PlayerComponent) handleMovement() {
	var isTouching = pc.checkTouching()

	var x = pc.Position.X
	var y = pc.Position.Y

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
	newX := x + dx
	newY := y + dy
	pc.Position = positionservice.GetRectanglePosition(newX, newY, pc.width, pc.height)

	var willCollide = pc.checkCollision()

	if isTouching && willCollide {
		pc.Position = positionservice.GetRectanglePosition(x, y, pc.width, pc.height)
		return

	}

	pc.op.GeoM.Translate(float64(dx), float64(dy))
}

func (pc *PlayerComponent) checkCollision() bool {

	for _, n := range pc.npcs {
		if collisionservice.AreComponentsColliding(pc, n) {
			return true
		}

	}
	return false
}

func (pc *PlayerComponent) checkTouching() bool {

	for _, n := range pc.npcs {

		if collisionservice.AreComponentsTouching(pc, n) {
			return true

		}

	}
	return false
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
