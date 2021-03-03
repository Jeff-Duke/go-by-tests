package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertRepeated := func(t testing.TB, expected, repeated string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeats a single char", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertRepeated(t, expected, repeated)
	})

	t.Run("repeats any char", func(t *testing.T) {
		repeated := Repeat("d", 5)
		expected := "ddddd"

		assertRepeated(t, expected, repeated)
	})

	t.Run("repeats any string", func(t *testing.T) {
		repeated := Repeat("ha", 5)
		expected := "hahahahaha"

		assertRepeated(t, expected, repeated)
	})

	t.Run("repeats by count provided", func(t *testing.T) {
		repeated := Repeat("lol", 3)
		expected := "lollollol"

		assertRepeated(t, expected, repeated)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
