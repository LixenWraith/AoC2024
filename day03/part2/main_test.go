package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name          string
		content       []byte
		parsedContent [][]int
	}{
		{
			name:          "Example",
			content:       []byte("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"),
			parsedContent: [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}},
		},
		{
			name:          "Example 2",
			content:       []byte("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"),
			parsedContent: [][]int{{2, 4}, {8, 5}},
		},
		{
			name:          "Bad",
			content:       []byte("who(1,2)$#mul2,5)aamul()mul(23 ,8)mul( 3,22)"),
			parsedContent: [][]int{},
		},
		{
			name:          "Simple bad",
			content:       []byte("x"),
			parsedContent: [][]int{},
		},
		{
			name:          "Simple",
			content:       []byte("mul(1,2)"),
			parsedContent: [][]int{{1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parse(tt.content)

			if !reflect.DeepEqual(result, tt.parsedContent) {
				t.Errorf("result = %v, expected = %v", result, tt.parsedContent)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	tests := []struct {
		name          string
		parsedContent [][]int
		expected      int
	}{

		{
			name:          "Example",
			parsedContent: [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}},
			expected:      161,
		},
		{
			name:          "Simple",
			parsedContent: [][]int{},
			expected:      0,
		},
		{
			name:          "Simple",
			parsedContent: [][]int{{1, 2}},
			expected:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc(tt.parsedContent)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("result = %v, expected = %v", result, tt.expected)
			}
		})
	}
}
