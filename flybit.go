package flybit

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

func (g *Game) Load() {
	for _, s := range g.App.schedule {
		if s.ScheduleLabel == LOAD && s.State == 0 {
			s.System(g.App.world)
		}
	}
}

func (g *Game) Update() error {
	for _, s := range g.App.schedule {
		if s.ScheduleLabel == UPDATE && (s.State == 0 || s.State == g.App.appState) {
			s.System(g.App.world)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Canvas = screen
	for _, s := range g.App.schedule {
		if s.ScheduleLabel == DRAW && (s.State == 0 || s.State == g.App.appState) {
			s.System(g.App.world)
		}
	}
}

func (g *Game) Exit() {
	for _, s := range g.App.schedule {
		if s.ScheduleLabel == EXIT && s.State == 0 {
			s.System(g.App.world)
		}
	}
}

func New(game ebiten.Game) *App {
	world := ecs.NewWorld()
	app := &App{
		world: &world,
	}

	return app
}

func (a *App) GetWorld() *ecs.World {
	return a.world
}

func (a *App) GetSystems() []System {
	return a.schedule
}

func (a *App) SetState(state uint8) {
	if a.appState != state {
		for _, s := range a.schedule {
			if s.ScheduleLabel == EXIT && s.State == a.appState {
				s.System(a.world)
			}
		}
		for _, s := range a.schedule {
			if s.ScheduleLabel == LOAD && s.State == state {
				s.System(a.world)
			}
		}

		a.appState = state
	}
}

func (a *App) AddSystems(scheduleLabel uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: 0, ScheduleLabel: scheduleLabel, System: s})
	}
}

func (a *App) AddSystemsOnEnter(state uint8, scheduleLabel uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: scheduleLabel, System: s})
	}
}
