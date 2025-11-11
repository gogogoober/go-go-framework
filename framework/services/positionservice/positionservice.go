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

type PositionGetter interface {
	GetPosition() *Position
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

	rows := gh / ggh
	cols := gw / ggw
	if rows <= 0 || cols <= 0 {
		return nil
	}

	positions := make([]GridPosisition, 0, rows*cols)
	for y := 0; y < gh; y += ggh {
		for x := 0; x < gw; x += ggw {
			positions = append(positions, GridPosisition{X: x, Y: y})
		}
	}
	return positions
}

func GetNewPosition[T PositionGetter](pp []GridPosisition, players []T) GridPosisition {
	blocked := make(map[GridPosisition]struct{}, len(players))
	available := make([]GridPosisition, 0, len(pp))

	for _, p := range players {
		pos := p.GetPosition()
		blocked[GridPosisition{X: pos.X, Y: pos.Y}] = struct{}{}
	}

	for _, cell := range pp {
		if _, used := blocked[cell]; !used {
			available = append(available, cell)
		}
	}

	if len(available) == 0 {
		return GridPosisition{}
	}
	return available[rand.Intn(len(available))]
}
