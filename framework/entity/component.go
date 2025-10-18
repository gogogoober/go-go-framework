package entity

import "github.com/hajimehoshi/ebiten/v2"

type Component interface {
	SetSprite(sprite *ebiten.Image)
	GetSprite() *ebiten.Image
	GetOptions() *ebiten.DrawImageOptions
	Update()
}
