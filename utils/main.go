package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
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

// ReadStrings reads whitespace-separated strings from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadStrings(r io.Reader, split bufio.SplitFunc) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(split)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

// SplitComma splits comma separated data
func SplitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return SplitByChar(',', data, atEOF)
}

// SplitParenthesis splits comma separated data
func SplitParenthesis(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return SplitByChar(')', data, atEOF)
}

// SplitByChar splits data by a char
func SplitByChar(char byte, data []byte, atEOF bool) (advance int, token []byte, err error) {
	commaidx := bytes.IndexByte(data, char)
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

// FmtIntSlice returns a formatted int slice into a string comma separated
func FmtIntSlice(values []int) string {
	valuesText := []string{}
	for i := range values {
		number := values[i]
		text := "[" + fmt.Sprintf("%5d", i) + "] " + fmt.Sprintf("%5d", number)

		if (i+1)%10 == 0 {
			text = text + "\n"
		}
		valuesText = append(valuesText, text)
	}
	return strings.Join(valuesText, ", ")
}

// FmtStrSlice returns a formatted str slice into comma separated
func FmtStrSlice(values []string) string {
	valuesText := []string{}
	for i := range values {
		valuesText = append(valuesText, values[i])
	}
	return strings.Join(valuesText, ",")
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Pow exponential for ints
func Pow(a, b int) int {
	result := 1

	for 0 != b {
		if 0 != (b & 1) {
			result *= a

		}
		b >>= 1
		a *= a
	}

	return result
}

// Unique Remove duplicate values from Slice
func Unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// RemoveDuplicatesUnordered from string array
func RemoveDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

// Reverse ...
func Reverse(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}
