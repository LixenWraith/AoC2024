package main

import "fmt"

type Simulator struct {
	original *Map
	sandbox  *Map
	puppet   *Guard
}

func NewSimulator(o *Map) *Simulator {
	return &Simulator{
		original: o,
		sandbox: &Map{
			Level:   o.Level,
			Visited: make(map[int]map[int]map[byte]bool),
			XMax:    o.XMax,
			YMax:    o.YMax,
		},
		puppet: &Guard{},
	}
}

func (s *Simulator) reset() {
	for x := range s.original.Level {
		s.sandbox.Level[x] = s.original.Level[x]
	}
	s.sandbox = &Map{
		Level:   s.original.Level,
		Visited: make(map[int]map[int]map[byte]bool),
		XMax:    s.original.XMax,
		YMax:    s.original.YMax,
	}
	s.sandbox.ClearMarks()
	s.puppet = &Guard{}
}

func (s *Simulator) simulate(x, y int) (failures int) {
	/*
		if x == 4 && y == 4 {
			visualize = true
		}

	*/
	passes := s.original.VisitLog(x, y)
	var d byte
	for d = 0; d < 4; d++ {
		if _, ok := passes[d]; ok {
			s.reset()

			s.puppet.X, s.puppet.Y = s.puppet.TraceBack(x, y, d)
			s.puppet.InBounds = true
			s.puppet.Direction = d

			s.sandbox.Level[x][y] = '#'

			if visualize {
				fmt.Printf("== Starting patrol: ( %d , %d ) <- ( X %d , Y %d ) dir %2b\n", s.puppet.X, s.puppet.Y, x, y, d)
			}

			if err := s.puppet.Patrol(s.sandbox); err != nil {
				fmt.Printf("== Error patrol ( %d , %d )\n", x, y)
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
	for x := 0; x < s.original.XMax; x++ {
		for y := 0; y < s.original.YMax; y++ {
			l += s.simulate(x, y)
		}
	}

	return l
}