package collisionservice

import (
	"go-go-Framework/framework/entity"
)

func AreComponentsColliding(c1X entity.Component, c2X entity.Component) bool {
	// If the x + W && y + h overlap, its a collision

	c1 := c1X.GetPosition()
	c2 := c2X.GetPosition()

	if c1.X <= c2.X && c2.X <= c1.X+c1.Width {
		if c1.Y <= c2.Y && c2.Y <= (c1.Y+c1.Hight) {
			// fmt.Print("Collide")
			return true
		}
	}

	if c2.X <= c1.X && c1.X <= c2.X+c2.Width {
		if c2.Y <= c1.Y && c1.Y <= (c2.Y+c2.Hight) {
			// fmt.Print("Collide")
			return true
		}
	}

	if c1.X <= c2.X+c2.Width && c2.X+c2.Width <= c1.X+c1.Width {
		if c1.Y <= c2.Y && c2.Y <= (c1.Y+c1.Hight) {
			// fmt.Print("Collide")
			return true
		}
	}

	if c1.X <= c2.X && c2.X <= c1.X+c1.Width {
		if c1.Y <= c2.Y+c2.Hight && c2.Y+c2.Hight <= c1.Y+c1.Hight {
			// fmt.Print("Collide")
			return true
		}
	}

	return false
}
