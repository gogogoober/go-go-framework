package utils

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// NewCircle creates a filled circle image of radius r.
func NewCircle(r float32, fill color.Color) *ebiten.Image {
	// small padding to avoid edge clipping from AA
	const pad = float32(1)
	w := int(math.Ceil(float64(2*r + 2*pad)))
	h := w
	img := ebiten.NewImage(w, h)
	cx, cy := r+pad, r+pad
	vector.DrawFilledCircle(img, cx, cy, r, fill, true)
	return img
}

// NewRect creates a filled rectangle image of size w×h.
func NewRect(w, h float32, fill color.Color) *ebiten.Image {
	wi := int(math.Ceil(float64(w)))
	hi := int(math.Ceil(float64(h)))
	if wi < 1 {
		wi = 1
	}
	if hi < 1 {
		hi = 1
	}
	img := ebiten.NewImage(wi, hi)
	vector.DrawFilledRect(img, 0, 0, w, h, fill, true)
	return img
}
