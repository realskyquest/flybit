package flybit

import (
	"github.com/mlange-42/arche/ecs"
)

// Runs systems
func (a *App) AddSystems(scheduleLabel uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: 0, ScheduleLabel: scheduleLabel, RunCondition: NO_CONDITION, Run: s})
	}
}
func (a *SubApp) AddSystems(scheduleLabel uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: 0, ScheduleLabel: scheduleLabel, RunCondition: NO_CONDITION, Run: s})
	}
}

// Runs systems only on load
func (a *App) AddSystemsOnLoad(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ON_LOAD, RunCondition: NO_CONDITION, Run: s})
	}
}
func (a *SubApp) AddSystemsOnLoad(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ON_LOAD, RunCondition: NO_CONDITION, Run: s})
	}
}

// Runs systems only on transition when state changes, after on load and before on exit
func (a *App) AddSystemsOnTransition(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ON_TRANSITION, RunCondition: NO_CONDITION, Run: s})
	}
}
func (a *SubApp) AddSystemsOnTransition(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ON_TRANSITION, RunCondition: NO_CONDITION, Run: s})
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
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ON_EXIT, RunCondition: NO_CONDITION, Run: s})
	}
}
func (a *SubApp) AddSystemsOnExit(state uint8, systems ...func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: ON_EXIT, RunCondition: NO_CONDITION, Run: s})
	}
}
