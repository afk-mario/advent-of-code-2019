package main

import (
	"advent/logger"
	"advent/utils"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type opts struct {
	filePath string
	debug    bool
}

type planet struct {
	data   string
	orbits []*planet
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

	result := valueByInput(_o.filePath)
	fmt.Println("Result: ", result)
}

func valueByInput(file string) int {
	fmt.Println("\nPrinting value by file ---", file)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	input, err := utils.ReadStrings(f, bufio.ScanLines)
	if err != nil {
		panic(err)
	}

	planets := makeOrbits(input)

	san := findPlanet(planets, "SAN")
	you := findPlanet(planets, "YOU")
	a := pathToRoot(san)
	b := pathToRoot(you)

	logger.Debug("san -> root %d", len(a))
	logger.Debug("you -> root %d", len(b))
	return mergePaths(a, b)
}

func countOrbits(arr []*planet) int {
	cOrbits := 0

	for _, p := range arr {
		cOrbits += countP(p)
	}

	return cOrbits
}

func countP(p *planet) int {
	acc := 0
	for _, a := range p.orbits {
		acc++
		acc += countP(a)
	}
	return acc
}

func makeOrbits(s []string) []*planet {
	arr := make([]*planet, 0)
	instructions := make([]string, 0)

	for _, e := range s {
		f := strings.NewReader(e)
		a, err := utils.ReadStrings(f, utils.SplitParenthesis)
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, a...)

	}

	instructions = utils.RemoveDuplicatesUnordered(instructions)

	for _, e := range instructions {
		node := planet{e, make([]*planet, 0)}
		arr = append(arr, &node)
	}

	for _, e := range s {
		f := strings.NewReader(e)
		a, err := utils.ReadStrings(f, utils.SplitParenthesis)
		if err != nil {
			panic(err)
		}

		// fmt.Println("Adding Orbits", e)
		p := findPlanet(arr, a[1])
		o := findPlanet(arr, a[0])

		has := hasPlanet(p.orbits, o)
		if !has {
			p.orbits = append(p.orbits, o)
		}

	}

	if _o.debug {
		debugPlanets(arr)
	}

	return arr
}

func debugPlanets(arr []*planet) {
	for _, e := range arr {
		fmt.Println("node", e.data, len(e.orbits))
		for _, o := range e.orbits {
			fmt.Println("	orbits", o.data)
		}
		fmt.Println("")
	}
}

func hasPlanet(orbits []*planet, p *planet) bool {
	for _, o := range orbits {
		if o.data == p.data {
			return true
		}
	}
	return false
}

func findPlanet(arr []*planet, data string) *planet {
	for _, a := range arr {
		if a.data == data {
			return a
		}
	}
	return nil
}

func findInOrbit(arr []*planet, data string) *planet {
	for _, a := range arr {
		for _, b := range a.orbits {
			if b.data == data {
				return b
			}
		}
	}
	return nil
}

func pathToRoot(p *planet) []*planet {
	arr := make([]*planet, 0)
	logger.Debug("[%s] | %d", p.data, len(p.orbits))
	if len(p.orbits) < 1 {
		return arr
	}
	n := p.orbits[0]
	arr = append(arr, n)
	arr = append(arr, pathToRoot(n)...)
	return arr
}

func mergePaths(a []*planet, b []*planet) int {
	acc := 0
	intersection := findIntersection(a, b)
	logger.Debug("intersection: [%s]", intersection.data)

	for _, p := range a {
		if p == intersection {
			break
		}
		acc++
	}

	logger.Debug("%s -> [%s] %d", a[0].data, intersection.data, len(a))

	for _, p := range b {
		if p == intersection {
			break
		}
		acc++
	}

	logger.Debug("%s -> [%s] %d", b[0].data, intersection.data, len(a))

	// acc += len(pathToRoot(intersection)) - 1

	return acc
}

func findIntersection(a []*planet, b []*planet) *planet {
	for _, _a := range b {
		for _, _b := range a {
			if _a.data == _b.data {
				return _a
			}
		}
	}
	return nil
}
