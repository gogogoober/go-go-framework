package entity

import (
	"go-go-Framework/framework/services/positionservice"

	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	GetId() string
	GetSprite() *ebiten.Image
	GetOptions() *ebiten.DrawImageOptions
	GetPosition() *positionservice.Position
	Update(tick int)
}
