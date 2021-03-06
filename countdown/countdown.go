package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper handles sleeping a function for a bit of time
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper creates a sleeper interface to use for the countdown
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep for the given duration
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown counts down and prints a message
func Countdown(out io.Writer, sleeper Sleeper, countdownStart int, finalWord string) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper, 5, "Ahoy!")
}
