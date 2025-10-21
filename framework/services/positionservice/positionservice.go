package positionservice

type Position struct {
	X  int
	Y  int
	X2 int
	Y2 int
}

func GetRectanglePosition(x, y, width, height int) Position {
	return Position{x, y, x + width, y + height}
}
