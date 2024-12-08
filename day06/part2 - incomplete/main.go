// AoC 2024, Day 6, Part 2, Lixen Wraith
package main

import (
	"fmt"
	"os"
	"strings"
)

// for fun!
var visualize bool

// parse the input into slice (lines) of int slice (elements in each line)
func parse(input []byte) (*Map, *Guard) {
	content := strings.Split(string(input), "\n")
	// traversed map cuz I like x,y coords
	output := make([][]rune, len(content[0]))
	for i := 0; i < len(output); i++ {
		output[i] = make([]rune, len(content))
	}
	g := &Guard{}

	for y := 0; y < len(content); y++ {
		for x := 0; x < len(content[0]); x++ {
			c := rune(content[y][x])
			// set guard info and decouples guard from level:
			// extracts guard info from map and sets guard position to open space
			if c == '^' {
				g.InBounds = true
				g.X = x
				g.Y = y
				g.Direction = DirUp
				output[x][y] = '.'
			} else {
				output[x][y] = c
			}
		}
	}
	return &Map{
		Level:   output,
		Visited: make(map[int]map[int]map[byte]bool),
		XMax:    len(output[0]),
		YMax:    len(output),
	}, g
}

// read input file into a slice of byte
func inputFromFile(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// output to file
func writeToFile(filename string, m *Map) error {
	var content []byte

	for i := 0; i < len(m.Level); i++ {
		// Convert each rune in the row to bytes
		content = append(content, string(m.Level[i])...)
		// Add newline after each row except the last one
		if i < len(m.Level)-1 {
			content = append(content, '\n')
		}
	}

	if err := os.WriteFile(filename, content, 0666); err != nil {
		return err
	}
	return nil
}

func main() {
	// filename := "./input.txt"
	filename := "./example.txt"
	input, err := inputFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading input file: %s\n", err)
		os.Exit(1)
	}
	levelMap, guard := parse(input)
	if levelMap == nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}
	if !guard.InBounds {
		fmt.Printf("No guards found, all safe!\n")
		os.Exit(0)
	}

	// for visual debug
	// visualize = false
	// visualize = true
	guard.Patrol(levelMap)

	levelMap.MarkVisited()

	if err = writeToFile("./output.txt", levelMap); err != nil {
		fmt.Printf("Error writing output file: %s\n", err)
	}

	fmt.Printf("Part 1: Guard visited positions = %d\n", guard.Visited)

	// for visual debug
	// visualize = false
	// visualize = true

	infiniteLoops := NewSimulator(levelMap)

	fmt.Printf("Part 2: Infinite loop posibilities = %d\n", infiniteLoops.Loops())
}