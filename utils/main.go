package utils

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader, split bufio.SplitFunc) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(split)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

// SplitComma splits comma separated data
func SplitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	commaidx := bytes.IndexByte(data, ',')
	if commaidx > 0 {
		// we need to return the next position
		buffer := data[:commaidx]
		return commaidx + 1, bytes.TrimSpace(buffer), nil
	}

	// if we are at the end of the string, just return the entire buffer
	if atEOF {
		// but only do that when there is some data. If not, this might mean
		// that we've reached the end of our input CSV string
		if len(data) > 0 {
			return len(data), bytes.TrimSpace(data), nil
		}
	}

	// when 0, nil, nil is returned, this is a signal to the interface to read
	// more data in from the input reader. In this case, this input is our
	// string reader and this pretty much will never occur.
	return 0, nil, nil
}
