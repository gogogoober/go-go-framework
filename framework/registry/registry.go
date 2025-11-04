package registry

import (
	"fmt"
	"go-go-Framework/framework/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type Registry struct {
	Scene      *ebiten.Image
	Players    []entity.Component //Component stores an abstract type, only use pointers for concrete types
	Npcs       []entity.Component
	Components map[string]map[string]entity.Component2
}

func NewRegistry() *Registry {
	return &Registry{
		Players:    make([]entity.Component, 0),
		Npcs:       make([]entity.Component, 0),
		Components: make(map[string]map[string]entity.Component2),
	}
}

func (r *Registry) SetScene(s *ebiten.Image) {
	r.Scene = s
}

func (r *Registry) AddComponent(c entity.Component2, group string) {
	fmt.Println("AddComponent")
	if _, ok := r.Components[group]; !ok {
		r.Components[group] = make(map[string]entity.Component2)
	}
	r.Components[group][c.GetId()] = c
}

func (r *Registry) RemoveComponentById(id string, group string) {
	if groupMap, ok := r.Components[group]; ok {
		delete(groupMap, id)
	}
}

func (r *Registry) GetGroupArray(group string) []entity.Component2 {
	var collection []entity.Component2
	if g, ok := r.Components[group]; ok {
		for _, c := range g {
			collection = append(collection, c)
		}
	}
	return collection
}

func (r *Registry) AddPlayer(p entity.Component) {
	r.Players = append(r.Players, p)
}

func (r *Registry) RemovePlayerById(id string) {
	for i := range r.Players {
		if r.Players[i].GetId() == id {
			r.Players = append(r.Players[:i], r.Players[i+1:]...)
			return
		}
	}

}

func (r *Registry) AddNpc(n entity.Component) {
	r.Npcs = append(r.Npcs, n)
}
