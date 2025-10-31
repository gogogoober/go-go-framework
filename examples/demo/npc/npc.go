package npc

import (
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/positionservice"
	"go-go-Framework/framework/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type NpcComponent struct {
	sprite *ebiten.Image
	op     ebiten.DrawImageOptions
	npcs   []entity.Component
	entity.Component
	positionservice.Position
}

func NewNpcComponent() *NpcComponent {
	rect := utils.NewRect(120, 80, color.RGBA{60, 225, 255, 255})
	var w = rect.Bounds()

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(100, 100)

	return &NpcComponent{
		sprite:   rect,
		op:       op,
		npcs:     make([]entity.Component, 0),
		Position: positionservice.GetRectanglePosition(100, 100, w.Size().X, w.Size().Y),
	}
}

func (np *NpcComponent) SetSprite(sprite *ebiten.Image) {
	np.sprite = sprite
}
func (np *NpcComponent) GetSprite() *ebiten.Image {
	return np.sprite
}
func (np *NpcComponent) GetOptions() *ebiten.DrawImageOptions {
	return &np.op
}
func (np *NpcComponent) GetPosition() *positionservice.Position {
	return &np.Position
}
func (np *NpcComponent) SetNpcs(npc []entity.Component) {
	np.npcs = npc
}
func (np *NpcComponent) Update(tick int) {}
