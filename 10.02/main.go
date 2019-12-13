package main

import (
	"advent/logger"
	"advent/utils"
	"advent/vect"
	"bufio"
	"flag"
	"fmt"
	"os"
)

type opts struct {
	path  string
	debug bool
	r     int
}

type board struct {
	w int
	h int
}

var _o opts
var _b board

func main() {
	path := flag.String("path", "./input.txt", "input file path")
	debug := flag.Bool("d", false, "debug text")
	revolutions := flag.Int("r", 200, "number of revolutions")
	flag.Parse()

	_o = opts{
		*path,
		*debug,
		*revolutions,
	}

	if _o.debug {
		logger.EnableDebug()
	}

	result := resultByFile(_o.path)
	fmt.Println("Result: ", result)
	fmt.Println("Result: ", result.X*100+result.Y)
}

func resultByFile(file string) vect.Vector2 {
	fmt.Println("\nPrinting value by file ---", file)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	input, err := utils.ReadStrings(f, bufio.ScanLines)

	_b.h = len(input)
	_b.w = len(input[0])

	if err != nil {
		panic(err)
	}

	grid := makeGrid(input)
	connections := makeConnections(grid)
	vect.SortByZ(connections)

	laser := connections[len(connections)-1].ToVec2()

	angles := makeAngles(grid, laser)
	vect.SortByZ(angles)
	gAngles := groupAngles(angles, laser)

	if _o.debug {
		fmt.Println("angles \n")
		for _, i := range angles {
			fmt.Println(i)
		}
		fmt.Println("\n gAngles \n")
		for _, i := range gAngles {
			fmt.Println(i)
		}
		fmt.Println("Laser: ", laser)
	}

	destroyed := getDesroyed(gAngles, laser, _o.r)

	if _o.debug {
		drawBoard(input, laser, destroyed)
	}

	var last vect.Vector2

	if len(destroyed) > 0 {
		last = destroyed[len(destroyed)-1]
	}

	return last
}

func isElement(arr []vect.Vector2, i, j int) bool {
	tmp := vect.Vector2{j, i}
	for _, e := range arr {
		if e.Equals(tmp) {
			return true
		}
	}
	return false
}

func drawBoard(arr []string, laser vect.Vector2, destroyed []vect.Vector2) {
	fmt.Println("")
	fmt.Println("")

	fmt.Printf("    ")

	for i := range arr[0] {
		fmt.Printf(" %02d ", i)
	}
	fmt.Printf("\n   ┌")
	for range arr[0] {
		fmt.Printf("────")
	}

	fmt.Printf("─┐\n")
	for i, a := range arr {
		fmt.Printf("%02d │ ", i)
		for j, b := range a {
			char := string(b)
			isD := isElement(destroyed, i, j)
			if isD {
				i := getIndexOfElementP(destroyed, vect.Vector2{j, i})
				char = fmt.Sprintf("%3d ", i+1)
			} else if laser.X == j && laser.Y == i {
				char = " \033[1;31m⌧\033[0m  "
			} else {
				if char == "#" {
					char = "⌬"
				}
				char = fmt.Sprintf(" %s  ", char)
			}
			fmt.Printf("%s", char)
		}
		fmt.Printf("│ %02d", i)
		fmt.Printf("\n")
	}
	fmt.Printf("   └")
	for range arr[0] {
		fmt.Printf("────")
	}
	fmt.Printf("─┘\n    ")
	for i := range arr[0] {
		fmt.Printf(" %02d ", i)
	}
	fmt.Println("\n")
}

func makeGrid(arr []string) []vect.Vector2 {
	result := make([]vect.Vector2, 0)
	for i, a := range arr {
		for j, b := range a {
			if b == '#' {
				result = append(result, vect.Vector2{j, i})
			}
		}
	}
	return result
}

func makeAngles(arr []vect.Vector2, laser vect.Vector2) []vect.Vector3 {
	result := make([]vect.Vector3, 0)
	for _, e := range arr {
		if !e.Equals(laser) {
			v := vect.Vector3{e.X, e.Y, e.Angle(laser)}
			result = append(result, v)
		}
	}
	return result
}

func groupAngles(arr []vect.Vector3, laser vect.Vector2) [][]vect.Vector3 {
	result := make([][]vect.Vector3, 0)
	result = append(result, make([]vect.Vector3, 0))
	acc := 0

	for i, e := range arr {
		result[acc] = append(result[acc], e)
		if i < len(arr)-1 && e.Z != arr[i+1].Z {
			result = append(result, make([]vect.Vector3, 0))
			acc++
		}
	}

	for i, x := range result {
		for j, y := range x {
			v2 := y.ToVec2()
			l := vect.Line{v2, laser}
			z := l.Magnitude()
			result[i][j] = vect.Vector3{v2.X, v2.Y, z}
		}

		vect.SortByZ(result[i])
	}

	return result
}

func makeConnections(arr []vect.Vector2) []vect.Vector3 {
	result := make([]vect.Vector3, 0)
	for i, a := range arr {
		acc := 0
		for j := 0; j < len(arr); j++ {
			b := arr[j]
			if i != j {
				can := canView(arr, a, b)
				test := vect.Vector2{5, 8}
				if can {
					acc++
				} else {
					if a.Equals(test) {
						logger.Debug("%v -|- %v", a, b)
					}
				}
			}
		}
		result = append(result, vect.Vector3{a.X, a.Y, acc})
	}
	return result
}

func doesContain(arr []vect.Vector2, a vect.Vector2) bool {
	for _, e := range arr {
		if e.Equals(a) {
			return true
		}
	}
	return false
}

func getDesroyed(arr [][]vect.Vector3, laser vect.Vector2, revolutions int) []vect.Vector2 {
	ded := make([]vect.Vector2, 0)
	i := 0

	for len(ded) < revolutions {

		// fmt.Println("ded0----", revolutions, len(ded), i)
		if len(arr[i]) > 0 {
			last := len(arr[i]) - 1
			last = 0
			// fmt.Println("ded----", arr[i][last], len(arr[i]))
			v := arr[i][last].ToVec2()
			ded = append(ded, v)
			arr[i] = remove(arr[i], arr[i][last])
			// fmt.Println("ded2----", arr[i], len(arr[i]))
			i++
		} else {
			i++
		}

		if i >= len(arr) {
			i = 0
		}

		// if len(ded) >= 35 {
		// 	return ded
		// }
	}

	return ded
}

func getIndexOfElement(arr []vect.Vector2, f *vect.Vector2) int {
	z := vect.Vector2{f.X, f.Y}
	for i, e := range arr {
		if e.Equals(z) {
			return i
		}
	}
	return -1
}

func getIndexOfElementP(arr []vect.Vector2, f vect.Vector2) int {
	z := vect.Vector2{f.X, f.Y}
	for i, e := range arr {
		if e.Equals(z) {
			return i
		}
	}
	return -1
}

func getFirstAsterioidInLine(arr []vect.Vector2, l vect.Line) *vect.Vector2 {
	collisions := getAllBetweenLine(arr, l)
	for _, collision := range collisions {
		if canView(arr, l.A, collision) {
			// logger.Debug("first in line %v", collision)
			return &collision
		}
	}

	return nil
}

func getAllBetweenLine(arr []vect.Vector2, l vect.Line) []vect.Vector2 {
	result := make([]vect.Vector2, 0)
	for _, e := range arr {
		if l.IsBetween(e) && !e.Equals(l.A) {
			result = append(result, e)
		}
	}
	return result
}

func canView(arr []vect.Vector2, a, b vect.Vector2) bool {
	line := vect.Line{a, b}
	for _, e := range arr {
		nLine := vect.Line{a, e}
		onPoint := line.IsBetween(e)
		notAoB := !e.Equals(a) && !e.Equals(b)
		isBetween := line.Magnitude() > nLine.Magnitude()

		if onPoint && notAoB && isBetween {
			test := vect.Vector2{5, 8}
			if a.Equals(test) {
				logger.Debug("%v -> %v | %v | %v", a, b, e, onPoint)
			}
			return false
		}
	}
	return true
}

func remove(items []vect.Vector3, item vect.Vector3) []vect.Vector3 {
	newitems := []vect.Vector3{}

	for _, i := range items {
		if !i.Equals(item) {
			newitems = append(newitems, i)
		}
	}

	return newitems
}
