package main

import "testing"

func TestHello(t *testing.T) {
	got := "Hallo welt"
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
