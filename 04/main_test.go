package main

import (
	"testing"
)

// TestComply should match the requiremeents
func TestCompy(t *testing.T) {
	got := doesComply(111111)
	if got != true {
		t.Errorf("DoesComply = %t; want true", got)
	}

	got = doesComply(223450)
	if got != false {
		t.Errorf("DoesComply = %t; want false", got)
	}

	got = doesComply(123789)
	if got != false {
		t.Errorf("DoesComply = %t; want false", got)
	}

	arr := []int{234444, 234445, 234446, 234447, 234448, 234449, 234455,
		234456, 234457, 234458, 234459, 234466, 234467, 234468,
		234469, 234477, 234478, 234479, 234488, 234489, 234499,
		234555, 234556, 234557, 234558, 234559, 234566, 234577,
		234588, 234599, 234666, 234667, 234668, 234669, 234677,
		234688, 234699, 234777, 234778, 234779, 234788, 234799,
		234888, 234889, 234899, 234999, 235555, 235556, 235557,
		235558}

	for _, n := range arr {
		got = doesComply(n)
		if got != true {
			t.Errorf("DoesComply = %t; want true", got)
		}

	}
}
