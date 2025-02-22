package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/realskyquest/flybit/v3"
)

const (
	DEFAULT uint8 = iota
	MENU
	INGAME
)

type MenuData struct {
	msg string
}

type Game struct {
	App flybit.App
	flybit.Game
}

var (
	GameRes generic.Resource[Game]
)

func (g *Game) Layout(outsideWidth, OutsideHeight int) (screenWidth, ScreenHeight int) {
	return outsideWidth, OutsideHeight
}

func (g *Game) Draw(screen *ebiten.Image) {}

func main() {
	game := &Game{}
	world := ecs.NewWorld()
	app := flybit.NewApp(MENU, &world, game)

	ecs.AddResource(app.GetWorld(), game)

	{

	}

	game.App = *app
	game.Load()

	ebiten.SetWindowTitle("event")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func test(world *ecs.World) {
	g := GameRes.Get()

	g.App.GetWorld()
}
