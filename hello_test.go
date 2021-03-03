package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectGreeting := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Jeff", "")
		want := "Hello, Jeff"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("says hello to the world if no name given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectGreeting(t, got, want)
	})
}
