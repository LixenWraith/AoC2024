package main

import (
	"fmt"
	"time"
)

// guard info
type Guard struct {
	InBounds      bool
	X, Y, Visited int
	Direction     byte
}

// movement directions, Gray code for simple bit-wise operation
const (
	DirUp    = 0b00
	DirRight = 0b01
	DirDown  = 0b10
	DirLeft  = 0b11
)

// guard position delta for one step move in any direciton
var (
	// indexed by direction bits
	dx = [4]int{0, 1, 0, -1} // Up, Right, Down, Left
	dy = [4]int{-1, 0, 1, 0} // Up, Right, Down, Left
)

// changes guard direction according to rules
func (g *Guard) rotate() {
	clockwise := (g.Direction + 1) & 0b11 // & drops bit2 for cycle
	g.Direction = clockwise
	if visualize {
		fmt.Printf("clockwise, new dir %2b\n", g.Direction)
	}
}

func (g *Guard) TraceBack(x, y int, d byte) (traceX, traceY int) {
	return x - dx[d], y - dy[d]
}

// returns false if can't move in the current direction, visit counted on moving away from location
func (g *Guard) move(m *Map) (bool, error) {
	newX, newY := g.X+dx[g.Direction], g.Y+dy[g.Direction]
	if visualize {
		fmt.Printf("( X %d , Y %d -> %d , %d ) , dir %2b , dx %d , dy %d\n",
			g.X, g.X, newX, newY, g.Direction, dx[g.Direction], dy[g.Direction])
	}

	if newX < 0 || newY < 0 || newX >= m.XMax || newY >= m.YMax {
		g.InBounds = false
		return true, nil
	} // position not changed for less code and avoiding potential out of bound slice referencing

	if m.Level[newX][newY] == '#' {
		if visualize {
			fmt.Printf("# wall @ ( %d , %d )\n", newX, newY)
			fmt.Printf("# current %c ahead %c visited %v\n", m.Level[g.X][g.Y], m.Level[newX][newY], g.Visited)
		}
		return false, nil
	}

	if visualize {
		fmt.Printf(" current %c ahead %c visited %v\n", m.Level[g.X][g.Y], m.Level[newX][newY], g.Visited)
	}

	g.X = newX
	g.Y = newY

	// count it if was not visited
	vl := m.VisitLog(g.X, g.Y)
	if vl == nil {
		g.Visited++
		m.SetVisited(g.X, g.Y, g.Direction)
	} else {
		if _, looped := vl[g.Direction]; looped {
			return false, fmt.Errorf("looped x %d y %d d %b", g.X, g.Y, g.Direction)
		}
	}

	if visualize {
		time.Sleep(500 * time.Millisecond)
	}

	return true, nil
}

// patrol the guard throughout the map till out of bounds
func (g *Guard) Patrol(m *Map) error {
	g.Visited++
	m.SetVisited(g.X, g.Y, g.Direction)

	for {
		if !g.InBounds {
			return nil
		}

		for {
			moved, err := g.move(m)
			if err != nil {
				return fmt.Errorf("patrol failed: %v", err)
			}
			if !moved {
				g.rotate()
			} else {
				break
			}
		}
	}
}