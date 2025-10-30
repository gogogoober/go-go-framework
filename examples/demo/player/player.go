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

const moveSpeed = 5

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
	var dx = int(0)
	var dy = int(0)
	if inputservice.IsKeyStringPressed("w") {
		dy -= moveSpeed
	}

	if inputservice.IsKeyStringPressed("a") {
		dx -= moveSpeed
	}

	if inputservice.IsKeyStringPressed("s") {
		dy += moveSpeed
	}

	if inputservice.IsKeyStringPressed("d") {
		dx += moveSpeed
	}

	x := pc.Position.X + dx
	y := pc.Position.Y + dy
	x2 := pc.Position.X2 + dx
	y2 := pc.Position.Y2 + dy

	var newPosition = positionservice.Position{x, y, x2, y2}

	var willCollide = checkCollision(newPosition, pc.npcs)

	if willCollide {
		return
	}

	pc.Position = newPosition
	pc.op.GeoM.Translate(float64(dx), float64(dy))
}

func checkCollision(pcPosition positionservice.Position, npcs []entity.Component) bool {
	for _, n := range npcs {
		if collisionservice.AreComponentsColliding(pcPosition, *n.GetPosition()) {
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
