package main

import (
	"fmt"
	"strconv"
	"strings"
)

// map info
type Map struct {
	Level      [][]rune
	Visited    map[int]map[int]map[byte]bool // x, y -> y, x, direction, visitited in direction
	YMax, XMax int
}

// checks if guard has visited the coords before
func (m *Map) VisitLog(y, x int) map[byte]bool {
	return m.Visited[y][x]
}

// record guard visit
func (m *Map) SetVisited(y, x int, d byte, noMark ...bool) {
	if m.Visited[y] == nil {
		m.Visited[y] = make(map[int]map[byte]bool)
	}
	if m.Visited[y][x] == nil {
		m.Visited[y][x] = make(map[byte]bool)
	}

	if len(noMark) == 0 || (len(noMark) > 0 && !noMark[0]) {
		m.Level[y][x] = 'x'
	}

	m.Visited[y][x][d] = true
	if visualize {
		linesPrinted := 0
		for j := -11; j < 12; j++ {
			if y+j >= 0 && y+j < m.YMax {
				for i := 0; i < m.XMax; i++ {
					fmt.Printf("%c", m.Level[y+j][i])
				}
				linesPrinted++
				fmt.Println()
			}
		}
		linesPrinted++
		fmt.Println()
		// go to the beginning of map print
		fmt.Print(strings.Join([]string{"\033[", strconv.Itoa(linesPrinted + 2), "A"}, ""))
	}
}

// visual output
func (m *Map) MarkVisited() {
	for y := range m.Visited {
		for x := range m.Visited[y] {
			m.Level[y][x] = 'X'
		}
	}
}

// clear guard marks from map
func (m *Map) ClearMarks() {
	for y := 0; y < m.YMax; y++ {
		for x := 0; x < m.XMax; x++ {
			// fmt.Printf("y %d x %d m.YMax %d m.XMax %d \n", y, x, m.YMax, m.XMax)
			if strings.ToLower(string(m.Level[y][x])) == "x" {
				m.Level[y][x] = '.'
			}
		}
	}
}