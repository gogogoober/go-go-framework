package player

import (
	"fmt"
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/collisionservice"
	"go-go-Framework/framework/services/inputservice"
	"go-go-Framework/framework/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerComponent struct {
	IsPlayerControled bool
	sprites           []*ebiten.Image
	op                ebiten.DrawImageOptions
	entity.Component
	entity.Position
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

	startP := entity.Position{X: 0, Y: 0, Width: w.Size().X, Hight: w.Size().Y}

	if !playerParams.IsPlayerControled {
		op.GeoM.Translate(100, 100)
		startP.X = 100
		startP.Y = 100
	}
	fmt.Print(w.Size())

	return &PlayerComponent{
		sprites:           []*ebiten.Image{rect},
		op:                op,
		Position:          startP,
		IsPlayerControled: playerParams.IsPlayerControled,
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

	if pc.IsPlayerControled {
		pc.handleMovement(pc.npcs)
	}
	// pc.npc = make([]entity.Component, 0)
}

func (pc *PlayerComponent) SetNpcs(npcs []entity.Component) {
	pc.npcs = npcs

}

func (pc *PlayerComponent) handleMovement(npcs []entity.Component) {
	var isColliding = false
	if pc.IsPlayerControled {
		for _, n := range npcs {
			if collisionservice.AreComponentsColliding(pc, n) {
				fmt.Print("colliding")
				isColliding = true

			}
		}
	}
	if isColliding {
		return

	}

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
