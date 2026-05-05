package entity

import (
	"fmt"
	"go-go-Framework/ecs/component"
)

func addComponent(component component.Component) bool {

	fmt.Printf(component.GetName())
	return true
}
