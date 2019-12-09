package main

import (
	"advent/logger"
	"advent/utils"
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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

	layers := makeLayers(input[0], 25, 6)
	// layers := makeLayers(input[0], 2, 3)
	layers = sortLayers(layers)
	if _o.debug {
		for _, e := range layers {
			fmt.Println(e)
		}
	}
	logger.Debug("Len layers %d", len(layers))
	if _o.debug {
		fmt.Println(layers[0])
	}
	a := ammountOf(layers[0], 1)
	b := ammountOf(layers[0], 2)
	return a * b
}

func makeLayers(s string, w, h int) [][]int {
	logger.Debug("w: %d h: %d len: %d s:%s|", w, h, len(s), s)
	arr := make([][]int, 0)
	for i, c := range s {
		logger.Debug("i: %d m: %d res: %d", i, w*h, i%(w*h))
		if i%(w*h) == 0 {
			arr = append(arr, make([]int, 0))
		}
		arr[len(arr)-1] = append(arr[len(arr)-1], int(c-'0'))
	}
	return arr
}

func sortLayers(s [][]int) [][]int {
	sort.SliceStable(s, func(i, j int) bool {
		return ammountOf(s[i], 0) < ammountOf(s[j], 0)
	})
	return s
}

func ammountOf(s []int, c int) int {
	acc := 0
	for _, e := range s {
		if e == c {
			acc++
		}
	}
	return acc
}
