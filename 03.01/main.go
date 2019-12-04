package main

import (
	"advent/utils"
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type vector2 struct {
	x int
	y int
}

func (a *vector2) add(b vector2) vector2 {
	return vector2{a.x + b.x, a.y + b.y}
}

func (a *vector2) equals(b vector2) bool {
	return a.x == b.x && a.y == b.y
}

func (a *vector2) mDistance() int {
	c := vector2{0, 0}
	return utils.Abs(a.x-c.x) + utils.Abs(a.y-c.y)
}

func (a *vector2) cross(b vector2) int {
	return a.x*b.y - b.x*a.y
}

type line struct {
	a vector2
	b vector2
}

func (l *line) isPointOn(p vector2) bool {
	aTmp := line{vector2{0, 0}, vector2{l.b.x - l.a.x, l.b.y - l.a.y}}
	bTmp := vector2{p.x - l.a.x, p.y - l.a.y}
	r := aTmp.b.cross(bTmp)
	return utils.Abs(r) == 0
}

func (l *line) isPointRight(b vector2) bool {
	aTmp := line{vector2{0, 0}, vector2{l.b.x - l.a.x, l.b.y - l.a.y}}
	bTmp := vector2{b.x - l.a.x, b.y - l.a.y}
	return aTmp.b.cross(bTmp) < 0
}

func (l *line) touchOrCross(b line) bool {
	return l.isPointOn(b.a) || l.isPointOn(b.b) || (l.isPointRight(b.a) != l.isPointRight(b.b))
}

func (l *line) intersects(b line) bool {
	return l.touchOrCross(b) && b.touchOrCross(*l)
}

func drawCables(a []vector2, b []vector2, c []vector2) {

	allPos := append(a, b...)
	sortByX(allPos)
	maxX := allPos[len(allPos)-1].x + 1
	sortByY(allPos)
	maxY := allPos[len(allPos)-1].y + 1

	fmt.Println("\n")

	// linesA := positionsToLineSegments(positionsA)
	// linesB := positionsToLineSegments(positionsB)

	// empty grid
	grid := make([][]int, maxY)
	for i := 0; i < maxY; i++ {
		grid[i] = make([]int, maxX)
		for j := 0; j < maxX; j++ {
			grid[i][j] = 0 // .
		}
	}

	// add cables from a
	for _, pos := range a {
		grid[pos.y][pos.x] = 1 // +
	}

	// add cables from a
	for _, pos := range b {
		grid[pos.y][pos.x] = 10 // +
	}

	for _, pos := range c {
		grid[pos.y][pos.x] = 20 // x
	}

	grid[0][0] = -1

	var acc strings.Builder
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			acc.WriteString(" ")
			if grid[i][j] == -1 {
				acc.WriteString("o")
			} else if grid[i][j] == 0 {
				acc.WriteString(".")
			} else if grid[i][j] == 1 {
				acc.WriteString("a")
				// acc.WriteString(strconv.Itoa(grid[i][j]))
			} else if grid[i][j] == 10 {
				acc.WriteString("b")
			} else if grid[i][j] == 20 {
				acc.WriteString("x")
			}
			acc.WriteString(" ")
		}
		acc.WriteString("\n")
	}
	fmt.Println(acc.String())
}

func (l *line) getIntersection(b line) vector2 {
	x, y := -1, -1

	if l.a.x == l.b.x {
		x = l.a.x
	} else if b.a.x == b.b.x {
		x = b.a.x
	}

	if l.a.y == l.b.y {
		y = l.a.y
	} else if b.a.y == b.b.y {
		y = b.a.y
	}

	return vector2{x, y}
}

func sortByMDistance(s []vector2) {
	sort.Slice(s, func(i, j int) bool {
		d1 := s[i].mDistance()
		d2 := s[j].mDistance()
		return d1 < d2
	})
}

func sortByX(s []vector2) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].x < s[j].x
	})
}

func sortByY(s []vector2) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].y < s[j].y
	})
}

func main() {
	inputPath := flag.String("input", "./input.txt", "input file path")
	// verbose := flag.Bool("v", false, "Extra prints")
	flag.Parse()

	if len(flag.Args()) > 0 {
		valueByArg(flag.Args()[0], flag.Args()[1])
	}

	fmt.Println(*inputPath)
	valueByInput(*inputPath)
}

func valueByInput(file string) {
	fmt.Println("\nPrinting value by file ---", file)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	input, err := utils.ReadStrings(f, bufio.ScanWords)
	if err != nil {
		panic(err)
	}

	valueByArg(input[0], input[1])
}

func valueByArg(_a string, _b string) {
	f := strings.NewReader(_a)
	a, err := utils.ReadStrings(f, utils.SplitComma)

	if err != nil {
		panic(err)
	}

	f = strings.NewReader(_b)
	b, err := utils.ReadStrings(f, utils.SplitComma)

	fmt.Println("\nPrinting value by arg ---", a, b)
	fmt.Println("")
	if err != nil {
		panic(err)
	}

	positionsA := instructionsToPositions(a)
	positionsB := instructionsToPositions(b)

	fmt.Println("\n")

	collisions := findCollisions(positionsA, positionsB)
	// drawCables(positionsA, positionsB, collisions)
	sortByMDistance(collisions)
	fmt.Println("\nResult:", collisions[0].mDistance())
}

func findCollisions(a []vector2, b []vector2) []vector2 {
	linesA := positionsToLineSegments(a)
	linesB := positionsToLineSegments(b)
	collisions := make([]vector2, 0)
	for _, lA := range linesA {
		for _, lB := range linesB {
			if lA.intersects(lB) {
				intersection := lA.getIntersection(lB)
				// fmt.Println("it intersects", lA, lB, intersection)
				collisions = append(collisions, intersection)
			}
		}
	}
	return collisions
}

func instructionsToPositions(input []string) []vector2 {
	arr := make([]vector2, 1)
	arr[0] = vector2{0, 0}

	for i, instruction := range input {
		pos := arr[i].add(instructionToVector(instruction))
		arr = append(arr, pos)
		// fmt.Println(pos)
	}
	// fmt.Println("----")
	return arr
}

func positionsToLineSegments(positions []vector2) []line {
	lines := make([]line, 0)
	for i := 1; i < len(positions)-1; i++ {
		lines = append(lines, line{positions[i], positions[i+1]})
	}
	return lines
}

func instructionToVector(instruction string) vector2 {
	dx := 0
	dy := 0
	if instruction[0] == 'U' {
		dy = 1
	} else if instruction[0] == 'D' {
		dy = -1
	} else if instruction[0] == 'R' {
		dx = 1
	} else if instruction[0] == 'L' {
		dx = -1
	}
	d, err := strconv.Atoi(instruction[1:len(instruction)])
	if err != nil {
		panic(err)
	}
	return vector2{dx * d, dy * d}
}
