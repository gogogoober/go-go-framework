package collisionservice

import (
	"go-go-Framework/framework/entity"
	"go-go-Framework/framework/services/positionservice"
)

func CheckCollision(pcPosition positionservice.Position, id string, components []entity.Component) []entity.Component {
	var collisions []entity.Component

	for _, n := range components {
		if AreComponentsColliding(pcPosition, *n.GetPosition()) {
			collisions = append(collisions, n)
		}
	}
	return collisions
}

func CheckCollision2(pcPosition positionservice.Position, id string, components []entity.Component2) []entity.Component2 {
	var collisions []entity.Component2

	for _, n := range components {
		if AreComponentsColliding(pcPosition, *n.GetPosition()) {
			collisions = append(collisions, n)
		}
	}
	return collisions
}

func AreComponentsColliding(p1, p2 positionservice.Position) bool {
	left1, right1 := p1.X, p1.X2
	top1, bottom1 := p1.Y, p1.Y2
	left2, right2 := p2.X, p2.X2
	top2, bottom2 := p2.Y, p2.Y2

	return left1 < right2 && right1 > left2 &&
		top1 < bottom2 && bottom1 > top2
}

func AreComponentsTouching(c1 entity.Component, c2 entity.Component) bool {
	p1 := c1.GetPosition()
	p2 := c2.GetPosition()

	left1, right1 := p1.X, p1.X2
	top1, bottom1 := p1.Y, p1.Y2
	left2, right2 := p2.X, p2.X2
	top2, bottom2 := p2.Y, p2.Y2

	horizontalTouch := right1 == left2 || right2 == left1
	verticalTouch := bottom1 == top2 || bottom2 == top1
	horizontalOverlap := bottom1 > top2 && top1 < bottom2
	verticalOverlap := right1 > left2 && left1 < right2

	return (horizontalTouch && horizontalOverlap) ||
		(verticalTouch && verticalOverlap)
}
