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

func MovePosition(position Position, dx int, dy int) Position {
	x := position.X + dx
	x2 := position.X2 + dx
	y := position.Y + dy
	y2 := position.Y2 + dy

	return Position{x, y, x2, y2}
}
