package main

import (
	"advent/logger"
	"advent/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
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

	resultByFile(_o.filePath)
}

func resultByFile(file string) {
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

	fmt.Println("\nInput:", utils.FmtIntSlice(input))
	result := readOpcodes(input)
	fmt.Println("\nResult:", utils.FmtIntSlice(result))
}

func readOpcodes(input []int) []int {
	result := append([]int(nil), input...)
	pointer := 5
	for i := 0; i < len(result); i += pointer {
		instruction := -1
		t := strconv.Itoa(result[i])
		l := len(t)
		modA := 0
		modB := 0
		modC := 0

		// Instruction is more than 2
		if l > 2 {
			instruction, _ = strconv.Atoi(t[l-2 : len(t)])
			if l > 2 {
				modA, _ = strconv.Atoi(t[l-3 : l-2])
			}

			if l > 3 {
				modB, _ = strconv.Atoi(t[l-4 : l-3])
			}
			if l > 4 {
				modC, _ = strconv.Atoi(t[l-5 : l-4])
			}
		} else {
			instruction = result[i]
		}

		logger.Debug("instruction: %d - a: %d b: %d c: %d", instruction, modA, modB, modC)

		val := 0
		where := 0

		a := result[i+1]
		b := result[i+2]
		c := result[i+3]

		if modA == 0 {
			a = result[a]
		}
		if modB == 0 {
			b = result[b]
		}
		if modC == 0 {
			c = result[c]
		}

		switch instruction {
		case 1:
			val = a + b
			where = c
			pointer = 5
			break
		case 2:
			val = a * b
			where = c
			pointer = 5
			break
		case 3:
			val = b
			where = a
			pointer = 4
			break
		case 4:
			val = a
			where = b
			pointer = 4
			break
		case 99:
			return result
			break
		}

		result[where] = val

	}
	return result
}
