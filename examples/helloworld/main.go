package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
	"github.com/realskyquest/flybit/v3"
)

const (
	MainMenu AppState = iota
	InGame
)

type AppState uint8

type Game struct {
	flybit.Game
	myMsg string
}

func (g *Game) Layout(outsideWidth, OutsideHeight int) (screenWidth, ScreenHeight int) {
	return outsideWidth, OutsideHeight
}

func main() {
	ebiten.SetWindowTitle("helloworld")

	game := &Game{}
	app := flybit.New(game)
	app.AddSystems(flybit.LOAD, Hello)

	game.App = *app
	game.Load()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func Hello(world *ecs.World) {
	fmt.Println("HEllo")
}

func Die(world *ecs.World) {
	fmt.Println("DIE")
}
