package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
)

type AntennaMap struct {
	AntennasCoordinates     map[string][][2]int
	UniqueAntinodeLocations [][2]int
	width                   int
	height                  int
}

func NewAntennaMap(rd io.Reader) *AntennaMap {
	antennaMap := &AntennaMap{
		AntennasCoordinates: make(map[string][][2]int),
	}
	scanner := bufio.NewScanner(rd)
	lineNumber := -1
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		antennaMap.width = len(line) - 1
		regex := regexp.MustCompile(`[A-Z]|[a-z]|[0-9]`)
		matches := regex.FindAllStringIndex(line, -1)
		for _, match := range matches {
			letter := line[match[0]:match[1]]
			antennaMap.AntennasCoordinates[letter] = append(antennaMap.AntennasCoordinates[letter], [2]int{match[0], lineNumber})
		}
	}
	antennaMap.height = lineNumber
	return antennaMap
}

func (am *AntennaMap) FindAntinodeLocations() {
	for _, antennaCoords := range am.AntennasCoordinates {
		// Make unique pairs of antenna coordinates
		pairs := [][2][2]int{}
		for i := 0; i < len(antennaCoords); i++ {
			for j := i + 1; j < len(antennaCoords); j++ {
				pair := [2][2]int{antennaCoords[i], antennaCoords[j]}
				pairs = append(pairs, pair)
			}
		}
		// Draw and place nodes
		for _, pair := range pairs {
			am.DrawAndPlaceNodes(pair)
		}
	}
}

func (am *AntennaMap) DrawAndPlaceNodes(pair [2][2]int) {
	width := pair[1][0] - pair[0][0]
	height := pair[1][1] - pair[0][1]

	higherLocations := [][2]int{}
	for {
		location := [2]int{pair[1][0] - width*(len(higherLocations)+1), pair[1][1] - height*(len(higherLocations)+1)}
		if valid := am.CheckBounds(location); valid {
			higherLocations = append(higherLocations, location)
		} else {
			break
		}
	}
	lowerLocations := [][2]int{}
	for {
		location := [2]int{pair[0][0] + width*(len(lowerLocations)+1), pair[0][1] + height*(len(lowerLocations)+1)}
		if valid := am.CheckBounds(location); valid {
			lowerLocations = append(lowerLocations, location)
		} else {
			break
		}
	}

	// Append unique antinode locations
	for _, location := range higherLocations {
		if !slices.Contains(am.UniqueAntinodeLocations, location) {
			am.UniqueAntinodeLocations = append(am.UniqueAntinodeLocations, location)
		}
	}
	for _, location := range lowerLocations {
		if !slices.Contains(am.UniqueAntinodeLocations, location) {
			am.UniqueAntinodeLocations = append(am.UniqueAntinodeLocations, location)
		}
	}
}

func (am *AntennaMap) CheckBounds(location [2]int) bool {
	// Check not -1 or greater than the length of the map
	if location[0] < 0 || location[0] > am.width {
		return false
	}
	if location[1] < 0 || location[1] > am.height {
		return false
	}
	return true
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	antennaMap := NewAntennaMap(file)
	antennaMap.FindAntinodeLocations()
	fmt.Println(len(antennaMap.UniqueAntinodeLocations))
}
