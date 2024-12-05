package main

import (
	"testing"
)

func TestIsCorrectProductionOrder(t *testing.T) {
	tests := []struct {
		name         string
		orders       map[int][]int
		productOrder []int
		want         bool
	}{
		{
			name: "valid order",
			orders: map[int][]int{
				1: {2, 3},
				2: {3},
				3: {},
			},
			productOrder: []int{1, 59, 2, 66, 5, 74, 3, 4},
			want:         true,
		},
		{
			name: "invalid order",
			orders: map[int][]int{
				1: {2, 3},
				2: {3},
				3: {},
			},
			productOrder: []int{1, 44, 3, 2, 66},
			want:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isCorrectProductionOrder(tt.orders, tt.productOrder)
			if got != tt.want {
				t.Errorf("isCorrectProductionOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		input2 string
		want   int
	}{
		{
			name: "Example",
			input1: `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13`,
			input2: `75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`,
			want: 143,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orders := orderMap(parse([]byte(tt.input1), "|"))
			productionOrders := parse([]byte(tt.input2), ",")

			got := midpageSum(orders, productionOrders)
			if got != tt.want {
				t.Errorf("isCorrectProductionOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}