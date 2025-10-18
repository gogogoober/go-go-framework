package registry

import "github.com/hajimehoshi/ebiten/v2"

type Registry struct {
	Scene   *ebiten.Image
	Players []Component //Component stores an abstract type, only use pointers for concrete types
	Npcs    []*ebiten.Image
}

type Component interface {
	SetSprite(sprite *ebiten.Image)
	GetSprite() *ebiten.Image
	GetOptions() *ebiten.DrawImageOptions
	Update()
}

func NewRegistry() *Registry {
	return &Registry{
		Players: make([]Component, 0),
		Npcs:    make([]*ebiten.Image, 0),
	}
}

func (r *Registry) SetScene(s *ebiten.Image) {
	r.Scene = s
}

func (r *Registry) AddPlayer(p Component) {
	r.Players = append(r.Players, p)
}

func (r *Registry) AddNpc(n *ebiten.Image) {
	r.Npcs = append(r.Npcs, n)
}
