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
	noun := flag.Int("noun", -1, "the noun of the problem")
	verb := flag.Int("verb", -1, "the verb of the problem")
	flag.Parse()

	if *noun == -1 || *verb == -1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(flag.Args()) > 0 {
		valueByArg(flag.Args()[0])
	}

	valueByInput(*inputPath, *noun, *verb)
}

func valueByInput(file string, noun int, verb int) {
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

	input[1] = noun
	input[2] = verb

	result := transformData(input)
	fmtResult := utils.FmtIntSlice(result)
	fmt.Println("\nPrinting value by input ---", file)
	fmt.Println("\nResult:", fmtResult)

	f, err = os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(fmtResult)
	f.Sync()
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
