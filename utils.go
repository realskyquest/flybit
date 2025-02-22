package flybit

import "github.com/mlange-42/arche/ecs"

func runScheduleOnce(app *App, scheduleLabelA, ScheduleLabelB uint8) {
	for _, s := range app.schedulePtr {
		if (*s.statePtr == 0 || *s.statePtr == *app.appStatePtr) && (*s.scheduleLabelPtr == scheduleLabelA || *s.scheduleLabelPtr == ScheduleLabelB) && *s.runConditionPtr == NO_CONDITION {
			s.run(app.worldPtr)

		}
	}

	for _, sa := range app.subAppsPtr {
		for _, s := range sa.schedulePtr {
			if (*s.statePtr == 0 || *s.statePtr == *app.appStatePtr) && (*s.scheduleLabelPtr == scheduleLabelA || *s.scheduleLabelPtr == ScheduleLabelB) && *s.runConditionPtr == NO_CONDITION {
				s.run(app.worldPtr)
			}
		}
	}
}

func runSchedule(app *App, scheduleLabel uint8) {
	for _, s := range app.schedulePtr {
		if (*s.statePtr == 0 || *s.statePtr == *app.appStatePtr) && *s.scheduleLabelPtr == scheduleLabel && (*s.runConditionPtr == NO_CONDITION || *s.runConditionPtr == IN_STATE) {
			s.run(app.worldPtr)
		}
	}
	// sub
	for _, sa := range app.subAppsPtr {
		for _, s := range sa.schedulePtr {
			if (*s.statePtr == 0 || *s.statePtr == *app.appStatePtr) && *s.scheduleLabelPtr == scheduleLabel && (*s.runConditionPtr == NO_CONDITION || *s.runConditionPtr == IN_STATE) {
				s.run(app.worldPtr)
			}
		}
	}
}

func runScheduleOnceStateChanged(app *App, state, scheduleLabel, runCondition uint8) {
	for _, s := range app.schedulePtr {
		if *s.statePtr == state && *s.scheduleLabelPtr == scheduleLabel && *s.runConditionPtr == runCondition {
			s.run(app.worldPtr)
		}
	}
}

func runAppAddSystems(a *App, state, scheduleLabel, runCondition uint8, systems []func(world *ecs.World)) {
	for _, s := range systems {
		a.schedulePtr = append(a.schedulePtr, &System{statePtr: &state, scheduleLabelPtr: &scheduleLabel, runConditionPtr: &runCondition, run: s})
	}
}

func runSubAppAddSystems(a *SubApp, state, scheduleLabel, runCondition uint8, systems []func(world *ecs.World)) {
	for _, s := range systems {
		a.schedulePtr = append(a.schedulePtr, &System{statePtr: &state, scheduleLabelPtr: &scheduleLabel, runConditionPtr: &runCondition, run: s})
	}
}
