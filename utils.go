package flybit

import "github.com/mlange-42/arche/ecs"

func runScheduleOnce(app *App, scheduleLabelA, ScheduleLabelB uint8) {
	for _, s := range app.schedule {
		if (s.State == 0 || s.State == app.appState) && (s.ScheduleLabel == scheduleLabelA || s.ScheduleLabel == ScheduleLabelB) && s.RunCondition == NO_CONDITION {
			s.Run(app.world)
		}
	}

	for _, sa := range app.subApps {
		for _, s := range sa.schedule {
			if (s.State == 0 || s.State == app.appState) && (s.ScheduleLabel == scheduleLabelA || s.ScheduleLabel == ScheduleLabelB) && s.RunCondition == NO_CONDITION {
				s.Run(app.world)
			}
		}
	}
}

func runSchedule(app *App, scheduleLabel uint8) {
	for _, s := range app.schedule {
		if (s.State == 0 || s.State == app.appState) && s.ScheduleLabel == scheduleLabel && (s.RunCondition == NO_CONDITION || s.RunCondition == IN_STATE) {
			s.Run(app.world)
		}
	}
	// sub
	for _, sa := range app.subApps {
		for _, s := range sa.schedule {
			if (s.State == 0 || s.State == app.appState) && s.ScheduleLabel == scheduleLabel && (s.RunCondition == NO_CONDITION || s.RunCondition == IN_STATE) {
				s.Run(app.world)
			}
		}
	}
}

func runScheduleOnceStateChanged(app *App, state, scheduleLabel, runCondition uint8) {
	for _, s := range app.schedule {
		if s.State == state && s.ScheduleLabel == scheduleLabel && s.RunCondition == runCondition {
			s.Run(app.world)
		}
	}
}

func runAppAddSystems(a *App, state, scheduleLabel, runCondition uint8, systems []func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: scheduleLabel, RunCondition: runCondition, Run: s})
	}
}

func runSubAppAddSystems(a *SubApp, state, scheduleLabel, runCondition uint8, systems []func(world *ecs.World)) {
	for _, s := range systems {
		a.schedule = append(a.schedule, System{State: state, ScheduleLabel: scheduleLabel, RunCondition: runCondition, Run: s})
	}
}
