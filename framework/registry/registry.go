package registry

import (
	"go-go-Framework/framework/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type Registry struct {
	Scene              *ebiten.Image
	components         map[string]map[string]entity.Component
	ComponentsSnapshot map[string]map[string]entity.Component
	RegistryQueues
}

type RegistryQueues struct {
	addComponent    map[string]map[string]entity.Component
	removeComponent map[string]map[string]entity.Component
	removeGroup     []string
}

func NewRegistry() *Registry {
	return &Registry{
		components:         make(map[string]map[string]entity.Component),
		ComponentsSnapshot: make(map[string]map[string]entity.Component),
		RegistryQueues: RegistryQueues{
			addComponent:    make(map[string]map[string]entity.Component),
			removeComponent: make(map[string]map[string]entity.Component),
			removeGroup:     []string{},
		},
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
	for _, group := range r.RegistryQueues.removeGroup {
		delete(r.components, group)
	}

	for groupKey, group := range r.RegistryQueues.removeComponent {
		for _, c := range group {
			if groupMap, ok := r.components[groupKey]; ok {
				delete(groupMap, c.GetId())
			}
		}
	}

	for groupKey, group := range r.RegistryQueues.addComponent {
		for _, c := range group {
			if _, ok := r.components[groupKey]; !ok {
				r.components[groupKey] = make(map[string]entity.Component)
			}
			r.components[groupKey][c.GetId()] = c
		}
	}

}

func (r *Registry) InitAddedComponents() {
	for groupKey, group := range r.RegistryQueues.addComponent {
		if _, ok := r.components[groupKey]; ok {
			for compKey, _ := range group {
				if _, ok := r.components[groupKey][compKey]; ok {
					r.components[groupKey][compKey].Init()
				}
			}
		}
	}
}

func (r *Registry) ResetQueues() {
	r.RegistryQueues = RegistryQueues{
		addComponent:    make(map[string]map[string]entity.Component),
		removeComponent: make(map[string]map[string]entity.Component),
		removeGroup:     []string{},
	}
}

func (r *Registry) SetScene(s *ebiten.Image) {
	r.Scene = s
}

func (r *Registry) AddComponent(group string, c entity.Component) {
	if _, ok := r.RegistryQueues.addComponent[group]; !ok {
		r.RegistryQueues.addComponent[group] = make(map[string]entity.Component)
	}
	r.RegistryQueues.addComponent[group][c.GetId()] = c
}

func (r *Registry) RemoveComponent(group string, c entity.Component) {
	if _, ok := r.RegistryQueues.removeComponent[group]; !ok {
		r.RegistryQueues.removeComponent[group] = make(map[string]entity.Component)
	}
	r.RegistryQueues.removeComponent[group][c.GetId()] = c
}

func (r *Registry) RemoveGroupById(groupId string) {
	r.RegistryQueues.removeGroup = append(r.RegistryQueues.removeGroup, groupId)
}

func (r *Registry) GetLiveComponents() map[string]map[string]entity.Component {
	return r.ComponentsSnapshot
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
