package intcode

import (
	"advent/logger"
	"advent/utils"
	"fmt"
	"strconv"
)

// Intcode computer
type Intcode struct {
	program    []int
	pointer    int
	output     int
	inputIndex int
	inputArr   []int
}

// DrawOutput draws the output from the program
func DrawOutput(val int) {
	fmt.Printf("┌─────────┐")
	fmt.Printf("\n│%9d│", val)
	fmt.Printf("\n└─────────┘\n")
}

// GetOperation from an int 10001 = 1
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

// getModes from an instruction 10001 A: 0 B: 0 C: 1
func getModes(instruction string, l int) []int {
	// If the instruction has inly two digits then it's default values for each mode
	if len(instruction) < 3 {
		return []int{0, 0, 0}
	}

	modes := []int{}
	// remove last two charts from the instruction
	instruction = instruction[:len(instruction)-2]
	for _, c := range instruction {
		mode := c - '0'
		modes = append(modes, int(mode))
	}

	// if the modes already have the desired length
	// return modes as is
	if len(modes) == l {
		return modes
	}

	// If not we need to append to the end default values = 0
	modes = utils.Reverse(modes)

	for i := len(modes); i < l; i++ {
		modes = append(modes, 0)
	}
	return modes
}

// TODO: refactor into dict
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
	default:
		return 0
	}
}

// GetTotalOutput ...
// func (intcode *Intcode) GetTotalOutput() int {
// 	acc := 0

// 	for _, e := range intcode.output {
// 		acc += e
// 	}

// 	logger.Debug("Total Output: %d %v", acc, intcode.output)

// 	return acc
// }

// Normally operations take the first two params to do something and the last to output the operation
func (intcode *Intcode) getParams(op int, modes []int) []int {
	size := getSizeByOP(op)
	params := []int{}

	for i := 0; i < size; i++ {
		a := intcode.program[intcode.pointer+i+1]
		params = append(params, a)
	}

	params = intcode.setParamsByOp(op, params, modes)

	return params
}

// Some operations have the ourput at the end therefore the mode can't be inmediate mode  = 1
func (intcode *Intcode) setParamsByOp(op int, params, modes []int) []int {
	l := len(params)
	switch op {
	case 1:
		l--
		break
	case 2:
		l--
		break
	case 3:
		l--
		break
	case 7:
		l--
		break
	case 8:
		l--
		break
	default:
		break
	}

	for i := range params {
		if modes[i] == 0 && i < l {
			params[i] = intcode.program[params[i]]
		}
	}

	return params
}

func (intcode *Intcode) setPointerByOp(op int, params, modes []int) int {
	switch op {
	case 1:
		intcode.pointer += getSizeByOP(op) + 1
		break
	case 2:
		intcode.pointer += getSizeByOP(op) + 1
		break
	case 3:
		intcode.pointer += getSizeByOP(op) + 1
		break
	case 4:
		intcode.pointer += getSizeByOP(op) + 1
		break
		// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
	case 5:
		if params[0] != 0 {
			intcode.pointer = params[1]
		} else {
			intcode.pointer += getSizeByOP(op) + 1
		}
		break

		// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
	case 6:
		if params[0] == 0 {
			intcode.pointer = params[1]
		} else {
			intcode.pointer += getSizeByOP(op) + 1
		}
		break
		// Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	case 7:
		intcode.pointer += getSizeByOP(op) + 1
		// Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	case 8:
		intcode.pointer += getSizeByOP(op) + 1

	}
	return 0
}

func (intcode *Intcode) operate(op int, params, modes []int) {
	arr := intcode.program
	switch op {
	case 1:
		arr[params[2]] = params[0] + params[1]
		break
	case 2:
		arr[params[2]] = params[0] * params[1]
		break
	case 3:
		arr[params[0]] = intcode.inputArr[intcode.inputIndex]
		break
	case 4:
		intcode.output += params[0]
		fmt.Println("Poop", intcode.output)
		DrawOutput(params[0])
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
}

// DoOperation ...
func (intcode *Intcode) DoOperation() int {
	arr := intcode.program
	nInstruction := arr[intcode.pointer]
	sInstruction := strconv.Itoa(arr[intcode.pointer])

	operation := getOperation(sInstruction)
	debugtxt := fmt.Sprintf("DO\n instruction: %04d | op: %2d | pointer: %2d", nInstruction, operation, intcode.pointer)

	if operation == 99 {
		logger.Debug(debugtxt)
		fmt.Println("\nFinish at: ", intcode.pointer)
		return intcode.output
	}

	if operation > 8 || operation < 1 {
		logger.Debug("%s \n", debugtxt)
		panic(fmt.Sprintf("Operation not supported %d", operation))
	}

	modes := getModes(sInstruction, 3)
	params := intcode.getParams(operation, modes)
	intcode.setPointerByOp(operation, params, modes)

	intcode.operate(operation, params, modes)
	debugtxt += fmt.Sprintf(" | modes: %v | params %v | jump %d | input %v | inputIndex %d | output %d", modes, params, intcode.pointer, intcode.inputArr, intcode.inputIndex, intcode.output)
	logger.Debug(debugtxt)

	if intcode.inputIndex > len(intcode.inputArr)-1 {
		logger.Debug("Input index out of bounds")
		return intcode.output
	}

	output := intcode.DoOperation()
	intcode.output += output

	if operation == 3 {
		intcode.inputIndex++
	}

	return intcode.output
}
