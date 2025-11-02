package player

import (
	"fmt"
	"go-go-Framework/framework"
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/collisionservice"
	"go-go-Framework/framework/services/inputservice"
	"go-go-Framework/framework/services/positionservice"
	"go-go-Framework/framework/utils"
	"image/color"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerComponent struct {
	Index             int
	Id                uuid.UUID
	Count             int
	IsDead            bool
	IsPlayerControled bool
	lastKeyPressed    string
	sprites           []*ebiten.Image
	body              []*entity.Component
	op                ebiten.DrawImageOptions
	width             int
	height            int
	moveDistance      int
	entity.Component
	positionservice.Position
	npcs      []entity.Component
	players   []entity.Component
	Framework *framework.GoGoFramework
}

type NewPlayerParams struct {
	Index             int
	IsPlayerControled bool
	lastKeyPressed    string
	X                 int
	Y                 int
	Width             int
	Height            int
	Framework         *framework.GoGoFramework
}

func NewPlayerComponent(playerParams NewPlayerParams) *PlayerComponent {
	position := positionservice.GetRectanglePosition(playerParams.X, playerParams.Y, playerParams.Width, playerParams.Height)

	if playerParams.lastKeyPressed == "" {
		playerParams.lastKeyPressed = "a"
	}

	rect := utils.NewRect(float32(playerParams.Width-1), float32(playerParams.Height-1), color.RGBA{60, 160, 255, 255})
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(position.X), float64(position.Y))

	return &PlayerComponent{
		Index:             playerParams.Index,
		sprites:           []*ebiten.Image{rect},
		lastKeyPressed:    playerParams.lastKeyPressed,
		op:                op,
		width:             playerParams.Width,
		height:            playerParams.Height,
		moveDistance:      playerParams.Width,
		Position:          position,
		IsPlayerControled: playerParams.IsPlayerControled,
		Framework:         playerParams.Framework,
		Count:             4,
		IsDead:            false,
		Id:                uuid.New(),
	}
}

func (pc *PlayerComponent) GetId() uuid.UUID {
	return pc.Id
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

func (pc *PlayerComponent) Update(tick int) {
	pc.lastKeyPressed = lastKeyPressed(pc.lastKeyPressed)

	if tick%60 == 0 {
		if pc.IsPlayerControled {
			pc.handleMovement(pc.lastKeyPressed)
		}
		if pc.Count <= 0 {
			pc.Framework.Registry.RemovePlayerById(pc.Id)
		}
		pc.Count--
	}
	pc.npcs = make([]entity.Component, 0)

}

func (pc *PlayerComponent) SetNpcs(npcs []entity.Component) {
	pc.npcs = npcs
}

func (pc *PlayerComponent) SetPlayers(players []entity.Component) {
	pc.players = players
}

func lastKeyPressed(lastKey string) string {
	if inputservice.IsKeyStringPressed("d") {
		return "d"
	}
	if inputservice.IsKeyStringPressed("a") {
		return "a"
	}
	if inputservice.IsKeyStringPressed("w") {
		return "w"
	}
	if inputservice.IsKeyStringPressed("s") {
		return "s"
	}

	return lastKey
}

func (pc *PlayerComponent) handleMovement(lastKeyPressed string) {
	var dx = int(0)
	var dy = int(0)

	var collidable = pc.players

	if inputservice.IsKeyStringPressed("d") || lastKeyPressed == "d" {
		dx += pc.moveDistance
	}
	if inputservice.IsKeyStringPressed("a") || lastKeyPressed == "a" {
		dx -= pc.moveDistance
	}

	if inputservice.IsKeyStringPressed("w") || lastKeyPressed == "w" {
		dy -= pc.moveDistance
	}
	if inputservice.IsKeyStringPressed("s") || lastKeyPressed == "s" {
		dy += pc.moveDistance
	}

	if collisionservice.CheckCollision(positionservice.MovePosition(pc.Position, dx, dy), collidable) {
		fmt.Println("Game Over x")
		panic("gg")
	}

	var newPosition = positionservice.MovePosition(pc.Position, dx, dy)

	var newBody = NewPlayerComponent(NewPlayerParams{
		Index:             pc.Index + 1,
		IsPlayerControled: true,
		X:                 newPosition.X,
		Y:                 newPosition.Y,
		Width:             pc.width,
		Height:            pc.height,
		Framework:         pc.Framework,
		lastKeyPressed:    pc.lastKeyPressed,
	})

	pc.Framework.Registry.AddPlayer(newBody)
	pc.IsPlayerControled = false
}
