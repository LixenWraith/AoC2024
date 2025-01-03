// AoC 2024, Day 8, Part 1, Lixen Wraith
package main

import (
	"fmt"
	"os"
	"strings"
)

// colors for visualization
const (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	reset  = "\033[0m"
)

// data types
type grid = [][]rune   // [y][x]f
type gridLine = []rune // [x]f

type position struct {
	y, x int
}
type frequencyMap = map[int]map[int]position // [fn][index]position(f)
type antinodeMap = map[position]rune         // [y][x]f

// map instance data and operations
// using a separate instance of input for potential need of original in part2
type AntennaFrequencyMap struct {
	Map         grid
	Frequencies frequencyMap
	Antinodes   antinodeMap
	Overlaps    antinodeMap
	yMax        int
	xMax        int
}

// init an instance
func NewAntennaFrequencyMap(g grid) *AntennaFrequencyMap {
	gridCopy := make(grid, len(g))
	copy(gridCopy, g)
	return &AntennaFrequencyMap{
		Map:         gridCopy,
		Frequencies: make(frequencyMap),
		Antinodes:   make(antinodeMap),
		Overlaps:    make(antinodeMap),
		yMax:        len(g),
		xMax:        len(g[0]),
	}
}

// put the position of each antenna in the Frequencies map
func (m *AntennaFrequencyMap) ScanFrequencies() {
	for y, l := range m.Map {
		for x := range l {

			f := m.Map[y][x]
			if f != '.' {
				fn := int(f)
				if m.Frequencies[fn] == nil {
					m.Frequencies[fn] = make(map[int]position)
				}
				i := len(m.Frequencies[fn])
				m.Frequencies[fn][i] = position{y, x}
			}

		}
	}
}

// put valid antinodes in Antinodes map
func (m *AntennaFrequencyMap) DetectAntinodes() {
	for fn := range m.Frequencies {
		m.detectFrequencyAntiNodes(fn)
	}
	m.removeInvalidAntinodes()
}

// counts antinodes
func (m *AntennaFrequencyMap) AntinodeCount() int {
	return len(m.Antinodes)
}

// traverse over the frequency's antenna position pairs and calculate the antinodes
func (m *AntennaFrequencyMap) detectFrequencyAntiNodes(fn int) {
	antennaPositions := len(m.Frequencies[fn])
	// first antenna iteraion
	for a1 := 0; a1 < antennaPositions; a1++ {
		for a2 := 0; a2 < antennaPositions; a2++ {
			if a1 == a2 {
				continue
			}
			anti1, anti2 := m.getAntinodes(m.Frequencies[fn][a1], m.Frequencies[fn][a2])
			if _, ok := m.Antinodes[anti1]; ok {
				m.Antinodes[anti1] = rune(fn)
			}
			if _, ok := m.Antinodes[anti2]; !ok {
				m.Antinodes[anti2] = rune(fn)
			}
		}
	}
}

// return antinode positions of 2 antennas
func (m *AntennaFrequencyMap) getAntinodes(antenna1, antenna2 position) (position, position) {
	dy := antenna2.y - antenna1.y
	dx := antenna2.x - antenna1.x

	antinode1 := position{
		y: antenna2.y + dy,
		x: antenna2.x + dx,
	}

	antinode2 := position{
		y: antenna1.y - dy,
		x: antenna1.x - dx,
	}

	return antinode1, antinode2
}

// puts the antinodes on the map for visualizations
func (m *AntennaFrequencyMap) SetAntinodes(r ...rune) {
	for p := range m.Antinodes {
		if len(r) > 0 {
			if m.Antinodes[p] == r[0] {
				continue
			}
		}
		m.Map[p.y][p.x] = '#'
	}
}

// apply 2 rules to keep only valid antinodes
func (m *AntennaFrequencyMap) removeInvalidAntinodes() {
	for p := range m.Antinodes {
		// remove out of boundary
		if p.y < 0 || p.y >= m.yMax || p.x < 0 || p.x >= m.xMax {
			delete(m.Antinodes, p)
			continue
		}
		// remove overlaps
		if _, ok := m.Overlaps[p]; ok {
			delete(m.Antinodes, p)
		}
	}
}

// visualize
func (m *AntennaFrequencyMap) PrintMap() {
	for y := range m.Map {
		for x := range m.Map[y] {
			switch m.Map[y][x] {
			case '#':
				fmt.Printf("%s%c%s", green, m.Map[y][x], reset)
			case '.':
				fmt.Printf("%c", m.Map[y][x])
			default:
				fmt.Printf("%s%c%s", red, m.Map[y][x], reset)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// parse the input into slice (lines) of int slice (elements in each line)
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

	antennaFrequencyMap := NewAntennaFrequencyMap(parse(input))

	antennaFrequencyMap.PrintMap()
	antennaFrequencyMap.ScanFrequencies()
	antennaFrequencyMap.DetectAntinodes()
	antennaFrequencyMap.SetAntinodes()
	antennaFrequencyMap.PrintMap()

	fmt.Printf("Part 1 : Number of valid antinodes = %d\n", antennaFrequencyMap.AntinodeCount())
}