package main

import (
	"advent/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
)

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

	input, err := utils.ReadInts(f, utils.SplitComma)
	if err != nil {
		panic(err)
	}

	fmt.Println(input)

}

func valueByArg(input string) {
	arg, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Printing value by arg ---", arg)
}
