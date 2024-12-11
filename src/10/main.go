package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

type TopographicMap struct {
	Topography [][]int
	TrailHeads []*TrailHead
}

type TrailHead struct {
	X         int
	Y         int
	endsFound [][2]int
	rating    int
}

func NewTopographicMap(rd io.Reader) *TopographicMap {
	scanner := bufio.NewScanner(rd)

	var topography [][]int

	lineNum := 0
	for scanner.Scan() {
		var topoLine []int
		line := scanner.Text()
		for _, c := range line {
			char, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			topoLine = append(topoLine, char)
		}
		topography = append(topography, topoLine)
		lineNum++
	}
	return &TopographicMap{Topography: topography}
}

func (t *TopographicMap) FindTrailHeads() {
	for i, row := range t.Topography {
		for j, point := range row {
			if point == 0 {
				trailhead := &TrailHead{X: j, Y: i, endsFound: [][2]int{}, rating: 0}
				t.TrailHeads = append(t.TrailHeads, trailhead)
				// Have found the start of a trail
				// Follow the trail
				t.FollowTrail(trailhead, [2]int{j, i}, 1)
			}
		}
	}
}

func (t *TopographicMap) FollowTrail(trailHead *TrailHead, pointFrom [2]int, stepToCheckFor int) {
	// Follow the trail from the starting point (x, y)
	possibleDirections := [][2]int{}
	// Look up, down, left, right for i, and add to possibleDirections
	if pointFrom[1]-1 >= 0 && t.Topography[pointFrom[1]-1][pointFrom[0]] == stepToCheckFor {
		if stepToCheckFor == 9 {
			trailHead.AddEndFound([2]int{pointFrom[0], pointFrom[1] - 1})
			trailHead.rating++
		} else {
			// Look up
			possibleDirections = append(possibleDirections, [2]int{pointFrom[0], pointFrom[1] - 1})
		}
	}
	if pointFrom[1]+1 < len(t.Topography) && t.Topography[pointFrom[1]+1][pointFrom[0]] == stepToCheckFor {
		if stepToCheckFor == 9 {
			trailHead.AddEndFound([2]int{pointFrom[0], pointFrom[1] + 1})
			trailHead.rating++
		} else {
			// Look down
			possibleDirections = append(possibleDirections, [2]int{pointFrom[0], pointFrom[1] + 1})
		}
	}
	if pointFrom[0]-1 >= 0 && t.Topography[pointFrom[1]][pointFrom[0]-1] == stepToCheckFor {
		if stepToCheckFor == 9 {
			trailHead.AddEndFound([2]int{pointFrom[0] - 1, pointFrom[1]})
			trailHead.rating++
		} else {
			// Look left
			possibleDirections = append(possibleDirections, [2]int{pointFrom[0] - 1, pointFrom[1]})
		}
	}
	if pointFrom[0]+1 < len(t.Topography[0]) && t.Topography[pointFrom[1]][pointFrom[0]+1] == stepToCheckFor {
		if stepToCheckFor == 9 {
			trailHead.AddEndFound([2]int{pointFrom[0] + 1, pointFrom[1]})
			trailHead.rating++
		} else {
			// Look right
			possibleDirections = append(possibleDirections, [2]int{pointFrom[0] + 1, pointFrom[1]})
		}
	}

	// Follow every possible direction
	for _, direction := range possibleDirections {
		t.FollowTrail(trailHead, direction, stepToCheckFor+1)
	}
}

func (h *TrailHead) AddEndFound(end [2]int) {
	// Add to trailhead ends found if not already found
	if !slices.Contains(h.endsFound, end) {
		h.endsFound = append(h.endsFound, end)
	}
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	topographicMap := NewTopographicMap(file)
	topographicMap.FindTrailHeads()
	totalScore := 0
	totalRating := 0
	for _, trailHead := range topographicMap.TrailHeads {
		totalScore += len(trailHead.endsFound)
		totalRating += trailHead.rating
	}
	fmt.Println(totalRating)
	fmt.Println(totalScore)
}
