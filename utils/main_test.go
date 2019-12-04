package utils

import (
	"testing"
)

// TestAbs should be -1 = 1 1 = 1
func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

// TestPow should be 2 ^ 10 = 1024
func TestPow(t *testing.T) {
	got := Pow(2, 10)
	if got != 1024 {
		t.Errorf("Pow(2,10) = %d; want 1024", got)
	}
}
