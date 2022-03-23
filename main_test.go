package main

import (
	"lemin/function"
	"testing"
)

func TestDistance(t *testing.T) {
	got := function.Distance(9, 5, 16, 5)
	want := 7

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	got = function.Distance(9, 5, 1, 5)
	want = 8

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	got = function.Distance(23, 3, 16, 3)
	want = 7

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	got = function.Distance(16, 5, 16, 3)
	want = 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	got = function.Distance(9, 3, 1, 5)
	want = 8

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	got = function.Distance(4, 8, 1, 5)
	want = 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func TestValidateFile(t *testing.T) {
	s := []string{"example00.txt", "example01.txt", "example02.txt", "example03.txt", "example04.txt", "example05.txt", "example06.txt", "example07.txt", "example08.txt", "example09.txt", "example10.txt", "badexample00.txt", "badexample01.txt"}
	f := []bool{true, true, true, true, true, true, true, true, false, false, false, false, false} // the badexample01.txt haven't take into consideration yet
	for i := range s {
		got, _ := function.ValidateFile("example/" + s[i])
		want := f[i]
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}
}
