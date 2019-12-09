package main

import (
	"testing"
)

func TestCountBasic(t *testing.T) {
	_o = opts{
		"./test.txt",
		true,
	}
	got := valueByInput(_o.filePath)
	want := 4

	if got != want {
		t.Errorf("Count orbits = %d; want %d", got, want)
	}
}

func TestCountInput(t *testing.T) {
	_o = opts{
		"./input.txt",
		false,
	}
	got := valueByInput(_o.filePath)
	want := 1920

	if got != want {
		t.Errorf("Count orbits = %d; want %d", got, want)
	}
}
