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
	entity.Component
	Framework *framework.GoGoFramework

	IsPlayerControled    bool
	previouseKeyPresssed string
	moveDistance         int
	Count                int
	Height               int
}

type NewSnakeOptions struct {
	Height               int
	Framework            *framework.GoGoFramework
	previouseKeyPresssed string
	Count                int
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
		Count:                options.Count,
		Height:               options.Height,
	}
}

func (s *Snake) GetId() string {
	return s.Id
}
func (s *Snake) GetSprite(screen *ebiten.Image) *ebiten.Image {
	return s.Sprite
}
func (s *Snake) GetOptions() *ebiten.DrawImageOptions {
	return s.Op
}
func (s *Snake) GetPosition() *positionservice.Position {
	return &s.Position
}
func (s *Snake) Update(tick int) {
	s.previouseKeyPresssed = getKeyJustPressed(s.previouseKeyPresssed)

	if s.previouseKeyPresssed == "" {
		return
	}

	if tick%8 == 0 {

		if s.IsPlayerControled {
			s.handleMovement(s.previouseKeyPresssed)
		}
		if s.Count == 0 {
			s.Framework.Registry.RemoveComponentById(s.Id, "snake")
		}
		s.Count = s.Count - 1
	}
}

func getKeyJustPressed(previouseKeyPresssed string) string {
	if inputservice.IsKeyStringJustPressed("d") && previouseKeyPresssed != "a" {
		return "d"
	}
	if inputservice.IsKeyStringJustPressed("a") && previouseKeyPresssed != "d" {
		return "a"
	}
	if inputservice.IsKeyStringJustPressed("w") && previouseKeyPresssed != "s" {
		return "w"
	}
	if inputservice.IsKeyStringJustPressed("s") && previouseKeyPresssed != "w" {
		return "s"
	}

	return previouseKeyPresssed
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

func (s *Snake) handleMovement(keyPressed string) {
	var dx = int(0)
	var dy = int(0)

	var snakeBody = s.Framework.Registry.GetComponentGroupArray("snake")
	var apple = s.Framework.Registry.GetComponentGroupArray("apple")

	if keyPressed == "d" {
		dx += s.moveDistance
	}
	if keyPressed == "a" {
		dx -= s.moveDistance
	}

	if keyPressed == "w" {
		dy -= s.moveDistance
	}
	if keyPressed == "s" {
		dy += s.moveDistance
	}

	var newPosition = positionservice.MovePosition(s.Position, dx, dy)
	var isCollidingWithSelf = len(collisionservice.CheckCollision(newPosition, snakeBody)) != 0
	var isOutOfBounds = newPosition.X < 0 || newPosition.Y < 0 || newPosition.X2 > s.Framework.Options.Window.Height || newPosition.Y2 > s.Framework.Options.Window.Width

	if isOutOfBounds || isCollidingWithSelf {
		fmt.Println("Game Over x")
		reset(s.Framework, s.Height)
		return
	}

	if len(collisionservice.CheckCollision(newPosition, apple)) != 0 {
		s.Count++
	}

	s.IsPlayerControled = false

	var newBody = NewSnake(NewSnakeOptions{
		Height:               s.Sprite.Bounds().Size().X,
		Framework:            s.Framework,
		previouseKeyPresssed: keyPressed,
		Position:             newPosition,
		Count:                s.Count,
	})
	s.Framework.Registry.AddComponent(newBody, "snake")

}

func reset(framework *framework.GoGoFramework, height int) {
	framework.Registry.RemoveGroupById("snake")

	var newBody = NewSnake(NewSnakeOptions{
		Height:    height,
		Framework: framework,
	})
	framework.Registry.AddComponent(newBody, "snake")
}
