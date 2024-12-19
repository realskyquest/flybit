package gamestate

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/realskyquest/flybit/cloudbit"
)

type MyAppState struct {
	Cloud *cloudbit.Cloud
}

// Canvas resource for drawing.
type Canvas struct {
	// The screen image.
	Image *ebiten.Image
	// Current screen width.
	Width int
	// Current screen height.
	Height int
}

const (
	LoadingScreen cloudbit.State = iota
	MainMenu
	InGame
	Download
)
