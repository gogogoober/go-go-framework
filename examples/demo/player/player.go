package player

import (
	"fmt"
	"go-go-Framework/framework/registry"
	"go-go-Framework/framework/services/inputservice"
	"go-go-Framework/framework/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerComponent struct {
	registry.Component
	sprites []*ebiten.Image
}

func NewPlayerComponent() *PlayerComponent {
	rect := utils.NewRect(120, 80, color.RGBA{60, 160, 255, 255})
	return &PlayerComponent{
		sprites: []*ebiten.Image{rect},
	}
}

func (pc *PlayerComponent) GetSprite() *ebiten.Image {
	return pc.sprites[0]
}

func (pc *PlayerComponent) SetSprite(sprite *ebiten.Image) {

}

func (pc *PlayerComponent) Update() {
	if inputservice.IsKeyStringPressed("a") {
		fmt.Println("A")
	}
}
