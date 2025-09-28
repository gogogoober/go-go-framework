package gogoFramework

import (
	"go-go-framework/src/scheduler"

	"github.com/hajimehoshi/ebiten/v2"
)

type GoGoFramework struct {
	GameName string
	Scene    *ebiten.Image
	npc      []*ebiten.Image
	player   []*ebiten.Image
	Window   *GoGoWindow
}

type Component struct {
	Sprite *ebiten.Image
	X      int
	Y      int
}

type GoGoWindow struct {
	Width  int
	Height int
}

func (g *GoGoFramework) RegisterScene(scene *ebiten.Image) {
	g.Scene = scene
}

func (g *GoGoFramework) AddNpc(npc *ebiten.Image) {
	g.npc = append(g.npc, npc)
}

func NewGoGoFrameworkWithDefaults(gameName string) *GoGoFramework {

	// demoScene := ebiten.NewImage(50, 50)
	// demoScene.Fill(color.RGBA{0xff, 0, 0, 0xff})

	return &GoGoFramework{
		GameName: gameName,
		Window: &GoGoWindow{
			Width:  620,
			Height: 620,
		},
	}
}

func (g *GoGoFramework) Update() error {
	return nil
}

func (g *GoGoFramework) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	var images = scheduler.NewScheduler(g.Scene, g.npc, g.player)
	for _, image := range images.GetItemsToDraw() {
		screen.DrawImage(image, op)
	}
}

func (g *GoGoFramework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Width, g.Window.Height
}
