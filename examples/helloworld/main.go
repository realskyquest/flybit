package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/realskyquest/flybit/v3"
)

type Game struct {
	flybit.Game
	font       *text.GoTextFace
	myMsg      string
	msgX, msgY float64
}

var (
	GameRes generic.Resource[Game]
)

func (g *Game) Layout(outsideWidth, OutsideHeight int) (screenWidth, ScreenHeight int) {
	return outsideWidth, OutsideHeight
}

func main() {
	ebiten.SetWindowTitle("helloworld")

	game := &Game{}
	app := flybit.New(0, game)
	ecs.AddResource(app.GetWorld(), game)

	app.AddSystems(flybit.LOAD, LoadRes)
	app.AddSystems(flybit.UPDATE, UpdateTextPosition)
	app.AddSystems(flybit.DRAW, DrawText)

	game.App = *app
	game.Load()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func LoadRes(world *ecs.World) {
	GameRes = generic.NewResource[Game](world)
	g := GameRes.Get()

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource := s

	mplusNormalFace := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}

	g.font = mplusNormalFace
	g.myMsg = "Hello World!"
}

func UpdateTextPosition(world *ecs.World) {
}

func DrawText(world *ecs.World) {
	g := GameRes.Get()

	op := &text.DrawOptions{}
	text.Draw(g.Canvas, g.myMsg, g.font, op)
}
