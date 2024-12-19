package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/realskyquest/flybit"
	"github.com/realskyquest/flybit/examples/gamestate/gamestate"
)

type Game struct {
	AppState gamestate.MyAppState
	Flybit   flybit.App
	Canvas   gamestate.Canvas
}

// Load the game.
func (g *Game) Load() {
	ebiten.SetWindowTitle("Flybit - gamestate")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

// Handle game quit.
func (g *Game) Quit() {

}

// Update the game.
func (g *Game) Update() error {
	if g.Flybit.Quit() {
		g.Quit()
	}

	g.Flybit.Update()
	return nil
}

// Draw the game.
func (g *Game) Draw(screen *ebiten.Image) {
	g.Canvas.Image = screen
	g.Canvas.Width = screen.Bounds().Dx()
	g.Canvas.Height = screen.Bounds().Dy()

	g.Flybit.Draw()
}

// Layout the game.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.Monitor().DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}
