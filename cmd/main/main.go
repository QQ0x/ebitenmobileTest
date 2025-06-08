//go:build !mobile
// Non-mobile platforms only

package main

import (
	"example.com/go_astroids/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Astroids")

	if err := ebiten.RunGame(game.New()); err != nil {
		panic(err)
	}
}
