package entity

import "github.com/hajimehoshi/ebiten/v2"

type Position struct {
	X     int
	Y     int
	Width int
	Hight int
}

type Component interface {
	SetSprite(sprite *ebiten.Image)
	GetSprite() *ebiten.Image
	GetOptions() *ebiten.DrawImageOptions
	GetPosition() *Position
	SetNpcs([]Component)
	Update()
}
