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
	path  string
	debug bool
	w     int
	h     int
}

var _o opts

func main() {
	path := flag.String("input", "./input.txt", "input file path")
	debug := flag.Bool("debug", false, "debug text")
	w := flag.Int("w", 25, "width of image")
	h := flag.Int("h", 6, "height of image")
	flag.Parse()

	_o = opts{
		*path,
		*debug,
		*w,
		*h,
	}

	if _o.debug {
		logger.EnableDebug()
	}

	valueByInput(_o.path)
}

func valueByInput(file string) {
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

	layers := makeLayers(input[0], _o.w, _o.h)

	if _o.debug {
		// drawLayers(layers, _o.w, _o.h)
	}

	fmt.Println("++++++++++++++last\n")
	drawLayerWH(layers[len(layers)-1], _o.w, _o.h)
	fmt.Println("++++++++++++++first\n")
	drawLayerWH(layers[0], _o.w, _o.h)
	fmt.Println("++++++++++++++\n")

	result := mergeLayers(layers)

	fmt.Println("__________________\n")
	drawLayerWH(result, _o.w, _o.h)
	fmt.Println("__________________\n")
}

func makeLayers(s string, w, h int) [][]int {
	arr := make([][]int, 0)
	for i, c := range s {
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

func drawLayerWH(layer []int, w, h int) {
	for i, e := range layer {
		if i%w == 0 && i != 0 {
			// fmt.Printf("\n")
			fmt.Printf("\n")
		}
		fmt.Printf(drawPixel(e))
	}
	fmt.Printf("\n")
}

func drawLayers(layers [][]int, w, h int) {
	for _, layer := range layers {
		drawLayerWH(layer, w, h)
	}
}

func drawLayer(layer []int) {
	for _, e := range layer {
		fmt.Printf(drawPixel(e))
	}
	fmt.Printf("\n")
}

func drawPixel(i int) string {
	// c = " ▀"
	c := "  "
	switch i {
	case 0:
		c = " ."
		break
	case 1:
		c = " ▀"
	}
	return c
}

func mergeLayers(layers [][]int) []int {
	last := layers[len(layers)-1]
	arr := append(make([]int, 0), last...)

	for i := len(layers) - 1; i >= 0; i-- {
		layer := layers[i]
		for j, c := range layer {
			if c != 2 {
				arr[j] = c
			}
		}
		// fmt.Println("__________________\n")
		// drawLayerWH(layer, _o.w, _o.h)
	}

	return arr
}
