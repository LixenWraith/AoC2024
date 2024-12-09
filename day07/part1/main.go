// AoC 2024, Day 7, Part 1, Lixen Wraith
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type halfStructure = map[uint64]uint64            // half data structure for line operations
type fullStructure = map[uint64]map[uint64]uint64 // full data structure for input/output

const (
	OpAdd = iota
	OpMul
)

// sum of the keys for valid equations
func calc(opsMap fullStructure) uint64 {
	var sum uint64

	for k, _ := range opsMap { // opsMap only has valid equation keys
		sum += k
	}

	return sum
}

// generate a map of valid bit sequence representing the operations for correct euqations
func generateOpsMap(numbersMap fullStructure) fullStructure {
	opsMap := make(fullStructure)

	for k, v := range numbersMap {
		lo := lineOps(k, v)
		if len(lo) > 0 {
			for i := 0; i < len(lo); i++ {
				if opsMap[k] == nil {
					opsMap[k] = make(halfStructure)
				}
				opsMap[k][uint64(i)] = lo[i]
			}
		}
	}

	return opsMap
}

// iterate operators of line
func lineOps(testValue uint64, numbers halfStructure) []uint64 {
	lineOpsCount := uint64(len(numbers) - 1)
	// 2^lineOpsCount-1 : bits that represent valid operations for iteration
	opsMask := uint64(1<<lineOpsCount) - 1
	var opsIterator uint64

	validOps := make([]uint64, 0)
	for opsIterator = 0; opsIterator <= opsMask; opsIterator++ { // ops iteration represents a sequence of ops
		if isValidOps(testValue, numbers, opsIterator) {
			validOps = append(validOps, opsIterator)
		}
	}

	return validOps
}

// validate operators for test result
func isValidOps(testValue uint64, numbers halfStructure, ops uint64) bool {
	var i, equation uint64
	n := uint64(len(numbers))

	for i = 0; i < n; i++ {
		if i == 0 {
			equation += numbers[i]
		} else {
			// lowest bit is the next operation
			opBit := ops & 0b1
			ops = ops >> 1 // remove lowest bit
			switch opBit {
			case OpAdd:
				equation += numbers[i]
			case OpMul:
				equation *= numbers[i]
			}
		}
	}

	return equation == testValue
}

// parse the input into slice (lines) of int slice (elements in each line)
func parse(input []byte) fullStructure {
	content := strings.Split(string(input), "\n")
	numbers := make(fullStructure)

	for _, line := range content {
		stringSlice := strings.Split(line, " ")
		keyString := stringSlice[0]
		testResult, err := strconv.Atoi(keyString[:len(keyString)-1])
		if err != nil {
			log.Fatalln("Error converting key to int")
		}

		key := uint64(testResult)
		if _, ok := numbers[key]; ok {
			log.Fatalln("Duplicate Key ", key)
		}
		numbers[key] = make(map[uint64]uint64)
		for j, s := range stringSlice[1:] {
			n, _ := strconv.Atoi(s)
			numbers[key][uint64(j)] = uint64(n)
		}
	}

	return numbers
}

// read input file into a slice of byte
func inputFromFile(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func main() {
	filename := "./input.txt"
	// filename := "./example.txt"
	input, err := inputFromFile(filename)
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 : Sum of correct equations = %v\n", calc(generateOpsMap(parse(input))))
}