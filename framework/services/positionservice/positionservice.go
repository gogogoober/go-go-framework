package positionservice

import (
	"go-go-Framework/framework/window"
	"math/rand"
)

type Position struct {
	X  int
	Y  int
	X2 int
	Y2 int
}

type GridPosisition struct {
	X int
	Y int
}

type GridPosisitionI interface {
	GetPosition() Position
}

func GetRectanglePosition(x, y, width, height int) Position {
	return Position{x, y, x + width, y + height}
}

func MovePosition(position Position, dx int, dy int) Position {
	x := position.X + dx
	x2 := position.X2 + dx
	y := position.Y + dy
	y2 := position.Y2 + dy

	return Position{x, y, x2, y2}
}

func GetPositions(window window.GoGoWindow, gridSize window.GoGoWindow) []GridPosisition {
	gh := window.Height
	gw := window.Width
	ggh := gridSize.Height
	ggw := gridSize.Width

	var hg = (gh / ggh)
	// var wg = (gw / ggw)

	var positions = make([]GridPosisition, hg)
	var count = 0

	for i := 0; i < gh; i = i + ggh {
		positions[count].Y = i
		count++
	}
	count = 0
	for i := 0; i < gw; i = i + ggw {
		positions[count].X = i
		count++

	}
	return positions
}

func GetNewPosition(pp []GridPosisition, players []GridPosisitionI) GridPosisition {
	var newPosition []GridPosisition

	if len(players) == 0 {
		newPosition = pp
	} else {
		for i := range players {
			for j := range pp {
				var pPos = players[i].GetPosition()
				if pp[j].X != pPos.X && pp[j].Y != pPos.Y {
					newPosition = append(newPosition, pp[j])

				}
			}
		}

	}

	var posPick = rand.Intn(len(newPosition))
	var _np = newPosition[posPick]
	return _np
}
