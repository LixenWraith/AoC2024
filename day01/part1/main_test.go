package main

import (
	"reflect"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name          string
		lists         [][]int
		totalDistance int
	}{
		{
			name: "Example",
			lists: [][]int{{3, 4, 2, 1, 3, 3},
				{4, 3, 5, 3, 9, 3}},
			totalDistance: 11,
		},
		{
			name:          "Empty slices test",
			lists:         [][]int{{}, {}},
			totalDistance: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := distance(tt.lists)

			if !reflect.DeepEqual(result, tt.totalDistance) {
				t.Errorf("result = %v, expected = %v", result, tt.totalDistance)
			}
		})
	}
}
