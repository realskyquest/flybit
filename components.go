package flybit

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ANCHOR_CENTER uint8 = iota
	ANCHOR_BOTTOM_LEFT
	ANCHOR_BOTTOM_CENTER
	ANCHOR_BOTTOM_RIGHT
	ANCHOR_CENTER_LEFT
	ANCHOR_CENTER_RIGHT
	ANCHOR_TOP_LEFT
	ANCHOR_TOP_RIGHT
)

type Sprite struct {
	Image        *ebiten.Image
	Color        color.Color
	FlipX, FlipY bool
	Anchor       uint8
}
