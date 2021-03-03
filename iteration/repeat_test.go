package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertRepeated := func(t testing.TB, expected, repeated string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeats a single char", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		assertRepeated(t, expected, repeated)
	})

	t.Run("repeats any char", func(t *testing.T) {
		repeated := Repeat("d")
		expected := "ddddd"

		assertRepeated(t, expected, repeated)
	})

	t.Run("repeats any string", func(t *testing.T) {
		repeated := Repeat("ha")
		expected := "hahahahaha"

		assertRepeated(t, expected, repeated)
	})
}
