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

	resultByFile(_o.path)
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
	doOperation(input, 0)
	// if _o.debug {
	// 	fmt.Println("\nResult:", utils.FmtIntSlice(result))
	// }
}

func getOperation(s string) int {
	if len(s) < 3 {
		r, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return r
	}
	r, err := strconv.Atoi(s[len(s)-2 : len(s)])
	if err != nil {
		panic(err)
	}
	return r
}

func getModes(instruction string, l int) []int {
	if len(instruction) < 3 {
		return []int{0, 0, 0}
	}

	modes := make([]int, 0)
	instruction = instruction[:len(instruction)-2]
	for _, c := range instruction {
		mode := c - '0'
		modes = append(modes, int(mode))
	}

	if len(modes) == l {
		return modes
	}

	modes = utils.Reverse(modes)

	for i := len(modes); i < l; i++ {
		modes = append(modes, 0)
	}
	return modes

}

// Normally operations take the first two params to do something and the last to output the operation
func getParams(op, pointer int, modes []int, arr []int) []int {
	size := getSizeByOP(op)
	result := make([]int, 0)

	for i := 0; i < size; i++ {
		a := arr[pointer+i+1]
		result = append(result, a)
	}

	// Reference mode last param can't be reference
	for i := range result {
		if modes[i] == 0 && i < len(result)-1 {
			result[i] = arr[result[i]]
		}
	}

	return result
}

func getSizeByOP(op int) int {
	switch op {
	case 1:
		return 3
	case 2:
		return 3
	case 3:
		return 1
	case 4:
		return 1
	}
	return 0
}

func getJumpByOp(op int) int {
	switch op {
	case 1:
		return getSizeByOP(op) + 1
	case 2:
		return getSizeByOP(op) + 1
	case 3:
		return getSizeByOP(op) + 1
	case 4:
		return getSizeByOP(op) + 1
	}
	return 0
}

func output(val int) {
	fmt.Printf("\n┌─────────┐")
	fmt.Printf("\n│%9d│", val)
	fmt.Printf("\n└─────────┘\n\n")
}

func operate(op, pointer int, params, arr []int) []int {
	switch op {
	case 1:
		arr[params[2]] = params[0] + params[1]
		break
	case 2:
		arr[params[2]] = params[0] * params[1]
		break
	case 3:
		arr[params[0]] = _o.input
		break
	case 4:
		output(arr[params[0]])
		break
	}
	return arr
}

func doOperation(arr []int, pointer int) []int {
	nInstruction := arr[pointer]
	sInstruction := strconv.Itoa(arr[pointer])

	operation := getOperation(sInstruction)
	debugtxt := fmt.Sprintf("DO \n instruction: %04d | op: %2d | pointer: %2d", nInstruction, operation, pointer)

	if operation == 99 {
		fmt.Println("\nFinish at: ", pointer)
		return arr
	}

	if operation > 4 {
		return arr
	}

	modes := getModes(sInstruction, 3)
	params := getParams(operation, pointer, modes, arr)
	jump := getJumpByOp(operation)

	debugtxt += fmt.Sprintf("| modes: %v | params %v | jump %d", modes, params, jump)
	logger.Debug(debugtxt)
	// fmt.Println("operation", operation)
	// fmt.Println("modes", modes)
	// fmt.Println("params", params)
	// fmt.Println("size", params)

	arr = operate(operation, pointer, params, arr)

	return doOperation(arr, pointer+jump)
}

// func readOpcodes(input []int) []int {
// 	logger.Debug("\n\n")
// 	result := append([]int(nil), input...)
// 	for i := 0; i < len(result); i++ {
// 		instruction := strconv.Itoa(result[i])

// 		val := 0
// 		where := 0
// 		operator := 0

// 		a := result[i+1]
// 		b := result[i+2]
// 		c := result[i+3]

// 		logger.Debug("[%2d] instruction: %05d - | a: m %4d - i %4d | b: m %4d - i %4d | c: m %4d - i %4d", i, z, modA, a, modB, b, modC, c)

// 		switch operation {
// 		case 1:
// 			if modA == 0 {
// 				a = result[a]
// 			}

// 			if modB == 0 {
// 				b = result[b]
// 			}

// 			val = a + b
// 			where = c
// 			operator = 3
// 			break
// 		case 2:
// 			if modA == 0 {
// 				a = result[a]
// 			}

// 			if modB == 0 {
// 				b = result[b]
// 			}
// 			val = a * b
// 			where = c
// 			operator = 3
// 			break
// 		case 3:
// 			val = _o.input
// 			where = a
// 			operator = 1
// 			break
// 		case 4:
// 			where = a
// 			fmt.Printf("\n\n┌─────────┐")
// 			fmt.Printf("\n│%9d│", result[where])
// 			fmt.Printf("\n└─────────┘\n\n")
// 			operator = 1
// 			break
// 		case 99:
// 			return result
// 			break
// 		}
// 		i += operator
// 		result[where] = val
// 		logger.Debug("[%2d] instruction: %05d - | a: m %4d - i %4d | b: m %4d - i %4d | c: m %4d - i %4d \n", i, z, modA, a, modB, b, modC, c)

// 	}
// 	return result
// }
