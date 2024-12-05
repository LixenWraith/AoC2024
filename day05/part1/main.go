// AoC 2024, Day 5, Part 1, Lixen Wraith
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// merge all orders into a single order
func orderMap(orders [][]int) map[int][]int {
	result := map[int][]int{}
	for i := 0; i < len(orders); i++ {
		if _, exists := result[orders[i][0]]; !exists {
			result[orders[i][0]] = []int{orders[i][1]}
		} else {
			result[orders[i][0]] = append(result[orders[i][0]], orders[i][1])
		}
	}
	return result
}

// check if a production order has correct page order
func isCorrectProductionOrder(orders map[int][]int, productionOrder []int) bool {
	for i := 0; i < len(productionOrder); i++ {
		for j := 0; j < i; j++ {
			focus := productionOrder[i]
			compare := productionOrder[j]
			for _, t := range orders[focus] {
				if compare == t {
					return false
				}
			}
		}
	}
	return true
}

// sum the middle page number of the production orders that have the correct page order
func midpageSum(orders map[int][]int, productionOrder [][]int) int {
	midPageSum := 0
	for i := 0; i < len(productionOrder); i++ {
		focus := productionOrder[i]
		if isCorrectProductionOrder(orders, focus) {
			midPage := focus[len(focus)/2]
			midPageSum += midPage
		}
	}
	return midPageSum
}

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

	fmt.Printf("Part 1 : Mid-page sum of correct production orders = %v\n", midpageSum(orders, productionOrders))
}