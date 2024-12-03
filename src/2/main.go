package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Reactor struct {
	Reports []*ReactorReport
}

func (rr *Reactor) AddReport(report *ReactorReport) {
	rr.Reports = append(rr.Reports, report)
}

func (rr *Reactor) GetSafeReportCount() int {
	safeCount := 0

	for _, report := range rr.Reports {
		// Check the original report first, if its safe, increment the safeCount and continue
		safe := report.IsSafe()
		if safe {
			safeCount++
			continue
		}

		// Remove each element from the report, and then check if it's safe
		// If it's safe, we found a safe permutation, increment the safeCount
		for i := 0; i < len(report.OriginalLevel); i++ {
			report.RemoveLevelAtIndex(i)
			safe = report.IsSafe()
			if safe {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

type ReactorReport struct {
	OriginalLevel []int
	Levels        []int
	Direction     int
}

func NewReactorReport(levels []int) *ReactorReport {
	dir := levels[1] - levels[0]
	return &ReactorReport{Levels: levels, Direction: dir, OriginalLevel: levels}
}

func (rl *ReactorReport) DirectionIsDown() bool {
	return rl.Direction < 0
}

func (rl *ReactorReport) RecalculateDirection() {
	rl.Direction = rl.Levels[1] - rl.Levels[0]
}

func (rl *ReactorReport) DifferenceMoreThan(index1, index2, diff int) bool {
	return int(math.Abs(float64(rl.Levels[index1]-rl.Levels[index2]))) > diff
}

func (rl *ReactorReport) RemoveLevelAtIndex(index int) {
	copyOfOriginal := make([]int, len(rl.OriginalLevel))
	copy(copyOfOriginal, rl.OriginalLevel)
	rl.Levels = append(copyOfOriginal[:index], copyOfOriginal[index+1:]...)
	rl.RecalculateDirection()
}

func (rl *ReactorReport) IsSafe() bool {
	// Exit early, difference is 0 and thats not safe
	if rl.Direction == 0 {
		return false
	}
	// Exit early, difference is already more than 3
	if rl.DifferenceMoreThan(0, 1, 3) {
		return false
	}

	levelSafe := true

	for j, level := range rl.Levels {
		// Skip the first level, we already have the direction
		if j == 0 {
			continue
		}
		if (level == rl.Levels[j-1]) || rl.DifferenceMoreThan(j, j-1, 3) {
			// If the level is the same as the previous level, or the difference is more than 3
			// Then the level is unsafe, so break
			return false
		}
		if rl.DirectionIsDown() {
			// Direction is negative, so subsequent levels should be decreasing
			// So if it's not, mark the level as unsafe and break
			if level > rl.Levels[j-1] {
				return false
			}
		} else {
			// Direction is positive, so subsequent levels should be increasing
			// So if it's not, mark the level as unsafe and break
			if level < rl.Levels[j-1] {
				return false
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
	reactor := Reactor{}

	for scanner.Scan() {
		var levels []int
		line := scanner.Text()
		fields := strings.Fields(line)

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return
			}
			levels = append(levels, num)
		}
		reactor.AddReport(NewReactorReport(levels))
	}

	safeCount := reactor.GetSafeReportCount()

	fmt.Print(safeCount)
}

// 1 2 3 4 5 6
// [1, 2] 1 - 2 = -1 (indicating neg direction)
// [1, 1] 1 - 1 = 0 (no change, already unsafe)
// [9, 7] 9 - 7 = 2 (indicating positive direction)
