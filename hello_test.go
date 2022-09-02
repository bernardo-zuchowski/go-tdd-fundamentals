package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("John Roe")
	want := "Hello, John Roe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}