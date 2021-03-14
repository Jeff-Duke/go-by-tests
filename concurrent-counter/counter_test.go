package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter(0)
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("decrementing the counter to 0 leaves it at 0", func(t *testing.T) {
		counter := NewCounter(3)

		counter.Dec()
		counter.Dec()
		counter.Dec()

		assertCounter(t, counter, 0)
	})

	t.Run("cannot decrement past 0", func(t *testing.T) {
		counter := NewCounter(2)

		counter.Dec()
		counter.Dec()
		counter.Dec()

		assertCounter(t, counter, 0)
	})

	t.Run("it decrements safely concurrently", func(t *testing.T) {
		startCount := 1000
		counter := NewCounter(startCount)

		var wg sync.WaitGroup
		wg.Add(startCount)

		for i := 1000; i > 0; i-- {
			go func(w *sync.WaitGroup) {
				counter.Dec()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, 0)
	})

	t.Run("it increments safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter(0)

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

func NewCounter(val int) *Counter {
	counter := &Counter{}
	counter.val = val
	return counter
}
