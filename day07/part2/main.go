// AoC 2024, Day 7, Part 2, Lixen Wraith
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

// number of bits representing an operation in an operations bit sequence
const bitsPerOp uint64 = 2

// operation bit sequence
const (
	OpAdd    uint64 = 0b00
	OpMul    uint64 = 0b01
	OpConcat uint64 = 0b10
	OpUnused uint64 = 0b11
)

// bit mask for serialized bitmap of line operations
func getMask(opBits uint64, numOps uint64) uint64 {
	// 2^(opBits * numOps) - 1 : bits that represent valid operations for iteration
	// opBits = bitsPerOp , numOps = 2 -> 2 ^ (2 * 2) - 1 = 0b1111
	return uint64(1<<(opBits*numOps) - 1)
}

// checks if an operations bit sequence contains an instance of any unused operation bit sequence
func containsUnusedOp(opsSequence, opsSequenceMask, opBits uint64, unusedOps ...uint64) bool {
	if len(unusedOps) == 0 {
		return false
	}
	singleOpMask := getMask(opBits, 1)
	for _, unusedOp := range unusedOps {
		for validOpsSequenceBits := opsSequenceMask; validOpsSequenceBits != 0; validOpsSequenceBits = validOpsSequenceBits >> singleOpMask {
			if validOpsSequenceBits&singleOpMask == unusedOp {
				return true
			}
		}
	}
	return false
}

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

	// fmt.Println("== generateOpsMap, numbersMap from parse: ", numbersMap)
	for k, v := range numbersMap {
		// fmt.Println("-- Checking (result, numbers): ", k, v)
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
	var opsIterator uint64
	mask := getMask(bitsPerOp, lineOpsCount)

	validOps := make([]uint64, 0)
	for opsIterator = 0; opsIterator <= mask; opsIterator++ { // ops iteration represents a sequence of ops
		// filters out operations sequences that contain unused operation bits
		// TODO: performance optimization attempt, had no effect.
		//  Iteration loop can be changed instead to generate only valid sequences.
		if !containsUnusedOp(opsIterator, bitsPerOp, mask, OpUnused) {
			if isValidOps(testValue, numbers, opsIterator) {
				validOps = append(validOps, opsIterator)
			}
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
			// lowest n bits represent the next operation, n is bitsPerOp
			opBits := ops & getMask(bitsPerOp, 1) // apply mask for next operation filtering
			ops = ops >> bitsPerOp
			switch opBits {
			case OpAdd:
				equation += numbers[i]
			case OpMul:
				equation *= numbers[i]
			case OpConcat:
				e, _ := strconv.Atoi(strings.Join([]string{strconv.Itoa(int(equation)), strconv.Itoa(int(numbers[i]))}, ""))
				equation = uint64(e)
			default:
				continue
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

	fmt.Printf("Part 2 : Sum of correct equations = %v\n", calc(generateOpsMap(parse(input))))
}