package snake

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

type Snake struct {
	Id     string
	Sprite *ebiten.Image
	Op     *ebiten.DrawImageOptions
	positionservice.Position
	entity.Component2
	Framework *framework.GoGoFramework

	IsPlayerControled bool
	lastKeyPressed    string
	moveDistance      int
	Count             int
}

type NewSnakeOptions struct {
	Height         int
	Framework      *framework.GoGoFramework
	lastKeyPressed string
	positionservice.Position
	count int
}

func NewSnake(options NewSnakeOptions) *Snake {
	var pp = positionservice.GetPositions(*options.Framework.Options.Window, *options.Framework.Options.GridSize)
	var newPos = positionservice.GetNewPosition(pp, make([]positionservice.GridPosisitionI, 0))

	if options.lastKeyPressed == "" {
		options.lastKeyPressed = "w"
	}

	zeroPos := positionservice.Position{}

	if options.Position == zeroPos {
		position := positionservice.GetRectanglePosition(newPos.X, newPos.Y, options.Height, options.Height)
		options.Position = position
	}

	rect := utils.NewRect(float32(options.Height), float32(options.Height), color.RGBA{60, 160, 255, 255})
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(options.Position.X), float64(options.Position.Y))

	return &Snake{
		Id:                uuid.NewString(),
		Sprite:            rect,
		Op:                &op,
		Position:          options.Position,
		Framework:         options.Framework,
		IsPlayerControled: true,
		lastKeyPressed:    options.lastKeyPressed,
		moveDistance:      options.Height,
		Count:             options.count,
	}
}

func (s *Snake) GetId() string {
	return s.Id
}
func (s *Snake) GetSprite() *ebiten.Image {
	return s.Sprite
}
func (s *Snake) GetOptions() *ebiten.DrawImageOptions {
	return s.Op
}
func (s *Snake) GetPosition() *positionservice.Position {
	return &s.Position
}
func (s *Snake) Update(tick int) {
	s.lastKeyPressed = lastKeyPressed(s.lastKeyPressed)

	if tick%60 == 0 {
		if s.IsPlayerControled {
			s.handleMovement(s.lastKeyPressed)
		}
		if s.Count == 0 {
			s.Framework.Registry.RemoveComponentById(s.Id, "snake")
		}
		s.Count = s.Count - 1

	}
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

func (pc *Snake) handleMovement(lastKeyPressed string) {
	var dx = int(0)
	var dy = int(0)

	var snakeBody = pc.Framework.Registry.GetComponentGroupArray("snake")
	var apple = pc.Framework.Registry.Npcs

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

	if len(collisionservice.CheckCollision2(positionservice.MovePosition(pc.Position, dx, dy), pc.Id, snakeBody)) != 0 {
		fmt.Println("Game Over x")
		panic("gg")
	}

	var food = collisionservice.CheckCollision(positionservice.MovePosition(pc.Position, dx, dy), pc.Id, apple)
	if len(food) != 0 {
		pc.Count++
	}

	var newPosition = positionservice.MovePosition(pc.Position, dx, dy)

	pc.IsPlayerControled = false

	var newBody = NewSnake(NewSnakeOptions{
		Height:         pc.Sprite.Bounds().Size().X,
		Framework:      pc.Framework,
		lastKeyPressed: pc.lastKeyPressed,
		Position:       newPosition,
		count:          pc.Count,
	})
	pc.Framework.Registry.AddComponent(newBody, "snake")
}
