package game

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/realskyquest/flybit"
	"github.com/realskyquest/flybit/cloudbit"
	"github.com/realskyquest/flybit/examples/gamestate/gamestate"
	"github.com/realskyquest/flybit/examples/gamestate/gamestate/game/system"
)

func Run() {
	runner := new(Game)

	state := cloudbit.New(
		cloudbit.State(gamestate.LoadingScreen),
		cloudbit.State(gamestate.MainMenu),
		cloudbit.State(gamestate.InGame),
		cloudbit.State(gamestate.Download),
	)
	state.SwitchTo(gamestate.LoadingScreen)
	runner.AppState = gamestate.MyAppState{
		Cloud: state,
	}

	world := ecs.NewWorld()
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
