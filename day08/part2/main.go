// AoC 2024, Day 8, Part 2, Lixen Wraith
package main

import (
	"fmt"
	"math"
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

// counts antinodes
func (m *AntennaFrequencyMap) AntinodeCount() int {
	return len(m.Antinodes)
}

// put valid antinodes in Antinodes map
func (m *AntennaFrequencyMap) DetectAntinodes() {
	for fn := range m.Frequencies {
		m.detectFrequencyAntiNodes(fn)
	}

	m.removeInvalidAntinodes()
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

			antinodeHarmonics := m.getAntinodes(m.Frequencies[fn][a1], m.Frequencies[fn][a2])
			for _, a := range antinodeHarmonics {
				if _, ok := m.Antinodes[a]; !ok && m.inBounds(a) {
					m.Antinodes[a] = rune(fn)
				}
			}

		}
	}
}

// return antinode positions of 2 antennas
func (m *AntennaFrequencyMap) getAntinodes(antenna1, antenna2 position) (atinodeHarmonics []position) {
	dy := antenna2.y - antenna1.y
	dx := antenna2.x - antenna1.x
	harmonicRangeY := int(math.Abs(float64(m.yMax/dy))) + 1
	harmonicRangeX := int(math.Abs(float64(m.xMax/dx))) + 1
	ah := []position{}

	for i := min(-harmonicRangeY, -harmonicRangeX); i < max(harmonicRangeY, harmonicRangeX); i++ {
		hy := antenna1.y - (i * dy)
		hx := antenna1.x - (i * dx)
		if m.inBounds(position{hy, hx}) {
			ah = append(ah, position{hy, hx})
		}
	}

	return ah
}

func (m *AntennaFrequencyMap) inBounds(p position) bool {
	if p.y < 0 || p.y >= m.yMax || p.x < 0 || p.x >= m.xMax {
		return false
	}
	return true
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
		if !m.inBounds(p) {
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

	fmt.Printf("Part 2 : Number of valid antinode harmonics = %d\n", antennaFrequencyMap.AntinodeCount())
}