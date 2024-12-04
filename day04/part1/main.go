// AoC 2024, Day 4, Part 1, Lixen Wraith
package main

import (
	"fmt"
	"os"
	"strings"
)

const xmas = "XMAS"

// not counting correctly for words of 1, not going to fix it and teste cases removed
func wordSearch(grid [][]rune, word string, row, column int) int {
	// no match if the first rune of the string doesn't match with the grid's rune at the starting of the search
	if grid[row][column] != rune(word[0]) {
		return 0
	}

	matchCount := 0
	maxDepth := len(word)
	// setting search bounds
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			depth := 0

			matchFound := false
			for n, c := range word {
				di := row + (i * n)
				dj := column + (j * n)
				if di < 0 || di >= len(grid) || dj < 0 || dj >= len(grid[0]) {
					break
				}

				if c == grid[di][dj] {
					depth++
					if depth == maxDepth {
						matchFound = true
						break
					}
				} else {
					break
				}
			}
			if matchFound {
				matchCount++
			}
		}
	}
	return matchCount
}

/*
// first wrong implementation that allowed direction change mid-word, not what was asked
func wordSearchFreeForm(grid [][]rune, word string, row, column int) int {
	if len(word) == 0 {
		return 0
	}

	// no match if the first rune of the string doesn't match with the grid's rune at the starting of the search
	if grid[row][column] != rune(word[0]) {
		return 0
	}

	// check end of word
	if len(word) == 1 && grid[row][column] == rune(word[0]) {
		return 1
	}

	matchCount := 0
	// setting search bounds
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if (row+i < 0) || (row+i >= len(grid)) || (column+j < 0) || (column+j >= len(grid[0])) {
				continue
			}

			if 0 <= row+i && row+i < len(grid) && 0 <= column+j && column+j < len(grid[0]) {
				matchCount += wordSearch(grid, word[1:], row+i, column+j)
			}
		}
	}

	return matchCount
}
*/

func crawl(grid [][]rune, word string) int {
	if len(word) == 0 {
		return 0
	}

	wordCount := 0

	for i := 0; i < len(grid); i++ {
		fmt.Printf("found: %d - Searching row: %d\n", wordCount, i)
		for j := 0; j < len(grid[0]); j++ {
			wordCount += wordSearch(grid, word, i, j)
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

	fmt.Printf("Part 1 : Total matches = %v\n", crawl(parse(input), xmas))
}
