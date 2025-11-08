package apple

import (
	"go-go-Framework/framework"
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/positionservice"

	"github.com/hajimehoshi/ebiten/v2"
)

type Apple struct {
	Id                 string
	sprite             *ebiten.Image
	op                 ebiten.DrawImageOptions
	possiblePossitions []GridPosisition
	entity.Component
	positionservice.Position
	Framework *framework.GoGoFramework
}

type GridPosisition struct {
	x int
	y int
}
