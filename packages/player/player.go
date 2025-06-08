package player

import (
	"example.com/go_astroids/assets"
	"example.com/go_astroids/internal/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const rotationPerSecond = math.Pi

type Player struct {
	sprite   *ebiten.Image
	world    engine.World
	rotation float64
}

// Create player instance
func NewPlayer(world engine.World) *Player {
	sprite := assets.PlayerSprite

	p := &Player{
		sprite: sprite,
		world:  world,
	}

	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Calculate sprite center
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	// Move to origin
	op.GeoM.Translate(-halfW, -halfH)
	// Rotate around center
	op.GeoM.Rotate(p.rotation)
	// Move to world center
	centerX := float64(p.world.GetWidth()) / 2
	centerY := float64(p.world.GetHeight()) / 2
	op.GeoM.Translate(centerX, centerY)

	// Draw sprite
	screen.DrawImage(p.sprite, op)
}

func (p *Player) Update() error {
	// Handle player movement

	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}

	return nil
}
