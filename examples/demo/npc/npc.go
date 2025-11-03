package npc

import (
	"go-go-Framework/framework"
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/positionservice"
	"go-go-Framework/framework/utils"
	"image/color"
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type NpcComponent struct {
	Id                 string
	sprite             *ebiten.Image
	op                 ebiten.DrawImageOptions
	npcs               []entity.Component
	players            []entity.Component
	possiblePossitions []GridPosisition
	contact            bool
	entity.Component
	positionservice.Position
	Framework *framework.GoGoFramework
}

type NewNPCOptions struct {
	Width     int
	Height    int
	Framework *framework.GoGoFramework
}

type GridPosisition struct {
	x int
	y int
}

func NewNpcComponent(options NewNPCOptions) *NpcComponent {
	rect := utils.NewRect(float32(options.Width), float32(options.Height), color.RGBA{225, 0, 0, 1})
	op := ebiten.DrawImageOptions{}

	var pp = getPositions(options.Framework)
	var newPos = getNewPosition(pp, make([]entity.Component, 0))
	op.GeoM.Translate(float64(newPos.x), float64(newPos.y))

	return &NpcComponent{
		Id:                 uuid.NewString(),
		sprite:             rect,
		op:                 op,
		possiblePossitions: pp,
		contact:            false,
		npcs:               make([]entity.Component, 0),
		Position:           positionservice.GetRectanglePosition(newPos.x, newPos.y, rect.Bounds().Size().X, rect.Bounds().Size().Y),
		Framework:          options.Framework,
	}
}

func (np *NpcComponent) SetContact() {
	np.contact = true
}

func (np *NpcComponent) SetSprite(sprite *ebiten.Image) {

}

func (np *NpcComponent) GetId() string {
	return np.Id
}

func (np *NpcComponent) GetSprite() *ebiten.Image {
	return np.sprite
}
func (np *NpcComponent) GetOptions() *ebiten.DrawImageOptions {
	return &np.op
}
func (np *NpcComponent) GetPosition() *positionservice.Position {
	return &np.Position
}
func (np *NpcComponent) SetNpcs(npc []entity.Component) {
	np.npcs = npc
}
func (np *NpcComponent) SetPlayers(players []entity.Component) {
	np.players = players
}
func (np *NpcComponent) Update(tick int) {
	if tick%60 == 0 && np.contact {
		var newPos = getNewPosition(np.possiblePossitions, np.players)
		position := positionservice.GetRectanglePosition(newPos.x, newPos.y, np.sprite.Bounds().Size().X, np.sprite.Bounds().Size().Y)

		np.Position = position
		np.op.GeoM.Reset() // important
		np.op.GeoM.Translate(float64(position.X), float64(position.Y))
		np.contact = false
	}
}

func getPositions(Framework *framework.GoGoFramework) []GridPosisition {
	gh := Framework.Options.Window.Height
	gw := Framework.Options.Window.Width
	ggh := Framework.Options.GridSize.Height
	ggw := Framework.Options.GridSize.Width

	var hg = (gh / ggh)
	// var wg = (gw / ggw)

	var positions = make([]GridPosisition, hg)
	var count = 0

	for i := 0; i < gh; i = i + ggh {
		positions[count].y = i
		count++
	}
	count = 0
	for i := 0; i < gw; i = i + ggw {
		positions[count].x = i
		count++

	}
	return positions
}

func getNewPosition(pp []GridPosisition, players []entity.Component) GridPosisition {
	var newPosition []GridPosisition

	if len(players) == 0 {
		newPosition = pp
	} else {
		for i := range players {
			for j := range pp {
				var pPos = players[i].GetPosition()
				if pp[j].x != pPos.X && pp[j].y != pPos.Y {
					newPosition = append(newPosition, pp[j])

				}
			}
		}

	}

	var posPick = rand.Intn(len(newPosition))
	var _np = newPosition[posPick]
	return _np
}
