package main

import (
	"advent/utils"
	"advent/vect"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func drawCables(a []vect.Vector2, b []vect.Vector2, c []vect.Vector3) {

	allPos := append(a, b...)
	vect.SortByX(allPos)
	maxX := allPos[len(allPos)-1].X + 1
	vect.SortByY(allPos)
	maxY := allPos[len(allPos)-1].Y + 1

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
		grid[pos.Y][pos.X] = 1 // +
	}

	// add cables from a
	for _, pos := range b {
		grid[pos.Y][pos.X] = 10 // +
	}

	for _, pos := range c {
		grid[pos.Y][pos.X] = 20 // x
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
	vect.SortByZ(collisions)
	fmt.Println("\nResult:", collisions[0].Z)
}

func findCollisions(a []vect.Vector2, b []vect.Vector2) []vect.Vector3 {
	linesA := positionsToLineSegments(a)
	linesB := positionsToLineSegments(b)
	collisions := make([]vect.Vector3, 0)
	for i, lA := range linesA {
		for j, lB := range linesB {
			if lA.Intersects(lB) {
				intersection := lA.GetIntersection(lB)
				steps := getSteps(a, linesA, i-1, intersection)
				steps += getSteps(b, linesB, j-1, intersection)
				// fmt.Printf("it Intersects i: %d j: %d [%d, %d] steps: %d \n", i, j, intersection.X, intersection.Y, steps)
				collisions = append(collisions, vect.Vector3{intersection.X, intersection.Y, steps})
			}
		}
	}
	return collisions
}

func getSteps(pos []vect.Vector2, lines []vect.Line, at int, point vect.Vector2) int {
	// fmt.Println("\ngetSteps at: ", at)
	first := vect.Line{pos[0], pos[1]}
	steps := first.Magnitude()
	lineAToInter := vect.Line{lines[at].A, point}
	lineBToInter := vect.Line{lines[at].B, point}

	for i := 0; i <= at; i++ {
		steps += lines[i].Magnitude()
	}

	if lineAToInter.Magnitude() < lineBToInter.Magnitude() {
		steps += lineAToInter.Magnitude()
	} else {
		steps += lineBToInter.Magnitude()
	}

	// fmt.Printf("steps: %d \n", steps)
	return steps
}

func instructionsToPositions(input []string) []vect.Vector2 {
	arr := make([]vect.Vector2, 1)
	arr[0] = vect.Vector2{0, 0}

	for i, instruction := range input {
		pos := arr[i].Add(instructionToVector(instruction))
		arr = append(arr, pos)
		// fmt.Println(pos)
	}
	// fmt.Println("----")
	return arr
}

func positionsToLineSegments(positions []vect.Vector2) []vect.Line {
	lines := make([]vect.Line, 0)
	for i := 1; i < len(positions)-1; i++ {
		lines = append(lines, vect.Line{positions[i], positions[i+1]})
	}
	return lines
}

func instructionToVector(instruction string) vect.Vector2 {
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
	return vect.Vector2{dx * d, dy * d}
}
