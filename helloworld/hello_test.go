package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Connor")
	want := "Hello, Connor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
