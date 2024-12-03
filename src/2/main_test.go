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
		{"Increase then decrease", []int{1, 3, 2, 4, 5}, true},
		{"Decrease then stay the same", []int{8, 6, 4, 4, 1}, true},
		{"Increasing", []int{1, 3, 6, 7, 9}, true},
		{"Same at start", []int{1, 1, 2, 3, 4}, true},
		{"Same at start", []int{1, 1, 6, 7, 8}, false},
		{"1", []int{4, 6, 5, 4, 3}, true},
	}

	for _, tt := range testTable {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			reactor := Reactor{}
			reactor.AddReport(NewReactorReport(tt.in))
			count := reactor.GetSafeReportCount()

			if tt.expected && count != 1 {
				t.Errorf("Expected 1, got %d", count)
			} else if !tt.expected && count != 0 {
				t.Errorf("Expected 0, got %d", count)
			}
		})
		fmt.Println()
	}
}
