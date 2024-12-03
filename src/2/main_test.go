package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	testTable := []struct {
		name     string
		in       []int
		expected bool
	}{
		{"Decreasing", []int{7, 6, 4, 2, 1}, true},
		{"Increase more than 3", []int{1, 2, 7, 8, 9}, false},
		{"Decrease more than 3", []int{9, 7, 6, 2, 1}, false},
		{"Increase then decrease", []int{1, 3, 2, 4, 5}, false},
		{"Decrease then stay the same", []int{8, 6, 4, 4, 1}, false},
		{"Increasing", []int{1, 3, 6, 7, 9}, true},
	}

	for _, tt := range testTable {
		fmt.Print(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			rl := NewReactorLevel(tt.in)
			got := rl.IsSafe()
			if got != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, got)
			}
		})
	}
}
