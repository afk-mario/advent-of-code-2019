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
}

var _o opts

func main() {
	path := flag.String("path", "./input.txt", "input file path")
	debug := flag.Bool("d", false, "debug text")
	flag.Parse()

	_o = opts{
		*path,
		*debug,
	}

	if _o.debug {
		logger.EnableDebug()
	}

	result := resultByFile(_o.path)
	fmt.Println("Result: ", result)
}

func resultByFile(file string) int {
	fmt.Println("\nPrinting value by file ---", file)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	input, err := utils.ReadStrings(f, bufio.ScanLines)

	if _o.debug {
		fmt.Println("")
		fmt.Printf("    ")
		for i := range input {
			fmt.Printf(" %02d ", i)
		}
		fmt.Printf("\n    ")
		for range input {
			fmt.Printf("────")
		}
		fmt.Printf("\n")
		for i, a := range input {
			fmt.Printf("%02d │ ", i)
			for _, b := range a {
				fmt.Printf(" %s  ", string(b))
			}
			fmt.Printf("\n")
		}
		fmt.Println("")
	}

	if err != nil {
		panic(err)
	}

	grid := makeGrid(input)
	connections := makeConnections(grid)
	vect.SortByZ(connections)

	if _o.debug {
		for _, e := range connections {
			fmt.Println(e)
		}
	}

	return connections[len(connections)-1].Z
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
