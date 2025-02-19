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
	running bool
	Canvas  *ebiten.Image
	App     App
}

type App struct {
	appState uint8
	world    *ecs.World
	schedule []System
	subApps  []SubApp
}

type SubApp struct {
	world    *ecs.World
	schedule []System
}

type System struct {
	State         uint8
	ScheduleLabel uint8
	RunCondition  uint8
	Run           func(world *ecs.World)
}
