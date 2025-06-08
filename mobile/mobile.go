//go:build ios || android
// Mobile platforms only

package mobile

import (
	"example.com/go_astroids/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/hajimehoshi/ebiten/v2/mobile/ebitenmobileview"
)

// Init mobile view
func init() {
	// Ensure proper initialization

	// Keep package reference
	_ = ebitenmobileview.DeviceScale
}

// Mobile entry point
func Start() {
	// Create game instance
	g := game.New()

	// Set mobile game
	mobile.SetGame(g)

	// Setup iOS rendering
	ebitenmobileview.SetGame(g, nil)

	// Force initial render
	ebitenmobileview.Update()
}

// Mobile binding interface
type GameController interface {
	Update() error
}

// Create controller
func NewGameController() GameController {
	return game.New()
}

// Ebiten game alias
type Game = ebiten.Game

// Create game instance
func NewEbitenGame() Game {
	return game.New()
}

// Required export function
func DummyFunc() {
	// For mobile tool export
}

// Additional export function
func AnotherDummyFunc() int {
	return 0
}

// Exported view functions
var (
	// Renderer setup function
	SetRenderer = ebitenmobileview.SetRenderer
	// Layout function
	Layout = ebitenmobileview.Layout
	// Update function
	Update = ebitenmobileview.Update
	// Lifecycle functions
	Suspend     = ebitenmobileview.Suspend
	Resume      = ebitenmobileview.Resume
	DeviceScale = ebitenmobileview.DeviceScale
	// Context handling
	OnContextLost = ebitenmobileview.OnContextLost
	UsesStrictContextRestoration = ebitenmobileview.UsesStrictContextRestoration
	// iOS rendering setup
	SetGameEbitenMobileView = ebitenmobileview.SetGame
	SetSetGameNotifier = ebitenmobileview.SetSetGameNotifier
)

// Mobile package exports
var (
	// Game setup function
	SetGame = mobile.SetGame
)

// Game package exports
var (
	// Game creation function
	NewGame = game.New
)

// Ebiten package exports
var (
	// Window size control
	SetWindowSize = ebiten.SetWindowSize
	// Window title control
	SetWindowTitle = ebiten.SetWindowTitle
	// Game runner function
	RunGame = ebiten.RunGame
)
