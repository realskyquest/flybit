package cloudbit

type State int

// Struct for `Cloud`.
type Cloud struct {
	stack []State // Stores states.
}

// Creates a new `Cloud` with states.
func New(states ...State) *Cloud {
	return &Cloud{
		stack: states,
	}
}

// -- Cloud --

// Returns current `State` from stack.
func (c *Cloud) Current() State {
	return c.stack[len(c.stack)-1]
}

// Returns stack.
func (c *Cloud) Stack() []State {
	return c.stack
}

// Switch to `State`.
func (c *Cloud) SwitchTo(state State) {
	// Find the state index.
	for i, s := range c.stack {
		if s == state {
			// Remove the state from its current position.
			c.stack = append(c.stack[:i], c.stack[i+1:]...)
			break
		}
	}

	// Append the new state to the top of the stack.
	c.stack = append(c.stack, state)
}
