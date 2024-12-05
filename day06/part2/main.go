// AoC 2024, Day 6, Part 2, Lixen Wraith
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Placeholder

// parse the input into slice (lines) of int slice (elements in each line)
func parse(input []byte, separator string) [][]int {
	content := strings.Split(string(input), "\n")
	output := make([][]int, len(content))

	for i, line := range content {
		stringSlice := strings.Split(line, separator)
		output[i] = make([]int, len(stringSlice))
		for j, s := range stringSlice {
			output[i][j], _ = strconv.Atoi(s)
		}
	}
	return output
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
	input, err := inputFromFile("./input1.txt")
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}
	orders := orderMap(parse(input, "|"))

	input, err = inputFromFile("./input2.txt")
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}
	productionOrders := parse(input, ",")

	fmt.Printf("Part 2 : Mid-page sum of incorrect production orders = %v\n",
		midpageSum(orders, getIncorrectProductionOrders(orders, productionOrders)))
}

