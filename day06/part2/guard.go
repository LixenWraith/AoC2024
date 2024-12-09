package main

import (
	"fmt"
	"time"
)

// guard info
type Guard struct {
	InBounds      bool
	Y, X, Visited int
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
	dy = [4]int{-1, 0, 1, 0} // Up, Right, Down, Left
	dx = [4]int{0, 1, 0, -1} // Up, Right, Down, Left
)

// changes guard direction according to rules
func (g *Guard) rotate() {
	clockwise := (g.Direction + 1) & 0b11 // & drops bit2 for cycle
	g.Direction = clockwise
	if visualize {
		fmt.Printf("clockwise, new dir %2b\n", g.Direction)
	}
}

func (g *Guard) TraceBack(y, x int, d byte) (traceY, traceX int) {
	return y - dy[d], x - dx[d]
}

// returns false if can't move in the current direction, visit counted on moving away from location
func (g *Guard) move(m *Map) (bool, error) {
	newY, newX := g.Y+dy[g.Direction], g.X+dx[g.Direction]
	if visualize {
		fmt.Printf("( Y %d , X %d -> %d , %d ) , dir %2b , dy %d , dx %d\n",
			g.Y, g.X, newY, newX, g.Direction, dy[g.Direction], dx[g.Direction])
	}

	if newY < 0 || newY >= m.YMax || newX < 0 || newX >= m.XMax {
		g.InBounds = false
		return true, nil
	} // position not changed for less code and avoiding potential out of bound slice referencing

	if m.Level[newY][newX] == '#' {
		if visualize {
			fmt.Printf("# wall @ ( %d , %d )\n", newY, newX)
			fmt.Printf("# current %c ahead %c visited %v\n", m.Level[g.Y][g.X], m.Level[newY][newX], g.Visited)
		}
		return false, nil
	}

	if visualize {
		fmt.Printf(" current %c ahead %c visited %v\n", m.Level[g.Y][g.X], m.Level[newY][newX], g.Visited)
	}

	g.Y = newY
	g.X = newX

	// count it if was not visited
	vl := m.VisitLog(g.Y, g.X)
	if vl == nil {
		g.Visited++
		m.SetVisited(g.Y, g.X, g.Direction)
	} else {
		if _, looped := vl[g.Direction]; looped {
			return false, fmt.Errorf("looped y %d x %d d %b", g.Y, g.X, g.Direction)
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
	m.SetVisited(g.Y, g.X, g.Direction)

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