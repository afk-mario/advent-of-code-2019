package main

import (
	"advent/utils"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module,
// take its mass,
// divide by three,
// round down,
// and subtract 2.

func main() {
	inputPath := flag.String("input", "./input.txt", "input file path")
	flag.Parse()

	if len(flag.Args()) > 0 {
		valueByArg(flag.Args()[0])
	}

	valueByInput(*inputPath)
}

func valueByInput(file string) {
	fmt.Println("Printing fuel by file ---", file)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, err := utils.ReadInts(f, bufio.ScanWords)
	if err != nil {
		panic(err)
	}

	acc := 0

	for i := 0; i < len(input); i++ {
		acc += getFuel(input[i])
	}

	fmt.Println("\nResult:", acc)
	fmt.Println("")
}

func valueByArg(input string) {
	arg, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Printing fuel by arg ---", arg)
	fmt.Println("\nResult:", getFuel(arg))
	fmt.Println("")
}

func getFuel(mass int) int {
	fuel := mass/3 - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + getFuel(fuel)
}
