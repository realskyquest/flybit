package system

import (
	"fmt"
	"image/color"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/realskyquest/flybit/examples/gamestate/gamestate"
)

// Handle all rendering

type Render struct {
	appState generic.Resource[gamestate.MyAppState]
	canvas   generic.Resource[gamestate.Canvas]
}

func (s *Render) Load(world *ecs.World) {
	s.appState = generic.NewResource[gamestate.MyAppState](world)
	s.canvas = generic.NewResource[gamestate.Canvas](world)
}

func (s *Render) Update(world *ecs.World) {
}

func (s *Render) Draw(world *ecs.World) {
	appState := s.appState.Get()
	canvas := s.canvas.Get()

	fmt.Println(appState.Cloud.Stack())
	switch appState.Cloud.Current() {
	case gamestate.LoadingScreen:
		canvas.Image.Fill(color.White)
	case gamestate.MainMenu:
		clr := color.RGBA{}
		clr.R = 200
		clr.G = 0
		clr.B = 0
		clr.A = 255

		canvas.Image.Fill(clr)
	case gamestate.InGame:
		clr := color.RGBA{}
		clr.R = 0
		clr.G = 200
		clr.B = 0
		clr.A = 255

		canvas.Image.Fill(clr)
	case gamestate.Download:
		clr := color.RGBA{}
		clr.R = 0
		clr.G = 0
		clr.B = 200
		clr.A = 255

		canvas.Image.Fill(clr)
	}
}
