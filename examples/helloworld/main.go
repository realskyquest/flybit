package main

import (
	"flybit"
	"flybit/examples/helloworld/helloworld"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

type Game struct {
	Flybit flybit.App
	Canvas helloworld.Canvas
}

func (g *Game) Load() {
	ebiten.SetWindowTitle("Flybit - helloworld")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

func (g *Game) Quit() {

}

// Update the game.
func (g *Game) Update() error {
	if g.Flybit.Quit() {
		g.Quit()
	}

	if info, ok := g.Flybit.FileDropped(); ok {
		fmt.Println(info)
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

func main() {
	game := new(Game)
	game.Canvas = helloworld.Canvas{Image: nil, Width: 0, Height: 0}
	ecs := ecs.NewWorld()
	systems := []flybit.System{
		new(helloworld.Render),
	}

	app := flybit.New(&ecs, systems, game)
	game.Flybit = *app
	run(game)

	game.Load()
	app.Load()
	app.Run()
}

func run(g *Game) {
	ecs.AddResource(g.Flybit.World, &g.Canvas)

	gop := helloworld.Gopher{Image: helloworld.GophersImg}
	ecs.AddResource(g.Flybit.World, &gop)

	test := helloworld.Test{Image: helloworld.TestImg}
	ecs.AddResource(g.Flybit.World, &test)
}
