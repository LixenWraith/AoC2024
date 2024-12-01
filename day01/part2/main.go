// AoC 2024, Day 1, Part 2, Lixen Wraith
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// distance calculates sum of the distances of paired elements by index in 2 sorted lists
func distance(lists [][]int) int {
	// sort lists
	sortedOne := lists[0]
	sortedTwo := lists[1]
	sort.Ints(sortedOne)
	sort.Ints(sortedTwo)

	// calculate distance of numbers in the list with same index
	var totalDistance int
	for i := 0; i < len(sortedOne); i++ {
		totalDistance += int(math.Abs(float64(sortedOne[i] - sortedTwo[i])))
	}

	return totalDistance
}

// parse converts the strictly structured test file into 2 int slices
func parse(filename string) ([][]int, error) {
	// read file into []byte
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	list1 := make([]int, 0)
	list2 := make([]int, 0)
	// parse each line into 2 int
	const numLines = 1000
	const lineBytes = 14 // 5 digits, 3 space, 5 digits, \n
	for l := 0; l < numLines; l++ {
		i, err := strconv.Atoi(string(content[l*lineBytes : l*lineBytes+5]))
		if err != nil {
			return nil, err
		}
		j, err := strconv.Atoi(string(content[l*lineBytes+8 : l*lineBytes+13]))
		if err != nil {
			return nil, err
		}
		list1 = append(list1, i)
		list2 = append(list2, j)
	}

	return [][]int{list1, list2}, nil
}

// frequency of each location repetition in each list
func frequency(input [][]int) []map[int]int {
	frequency1 := map[int]int{}
	frequency2 := map[int]int{}
	// traverse the lists and populate the frequency maps
	for i := 0; i < len(input[0]); i++ {
		frequency1[input[0][i]]++
		frequency2[input[1][i]]++
	}
	return []map[int]int{frequency1, frequency2}
}

// similarity calculation by adding products of
// a key in the left(first) map and frequency of that key in the other(right) map
func similarity(lf []map[int]int) int {
	var s int
	// left list is traversed on according to example
	for k := range lf[0] {
		if f, ok := lf[1][k]; ok {
			s += k * f
		}
	}
	return s
}

func main() {
	lists, err := parse("./input.txt")
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Part1: Total Distance = ", distance(lists))
	fmt.Println("Part2: Similarity = ", similarity(frequency(lists)))
}
