package flybit

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

func (g *Game) Load() {
	runScheduleOnce(&g.App, LOAD, ON_LOAD)
}

func (g *Game) Update() error {
	runSchedule(&g.App, UPDATE)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Canvas = screen
	runSchedule(&g.App, DRAW)
}

func (g *Game) Exit() {
	runScheduleOnce(&g.App, EXIT, ON_EXIT)
}

func NewApp(state uint8, world *ecs.World, game ebiten.Game) *App {
	app := &App{
		appState: state,
		world:    world,
	}

	return app
}

func NewSubApp(world *ecs.World) *SubApp {
	subApp := &SubApp{
		world: world,
	}

	return subApp
}

// Runs app at 60 TPS headless without any rendering
func (a *App) HeadlessRun() {
	lastTime := time.Now()

	for {
		// Calculate how much time has passed since the last frame
		now := time.Now()
		delta := now.Sub(lastTime)

		// If the time passed is greater than or equal to the timestep, we perform the update
		if delta >= headlessTimestep {
			runSchedule(a, UPDATE)

			// Adjust the lastTime to match the fixed timestep
			lastTime = now
		}

		// Sleep for a short time to avoid maxing out the CPU
		time.Sleep(1 * time.Millisecond)
	}
}

func (a *App) GetWorld() *ecs.World {
	return a.world
}

func (a *App) SetState(state uint8) {
	if a.appState != state {
		// Runs (ONEXIT schedule) systems with the current state
		runScheduleOnceStateChanged(a, a.appState, ON_EXIT, NO_CONDITION)
		// Runs (ONTRANSITION schedule) systems with the next state
		runScheduleOnceStateChanged(a, state, ON_TRANSITION, NO_CONDITION)
		// Runs (ONLOAD schedule) systems with the next state
		runScheduleOnceStateChanged(a, state, ON_LOAD, NO_CONDITION)

		a.appState = state
		// Runs systems when state is changed
		runScheduleOnceStateChanged(a, 0, UPDATE, STATE_CHANGED)
	}
}

func (a *App) AddSubApps(subApps ...*SubApp) {
	for _, sa := range subApps {
		a.subApps = append(a.subApps, *sa)
	}
}
