package arrsum

import "testing"

func TestSum(t *testing.T) {
	assertSum := func(t testing.TB, got, want int, numbers [5]int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	}

	t.Run("adds a list of numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertSum(t, got, want, numbers)
	})

	t.Run("adds any list of 5 numbers", func(t *testing.T) {
		numbers := [5]int{2, 4, 6, 8, 10}
		got := Sum(numbers)
		want := 30

		assertSum(t, got, want, numbers)
	})

}
