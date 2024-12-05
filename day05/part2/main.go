// AoC 2024, Day 5, Part 2, Lixen Wraith
package main

import (
	"fmt"
	"os"
	"sort"
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

// returns a slice of production order slices that have incorrect page order
func getIncorrectProductionOrders(orders map[int][]int, productionOrders [][]int) [][]int {
	incorrectProductionOrders := [][]int{}
	for i := 0; i < len(productionOrders); i++ {
		focus := productionOrders[i]
		if !isCorrectProductionOrder(orders, focus) {
			incorrectProductionOrders = append(incorrectProductionOrders, focus)
		}
	}
	return incorrectProductionOrders
}

// sum the middle page number of the production orders that have incorrect page order
func midpageSum(orders map[int][]int, incorrectProductionOrders [][]int) int {
	midPageSum := 0
	for i := 0; i < len(incorrectProductionOrders); i++ {
		focus := incorrectProductionOrders[i]
		midPage := customSort(orders, focus)[len(focus)/2]
		midPageSum += midPage
	}
	return midPageSum
}

// slice sort based on orders map custom rule
func customSort(orders map[int][]int, incorrectProductionOrder []int) []int {
	productionOrder := incorrectProductionOrder
	sort.Slice(productionOrder, func(i, j int) bool {
		a, b := productionOrder[i], productionOrder[j]

		// check if b must come after a
		if allowed, exists := orders[a]; exists {
			for _, num := range allowed {
				if num == b {
					return false // b must come after a
				}
			}
		}

		// check if a must come after b
		if allowed, exists := orders[b]; exists {
			for _, num := range allowed {
				if num == a {
					return true // a must come after b
				}
			}
		}

		// default to normal ordering (no change) when order is not defined
		return a < b
	})
	return productionOrder
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

	fmt.Printf("Part 2 : Mid-page sum of incorrect production orders = %v\n",
		midpageSum(orders, getIncorrectProductionOrders(orders, productionOrders)))
}