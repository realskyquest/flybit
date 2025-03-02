package signal

import (
	"errors"

	"github.com/mlange-42/arche/ecs"
)

type SignalID uint16

type Signal struct {
	signals []SignalID
	slots   []func(world *ecs.World)
}

func (e *Signal) Register(signal SignalID, slot func(world *ecs.World)) error {
	for _, sig := range e.signals {
		if sig == signal {
			return errors.New("signal is already registered")
		}
	}

	e.signals = append(e.signals, signal)
	e.slots = append(e.slots, slot)
	return nil
}

func (e *Signal) Emit(world *ecs.World, signal SignalID) error {
	for i, v := range e.signals {
		if v == signal {
			e.slots[i](world)
			return nil
		}
	}
	return errors.New("signal not found")
}

func (e *Signal) Remove(signal SignalID) error {
	for i, v := range e.signals {
		if v == signal {
			// Swap the element with the last one
			e.signals[i] = e.signals[len(e.signals)-1]
			e.slots[i] = e.slots[len(e.slots)-1]

			// Truncate the slices to remove the last element (now swapped)
			e.signals = e.signals[:len(e.signals)-1]
			e.slots = e.slots[:len(e.slots)-1]
			return nil
		}
	}
	return errors.New("signal not found")
}

func New() *Signal {
	return &Signal{
		signals: make([]SignalID, 0, 256),
		slots:   make([]func(world *ecs.World), 0, 256),
	}
}
