package main

import (
	"advent/logger"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type opts struct {
	filePath string
	debug    bool
}

var _o opts

func main() {
	inputPath := flag.String("input", "./input.txt", "input file path")
	debug := flag.Bool("debug", false, "debug text")
	flag.Parse()

	_o = opts{
		*inputPath,
		*debug,
	}

	if _o.debug {
		logger.EnableDebug()
	}

	a := 234208
	b := 765869
	nums := findNumbers(a, b)
	logger.Debug("%d - %d", a, b)
	logger.Debug("%v", nums)
	fmt.Println("Result: ", len(nums))
}

// It is a six-digit number.
// The value is within the range given in your puzzle input.
// Two adjacent digits are the same (like 22 in 122345).
// Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
// the two adjacent matching digits are not part of a larger group of matching digits.

func findNumbers(a int, b int) []int {

	logger.Debug("%d - %d", a, b)
	result := make([]int, 0)
	d := b - a
	for i := 0; i < d; i++ {

		n := a + i
		if doesComply(n) {
			result = append(result, n)
		}

	}
	return result
}

func doesComply(a int) bool {
	return isSixDigit(a) && neverDecreases(a) && hasAdjacent(a) && hasAdjecentTwo(a)
}

func isSixDigit(a int) bool {
	return a/100000 > 0
}

func hasAdjacent(a int) bool {
	b := strconv.Itoa(a)
	for i := 0; i < len(b)-1; i++ {
		if b[i] == b[i+1] {
			return true
		}
	}
	return false
}

func hasAdjecentTwo(a int) bool {

	b := strconv.Itoa(a)
	arr := make([]strings.Builder, len(b))
	var bu strings.Builder
	arr = append(arr, bu)
	j := 0
	for i := 0; i < len(b)-1; i++ {
		if b[i] == b[i+1] {
			arr[j].WriteByte(b[i])
		} else {
			var acc strings.Builder
			j++
			arr = append(arr, acc)
		}
	}

	for _, s := range arr {
		if s.Len() == 1 {
			return true
		}
	}
	return false
}

func neverDecreases(a int) bool {
	b := strconv.Itoa(a)
	for i := 0; i < len(b)-1; i++ {
		// x, _ := strconv.Atoi(string(b[i]))
		// y, _ := strconv.Atoi(string(b[i+1]))
		x := b[i]
		y := b[i+1]
		if x > y {
			return false
		}
	}
	return true
}
