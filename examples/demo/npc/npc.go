package npc

import (
	"go-go-Framework/framework"
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/collisionservice"
	"go-go-Framework/framework/services/positionservice"
	"go-go-Framework/framework/utils"
	"image/color"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type NpcComponent struct {
	Id                 string
	sprite             *ebiten.Image
	op                 ebiten.DrawImageOptions
	npcs               []entity.Component
	players            []entity.Component
	possiblePossitions []positionservice.GridPosisition
	contact            bool
	entity.Component
	positionservice.Position
	Framework *framework.GoGoFramework
}

type NewNPCOptions struct {
	Width     int
	Height    int
	Framework *framework.GoGoFramework
}

func NewNpcComponent(options NewNPCOptions) *NpcComponent {
	rect := utils.NewRect(float32(options.Width), float32(options.Height), color.RGBA{225, 0, 0, 1})
	op := ebiten.DrawImageOptions{}

	var pp = positionservice.GetPositions(*options.Framework.Options.Window, *options.Framework.Options.GridSize)
	var snakeBody = options.Framework.Registry.GetComponentGroupArray("snake")
	var newPos, _ = positionservice.GetNewPosition(pp, snakeBody)
	op.GeoM.Translate(float64(newPos.X), float64(newPos.Y))

	return &NpcComponent{
		Id:                 uuid.NewString(),
		sprite:             rect,
		op:                 op,
		possiblePossitions: pp,
		contact:            false,
		npcs:               make([]entity.Component, 0),
		Position:           positionservice.GetRectanglePosition(newPos.X, newPos.Y, rect.Bounds().Size().X, rect.Bounds().Size().Y),
		Framework:          options.Framework,
	}
}

func (s *NpcComponent) Init() {}

func (np *NpcComponent) GetId() string {
	return np.Id
}

func (np *NpcComponent) GetSprite(screen *ebiten.Image) *ebiten.Image {
	return np.sprite
}
func (np *NpcComponent) GetOptions() *ebiten.DrawImageOptions {
	return &np.op
}
func (np *NpcComponent) GetPosition() *positionservice.Position {
	return &np.Position
}
func (np *NpcComponent) Update(tick int) {
	var snake = np.Framework.Registry.GetComponentGroupArray("snake")
	var collisionSnake = collisionservice.CheckCollision(np.Position, snake)
	if len(collisionSnake) > 0 {
		var newPos, _ = positionservice.GetNewPosition(np.possiblePossitions, snake)
		position := positionservice.GetRectanglePosition(newPos.X, newPos.Y, np.sprite.Bounds().Size().X, np.sprite.Bounds().Size().Y)

		np.Position = position
		np.op.GeoM.Reset() // important
		np.op.GeoM.Translate(float64(position.X), float64(position.Y))
	}

}
