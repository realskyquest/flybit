package flybit

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Load() {
	runScheduleOnce(g.appPtr, LOAD, ON_LOAD)
}

func (g *Game) Update() error {
	runSchedule(g.appPtr, UPDATE)

	for _, appSubState := range g.appPtr.appSubStatesPtr {
		if *appSubState.stateParentPtr == *g.appPtr.appStatePtr {
			for _, s := range appSubState.schedulePtr {
				if *s.statePtr == *appSubState.statePtr {
					s.run(g.appPtr.worldPtr)
				}
			}
		}
	}

	return nil
}

func (g *Game) Exit() {
	runScheduleOnce(g.appPtr, EXIT, ON_EXIT)
}

func (g *Game) GetCanvas() *ebiten.Image {
	return g.canvasPtr
}

func (g *Game) GetApp() *App {
	return g.appPtr
}

func (g *Game) SetCanvas(screen *ebiten.Image) {
	g.canvasPtr = screen
}

func (g *Game) SetApp(app *App) {
	g.appPtr = app
}
