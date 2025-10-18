package registry

import (
	"go-go-Framework/framework/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type Registry struct {
	Scene   *ebiten.Image
	Players []entity.Component //Component stores an abstract type, only use pointers for concrete types
	Npcs    []*ebiten.Image
}

func NewRegistry() *Registry {
	return &Registry{
		Players: make([]entity.Component, 0),
		Npcs:    make([]*ebiten.Image, 0),
	}
}

func (r *Registry) SetScene(s *ebiten.Image) {
	r.Scene = s
}

func (r *Registry) AddPlayer(p entity.Component) {
	r.Players = append(r.Players, p)
}

func (r *Registry) AddNpc(n *ebiten.Image) {
	r.Npcs = append(r.Npcs, n)
}
