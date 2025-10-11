package registry

import "github.com/hajimehoshi/ebiten/v2"

type Registry struct {
	Scene   *ebiten.Image
	Players []*ebiten.Image
	Npcs    []*ebiten.Image
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) SetScene(s *ebiten.Image) {
	r.Scene = s
}

func (r *Registry) AddPlayer(p *ebiten.Image) {
	r.Players = append(r.Players, p)
}

func (r *Registry) AddNpc(n *ebiten.Image) {
	r.Npcs = append(r.Npcs, n)
}
