package main

import (
	"reflect"
	"testing"
)

func TestFrequency(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []map[int]int
	}{
		{
			name:     "Mixed",
			input:    [][]int{{1, 2, 1}, {3, 3, 3}},
			expected: []map[int]int{{1: 2, 2: 1}, {3: 3}},
		},
		{
			name:     "Same numbers",
			input:    [][]int{{5, 5, 5}, {5, 5, 5}},
			expected: []map[int]int{{5: 3}, {5: 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := frequency(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("result = %v, expected = %v", result, tt.expected)
			}
		})
	}
}
