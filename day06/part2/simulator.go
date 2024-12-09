package main

import "fmt"

type Simulator struct {
	original *Map
	sandbox  *Map
	puppet   *Guard
}

func NewSimulator(o *Map) *Simulator {
	originalCopy := make([][]rune, len(o.Level))
	for i, line := range o.Level {
		originalCopy[i] = make([]rune, len(line))
		copy(originalCopy[i], line)
	}

	return &Simulator{
		original: o,
		sandbox: &Map{
			Level:   o.Level,
			Visited: make(map[int]map[int]map[byte]bool),
			YMax:    o.YMax,
			XMax:    o.XMax,
		},
		puppet: &Guard{},
	}
}

func (s *Simulator) reset() {
	s.sandbox = &Map{
		Level:   make([][]rune, len(s.original.Level)),
		Visited: make(map[int]map[int]map[byte]bool),
		YMax:    s.original.YMax,
		XMax:    s.original.XMax,
	}
	for y := range s.original.Level {
		s.sandbox.Level[y] = make([]rune, len(s.original.Level[y]))
		copy(s.sandbox.Level[y], s.original.Level[y])
	}
	s.sandbox.ClearMarks()
	s.puppet = &Guard{}
}

func (s *Simulator) simulate(y, x int) (failures int) {
	/*
		if x == 4 && y == 4 {
			visualize = true
		}

	*/
	passes := s.original.VisitLog(y, x)
	var d byte
	for d = 0; d < 4; d++ {
		if _, ok := passes[d]; ok {
			s.reset()

			s.puppet.Y, s.puppet.X = s.puppet.TraceBack(y, x, d)
			s.puppet.InBounds = true
			s.puppet.Direction = d

			s.sandbox.Level[y][x] = '#'

			if visualize {
				fmt.Printf("== Starting patrol: ( %d , %d ) <- ( Y %d , X %d ) dir %2b\n", s.puppet.Y, s.puppet.X, y, x, d)
			}

			if err := s.puppet.Patrol(s.sandbox); err != nil {
				if visualize {
					fmt.Printf("== Error patrol ( %d , %d )\n", y, x)
				}
				if visualize {
					fmt.Println(err)
				}
				failures++
			}
		}
	}

	return failures
}

func (s *Simulator) Loops() int {
	l := 0
	for y := 0; y < s.original.YMax; y++ {
		for x := 0; x < s.original.XMax; x++ {
			l += s.simulate(y, x)
		}
	}

	return l
}