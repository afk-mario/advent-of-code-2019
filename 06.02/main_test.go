package main

import (
	"testing"
)

// TestCountBasic should match the requiremeents
func TestCountBasic(t *testing.T) {
	_o = opts{
		"./input.txt",
		true,
	}
	got := valueByInput(_o.filePath)
	want := 4

	if got != want {
		t.Errorf("Count orbits = %d; want %d", got, want)
	}
}

// TestCountBasic should match the requiremeents
func TestCountInputFile(t *testing.T) {
	_o = opts{
		"./final.txt",
		false,
	}
	got := valueByInput(_o.filePath)
	want := 484

	if got != want {
		t.Errorf("Count orbits = %d; want %d", got, want)
	}
}

// TestCountBruno should match the requiremeents
func TestCountBruno(t *testing.T) {
	_o = opts{
		"./bruno.txt",
		false,
	}
	got := valueByInput(_o.filePath)
	want := 472

	if got != want {
		t.Errorf("Count orbits = %d; want %d", got, want)
	}
}
