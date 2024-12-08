package main

import (
	"reflect"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name   string
		opsMap fullStructure
		want   uint64
	}{
		{
			name: "sum of valid equation results",
			opsMap: fullStructure{
				190: halfStructure{
					0: 1, // 10 * 19
				},
				3267: halfStructure{
					0: 1, // 81 * 40 + 27
					1: 2, // 81 + 40 * 27
				},
				292: halfStructure{
					0: 5, // 11 + 6 * 16 + 20
				},
			},
			want: 3749, // 190 + 3267 + 292
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc(tt.opsMap)
			if got != tt.want {
				t.Errorf("calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateOpsMap(t *testing.T) {
	input := fullStructure{
		190: halfStructure{
			0: 10,
			1: 19,
		},
		3267: halfStructure{
			0: 81,
			1: 40,
			2: 27,
		},
		292: halfStructure{
			0: 11,
			1: 6,
			2: 16,
			3: 20,
		},
	}

	want := fullStructure{
		190: halfStructure{
			0: 0b1, // *
		},
		3267: halfStructure{
			0: 0b01, // + *
			1: 0b10, // * +
		},
		292: halfStructure{
			0: 0b010, // + * +
		},
	}

	got := generateOpsMap(input)

	if len(got) != len(want) {
		t.Errorf("generateOpsMap() returned %d valid equations, want %d", len(got), len(want))
	}

	for k, v := range want {
		if gotOps, exists := got[k]; !exists {
			t.Errorf("generateOpsMap() missing key %d", k)
		} else if len(gotOps) != len(v) {
			t.Errorf("generateOpsMap() for key %d = %b, want %b", k, gotOps, v)
		}
	}
}

func TestLineOps(t *testing.T) {
	tests := []struct {
		name      string
		testValue uint64
		numbers   halfStructure
		want      []uint64
	}{
		{
			name:      "two numbers multiply (190 = 10 * 19)",
			testValue: 190,
			numbers: halfStructure{
				0: 10,
				1: 19,
			},
			want: []uint64{0b1},
		},
		{
			name:      "three numbers with two ops (3267 = 81 + 40 * 27)",
			testValue: 3267,
			numbers: halfStructure{
				0: 81,
				1: 40,
				2: 27,
			},
			want: []uint64{0b01, 0b10},
		},
		{
			name:      "four numbers with three ops (292 = 11 + 6 * 16 + 20)",
			testValue: 292,
			numbers: halfStructure{
				0: 11,
				1: 6,
				2: 16,
				3: 20,
			},
			want: []uint64{0b010},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lineOps(tt.testValue, tt.numbers)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lineOps() = %b, want %b", got, tt.want)
			}
		})
	}
}

func TestIsValidOps(t *testing.T) {
	tests := []struct {
		name      string
		testValue uint64
		numbers   halfStructure
		op        uint64
		want      bool
	}{
		{
			name:      "valid two numbers multiply (190 = 10 * 19)",
			testValue: 190,
			numbers: halfStructure{
				0: 10,
				1: 19,
			},
			op:   0b1,
			want: true,
		},
		{
			name:      "valid three numbers (3267 = 81 + 40 * 27)",
			testValue: 3267,
			numbers: halfStructure{
				0: 81,
				1: 40,
				2: 27,
			},
			op:   0b01,
			want: true,
		},
		{
			name:      "valid four numbers (292 = 11 + 6 * 16 + 20)",
			testValue: 292,
			numbers: halfStructure{
				0: 11,
				1: 6,
				2: 16,
				3: 20,
			},
			op:   0b010,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidOps(tt.testValue, tt.numbers, tt.op)
			if got != tt.want {
				t.Errorf("isValidOps() = %v, want %v", got, tt.want)
			}
		})
	}
}