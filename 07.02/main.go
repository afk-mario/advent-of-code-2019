package main

import (
	"advent/logger"
	"advent/utils"
	"flag"
	"fmt"
	"os"
	"sort"

	prmt "github.com/gitchander/permutation"
)

type opts struct {
	path  string
	debug bool
	input int
}

var _o opts

func main() {
	path := flag.String("path", "./input.txt", "input file path")
	debug := flag.Bool("d", false, "debug text")
	input := flag.Int("i", 1, "input code")
	flag.Parse()

	_o = opts{
		*path,
		*debug,
		*input,
	}

	if _o.debug {
		logger.EnableDebug()
	}

	result := resultByFile(_o.path)
	fmt.Println("\n______")
	fmt.Println("result", result)
	fmt.Println("______\n")
}

func resultByFile(file string) int {
	fmt.Println("\nPrinting value by file ---", file)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	program, err := utils.ReadInts(f, utils.SplitComma)
	if err != nil {
		panic(err)
	}

	if _o.debug {
		fmt.Println("\nInput:\n", utils.FmtIntSlice(program))
	}

	a := []int{5, 6, 7, 8, 9}
	p := prmt.New(prmt.IntSlice(a))
	r := []int{}

	// result := append(make([]int, 0), program...)
	for p.Next() {
		acc := 0
		result := 0
		fmt.Printf("\n\n Permutation %v \n\n", a)
		for _, perm := range a {
			input := []int{perm, result}
			_, result = doOperation(program, input, 0, 0)
			acc++
		}

		r = append(r, result)
	}

	sort.Ints(r)
	return r[len(r)-1]
}
