package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	testTable := []struct {
		name string
		// in       []int
		// expected bool
	}{}

	for _, tt := range testTable {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {

		})
		fmt.Println()
	}
}
