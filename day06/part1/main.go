// AoC 2024, Day 6, Part 1, Lixen Wraith
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// for fun!
var visualize bool

// map info
type Map struct {
	level      [][]rune
	visited    map[int]map[int]bool // hash table for fast lookup
	xMax, yMax int
}

// checks if guard has visited the coords before
func (m *Map) wasVisited(x, y int) bool {
	return m.visited[x][y]
}

// record guard visit
func (m *Map) setVisited(x, y int, noMark ...bool) {
	if m.visited[x] == nil {
		m.visited[x] = make(map[int]bool)
	}

	if len(noMark) == 0 || (len(noMark) > 0 && !noMark[0]) {
		m.level[x][y] = 'x'
	}

	m.visited[x][y] = true
	if visualize {
		for j := -4; j < 5; j++ {
			for i := 0; i < m.xMax; i++ {
				fmt.Printf("%c", m.level[i][y+j])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

// visual output
func (m *Map) markVisited() {
	for x := range m.visited {
		for y := range m.visited[x] {
			m.level[x][y] = 'X'
		}
	}
}

// guard info
type Guard struct {
	inBounds      bool
	x, y, visited int
	direction     byte
}

// movement directions, Gray code for simple bit-wise operation
const (
	dirUp    = 0b00
	dirRight = 0b01
	dirDown  = 0b10
	dirLeft  = 0b11
)

// guard position delta for one step move in any direciton
var (
	// indexed by direction bits
	dx = [4]int{0, 1, 0, -1} // Up, Right, Down, Left
	dy = [4]int{-1, 0, 1, 0} // Up, Right, Down, Left
)

// changes guard direction according to rules
func (g *Guard) rotate() {
	clockwise := (g.direction + 1) & 0b11 // & drops bit2 for cycle
	g.direction = clockwise
	if visualize {
		fmt.Printf("clockwise, new dir %2b\n", g.direction)
	}
}

// returns false if can't move in the current direction, visit counted on moving away from location
func (g *Guard) move(m *Map) bool {
	// count it if was not visited
	if !m.wasVisited(g.x, g.y) {
		g.visited++
		m.setVisited(g.x, g.y)
	}

	newX, newY := g.x+dx[g.direction], g.y+dy[g.direction]
	if visualize {
		fmt.Printf("( X %d , Y %d -> %d , %d ) , dir %2b , dx %d , dy %d\n",
			g.x, g.y, newX, newY, g.direction, dx[g.direction], dy[g.direction])
	}

	if newX < 0 || newY < 0 || newX >= m.xMax || newY >= m.yMax {
		g.inBounds = false
		return true
	} // position not changed for less code and avoiding potential out of bound slice referencing

	if m.level[newX][newY] == '#' {
		if visualize {
			fmt.Printf("# wall @ ( %d , %d )\n", newX, newY)
			fmt.Printf("# current %c ahead %c visited %v\n", m.level[g.x][g.y], m.level[newX][newY], g.visited)
		}
		return false
	}

	if visualize {
		fmt.Printf(" current %c ahead %c visited %v\n", m.level[g.x][g.y], m.level[newX][newY], g.visited)
	}

	g.x = newX
	g.y = newY

	if visualize {
		time.Sleep(200 * time.Millisecond)
	}

	return true
}

// patrol the guard throughout the map till out of bounds
func (g *Guard) patrol(m *Map) {
	for {
		if !g.inBounds {
			return
		}

		for {
			if !g.move(m) {
				g.rotate()
			} else {
				break
			}
		}
	}
}

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
				g.inBounds = true
				g.x = x
				g.y = y
				g.direction = dirUp
				// g.visited = 1  // moved counter to the beginning of move method
				output[x][y] = '.'
			} else {
				output[x][y] = c
			}
		}
	}
	return &Map{
		level:   output,
		visited: make(map[int]map[int]bool),
		xMax:    len(output[0]),
		yMax:    len(output),
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

	for i := 0; i < len(m.level); i++ {
		// Convert each rune in the row to bytes
		content = append(content, string(m.level[i])...)
		// Add newline after each row except the last one
		if i < len(m.level)-1 {
			content = append(content, '\n')
		}
	}

	if err := os.WriteFile(filename, content, 0666); err != nil {
		return err
	}
	return nil
}

func main() {
	filename := "./input.txt"
	// filename := "./example.txt"
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
	if !guard.inBounds {
		fmt.Printf("No guards found, all safe!\n")
		os.Exit(0)
	}

	// for visual debug
	visualize = false
	guard.patrol(levelMap)

	levelMap.markVisited()

	if err = writeToFile("./output.txt", levelMap); err != nil {
		fmt.Printf("Error writing output file: %s\n", err)
	}

	fmt.Printf("Part 1: Guard visited positions = %d\n", guard.visited)
}