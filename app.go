package flybit

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

func NewApp(state uint8, world *ecs.World, game ebiten.Game) *App {
	app := &App{
		appStatePtr: &state,
		worldPtr:    world,
	}

	return app
}

func NewSubApp(world *ecs.World) *SubApp {
	subApp := &SubApp{
		worldPtr: world,
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

func (a *App) GetState() uint8 {
	return *a.appStatePtr
}

func (a *App) GetSubState(stateID uint8) uint8 {
	var state *uint8
	for _, appSubState := range a.appSubStatesPtr {
		if *appSubState.stateIDPtr == stateID {
			state = appSubState.statePtr
		}
	}
	return *state
}

func (a *App) GetWorld() *ecs.World {
	return a.worldPtr
}

func (a *App) SetState(state uint8) {
	if *a.appStatePtr != state {
		// Runs (ONEXIT schedule) systems with the current state
		runScheduleOnceStateChanged(a, *a.appStatePtr, ON_EXIT, NO_CONDITION)
		// Runs (ONTRANSITION schedule) systems with the next state
		runScheduleOnceStateChanged(a, state, ON_TRANSITION, NO_CONDITION)
		// Runs (ONLOAD schedule) systems with the next state
		runScheduleOnceStateChanged(a, state, ON_LOAD, NO_CONDITION)

		a.appStatePtr = &state
		// Runs systems when state is changed
		runScheduleOnceStateChanged(a, 0, UPDATE, STATE_CHANGED)
	}
}

func (a *App) SetSubState(stateID, state uint8) {
	for _, appSubState := range a.appSubStatesPtr {
		if *appSubState.stateIDPtr == stateID {
			appSubState.statePtr = &state
		}
	}
}

func (a *App) AddSubState(appState, stateID, state uint8) {
	a.appSubStatesPtr = append(a.appSubStatesPtr, &SubState{stateParentPtr: &appState, stateIDPtr: &stateID, statePtr: &state})
}

func (a *App) AddSubStateSystems(stateID, state uint8, systems ...func(world *ecs.World)) {
	for _, appSubState := range a.appSubStatesPtr {
		if *appSubState.stateIDPtr == stateID {
			for _, s := range systems {
				appSubState.schedulePtr = append(appSubState.schedulePtr, &System{statePtr: &state, run: s})
			}
		}
	}
}

func (a *App) AddSubApps(subApps ...*SubApp) {
	for _, sa := range subApps {
		a.subAppsPtr = append(a.subAppsPtr, sa)
	}
}
