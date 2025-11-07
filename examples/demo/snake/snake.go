package snake

import (
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
	entity.Component
	Framework *framework.GoGoFramework

	IsPlayerControled    bool
	previouseKeyPresssed string
	moveDistance         int
	Count                int
}

type NewSnakeOptions struct {
	Height               int
	Framework            *framework.GoGoFramework
	previouseKeyPresssed string
	count                int
	positionservice.Position
}

func NewSnake(options NewSnakeOptions) *Snake {
	var pp = positionservice.GetPositions(*options.Framework.Options.Window, *options.Framework.Options.GridSize)
	var newPos = positionservice.GetNewPosition(pp, make([]positionservice.GridPosisitionI, 0))

	zeroPos := positionservice.Position{}

	if options.Position == zeroPos {
		position := positionservice.GetRectanglePosition(newPos.X, newPos.Y, options.Height, options.Height)
		options.Position = position
	}

	rect := utils.NewRect(float32(options.Height), float32(options.Height), color.RGBA{60, 160, 255, 255})
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(options.Position.X), float64(options.Position.Y))

	return &Snake{
		Id:                   uuid.NewString(),
		Sprite:               rect,
		Op:                   &op,
		Position:             options.Position,
		Framework:            options.Framework,
		IsPlayerControled:    true,
		previouseKeyPresssed: options.previouseKeyPresssed,
		moveDistance:         options.Height,
		Count:                options.count,
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
	var keyPressed = getKeyPressed(s.previouseKeyPresssed)

	if keyPressed == "" {
		return
	}

	if tick%10 == 0 {
		if s.IsPlayerControled {
			s.handleMovement(keyPressed)
		}
		if s.Count == 0 {
			s.Framework.Registry.RemoveComponentById(s.Id, "snake")
		}
		s.Count = s.Count - 1
	}
}

func getKeyPressed(previouseKeyPresssed string) string {
	if inputservice.IsKeyStringPressed("d") && previouseKeyPresssed != "a" {
		return "d"
	}
	if inputservice.IsKeyStringPressed("a") && previouseKeyPresssed != "d" {
		return "a"
	}
	if inputservice.IsKeyStringPressed("w") && previouseKeyPresssed != "s" {
		return "w"
	}
	if inputservice.IsKeyStringPressed("s") && previouseKeyPresssed != "w" {
		return "s"
	}

	return previouseKeyPresssed
}

func (pc *Snake) handleMovement(keyPressed string) {
	var dx = int(0)
	var dy = int(0)

	var snakeBody = pc.Framework.Registry.GetComponentGroupArray("snake")
	var apple = pc.Framework.Registry.GetComponentGroupArray("apple")

	if keyPressed == "d" {
		dx += pc.moveDistance
	}
	if keyPressed == "a" {
		dx -= pc.moveDistance
	}

	if keyPressed == "w" {
		dy -= pc.moveDistance
	}
	if keyPressed == "s" {
		dy += pc.moveDistance
	}

	var newPosition = positionservice.MovePosition(pc.Position, dx, dy)

	if len(collisionservice.CheckCollision(newPosition, pc.Id, snakeBody)) != 0 {
		panic("Game Over x")
	}

	if len(collisionservice.CheckCollision(newPosition, pc.Id, apple)) != 0 {
		pc.Count++
	}

	pc.IsPlayerControled = false

	var newBody = NewSnake(NewSnakeOptions{
		Height:               pc.Sprite.Bounds().Size().X,
		Framework:            pc.Framework,
		previouseKeyPresssed: keyPressed,
		Position:             newPosition,
		count:                pc.Count,
	})
	pc.Framework.Registry.AddComponent(newBody, "snake")
}
