// AoC 2024, Day 4, Part 2, Lixen Wraith
package main

import (
	"fmt"
	"os"
	"strings"
)

// receives a 3x3 2D rune slice and checks if it's X-MAS
func isXMAS(block [][]rune) bool {
	if len(block) != len(block[0]) && len(block) != 3 {
		fmt.Printf("Bad block: %dx%d\n", len(block), len(block[0]))
		os.Exit(1)
	}

	// check center
	if block[1][1] != 'A' {
		return false
	}

	tokenize := func(r rune) int {
		if r == 'M' {
			return -1
		} else if r == 'S' {
			return 1
		} else {
			return 0
		}
	}

	c00 := tokenize(block[0][0])
	c01 := tokenize(block[0][2])
	c10 := tokenize(block[2][0])
	c11 := tokenize(block[2][2])

	if c00 == 0 || c01 == 0 || c10 == 0 || c11 == 0 {
		return false
	}

	if c00+c11 != 0 || c01+c10 != 0 {
		return false
	}

	return true
}

func crawl(grid [][]rune) int {
	// start form top left and scan lines left to right and content top to bottom
	wordCount := 0

	get3x3Block := func(grid [][]rune, row, col int) [][]rune {
		block := make([][]rune, 3)
		for i := range block {
			block[i] = make([]rune, 3)
			copy(block[i], grid[row+i][col:col+3])
		}
		return block
	}

	// -2 accounting for block size
	for i := 0; i < len(grid)-2; i++ {
		fmt.Printf("found: %d - Searching row: %d\n", wordCount, i)
		for j := 0; j < len(grid[0])-2; j++ {
			if isXMAS(get3x3Block(grid, i, j)) {
				wordCount++
			}
		}
	}
	return wordCount
}

// parse the input into a grid of runes.
func parse(input []byte) [][]rune {
	content := strings.Split(string(input), "\n")
	grid := make([][]rune, len(content))

	for i, line := range content {
		grid[i] = []rune(line)
	}
	return grid
}

func inputFromFile(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func main() {
	input, err := inputFromFile("./input.txt")
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part 2 : Total matches = %v\n", crawl(parse(input)))
}
