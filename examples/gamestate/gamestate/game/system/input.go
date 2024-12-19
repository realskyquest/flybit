package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/realskyquest/flybit/examples/gamestate/gamestate"
)

// Handle all rendering

type Input struct {
	appState generic.Resource[gamestate.MyAppState]
}

func (s *Input) Load(world *ecs.World) {
	s.appState = generic.NewResource[gamestate.MyAppState](world)
}

func (s *Input) Update(world *ecs.World) {
	appState := s.appState.Get()

	if inpututil.IsKeyJustPressed(ebiten.Key1) && appState.Cloud.Current() != gamestate.LoadingScreen {
		appState.Cloud.SwitchTo(gamestate.LoadingScreen)
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) && appState.Cloud.Current() != gamestate.MainMenu {
		appState.Cloud.SwitchTo(gamestate.MainMenu)
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) && appState.Cloud.Current() != gamestate.InGame {
		appState.Cloud.SwitchTo(gamestate.InGame)
	} else if inpututil.IsKeyJustPressed(ebiten.Key4) && appState.Cloud.Current() != gamestate.Download {
		appState.Cloud.SwitchTo(gamestate.Download)
	}
}

func (s *Input) Draw(world *ecs.World) {
}
