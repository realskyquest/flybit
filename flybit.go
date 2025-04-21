package flybit

import (
	"github.com/mlange-42/ark/ecs"
)

// ScheduleLabel defines when a system should be executed in the game loop
type ScheduleLabel uint8

const (
	LOAD          ScheduleLabel = iota // Execute during game loading
	UPDATE                             // Execute during game update
	EXIT                               // Execute during game exit
	ON_LOAD                            // Execute when entering a state
	ON_TRANSITION                      // Execute during state transitions
	ON_EXIT                            // Execute when exiting a state
)

// RunCondition defines under what conditions a system should run
type RunCondition uint8

const (
	NO_CONDITION  RunCondition = iota // Run without any conditions
	IN_STATE                          // Run only in specific state
	STATE_CHANGED                     // Run when state changes
)

// State represents a game state identifier
type State uint8

// SubState represents a child state that belongs to a parent state
type SubState struct {
	stateParent State
	stateID     uint8
	state       State
	schedulePtr []*System
}

// System represents a game system with its execution configuration
type System struct {
	state         State
	scheduleLabel ScheduleLabel
	runCondition  RunCondition
	run           func(world *ecs.World)
}

// Game manages the main game loop and system execution
type Game struct {
	appRes ecs.Resource[App]
}

// Load initializes the game and runs LOAD and ON_LOAD systems
func (g *Game) Load(app *App) {
	g.appRes = ecs.NewResource[App](app.worldPtr)
	runScheduleOnce(app, LOAD, ON_LOAD)
}

// Update runs all UPDATE systems and sub-state systems
func (g *Game) Update() error {
	app := g.appRes.Get()
	runSchedule(app, UPDATE)

	for _, ss := range app.subStatesPtr {
		if ss.stateParent == 0 || ss.stateParent == app.state {
			for _, s := range ss.schedulePtr {
				if s.state == ss.state {
					s.run(app.worldPtr)
				}
			}
		}
	}

	return nil
}

// Exit runs EXIT and ON_EXIT systems when the game is closing
func (g *Game) Exit() {
	app := g.appRes.Get()
	runScheduleOnce(app, EXIT, ON_EXIT)
}

// App manages game state and systems
type App struct {
	state        State
	subStatesPtr []*SubState
	worldPtr     *ecs.World
	schedulePtr  []*System
}

// GetState returns the current game state
func (a *App) GetState() State {
	return a.state
}

// GetSubState returns the state for a given sub-state ID
func (a *App) GetSubState(stateID uint8) State {
	var state State
	for _, ss := range a.subStatesPtr {
		if ss.stateID == stateID {
			state = ss.state
		}
	}
	return state
}

// GetWorld returns the ECS world
func (a *App) GetWorld() *ecs.World {
	return a.worldPtr
}

// SetState changes the current game state and triggers appropriate transition systems
func (a *App) SetState(state State) {
	if a.state != state {
		runScheduleOnceStateChanged(a, a.state, ON_EXIT, NO_CONDITION)
		runScheduleOnceStateChanged(a, state, ON_TRANSITION, NO_CONDITION)
		runScheduleOnceStateChanged(a, state, ON_LOAD, NO_CONDITION)
		a.state = state
		runScheduleOnceStateChanged(a, 0, UPDATE, STATE_CHANGED)
	}
}

// SetSubState updates the state of a specific sub-state
func (a *App) SetSubState(stateID uint8, state State) {
	for _, ss := range a.subStatesPtr {
		if ss.stateID == stateID {
			ss.state = state
		}
	}
}

// AddSystems adds systems to be run during the specified schedule
func (a *App) AddSystems(scheduleLabel ScheduleLabel, systems ...func(world *ecs.World)) {
	runAppAddSystems(a, 0, scheduleLabel, NO_CONDITION, systems)
}

// AddSystemsOnLoad adds systems to be run when entering a specific state
func (a *App) AddSystemsOnLoad(state State, systems ...func(world *ecs.World)) {
	runAppAddSystems(a, state, ON_LOAD, NO_CONDITION, systems)
}

// AddSystemsOnTransition adds systems to be run during state transitions
func (a *App) AddSystemsOnTransition(state State, systems ...func(world *ecs.World)) {
	runAppAddSystems(a, state, ON_TRANSITION, NO_CONDITION, systems)
}

// AddSystemsRunIf adds systems that only run in a specific state
func (a *App) AddSystemsRunIf(state State, systems ...func(world *ecs.World)) {
	runAppAddSystems(a, state, UPDATE, IN_STATE, systems)
}

// AddSystemsOnChange adds systems that run when the state changes
func (a *App) AddSystemsOnChange(systems ...func(world *ecs.World)) {
	runAppAddSystems(a, 0, UPDATE, STATE_CHANGED, systems)
}

// AddSystemsOnExit adds systems to be run when exiting a state
func (a *App) AddSystemsOnExit(state State, systems ...func(world *ecs.World)) {
	runAppAddSystems(a, state, ON_EXIT, NO_CONDITION, systems)
}

// AddSubState registers a new sub-state within the application's state hierarchy.
// A sub-state represents a collection of systems that should only run when a specific
// parent state is active. This allows for hierarchical organization of game states
// and their associated systems.
//
// Parameters:
//   - state: The parent state under which this sub-state will operate
//   - stateID: An 8-bit identifier for uniquely identifying this state
//   - subState: The State implementation containing systems to run in this sub-state
//
// Example:
//   app.AddSubState(GameState, 1, CombatState)
//   // CombatState systems will only run when GameState is active
func (a *App) AddSubState(state State, stateID uint8, subState State) {
	a.subStatesPtr = append(a.subStatesPtr, &SubState{stateParent: state, stateID: stateID, state: subState})
}

// AddSubStateSystems adds systems to a specific sub-state
func (a *App) AddSubStateSystems(stateID uint8, state State, systems ...func(world *ecs.World)) {
	for _, ss := range a.subStatesPtr {
		if ss.stateID == stateID {
			for _, s := range systems {
				ss.schedulePtr = append(ss.schedulePtr, &System{state: state, run: s})
			}
		}
	}
}

// New creates a new App instance with the specified initial state and ECS world
func New(state State, world *ecs.World) *App {
	return &App{
		state:        state,
		subStatesPtr: make([]*SubState, 0, 256),
		worldPtr:     world,
		schedulePtr:  make([]*System, 0, 256),
	}
}
