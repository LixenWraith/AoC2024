package main

import (
	"testing"
)

func TestCrawl(t *testing.T) {
	tests := []struct {
		name       string
		grid       [][]rune
		word       string
		matchCount int
	}{
		/*
			{
				name: "Example",
				grid: [][]rune{
					[]rune("MMMSXXMASM"),
					[]rune("MSAMXMSMSA"),
					[]rune("AMXSXMAAMM"),
					[]rune("MSAMASMSMX"),
					[]rune("XMASAMXAMM"),
					[]rune("XXAMMXXAMA"),
					[]rune("SMSMSASXSS"),
					[]rune("SAXAMASAAA"),
					[]rune("MAMMMXMMMM"),
					[]rune("MXMXAXMASX"),
				},
				word:       "XMAS",
				matchCount: 18,
			},

		*/
		{
			name:       "simple",
			grid:       [][]rune{{'A', 'B'}},
			word:       "AB",
			matchCount: 1,
		},
		{
			name: "simple horizontal match",
			grid: [][]rune{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			word:       "ABC",
			matchCount: 1,
		},
		{
			name: "no match",
			grid: [][]rune{
				{'X', 'Y', 'Z'},
				{'A', 'B', 'C'},
			},
			word:       "CAT",
			matchCount: 0,
		},
		{
			name: "multiple matches",
			grid: [][]rune{
				{'D', 'O', 'G'},
				{'O', 'O', 'O'},
				{'G', 'O', 'D'},
			},
			word:       "DOG",
			matchCount: 4,
		},
		{
			name: "empty word",
			grid: [][]rune{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:       "",
			matchCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := crawl(tt.grid, tt.word)

			if result != tt.matchCount {
				t.Errorf("result = %v, expected = %v", result, tt.matchCount)
			}
		})
	}
}
