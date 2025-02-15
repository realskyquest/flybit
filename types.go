package flybit

import (
	"github.com/mlange-42/arche/ecs"
)

const (
	LOAD uint8 = iota
	UPDATE
	DRAW
	EXIT
)

type System struct {
	ScheduleLabel uint8
	System        func(world *ecs.World)
}

type Game struct {
	App App
}

type App struct {
	world    *ecs.World
	schedule []System
}
