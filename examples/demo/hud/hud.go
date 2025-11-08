package hud

import (
	"fmt"
	"go-go-Framework/framework"
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
	framework *framework.GoGoFramework
}

func NewHud(framework *framework.GoGoFramework) *Hud {
	return &Hud{
		framework: framework,
	}
}

func (h *Hud) GetId() string {
	return ""
}
func (h *Hud) GetSprite(screen *ebiten.Image) *ebiten.Image {
	s := fmt.Sprintf("Time: %s", h.framework.Services.GetPrettySeconds())

	ebitenutil.DebugPrint(screen, s)
	return h.sprite
}
func (h *Hud) GetOptions() *ebiten.DrawImageOptions {
	return h.op
}
func (h *Hud) GetPosition() *positionservice.Position {
	return &h.Position
}
func (h *Hud) Update(tick int) {}
