package main

import (
	"advent/logger"
	"flag"
	"fmt"
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

	fmt.Println("Result: ", *inputPath)
}
