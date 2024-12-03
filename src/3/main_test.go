package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMain(t *testing.T) {

	testTable := []struct {
		name     string
		regex    *regexp.Regexp
		in       string
		expected string
	}{}

	for _, tt := range testTable {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {

		})
		fmt.Println()
	}
}
