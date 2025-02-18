package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	flybit.Game
}

var (
	GameRes generic.Resource[Game]
)

func (g *Game) Layout(outsideWidth, OutsideHeight int) (screenWidth, ScreenHeight int) {
	return outsideWidth, OutsideHeight
}

func main() {
	game := &Game{}
	world := ecs.NewWorld()
	app := flybit.NewApp(MENU, &world, game)

	ecs.AddResource(app.GetWorld(), game)

	{
		app.AddSystems(flybit.LOAD, setup)
		app.AddSystems(flybit.UPDATE, handleInput)

		app.AddSystemsRunIf(flybit.UPDATE, DEFAULT, flybit.STATE_CHANGED, handleStateChange)

		app.AddSystemsOnLoad(MENU, setupMenu)
		app.AddSystemsRunIf(flybit.UPDATE, MENU, flybit.IN_STATE, menu)
		app.AddSystemsOnExit(MENU, cleanupMenu)

		app.AddSystemsOnLoad(INGAME, setupGame)
		app.AddSystemsRunIf(flybit.UPDATE, INGAME, flybit.IN_STATE, movement, changeColor)
	}

	game.App = *app
	game.Load()

	ebiten.SetWindowTitle("event")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func setup(world *ecs.World) {
	fmt.Println("Setup")
	GameRes = generic.NewResource[Game](world)
}

func handleStateChange(world *ecs.World) {
	fmt.Println("STATE CHANGED")
}

func handleInput(world *ecs.World) {
	g := GameRes.Get()

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.App.SetState(MENU)
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.App.SetState(INGAME)
	}
}

func setupMenu(world *ecs.World) {
	fmt.Println("Setup Menu")
}

func menu(world *ecs.World) {
	fmt.Println("Menu")
}

func cleanupMenu(world *ecs.World) {
	fmt.Println("Cleanup Menu")
}

func setupGame(world *ecs.World) {
	fmt.Println("Setup Game")

}

func movement(world *ecs.World) {
	fmt.Println("movement")

}

func changeColor(world *ecs.World) {
	fmt.Println("change color")

}
