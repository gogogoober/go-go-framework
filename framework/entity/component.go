package entity

import (
	"go-go-Framework/framework/services/positionservice"

	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	GetId() string
	SetContact()
	SetSprite(sprite *ebiten.Image)
	GetSprite() *ebiten.Image
	GetOptions() *ebiten.DrawImageOptions
	GetPosition() *positionservice.Position
	SetNpcs([]Component)
	SetPlayers([]Component)
	Update(tick int)
}

type Component2 interface {
	GetId() string
	GetSprite() *ebiten.Image
	GetOptions() *ebiten.DrawImageOptions
	GetPosition() *positionservice.Position
	Update(tick int)
}
