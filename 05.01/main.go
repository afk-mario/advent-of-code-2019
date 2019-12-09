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

	if _o.debug {
		fmt.Println("\nInput:\n", utils.FmtIntSlice(input))
	}
	result := readOpcodes(input)
	if _o.debug {
		fmt.Println("\nResult:", utils.FmtIntSlice(result))
	}
}

func readOpcodes(input []int) []int {
	logger.Debug("\n\n")
	result := append([]int(nil), input...)
	for i := 0; i < len(result); i++ {

		instruction := -1
		z := result[i]
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

		val := 0
		where := 0
		operator := 0

		a := result[i+1]
		b := result[i+2]
		c := result[i+3]

		logger.Debug("[%2d] instruction: %05d - | a: m %4d - i %4d | b: m %4d - i %4d | c: m %4d - i %4d", i, z, modA, a, modB, b, modC, c)

		switch instruction {
		case 1:
			if modA == 0 {
				a = result[a]
			}

			if modB == 0 {
				b = result[b]
			}

			val = a + b
			where = c
			operator = 3
			break
		case 2:
			if modA == 0 {
				a = result[a]
			}

			if modB == 0 {
				b = result[b]
			}
			val = a * b
			where = c
			operator = 3
			break
		case 3:
			val = 1
			where = a
			operator = 1
			break
		case 4:
			where = a
			fmt.Printf("\n\n┌─────────┐")
			fmt.Printf("\n│%9d│", result[where])
			fmt.Printf("\n└─────────┘\n\n")
			operator = 1
			break
		case 99:
			return result
			break
		}
		i += operator
		result[where] = val
		logger.Debug("[%2d] instruction: %05d - | a: m %4d - i %4d | b: m %4d - i %4d | c: m %4d - i %4d \n", i, z, modA, a, modB, b, modC, c)

	}
	return result
}
