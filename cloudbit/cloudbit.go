package cloudbit

import "github.com/mlange-42/arche/ecs"

// Use this for states.
type State int

// Struct for `Droplet`.
type Droplet struct {
	State State
	Enter func(w *ecs.World)
	Leave func(w *ecs.World)
}

// Struct for `Cloud`.
type Cloud struct {
	stack []Droplet // Stores rain.
}

// Creates a new `Cloud` with droplets.
func New(droplets ...Droplet) *Cloud {
	return &Cloud{
		stack: droplets,
	}
}

// -- Cloud --

// Returns current `Droplet` from stack.
func (c *Cloud) Current() Droplet {
	return c.stack[len(c.stack)-1]
}

// Returns stack.
func (c *Cloud) Stack() []Droplet {
	return c.stack
}

// Execute Current `Droplet` Leave(), change state, execute New `Droplet` Enter().
func (c *Cloud) SwitchTo(w *ecs.World, state State) {
	// Check for nil in enter and leave.
	if c.Current().Leave != nil {
		c.Current().Leave(w)
	}
	c.changeState(state)
	if c.Current().Enter != nil {
		c.Current().Enter(w)
	}
}

// Moves droplet to top of stack.
func (c *Cloud) changeState(state State) {
	var NextState Droplet

	// Find the state index.
	for i, s := range c.stack {
		if s.State == state {
			NextState = s
			// Remove the state from its current position.
			c.stack = append(c.stack[:i], c.stack[i+1:]...)
			break
		}
	}

	// Append the new droplet to the top of the stack.
	c.stack = append(c.stack, NextState)
}
