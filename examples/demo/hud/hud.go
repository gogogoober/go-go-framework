package hud

import (
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/positionservice"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Hud struct {
	sprite *ebiten.Image
	op     *ebiten.DrawImageOptions
	positionservice.Position
	entity.Component
}

func NewHud() *Hud {
	return &Hud{}
}

func (h *Hud) GetId() string {
	return ""
}
func (h *Hud) GetSprite(screen *ebiten.Image) *ebiten.Image {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	return h.sprite
}
func (h *Hud) GetOptions() *ebiten.DrawImageOptions {
	return h.op
}
func (h *Hud) GetPosition() *positionservice.Position {
	return &h.Position
}
func (h *Hud) Update(tick int) {}
