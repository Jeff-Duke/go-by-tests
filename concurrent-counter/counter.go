package counter

import "sync"

// Counter creates a simple counter that can be executed concurrently
type Counter struct {
	val int
	mu  sync.Mutex
}

// Err is an error on the concurrent-counter
type Err string

//ErrValAtZero indicates we're trying to decrement past 0
const ErrValAtZero = (Err("cannot decrement past 0"))

func (e Err) Error() string {
	return string(e)
}

// Inc increments the counter
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

// Dec decrements the counter unless the counter is at 0
func (c *Counter) Dec() error {
	if c.val == 0 {
		return ErrValAtZero
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val--

	return nil
}

// Value returns the current value of the counter
func (c *Counter) Value() int {
	return c.val
}
