package intcode

import (
	"advent/logger"
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// TestGetOperation ...
func TestGetOperation(t *testing.T) {
	from := "1001"
	got := getOperation(from)
	want := 1
	if got != want {
		t.Errorf("Operation extracted from %s = %d; want %d", from, got, want)
	}

	from = "11102"
	got = getOperation(from)
	want = 2
	if got != want {
		t.Errorf("Operation extracted from %s = %d; want %d", from, got, want)
	}

	from = "103"
	got = getOperation(from)
	want = 3
	if got != want {
		t.Errorf("Operation extracted from %s = %d; want %d", from, got, want)
	}

	from = "10208"
	got = getOperation(from)
	want = 8
	if got != want {
		t.Errorf("Operation extracted from %s = %d; want %d", from, got, want)
	}
}

func TestGetModes(t *testing.T) {
	l := 3

	from := "1001"
	got := getModes(from, l)
	want := []int{0, 1, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Operation extracted from %s = %d; want %d", from, got, want)
	}

	from = "11102"
	got = getModes(from, l)
	want = []int{1, 1, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Operation extracted from %s = %d; want %d", from, got, want)
	}

	from = "103"
	got = getModes(from, l)
	want = []int{1, 0, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Operation extracted from %s = %d; want %d", from, got, want)
	}

	from = "10108"
	got = getModes(from, l)
	want = []int{1, 0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Operation extracted from %s = %d; want %d", from, got, want)
	}
}

func TestDay0701(t *testing.T) {

	logger.EnableDebug()

	a := []int{4, 3, 2, 1, 0}
	// p := prmt.New(prmt.IntSlice(a))
	r := []int{}

	// for p.Next() {
	result := 0
	fmt.Printf("\n\n Permutation %v \n\n", a)
	for _, perm := range a {
		fmt.Printf("result %d \n\n\n", result)
		intcode := Intcode{
			program: []int{
				3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0,
			},
			inputArr: []int{perm, result},
		}
		result = intcode.DoOperation()

	}
	r = append(r, result)
	// }

	sort.Ints(r)
	got := r[len(r)-1]
	fmt.Println(r)

	want := 43210
	if got != want {
		t.Errorf("Operate = %d; want %d", got, want)
	}

}
