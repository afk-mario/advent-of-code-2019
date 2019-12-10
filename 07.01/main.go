package main

import (
	"advent/logger"
	"advent/utils"
	"flag"
	"fmt"
	"os"
	"strconv"

	prmt "github.com/gitchander/permutation"
)

type opts struct {
	path  string
	debug bool
	input int
}

var _o opts

func output(val int) {
	fmt.Printf("┌─────────┐")
	fmt.Printf("\n│%9d│", val)
	fmt.Printf("\n└─────────┘\n")
}

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

	a := []int{0, 1, 2, 3, 4}
	p := prmt.New(prmt.IntSlice(a))
	// result := append(make([]int, 0), input...)
	for p.Next() {
		acc := 0
		z := 0
		result := append(make([]int, 0), input...)
		for _, i := range a {
			_o.input = i
			// fmt.Println("\nInput:\n", utils.FmtIntSlice(result))
			result, z = doOperation(result, 0)
			acc += z
			// fmt.Println("\nResult:\n", utils.FmtIntSlice(result))
		}
		fmt.Println("\n______")
		fmt.Println(a)
		fmt.Println("result", acc)
		fmt.Println("\n______")
	}

	// _o.input = 0
	// result, acc := doOperation(input, 0)
	// if _o.debug {
	// 	fmt.Println("\nResult:", utils.FmtIntSlice(result))
	// 	fmt.Println("\aacc:", acck
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

// Some operations have the ourput at the end therefore the mode can't be inmediate mode  = 1
func setParamsByOp(op int, arr, params, modes []int) []int {
	l := len(params)
	switch op {
	case 1:
		l--
	case 2:
		l--
	case 3:
		l--
	case 4:
		l = l
	case 5:
		l = l
	case 6:
		l = l
	case 7:
		l--
	case 8:
		l--
	}

	for i := range params {
		if modes[i] == 0 && i < l {
			params[i] = arr[params[i]]
		}
	}
	return params
}

// Normally operations take the first two params to do something and the last to output the operation
func getParams(op, pointer int, modes []int, arr []int) []int {
	size := getSizeByOP(op)
	params := make([]int, 0)

	for i := 0; i < size; i++ {
		a := arr[pointer+i+1]
		params = append(params, a)
	}

	params = setParamsByOp(op, arr, params, modes)

	return params
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
	case 5:
		return 2
	case 6:
		return 2
	case 7:
		return 3
	case 8:
		return 3
	}
	return 0
}

func getJumpByOp(op int, pointer int, params, modes []int, arr []int) int {
	switch op {
	case 1:
		return getSizeByOP(op) + pointer + 1
	case 2:
		return getSizeByOP(op) + pointer + 1
	case 3:
		return getSizeByOP(op) + pointer + 1
	case 4:
		return getSizeByOP(op) + pointer + 1
		// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
	case 5:
		if params[0] != 0 {
			return params[1]
		}
		return getSizeByOP(op) + pointer + 1
		// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
	case 6:
		if params[0] == 0 {
			return params[1]
		}
		return getSizeByOP(op) + pointer + 1
		// Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	case 7:
		return getSizeByOP(op) + pointer + 1
		// Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	case 8:
		return getSizeByOP(op) + pointer + 1

	}
	return 0
}

func operate(op, pointer int, params, modes, arr []int) ([]int, int) {
	acc := 0
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
		acc = params[0]
		output(params[0])
		break
		// Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	case 7:
		if params[0] < params[1] {
			arr[params[2]] = 1
		} else {
			arr[params[2]] = 0
		}
		// Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	case 8:
		if params[0] == params[1] {
			arr[params[2]] = 1
		} else {
			arr[params[2]] = 0
		}
	}
	return arr, acc
}

func doOperation(arr []int, pointer int) ([]int, int) {
	acc := 0
	nInstruction := arr[pointer]
	sInstruction := strconv.Itoa(arr[pointer])

	operation := getOperation(sInstruction)
	debugtxt := fmt.Sprintf("DO [%d]\n instruction: %04d | op: %2d | pointer: %2d", _o.input, nInstruction, operation, pointer)

	if operation == 99 {
		logger.Debug(debugtxt)
		fmt.Println("\nFinish at: ", pointer)
		return arr, acc
	}

	if operation > 8 || operation < 1 {
		fmt.Println("\nOperation not supported: ", operation)
		return arr, acc
	}

	modes := getModes(sInstruction, 3)
	params := getParams(operation, pointer, modes, arr)
	jump := getJumpByOp(operation, pointer, params, modes, arr)

	debugtxt += fmt.Sprintf(" | modes: %v | params %v | jump %d", modes, params, jump)
	logger.Debug(debugtxt)

	arr, a := operate(operation, pointer, params, modes, arr)

	acc += a

	x, y := doOperation(arr, jump)
	return x, y + acc
}
