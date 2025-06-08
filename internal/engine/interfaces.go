package engine

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type World interface {
	GetWidth() int
	GetHeight() int
}
