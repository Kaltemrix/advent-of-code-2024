package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

type Plotter struct {
	xDir                   int
	yDir                   int
	height                 int
	width                  int
	obstacles              [][2]int
	guardPos               [2]int
	distinctGuardPositions [][2]int

	originalObstacles [][2]int
	originalGuardPos  [2]int
	originalXDir      int
	originalYDir      int
}

func NewPlotter(rd io.Reader) (plot *Plotter) {
	scanner := bufio.NewScanner(rd)

	// [2]int{x, y}

	pl := &Plotter{
		xDir:                   0,
		yDir:                   1,
		obstacles:              [][2]int{},
		guardPos:               [2]int{},
		distinctGuardPositions: [][2]int{},
		width:                  0,
		height:                 0,
	}

	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if pl.width == 0 {
			pl.width = len(scanner.Text())
		}
		for i, char := range line {
			if char == '#' {
				// Add the position to the obstacles slice
				pl.obstacles = append(pl.obstacles, [2]int{i, lineNum})
			} else if char == '^' {
				// Add the position to the guardPos slice
				pl.guardPos = [2]int{i, lineNum}
				// Add the guardPos to the distinctGuardPositions slice
				pl.distinctGuardPositions = append(pl.distinctGuardPositions, pl.guardPos)
			}
		}
		lineNum++
	}

	pl.height = lineNum
	pl.originalObstacles = pl.obstacles
	pl.originalGuardPos = pl.guardPos
	pl.originalXDir = pl.xDir
	pl.originalYDir = pl.yDir

	return pl
}

func (pl *Plotter) Reset() {
	pl.obstacles = pl.originalObstacles
	pl.guardPos = pl.originalGuardPos
	pl.xDir = pl.originalXDir
	pl.yDir = pl.originalYDir
	pl.distinctGuardPositions = [][2]int{pl.originalGuardPos}
}

func (pl *Plotter) IsObstacle(pos [2]int) bool {
	for _, obstacle := range pl.obstacles {
		if pos == obstacle {
			return true
		}
	}
	return false
}

func (pl *Plotter) IsOffGrid(pos [2]int) bool {
	if pos[0] < 0 || pos[0] >= pl.width {
		return true
	}
	if pos[1] < 0 || pos[1] >= pl.height {
		return true
	}
	return false
}

func (pl *Plotter) AddDistinctGuardPosition(pos [2]int) {
	for _, distinctGuardPos := range pl.distinctGuardPositions {
		if distinctGuardPos == pos {
			return
		}
	}
	pl.distinctGuardPositions = append(pl.distinctGuardPositions, pos)
}

func (pl *Plotter) Move() (bool, error) {
	switch {
	case pl.xDir == 0 && pl.yDir == 1:
		// Going Up
		newPos := [2]int{pl.guardPos[0], pl.guardPos[1] - 1}
		if pl.IsObstacle(newPos) {
			// Turn Right
			pl.xDir = 1
			pl.yDir = 0
		} else if pl.IsOffGrid(newPos) {
			return false, fmt.Errorf("Guard has gone off the grid")
		} else {
			// Move Up
			pl.guardPos = newPos
			pl.AddDistinctGuardPosition(pl.guardPos)
		}
	case pl.xDir == 1 && pl.yDir == 0:
		// Going Right
		newPos := [2]int{pl.guardPos[0] + 1, pl.guardPos[1]}
		if pl.IsObstacle(newPos) {
			// Turn Down
			pl.xDir = 0
			pl.yDir = -1
		} else if pl.IsOffGrid(newPos) {
			return false, fmt.Errorf("Guard has gone off the grid")
		} else {
			// Move Right
			pl.guardPos = newPos
			pl.AddDistinctGuardPosition(pl.guardPos)
		}
	case pl.xDir == 0 && pl.yDir == -1:
		// Going Down
		newPos := [2]int{pl.guardPos[0], pl.guardPos[1] + 1}
		if pl.IsObstacle(newPos) {
			// Turn Left
			pl.xDir = -1
			pl.yDir = 0
		} else if pl.IsOffGrid(newPos) {
			return false, fmt.Errorf("Guard has gone off the grid")
		} else {
			// Move Down
			pl.guardPos = newPos
			pl.AddDistinctGuardPosition(pl.guardPos)
		}
	case pl.xDir == -1 && pl.yDir == 0:
		// Going Left
		newPos := [2]int{pl.guardPos[0] - 1, pl.guardPos[1]}
		if pl.IsObstacle(newPos) {
			// Turn Up
			pl.xDir = 0
			pl.yDir = 1
		} else if pl.IsOffGrid(newPos) {
			return false, fmt.Errorf("Guard has gone off the grid")
		} else {
			// Move Left
			pl.guardPos = newPos
			pl.AddDistinctGuardPosition(pl.guardPos)
		}
	default:
		panic("Invalid direction")
	}

	return true, nil
}

func (pl *Plotter) AddObstacle(pos [2]int) bool {
	// Check if the position is already an obstacle, it if is, return false
	// Check if the position is the guardPos, if it is, return false
	// Add the position to the obstacles slice, and note it on the tempObstacle
	// return true
	for _, obstacle := range pl.obstacles {
		if pos == obstacle {
			return false
		}
	}
	if pos == pl.guardPos {
		return false
	}
	pl.obstacles = append(pl.obstacles, pos)
	return true
}

func (pl *Plotter) RemoveObstacle(pos [2]int) {
	// Check if the tempObstacle is in the obstacles slice
	// If it is, remove it from the obstacles slice
	// Set the tempObstacle to nil
	for i, obstacle := range pl.obstacles {
		if pos == obstacle {
			pl.obstacles = append(pl.obstacles[:i], pl.obstacles[i+1:]...)
			return
		}
	}
}

func (pl *Plotter) Print() {
	for y := 0; y < pl.height; y++ {
		for x := 0; x < pl.width; x++ {
			if [2]int{x, y} == pl.guardPos {
				fmt.Print("^")
			} else if pl.IsObstacle([2]int{x, y}) {
				fmt.Print("#")
				// Print distinctGuardPositions as X
			} else if slices.Contains(pl.distinctGuardPositions, [2]int{x, y}) {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	loopCounter := 0
	maxMoves := 9000

	pl := NewPlotter(file)

	for y := 0; y < pl.width; y++ {
		for x := 0; x < pl.height; x++ {
			placed := pl.AddObstacle([2]int{x, y})
			if placed {
				for i := 0; i < maxMoves; i++ {
					fmt.Println(x, y, i)
					_, err := pl.Move()
					if err != nil {
						break
					}
					if i == maxMoves-1 {
						loopCounter++
					}
				}
				pl.RemoveObstacle([2]int{x, y})
			}
			pl.Reset()
		}
	}

	// fmt.Println(len(pl.distinctGuardPositions))
	fmt.Println("loops: ", loopCounter)

}
