package scheduler

import "github.com/hajimehoshi/ebiten/v2"

type GoGoScheduler struct {
	Scene  *ebiten.Image
	npc    []*ebiten.Image
	player []*ebiten.Image
}

func NewScheduler(scene *ebiten.Image, npc []*ebiten.Image, player []*ebiten.Image) *GoGoScheduler {
	return &GoGoScheduler{
		Scene:  scene,
		npc:    npc,
		player: player,
	}
}

func (gs *GoGoScheduler) GetItemsToDraw() []*ebiten.Image {
	var items = make([]*ebiten.Image, 0)

	items = append(items, gs.Scene)

	items = append(items, gs.player...)

	for _, n := range gs.npc {
		items = append(items, n)
	}

	return items
}
