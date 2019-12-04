package main

import (
	"testing"
)

// TestComply should match the requiremeents
func TestCompy(t *testing.T) {
	n := 112233
	got := doesComply(n)
	if got != true {
		t.Errorf("DoesComply %d = %t; want true", n, got)
	}

	n = 123444
	got = doesComply(n)
	if got != false {
		t.Errorf("DoesComply %d = %t; want false", n, got)
	}

	n = 111122
	got = doesComply(n)
	if got != true {
		t.Errorf("DoesComply %d = %t; want true", n, got)
	}
}
