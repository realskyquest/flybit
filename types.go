package flybit

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

const (
	LOAD uint8 = iota
	UPDATE
	DRAW
	EXIT
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
}

type System struct {
	State         uint8
	ScheduleLabel uint8
	System        func(world *ecs.World)
}
