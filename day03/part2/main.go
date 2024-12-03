// AoC 2024, Day 3, Part 2, Lixen Wraith
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func calc(numbers [][]int) int {
	sumProduct := 0
	for i := 0; i < len(numbers); i++ {
		sumProduct += numbers[i][0] * numbers[i][1]
	}
	return sumProduct
}

func isNumber(c rune) bool {
	return '0' <= c && c <= '9'
}

func inputFromFile(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

const (
	operation = "mul"
	enable    = "do()"
	disable   = "don't()"
)

const (
	stUnset = iota
	stOps
	stFirst
	stSecond
)

// parse the file into a slice of int pairs.
func parse(input []byte) [][]int {
	numbers := make([][]int, 0)

	content := string(bytes.ReplaceAll(input, []byte("\n"), []byte("")))

	emptyToken := []byte{}
	token := emptyToken
	state := stUnset
	opEnabled := true

	n1, n2 := -1, -1
	var j int

	for i, c := range content {
		if j > 0 {
			j--
			continue
		}

		if !opEnabled {
			if len(content)-i >= len(enable) {
				if content[i:i+len(enable)] == enable {
					opEnabled = true
					state = stUnset
					j = len(enable) - 1
				}
			}
			continue
		}

		if len(content)-i >= len(disable) {
			if content[i:i+len(disable)] == disable {
				opEnabled = false
				state = stUnset
				j = len(disable) - 1
			}
		}

		if state == stUnset && len(content)-i >= len(operation) {
			if content[i:i+len(operation)] == operation {
				state = stOps
				j = len(operation) - 1
				continue
			}
		}

		// non-operation parsing
		switch state {

		case stOps:
			if c == '(' {
				token = emptyToken
				state = stFirst
			} else {
				state = stUnset
			}

		case stFirst:
			if isNumber(c) {
				token = append(token, byte(c))
			} else if c == ',' && len(token) > 0 {
				n1, _ = strconv.Atoi(string(token))
				token = emptyToken
				state = stSecond
			} else if c != ',' {
				token = emptyToken
				state = stUnset
			}
		case stSecond:
			if isNumber(c) {
				token = append(token, byte(c))
			} else if c == ')' && len(token) > 0 {
				n2, _ = strconv.Atoi(string(token))
				numbers = append(numbers, []int{n1, n2})
				n1 = -1
				n2 = -1
				token = emptyToken
				state = stUnset
			} else if c != ')' {
				token = emptyToken
				state = stUnset
			}

		}

	}

	return numbers
}

func main() {
	input, err := inputFromFile("./input.txt")
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 : Calculation result = %v\n", calc(parse(input)))
}
