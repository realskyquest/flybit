package flybit

import (
	"flag"
	"io/fs"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tinne26/mipix"
)

var (
	debugPrefix = "[Flybit]"
	debug       = flag.Bool("flydebug", false, "enable debug mode for flybit")
)

// -- schedule --

// Interface for system.
type System interface {
	// Loads system.
	Load(w *ecs.World)
	// Update system.
	Update(w *ecs.World)
	// Draw system.
	Draw(w *ecs.World)
}

// -- App --

// Enables mipix by tinne26.
var MipixSupport = false

// Struct for `App`.
type App struct {
	World    *ecs.World  // Stores a world from arche ECS.
	Schedule []System    // Stores systems.
	runner   ebiten.Game // Stores a ebiten game.
}

// Creates a new `App` with a world, systems, runner.
func New(w *ecs.World, s []System, r ebiten.Game) *App {
	app := new(App)
	app.World = w
	app.Schedule = s
	app.runner = r

	if *debug {
		log.Info().Str(debugPrefix, "..:: Flybit ::..").Send()
		log.Info().Str(debugPrefix, "World is ready").Send()
		log.Info().Str(debugPrefix, "Schedule is ready").Send()
		log.Info().Str(debugPrefix, "Runner is ready").Send()
	}

	return app
}

// -- general --

// Runs the `App` by calling its `Runner`.
func (a *App) Run() {
	if *debug {
		log.Info().Str(debugPrefix, "Running Flybit App").Send()
	}

	if MipixSupport == true {
		if err := mipix.Run(a.runner); err != nil {
			panic(err)
		}
	} else {
		if err := ebiten.RunGame(a.runner); err != nil {
			panic(err)
		}
	}
}

// Loads systems in `Schedule`.
func (a *App) Load() {
	if *debug {
		log.Info().Str(debugPrefix, "App is loading").Send()
	}

	for _, s := range a.Schedule {
		s.Load(a.World)
	}
}

// Updates systems in `Schedule`.
func (a *App) Update() {
	for _, s := range a.Schedule {
		s.Update(a.World)
	}
}

// Draws systems in `Schedule`.
func (a *App) Draw() {
	for _, s := range a.Schedule {
		s.Draw(a.World)
	}
}

// Wrapper for `ebiten.IsWindowBeingClosed`, always returns false if the platform is not a desktop.
func (a *App) Quit() bool {
	return ebiten.IsWindowBeingClosed()
}

// -- window --

// Wrapper for `ebiten.DroppedFiles` (if info, ok := g.Flybit.FileDropped(); ok {}).
func (a *App) FileDropped() (fs.FS, bool) {
	file := ebiten.DroppedFiles()
	if file != nil {
		return file, true
	}

	return nil, false
}

// Wrapper for `ebiten.IsFocused`.
func (a *App) Focus() bool {
	return ebiten.IsFocused()
}

func init() {
	flag.Parse()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
