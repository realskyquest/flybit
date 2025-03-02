package flybit

import "github.com/mlange-42/arche/ecs"

// runScheduleOnce executes systems that match either of the two provided schedule labels
// and have no run conditions. Systems must either have state 0 (global) or match the app's current state.
func runScheduleOnce(app *App, scheduleLabelA, ScheduleLabelB ScheduleLabel) {
	for _, s := range app.schedulePtr {
		if (s.state == 0 || s.state == app.state) && (s.scheduleLabel == scheduleLabelA || s.scheduleLabel == ScheduleLabelB) && s.runCondition == NO_CONDITION {
			s.run(app.worldPtr)

		}
	}
}

// runSchedule executes all systems with the specified schedule label that either have no run conditions
// or are configured to run in the current state. Systems must either have state 0 (global) or match the app's current state.
func runSchedule(app *App, scheduleLabel ScheduleLabel) {
	for _, s := range app.schedulePtr {
		if (s.state == 0 || s.state == app.state) && s.scheduleLabel == scheduleLabel && (s.runCondition == NO_CONDITION || s.runCondition == IN_STATE) {
			s.run(app.worldPtr)
		}
	}
}

// runScheduleOnceStateChanged executes systems that match the specified state, schedule label and run condition.
// This is typically used for handling state transition events.
func runScheduleOnceStateChanged(app *App, state State, scheduleLabel ScheduleLabel, runCondition RunCondition) {
	for _, s := range app.schedulePtr {
		if s.state == state && s.scheduleLabel == scheduleLabel && s.runCondition == runCondition {
			s.run(app.worldPtr)
		}
	}
}

// runAppAddSystems adds multiple system functions to the app's schedule with the specified state,
// schedule label and run condition configurations.
func runAppAddSystems(a *App, state State, scheduleLabel ScheduleLabel, runCondition RunCondition, systems []func(world *ecs.World)) {
	for _, s := range systems {
		a.schedulePtr = append(a.schedulePtr, &System{state: state, scheduleLabel: scheduleLabel, runCondition: runCondition, run: s})
	}
}
