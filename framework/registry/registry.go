package registry

import (
	"go-go-Framework/framework/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type Registry struct {
	Scene              *ebiten.Image
	Players            []entity.Component
	Npcs               []entity.Component
	components         map[string]map[string]entity.Component2
	ComponentsSnapshot map[string]map[string]entity.Component2
}

func NewRegistry() *Registry {
	return &Registry{
		Players:            make([]entity.Component, 0),
		Npcs:               make([]entity.Component, 0),
		components:         make(map[string]map[string]entity.Component2),
		ComponentsSnapshot: make(map[string]map[string]entity.Component2),
	}
}

func (r *Registry) UpdateSnapshot() {
	snap := make(map[string]map[string]entity.Component2, len(r.components))
	for group, comps := range r.components {
		groupCopy := make(map[string]entity.Component2, len(comps))
		for id, comp := range comps {
			groupCopy[id] = comp
		}
		snap[group] = groupCopy
	}
	r.ComponentsSnapshot = snap
}

func (r *Registry) SetScene(s *ebiten.Image) {
	r.Scene = s
}

func (r *Registry) AddComponent(c entity.Component2, group string) {
	if _, ok := r.components[group]; !ok {
		r.components[group] = make(map[string]entity.Component2)
	}
	r.components[group][c.GetId()] = c
}

func (r *Registry) RemoveComponentById(id string, group string) {
	if groupMap, ok := r.components[group]; ok {
		delete(groupMap, id)
	}
}

func (r *Registry) GetComponents() map[string]map[string]entity.Component2 {
	return r.ComponentsSnapshot
}

func (r *Registry) GetComponentGroup(group string) (map[string]entity.Component2, bool) {
	if g, ok := r.ComponentsSnapshot[group]; ok {
		return g, true
	}
	return make(map[string]entity.Component2), false
}

func (r *Registry) GetComponentGroupArray(group string) []entity.Component2 {
	var collection []entity.Component2
	if g, ok := r.ComponentsSnapshot[group]; ok {
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
