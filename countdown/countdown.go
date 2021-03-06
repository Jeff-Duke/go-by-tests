package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Countdown counts down and prints a message
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
