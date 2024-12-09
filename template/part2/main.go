// AoC 2024, Day Y, Part X Lixen Wraith
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// colors for visualization
const (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	reset  = "\033[0m"
)

// number of bits representing an operation in an operations bit sequence
const bitsPerOp uint64 = 2

// data types
type grid = [][]rune   // [y][x]
type gridLine = []rune // [x]

type nestedMapStructure = map[uint64]map[uint64]uint64 // full data structure for input/output
type mapStructure = map[uint64]uint64                  // half data structure for line operations

// visualize
func visualize(g grid) {
	lineCount := 0
	for y := range g {
		for x := range y {
			switch g[y][x] {
			case '#':
				fmt.Printf("%s%c%s", green, g[y][x], reset)
			case '.':
				fmt.Printf("%c", g[y][x])
			default:
				fmt.Printf("%s%c%s", red, g[y][x], reset)
			}
		}
		fmt.Println()
		lineCount++
	}
	fmt.Println()
	lineCount++
	// go up equal to printed lines to keep the output grid stationary
	fmt.Print(strings.Join([]string{"\033[", strconv.Itoa(lineCount + 2), "A"}, ""))
}

// parse the input into a grid
func parse(input []byte) grid {
	content := strings.Split(string(input), "\n")
	g := make(grid, len(content))

	for y, line := range content {
		g[y] = make(gridLine, len(line))
		for x, r := range line {
			g[y][x] = r
		}
	}

	return g
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
	// filename := "./example1.txt"
	// filename := "./example2.txt"
	// filename := "./example3.txt"
	input, err := inputFromFile(filename)
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part 2 :  = %d\n", input)
}