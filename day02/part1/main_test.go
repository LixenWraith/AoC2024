package main

import (
	"reflect"
	"testing"
)

func TestSafe(t *testing.T) {
	tests := []struct {
		name         string
		reports      [][]int
		safeRerports int
	}{
		{
			name: "Example",
			reports: [][]int{{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			safeRerports: 2,
		},
		{
			name: "Unequal report sizes",
			reports: [][]int{{1}, // safe
				{4, 3, 2, 2},
				{9, 7}, // safe
				{1, 3, 2},
				{8, 6, 4, 4},
				{1, 3, 6, 7, 9}, // safe
			},
			safeRerports: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safe(tt.reports)

			if err != nil {
				t.Errorf("got error: %v\n", err)
				return
			}

			if !reflect.DeepEqual(result, tt.safeRerports) {
				t.Errorf("result = %v, expected = %v", result, tt.safeRerports)
			}
		})
	}
}
