// AoC 2024, Day 8, Part 1, Lixen Wraith
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type halfStructure = map[uint64]uint64
type fullStructure = map[uint64]map[uint64]uint64 // full data structure from file

// Placeholder

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

	// fmt.Println("== Parsed output: ", numbers)
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
	// filename := "./input.txt"
	filename := "./example.txt"
	input, err := inputFromFile(filename)
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 : Sum of correct equations = %v\n", parse(input))
}