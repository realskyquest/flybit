package flybit

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

func (g *Game) Load() {
	for _, s := range g.App.schedule {
		if (s.State == 0 || s.State == g.App.appState) && (s.ScheduleLabel == LOAD || s.ScheduleLabel == ONLOAD) && s.RunCondition == NOCONDITION {
			s.Run(g.App.world)
		}
	}
	// sub
	for _, sa := range g.App.subApps {
		for _, s := range sa.schedule {
			if (s.State == 0 || s.State == g.App.appState) && (s.ScheduleLabel == LOAD || s.ScheduleLabel == ONLOAD) && s.RunCondition == NOCONDITION {
				s.Run(g.App.world)
			}
		}
	}
}

func (g *Game) Update() error {
	for _, s := range g.App.schedule {
		if (s.State == 0 || s.State == g.App.appState) && s.ScheduleLabel == UPDATE && (s.RunCondition == NOCONDITION || s.RunCondition == INSTATE) {
			s.Run(g.App.world)
		}
	}
	// sub
	for _, sa := range g.App.subApps {
		for _, s := range sa.schedule {
			if (s.State == 0 || s.State == g.App.appState) && s.ScheduleLabel == UPDATE && (s.RunCondition == NOCONDITION || s.RunCondition == INSTATE) {
				s.Run(g.App.world)
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Canvas = screen
	for _, s := range g.App.schedule {
		if (s.State == 0 || s.State == g.App.appState) && s.ScheduleLabel == DRAW && (s.RunCondition == NOCONDITION || s.RunCondition == INSTATE) {
			s.Run(g.App.world)
		}
	}
	// sub
	for _, sa := range g.App.subApps {
		for _, s := range sa.schedule {
			if (s.State == 0 || s.State == g.App.appState) && s.ScheduleLabel == DRAW && (s.RunCondition == NOCONDITION || s.RunCondition == INSTATE) {
				s.Run(g.App.world)
			}
		}
	}
}

func (g *Game) Exit() {
	for _, s := range g.App.schedule {
		if (s.State == 0 || s.State == g.App.appState) && (s.ScheduleLabel == EXIT || s.ScheduleLabel == ONEXIT) && s.RunCondition == NOCONDITION {
			s.Run(g.App.world)
		}
	}
	// sub
	for _, sa := range g.App.subApps {
		for _, s := range sa.schedule {
			if (s.State == 0 || s.State == g.App.appState) && (s.ScheduleLabel == EXIT || s.ScheduleLabel == ONEXIT) && s.RunCondition == NOCONDITION {
				s.Run(g.App.world)
			}
		}
	}
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

func (a *App) GetWorld() *ecs.World {
	return a.world
}

func (a *App) SetState(state uint8) {
	if a.appState != state {
		// Runs (ONEXIT schedule) systems with the current state
		for _, s := range a.schedule {
			if s.State == a.appState && s.ScheduleLabel == ONEXIT && s.RunCondition == NOCONDITION {
				s.Run(a.world)
			}
		}
		// Runs (ONTRANSITION schedule) systems with the next state
		for _, s := range a.schedule {
			if s.State == state && s.ScheduleLabel == ONTRANSITION && s.RunCondition == NOCONDITION {
				s.Run(a.world)
			}
		}
		// Runs (ONLOAD schedule) systems with the next state
		for _, s := range a.schedule {
			if s.State == state && s.ScheduleLabel == ONLOAD && s.RunCondition == NOCONDITION {
				s.Run(a.world)
			}
		}

		a.appState = state
		// Runs systems when state is changed
		for _, s := range a.schedule {
			if s.State == 0 && s.ScheduleLabel == UPDATE && s.RunCondition == STATECHANGED {
				s.Run(a.world)
			}
		}
	}
}

func (a *App) AddSubApps(subApps ...*SubApp) {
	for _, sa := range subApps {
		a.subApps = append(a.subApps, *sa)
	}
}

// Runs systems
func (a *App) AddSystems(scheduleLabel uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: 0, ScheduleLabel: scheduleLabel, RunCondition: NOCONDITION, Run: s})
	}
}
func (a *SubApp) AddSystems(scheduleLabel uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: 0, ScheduleLabel: scheduleLabel, RunCondition: NOCONDITION, Run: s})
	}
}

// Runs systems only on load
func (a *App) AddSystemsOnLoad(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ONLOAD, RunCondition: NOCONDITION, Run: s})
	}
}
func (a *SubApp) AddSystemsOnLoad(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ONLOAD, RunCondition: NOCONDITION, Run: s})
	}
}

// Runs systems only on transition when state changes, after on load and before on exit
func (a *App) AddSystemsOnTransition(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ONTRANSITION, RunCondition: NOCONDITION, Run: s})
	}
}
func (a *SubApp) AddSystemsOnTransition(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ONTRANSITION, RunCondition: NOCONDITION, Run: s})
	}
}

// Runs systems only on certain state, with a run condition(no condition, in state, state changed), no condition is the default used by addsystems, addsystemsonload, addsystemsonexit. in state only triggers in update or draw, state changed only triggers in update and must use update scheduleLabel/1, state as DEFAULT/0
func (a *App) AddSystemsRunIf(scheduleLabel uint8, state uint8, runCondition uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: scheduleLabel, RunCondition: runCondition, Run: s})
	}
}
func (a *SubApp) AddSystemsRunIf(scheduleLabel uint8, state uint8, runCondition uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: scheduleLabel, RunCondition: runCondition, Run: s})
	}
}

// Runs systems only on exit
func (a *App) AddSystemsOnExit(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ONEXIT, RunCondition: NOCONDITION, Run: s})
	}
}
func (a *SubApp) AddSystemsOnExit(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ONEXIT, RunCondition: NOCONDITION, Run: s})
	}
}
