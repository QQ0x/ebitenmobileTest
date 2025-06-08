//go:build ios || android
// +build ios android

package mobile

// Import mobile packages with blank imports to ensure they're included in the build
import (
	_ "github.com/hajimehoshi/ebiten/v2/mobile"
	_ "github.com/hajimehoshi/ebiten/v2/mobile/ebitenmobileview"
)
