package entity

import (
	"go-go-Framework/framework/services/positionservice"

	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	SetSprite(sprite *ebiten.Image)
	GetSprite() *ebiten.Image
	GetOptions() *ebiten.DrawImageOptions
	GetPosition() *positionservice.Position
	SetNpcs([]Component)
	Update(tick int)
}
