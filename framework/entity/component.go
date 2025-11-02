package entity

import (
	"go-go-Framework/framework/services/positionservice"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	GetId() uuid.UUID
	SetSprite(sprite *ebiten.Image)
	GetSprite() *ebiten.Image
	GetOptions() *ebiten.DrawImageOptions
	GetPosition() *positionservice.Position
	SetNpcs([]Component)
	SetPlayers([]Component)
	Update(tick int)
}
