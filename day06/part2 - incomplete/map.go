package main

import (
	"fmt"
	"strings"
)

// map info
type Map struct {
	Level      [][]rune
	Visited    map[int]map[int]map[byte]bool // hash table for fast lookup
	XMax, YMax int
}

// checks if guard has visited the coords before
func (m *Map) VisitLog(x, y int) map[byte]bool {
	return m.Visited[x][y]
}

// record guard visit
func (m *Map) SetVisited(x, y int, d byte, noMark ...bool) {
	if m.Visited[x] == nil {
		m.Visited[x] = make(map[int]map[byte]bool)
	}
	if m.Visited[x][y] == nil {
		m.Visited[x][y] = make(map[byte]bool)
	}

	if len(noMark) == 0 || (len(noMark) > 0 && !noMark[0]) {
		m.Level[x][y] = 'x'
	}

	m.Visited[x][y][d] = true
	if visualize {
		for j := -11; j < 12; j++ {
			if y+j >= 0 && y+j < m.XMax {
				for i := 0; i < m.XMax; i++ {
					fmt.Printf("%c", m.Level[i][y+j])
				}
				fmt.Println()
			}
		}
		fmt.Println()
	}
}

// visual output
func (m *Map) MarkVisited() {
	for x := range m.Visited {
		for y := range m.Visited[x] {
			m.Level[x][y] = 'X'
		}
	}
}

// clear guard marks from map
func (m *Map) ClearMarks() {
	for x := 0; x < m.XMax; x++ {
		for y := 0; y < m.YMax; y++ {
			if strings.ToLower(string(m.Level[x][y])) == "x" {
				m.Level[x][y] = '.'
			}
		}
	}
}