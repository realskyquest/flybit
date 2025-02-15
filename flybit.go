package flybit

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

func (g *Game) Load() {
	for _, s := range g.App.schedule {
		if s.ScheduleLabel == LOAD {
			s.System(g.App.world)
		}
	}
}

func (g *Game) Update() error {
	for _, s := range g.App.schedule {
		if s.ScheduleLabel == UPDATE {
			s.System(g.App.world)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, s := range g.App.schedule {
		if s.ScheduleLabel == DRAW {
			s.System(g.App.world)
		}
	}
}

func (g *Game) Exit() {
	for _, s := range g.App.schedule {
		if s.ScheduleLabel == EXIT {
			s.System(g.App.world)
		}
	}
}

func New(game ebiten.Game) *App {
	world := ecs.NewWorld()
	app := &App{
		world: &world,
	}
	ecs.AddResource(&world, &game)

	return app
}

func (a *App) GetSystems() []System {
	return a.schedule
}

func (a *App) AddSystems(scheduleLabel uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{ScheduleLabel: scheduleLabel, System: s})
	}
}
