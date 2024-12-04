package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	testTable := []struct {
		name          string
		toTest        []string
		hExpected     int
		vExpected     int
		TRBLDExpected int
		TLBRDExpected int
	}{{
		name:          "1",
		toTest:        []string{"XMASSAMXMASAMXXXXMASSAMXMASAMX"},
		hExpected:     8,
		vExpected:     0,
		TRBLDExpected: 0,
		TLBRDExpected: 0,
	}, {
		name:          "2",
		toTest:        []string{"XXSS", "MMAA", "AAMM", "SSXX"},
		hExpected:     0,
		vExpected:     4,
		TRBLDExpected: 0,
		TLBRDExpected: 0,
	}, {
		name: "TopRightToBottomLeftAndReversed",
		toTest: []string{
			"XXXXXXX",
			"XMXXXXX",
			"XXAXXXX",
			"XXXSXXX",
			"XMXXAXX",
			"XXAXXMX",
			"XXXSXXX"},
		hExpected:     0,
		vExpected:     0,
		TRBLDExpected: 3,
		TLBRDExpected: 0,
	}, {
		name: "TopLeftToBottomRightAndReversed",
		toTest: []string{
			"XXXXXXX",
			"XXXXXMX",
			"XXXXAXX",
			"XXXSXXX",
			"XXAXXMX",
			"XMXXAXX",
			"XXXSXXX"},
		hExpected:     0,
		vExpected:     0,
		TRBLDExpected: 0,
		TLBRDExpected: 3,
	}}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			search := WordSearch{wordLines: tt.toTest}
			if tt.hExpected > 0 {
				search.FindAllXMASHorizontal()
				if search.xmasCount != tt.hExpected {
					t.Errorf("got %d, want %d", search.xmasCount, tt.hExpected)
				}
				search.xmasCount = 0
			}
			if tt.vExpected > 0 {
				search.FindAllXMASVertical()
				if search.xmasCount != tt.vExpected {
					t.Errorf("got %d, want %d", search.xmasCount, tt.vExpected)
				}
				search.xmasCount = 0
			}
			if tt.TRBLDExpected > 0 {
				search.FindAllXMASDiagonalTopLeftToBottomRight()
				if search.xmasCount != tt.TRBLDExpected {
					t.Errorf("got %d, want %d", search.xmasCount, tt.TRBLDExpected)
				}
				search.xmasCount = 0
			}
			if tt.TLBRDExpected > 0 {
				search.FindAllXMASDiagonalTopRightToBottomLeft()
				if search.xmasCount != tt.TLBRDExpected {
					t.Errorf("got %d, want %d", search.xmasCount, tt.TLBRDExpected)
				}
				search.xmasCount = 0
			}

		})
		fmt.Println()
	}
}
