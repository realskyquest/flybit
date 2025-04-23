package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/mlange-42/ark/ecs"
	"github.com/realskyquest/flybit/v3"
)

type HelloworldGame struct {
	canvas *ebiten.Image
	source *text.GoTextFaceSource
	font   *text.GoTextFace
}

// -- ecs resources --
var (
	AppRes  ecs.Resource[flybit.App]
	GameRes ecs.Resource[HelloworldGame]
)

type Game struct {
	flybit.Game
}

func (g *Game) Layout(outsideWidth, OutsideHeight int) (screenWidth, ScreenHeight int) {
	return outsideWidth, OutsideHeight
}

func (_ *Game) Draw(screen *ebiten.Image) {
	g := GameRes.Get()

	g.canvas = screen
	g.canvas.Fill(color.White)

	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.Black)
	text.Draw(g.canvas, "Hello World!", g.font, op)
}

func main() {
	game := &Game{}
	world := ecs.NewWorld()
	app := flybit.New(0, &world)

	ecs.AddResource(&world, app)
	ecs.AddResource(&world, &HelloworldGame{})

	app.AddSystems(flybit.LOAD, loadRes, loadFonts)

	game.Load(app)

	ebiten.SetWindowTitle("helloworld")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func loadRes(world *ecs.World) {
	AppRes = ecs.NewResource[flybit.App](world)
	GameRes = ecs.NewResource[HelloworldGame](world)
}

func loadFonts(world *ecs.World) {
	g := GameRes.Get()

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	g.source = s

	g.font = &text.GoTextFace{
		Source: s,
		Size:   24,
	}
}
