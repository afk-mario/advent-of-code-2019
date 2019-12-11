package vect

import (
	"testing"
)

// TestAngle ...
func TestAngle(t *testing.T) {
	a := Vector2{0, 0}
	b := Vector2{-10, -10}
	got := a.Angle(b)
	want := float64(0)
	if got != 0 {
		t.Errorf("Angle between %v - %v = %g; want %g", a, b, got, want)
	}
}
