package component

import "github.com/hajimehoshi/ebiten/v2"

type Component struct {
	Sprite *ebiten.Image
	X      int
	Y      int
}
