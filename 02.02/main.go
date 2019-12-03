package main

import (
	"advent/utils"
	"flag"
	"fmt"
	"os"
	"strings"
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
	fmt.Println("\nPrinting value by file ---", file)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	input, err := utils.ReadInts(f, utils.SplitComma)
	if err != nil {
		panic(err)
	}

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			arr := append([]int(nil), input...)
			arr[1] = i // noun
			arr[2] = j // verb
			result := transformData(arr)
			if result[0] == 19690720 {
				fmt.Printf("Found result noun=%d verb=%d \n\n", i, j)
				fmtResult := utils.FmtIntSlice(result)
				fmt.Println(fmtResult)
				return
			}
		}
	}

}

func valueByArg(arg string) {
	f := strings.NewReader(arg)
	input, err := utils.ReadInts(f, utils.SplitComma)

	if err != nil {
		panic(err)
	}

	result := transformData(input)
	fmt.Println("\nPrinting value by arg ---", input)
	fmt.Println("\nResult:", utils.FmtIntSlice(result))
}

func transformData(input []int) []int {
	result := append([]int(nil), input...)
	for i := range result {
		if i%4 == 0 {
			if result[i] == 1 {
				op := result[result[i+1]] + result[result[i+2]]
				where := result[i+3]
				result[where] = op
			} else if result[i] == 2 {
				op := result[result[i+1]] * result[result[i+2]]
				where := result[i+3]
				result[where] = op
			} else if result[i] == 99 {
				return result
			}
		}
	}
	return result
}
