package registry

import (
	"go-go-Framework/framework/entity"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Registry struct {
	Scene   *ebiten.Image
	Players []entity.Component //Component stores an abstract type, only use pointers for concrete types
	Npcs    []entity.Component
}

func NewRegistry() *Registry {
	return &Registry{
		Players: make([]entity.Component, 0),
		Npcs:    make([]entity.Component, 0),
	}
}

func (r *Registry) SetScene(s *ebiten.Image) {
	r.Scene = s
}

func (r *Registry) AddPlayer(p entity.Component) {
	r.Players = append(r.Players, p)
}

func (r *Registry) RemovePlayerById(id uuid.UUID) {
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
