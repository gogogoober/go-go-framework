package registry

import (
	"go-go-Framework/framework/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type Registry struct {
	Scene              *ebiten.Image
	Players            []entity.Component
	Npcs               []entity.Component
	components         map[string]map[string]entity.Component
	ComponentsSnapshot map[string]map[string]entity.Component
}

func NewRegistry() *Registry {
	return &Registry{
		Players:            make([]entity.Component, 0),
		Npcs:               make([]entity.Component, 0),
		components:         make(map[string]map[string]entity.Component),
		ComponentsSnapshot: make(map[string]map[string]entity.Component),
	}
}

func (r *Registry) UpdateSnapshot() {
	snap := make(map[string]map[string]entity.Component, len(r.components))
	for group, comps := range r.components {
		groupCopy := make(map[string]entity.Component, len(comps))
		for id, comp := range comps {
			groupCopy[id] = comp
		}
		snap[group] = groupCopy
	}
	r.ComponentsSnapshot = snap
}

func (r *Registry) UpdateComponents() {
	// for groupKey, group := range r.components {
	// 	for compKey, _ := range group {
	// 		group[compKey] = r.ComponentsSnapshot[groupKey][compKey]
	// 	}
	// }

	// snap := make(map[string]map[string]entity.Component, len(r.ComponentsSnapshot))
	// for group, comps := range r.ComponentsSnapshot {
	// 	groupCopy := make(map[string]entity.Component, len(comps))
	// 	for id, comp := range comps {
	// 		groupCopy[id] = comp
	// 	}
	// 	snap[group] = groupCopy
	// }
	// r.components = snap
}

func (r *Registry) SetScene(s *ebiten.Image) {
	r.Scene = s
}

func (r *Registry) AddComponent(c entity.Component, group string) {
	if _, ok := r.components[group]; !ok {
		r.components[group] = make(map[string]entity.Component)
	}
	r.components[group][c.GetId()] = c
}

func (r *Registry) RemoveComponentById(id string, group string) {
	if groupMap, ok := r.components[group]; ok {
		delete(groupMap, id)
	}
}

func (r *Registry) RemoveGroupById(group string) {
	delete(r.components, group)
}

func (r *Registry) GetComponents() map[string]map[string]entity.Component {
	return r.ComponentsSnapshot
}

func (r *Registry) GetComponentGroup(group string) (map[string]entity.Component, bool) {
	if g, ok := r.ComponentsSnapshot[group]; ok {
		return g, true
	}
	return make(map[string]entity.Component), false
}

func (r *Registry) GetComponentGroupArray(group string) []entity.Component {
	var collection []entity.Component
	if g, ok := r.ComponentsSnapshot[group]; ok {
		for _, c := range g {
			collection = append(collection, c)
		}
	}
	return collection
}
