package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/realskyquest/flybit/v3"
)

type StateGame struct {
	canvas *ebiten.Image
	source *text.GoTextFaceSource
	font   *text.GoTextFace
	msg    string
}

const (
	DEFAULT flybit.State = iota
	MENU
	INGAME
)

// -- ecs resources --
var (
	AppRes  generic.Resource[flybit.App]
	GameRes generic.Resource[StateGame]
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
	text.Draw(g.canvas, g.msg, g.font, op)
}

func main() {
	game := &Game{}
	world := ecs.NewWorld()
	app := flybit.New(MENU, &world)

	ecs.AddResource(&world, app)
	ecs.AddResource(&world, &StateGame{})

	app.AddSystems(flybit.LOAD, loadRes, loadFonts)
	app.AddSystems(flybit.UPDATE, handleInput)

	app.AddSystemsRunIf(MENU, menu)
	app.AddSystemsRunIf(INGAME, ingame)

	game.Load(app)

	ebiten.SetWindowTitle("state")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func loadRes(world *ecs.World) {
	AppRes = generic.NewResource[flybit.App](world)
	GameRes = generic.NewResource[StateGame](world)
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

func handleInput(world *ecs.World) {
	a := AppRes.Get()

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit1) {
		a.SetState(MENU)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDigit2) {
		a.SetState(INGAME)
	}
}

func menu(world *ecs.World) {
	g := GameRes.Get()

	g.msg = "menu"
}

func ingame(world *ecs.World) {
	g := GameRes.Get()

	g.msg = "ingame"
}
