package main

import (
	"advent/utils"
	"bufio"
	"strings"
	"testing"
)

// TestComply should match the requiremeents
func TestCompy(t *testing.T) {
	input := (`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`)

	f := strings.NewReader(input)
	a, err := utils.ReadStrings(f, bufio.ScanLines)

	if err != nil {
		t.Errorf("Couldn't read string %s", input)
	}

	planets := makeOrbits(a)
	got := countOrbits(planets)

	if got != 42 {
		t.Errorf("Count orbits = %d; want 42", got)
	}
}
