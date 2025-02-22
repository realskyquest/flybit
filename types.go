package flybit

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

// schedules
const (
	LOAD uint8 = iota
	UPDATE
	EXIT
	ON_LOAD
	ON_TRANSITION
	ON_EXIT
)

// run conditions
const (
	NO_CONDITION uint8 = iota
	IN_STATE
	STATE_CHANGED
)

const (
	headlessTps      = 60                        // Ticks per second (default is 60)
	headlessTimestep = time.Second / headlessTps // Duration for each tick (1000 ms / 60 = 16.666 ms)
)

type Game struct {
	running   bool // DOES NOTHING
	canvasPtr *ebiten.Image
	appPtr    *App
}

type App struct {
	appStatePtr     *uint8
	appSubStatesPtr []*SubState
	worldPtr        *ecs.World
	schedulePtr     []*System
	subAppsPtr      []*SubApp
}

type SubApp struct {
	worldPtr    *ecs.World
	schedulePtr []*System
}

type System struct {
	statePtr         *uint8
	scheduleLabelPtr *uint8
	runConditionPtr  *uint8
	run              func(world *ecs.World)
}

// appState is what state they belong to for identification
type SubState struct {
	stateParentPtr *uint8
	stateIDPtr     *uint8
	statePtr       *uint8
	schedulePtr    []*System
}
