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
	Width             int
	Height            int
}

const moveSpeed = 5

func NewPlayerComponent(playerParams NewPlayerParams) *PlayerComponent {

	position := positionservice.GetRectanglePosition(playerParams.X-(playerParams.Width/2), playerParams.Y-(playerParams.Height/2), playerParams.Width, playerParams.Height)

	rect := utils.NewRect(float32(playerParams.Width), float32(playerParams.Height), color.RGBA{60, 160, 255, 255})
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(position.X), float64(position.Y))
	return &PlayerComponent{
		sprites:           []*ebiten.Image{rect},
		op:                op,
		width:             playerParams.Width,
		height:            playerParams.Height,
		Position:          position,
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

	if inputservice.IsKeyStringPressed("d") {
		dx += moveSpeed
	}
	if inputservice.IsKeyStringPressed("a") {
		dx -= moveSpeed
	}

	if collisionservice.CheckCollision(positionservice.MovePosition(pc.Position, dx, dy), pc.npcs) {
		dx = 0
	}

	if inputservice.IsKeyStringPressed("w") {
		dy -= moveSpeed
	}
	if inputservice.IsKeyStringPressed("s") {
		dy += moveSpeed
	}

	if collisionservice.CheckCollision(positionservice.MovePosition(pc.Position, dx, dy), pc.npcs) {
		dy = 0
	}

	pc.Position = positionservice.MovePosition(pc.Position, dx, dy)
	pc.op.GeoM.Translate(float64(dx), float64(dy))
}
