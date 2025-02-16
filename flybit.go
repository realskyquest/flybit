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
}

func (g *Game) Update() error {
	for _, s := range g.App.schedule {
		if (s.State == 0 || s.State == g.App.appState) && s.ScheduleLabel == UPDATE && (s.RunCondition == NOCONDITION || s.RunCondition == INSTATE) {
			s.Run(g.App.world)
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
}

func (g *Game) Exit() {
	for _, s := range g.App.schedule {
		if (s.State == 0 || s.State == g.App.appState) && (s.ScheduleLabel == EXIT || s.ScheduleLabel == ONEXIT) && s.RunCondition == NOCONDITION {
			s.Run(g.App.world)
		}
	}
}

func New(state uint8, game ebiten.Game) *App {
	world := ecs.NewWorld()
	app := &App{
		appState: state,
		world:    &world,
	}

	return app
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

func (a *App) AddSystems(scheduleLabel uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: 0, ScheduleLabel: scheduleLabel, RunCondition: NOCONDITION, Run: s})
	}
}

func (a *App) AddSystemsOnLoad(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ONLOAD, RunCondition: NOCONDITION, Run: s})
	}
}

func (a *App) AddSystemsOnTransition(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ONTRANSITION, RunCondition: NOCONDITION, Run: s})
	}
}

func (a *App) AddSystemsRunIf(scheduleLabel uint8, state uint8, runCondition uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: scheduleLabel, RunCondition: runCondition, Run: s})
	}
}

func (a *App) AddSystemsOnExit(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ONEXIT, RunCondition: NOCONDITION, Run: s})
	}
}
