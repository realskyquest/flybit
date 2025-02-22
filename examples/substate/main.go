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

const (
	IS_PAUSED uint8 = iota
)

// IS_PAUSED
const (
	IS_PAUSED__PLAYING uint8 = iota
	IS_PAUSED__PAUSED
)

type Game struct {
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
	app.AddSubState(MENU, IS_PAUSED, IS_PAUSED__PLAYING)

	ecs.AddResource(app.GetWorld(), game)

	{
		app.AddSystems(flybit.LOAD, setup)
		app.AddSystems(flybit.UPDATE, handleInput)

		app.AddSubStateSystems(IS_PAUSED, IS_PAUSED__PLAYING, playingGame)
		app.AddSubStateSystems(IS_PAUSED, IS_PAUSED__PAUSED, pausedGame)

		app.AddSystemsRunIf(DEFAULT, flybit.STATE_CHANGED, handleStateChange)
	}

	game.SetApp(app)
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
		g.GetApp().SetState(MENU)
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.GetApp().SetState(INGAME)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		g.GetApp().SetSubState(IS_PAUSED, IS_PAUSED__PLAYING)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyY) {
		g.GetApp().SetSubState(IS_PAUSED, IS_PAUSED__PAUSED)
	}
}

func playingGame(world *ecs.World) {
	fmt.Println("Playing")
}

func pausedGame(world *ecs.World) {
	fmt.Println("Paused")
}
