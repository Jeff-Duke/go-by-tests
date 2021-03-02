package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jeff")
	want := "Hello, Jeff"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
