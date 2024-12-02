package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

type PairList struct {
	scanner *bufio.Scanner
	left    []int
	right   []int
}

func NewPairList(rd io.Reader) *PairList {
	return &PairList{scanner: bufio.NewScanner(rd)}
}

func (pl *PairList) ScanAndSortSides() {
	var left, right []int
	for pl.scanner.Scan() {
		var x, y int
		_, err := fmt.Sscanf(pl.scanner.Text(), "%d %d", &x, &y)
		if err != nil {
			panic(err)
		}
		left = append(left, x)
		right = append(right, y)
	}
	// Sort left and right smallest to largest
	sort.Ints(left)
	sort.Ints(right)

	pl.left = left
	pl.right = right
}

func (pl *PairList) GetTotalDistances() int {
	pl.ScanAndSortSides()

	pairs := make([][2]int, len(pl.left))
	for i := 0; i < len(pl.left); i++ {
		pairs[i] = [2]int{pl.left[i], pl.right[i]}
	}

	// get the absolute difference between each pair, and total them
	// Left or right can be bigger, so we need to get the absolute difference
	total := 0
	for _, pair := range pairs {
		total += int(math.Abs(float64(pair[0] - pair[1])))
	}
	return total
}

func (pl *PairList) CalculateSimilarity() int {
	pl.ScanAndSortSides()

	rightMap := make(map[int]int)
	for _, r := range pl.right {
		// Count the number of times each right value appears
		rightMap[r]++
	}

	similarity := 0

	for _, l := range pl.left {
		if rightMap[l] > 0 {
			similarity += l * rightMap[l]
		}
	}

	return similarity
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pairList := NewPairList(file)
	// value := pairList.GetTotalDistances()
	value := pairList.CalculateSimilarity()
	fmt.Println(value)
}
