package game

import (
	"example.com/go_astroids/packages/player"
	"github.com/hajimehoshi/ebiten/v2"
)

// Main game struct
type Game struct {
	player *player.Player
	width  int
	height int
}

func (g *Game) GetWidth() int {
	return g.width
}

func (g *Game) GetHeight() int {
	return g.height
}

func New() *Game {
	g := &Game{
		width:  640,
		height: 480,
	}
	g.player = player.NewPlayer(g)
	return g
}

func (g *Game) Update() error {
	return g.player.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Update game dimensions
	g.width = outsideWidth
	g.height = outsideHeight

	return outsideWidth, outsideHeight
}
