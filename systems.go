package flybit

import (
	"github.com/mlange-42/arche/ecs"
)

// Runs systems
func (a *App) AddSystems(scheduleLabel uint8, systems ...func(world *ecs.World)) {
	runAppAddSystems(a, 0, scheduleLabel, NO_CONDITION, systems)
}
func (a *SubApp) AddSystems(scheduleLabel uint8, systems ...func(world *ecs.World)) {
	runSubAppAddSystems(a, 0, scheduleLabel, NO_CONDITION, systems)
}

// Runs systems only on load
func (a *App) AddSystemsOnLoad(state uint8, systems ...func(world *ecs.World)) {
	runAppAddSystems(a, state, ON_LOAD, NO_CONDITION, systems)
}
func (a *SubApp) AddSystemsOnLoad(state uint8, systems ...func(world *ecs.World)) {
	runSubAppAddSystems(a, state, ON_LOAD, NO_CONDITION, systems)
}

// Runs systems only on transition when state changes, after on load and before on exit
func (a *App) AddSystemsOnTransition(state uint8, systems ...func(world *ecs.World)) {
	runAppAddSystems(a, state, ON_TRANSITION, NO_CONDITION, systems)
}
func (a *SubApp) AddSystemsOnTransition(state uint8, systems ...func(world *ecs.World)) {
	runSubAppAddSystems(a, state, ON_TRANSITION, NO_CONDITION, systems)
}

// Runs systems only on certain state, with a run condition(no condition, in state, state changed), no condition is the default used by addsystems, addsystemsonload, addsystemsonexit. in state only triggers in update or draw, state changed only triggers in update and must use update scheduleLabel/1, state as DEFAULT/0
func (a *App) AddSystemsRunIf(state uint8, runCondition uint8, systems ...func(world *ecs.World)) {
	runAppAddSystems(a, state, UPDATE, runCondition, systems)
}
func (a *SubApp) AddSystemsRunIf(state uint8, runCondition uint8, systems ...func(world *ecs.World)) {
	runSubAppAddSystems(a, state, UPDATE, runCondition, systems)
}

// Runs systems only on exit
func (a *App) AddSystemsOnExit(state uint8, systems ...func(world *ecs.World)) {
	runAppAddSystems(a, state, ON_EXIT, NO_CONDITION, systems)
}
func (a *SubApp) AddSystemsOnExit(state uint8, systems ...func(world *ecs.World)) {
	runSubAppAddSystems(a, state, ON_EXIT, NO_CONDITION, systems)
}
