package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/mlange-42/ark/ecs"
	"github.com/realskyquest/flybit/v3"
	"github.com/realskyquest/flybit/v3/signal"
)

type SignalGame struct {
	canvas  *ebiten.Image
	source  *text.GoTextFaceSource
	font    *text.GoTextFace
	signals *signal.Signal
}

const (
	SIGNAL_PLAYER_DETECTED signal.SignalID = iota
	SIGNAL_PLAYER_LOST
)

const (
	IS_PLAYER_DETECTED uint8 = iota
)

const (
	IS_PLAYER_DETECTED__NO flybit.State = iota
	IS_PLAYER_DETECTED__YES
	IS_PLAYER_DETECTED__SIGNAL_REMOVED
)

// -- ecs resources --
var (
	AppRes  ecs.Resource[flybit.App]
	GameRes ecs.Resource[SignalGame]
)

type Game struct {
	flybit.Game
}

func (g *Game) Layout(outsideWidth, OutsideHeight int) (screenWidth, ScreenHeight int) {
	return outsideWidth, OutsideHeight
}

func (_ *Game) Draw(screen *ebiten.Image) {
	a := AppRes.Get()
	g := GameRes.Get()

	g.canvas = screen
	g.canvas.Fill(color.White)

	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.Black)

	switch a.GetSubState(IS_PLAYER_DETECTED) {
	case IS_PLAYER_DETECTED__NO:
		text.Draw(g.canvas, "Player not found", g.font, op)
	case IS_PLAYER_DETECTED__YES:
		text.Draw(g.canvas, "Player is found", g.font, op)
	case IS_PLAYER_DETECTED__SIGNAL_REMOVED:
		text.Draw(g.canvas, "signal is removed", g.font, op)
	}
}

func main() {
	game := &Game{}
	world := ecs.NewWorld()
	app := flybit.New(0, &world)
	app.AddSubState(flybit.State(0), IS_PLAYER_DETECTED, IS_PLAYER_DETECTED__NO)

	ecs.AddResource(&world, app)
	ecs.AddResource(&world, &SignalGame{})

	app.AddSystems(flybit.LOAD, loadRes, loadFonts)
	app.AddSystems(flybit.UPDATE, handleInput)

	game.Load(app)

	ebiten.SetWindowTitle("signal")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func loadRes(world *ecs.World) {
	AppRes = ecs.NewResource[flybit.App](world)
	GameRes = ecs.NewResource[SignalGame](world)

	g := GameRes.Get()
	g.signals = signal.New()
	g.signals.Register(SIGNAL_PLAYER_DETECTED, playerDetected)
	g.signals.Register(SIGNAL_PLAYER_LOST, playerLost)
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
	g := GameRes.Get()

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.signals.Emit(world, SIGNAL_PLAYER_DETECTED)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		g.signals.Emit(world, SIGNAL_PLAYER_LOST)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.signals.Remove(SIGNAL_PLAYER_DETECTED)
		g.signals.Remove(SIGNAL_PLAYER_LOST)
		a.SetSubState(IS_PLAYER_DETECTED, IS_PLAYER_DETECTED__SIGNAL_REMOVED)
	}
}

func playerDetected(world *ecs.World) {
	a := AppRes.Get()

	a.SetSubState(IS_PLAYER_DETECTED, IS_PLAYER_DETECTED__YES)
}

func playerLost(world *ecs.World) {
	a := AppRes.Get()

	a.SetSubState(IS_PLAYER_DETECTED, IS_PLAYER_DETECTED__NO)
}
