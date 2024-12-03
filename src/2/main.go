package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ReactorLevel struct {
	Level     []int
	Direction int
}

func NewReactorLevel(level []int) *ReactorLevel {
	dir := level[1] - level[0]
	return &ReactorLevel{Level: level, Direction: dir}
}

func (rl *ReactorLevel) DirectionIsDown() bool {
	return rl.Direction < 0
}

func (rl *ReactorLevel) DifferenceMoreThan(index1, index2, diff int) bool {
	return int(math.Abs(float64(rl.Level[index1]-rl.Level[index2]))) > diff
}

func (rl *ReactorLevel) IsSafe() bool {
	// Exit early, difference is 0 and thats not safe
	if rl.Direction == 0 {
		return false
	}
	// Exit early, difference is already more than 3
	if rl.DifferenceMoreThan(1, 2, 3) {
		return false
	}

	levelSafe := true
	for j, report := range rl.Level {
		// Skip the first report, we already have the direction
		if j == 0 {
			continue
		}
		if (report == rl.Level[j-1]) || rl.DifferenceMoreThan(j, j-1, 3) {
			// If the report is the same as the previous report, or the difference is more than 3
			// Then the level is unsafe, so break
			levelSafe = false
			break
		}
		if rl.DirectionIsDown() {
			// Direction is negative, so subsequent reports should be decreasing
			// So if it's not, mark the level as unsafe and break
			if report > rl.Level[j-1] {
				levelSafe = false
				break
			}
		} else {
			// Direction is positive, so subsequent reports should be increasing
			// So if it's not, mark the level as unsafe and break
			if report < rl.Level[j-1] {
				levelSafe = false
				break
			}
		}
	}

	return levelSafe
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var levels [][]int
	for scanner.Scan() {
		var reports []int
		line := scanner.Text()
		fields := strings.Fields(line)

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return
			}
			reports = append(reports, num)
		}
		levels = append(levels, reports)

	}

	safeCount := 0

	for _, level := range levels {
		rl := NewReactorLevel(level)
		if rl.IsSafe() {
			safeCount++
		}
	}

	fmt.Print(safeCount)
}

// 1 2 3 4 5 6
// [1, 2] 1 - 2 = -1 (indicating neg direction)
// [1, 1] 1 - 1 = 0 (no change, already unsafe)
// [9, 7] 9 - 7 = 2 (indicating positive direction)
