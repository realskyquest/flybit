package game

import (
	"fmt"

	"github.com/mlange-42/arche/ecs"
	"github.com/realskyquest/flybit"
	"github.com/realskyquest/flybit/cloudbit"
	"github.com/realskyquest/flybit/examples/gamestate/gamestate"
	"github.com/realskyquest/flybit/examples/gamestate/gamestate/game/system"
)

func Run() {
	runner := new(Game)
	world := ecs.NewWorld()

	state := cloudbit.New(
		cloudbit.Droplet{State: cloudbit.State(gamestate.LoadingScreen), Enter: func(w *ecs.World) { fmt.Println("Enter 1") }, Leave: func(w *ecs.World) { fmt.Println("Leave 1") }},
		cloudbit.Droplet{State: cloudbit.State(gamestate.MainMenu), Enter: nil, Leave: nil},
		cloudbit.Droplet{State: cloudbit.State(gamestate.InGame), Enter: nil, Leave: nil},
		cloudbit.Droplet{State: cloudbit.State(gamestate.Download), Enter: nil, Leave: nil},
	)
	state.SwitchTo(&world, gamestate.LoadingScreen)

	runner.AppState = gamestate.MyAppState{
		Cloud: state,
	}

	schedule := []flybit.System{
		&system.Input{},
		&system.Render{},
	}
	runner.Canvas = gamestate.Canvas{Image: nil, Width: 0, Height: 0}

	app := flybit.New(&world, schedule, runner)

	runner.Flybit = *app
	run(runner)

	runner.Load()
	app.Load()
	app.Run()
}

func run(g *Game) {
	runLoad(g)
}

func runLoad(g *Game) {
	g.Flybit.World.Reset()

	ecs.AddResource(g.Flybit.World, &g.AppState)
	ecs.AddResource(g.Flybit.World, &g.Canvas)
}
