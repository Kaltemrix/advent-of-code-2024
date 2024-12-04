package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

type WordSearch struct {
	wordLines []string
	xmasCount int
	widerBy   int
}

func (ws *WordSearch) FindAllXMASInLines(lines []string) int {
	regex, err := regexp.Compile(`XMAS`)
	if err != nil {
		panic(err)
	}
	count := 0
	for _, line := range lines {
		lineCopy := line
		// Make a copy of the line, check for "XMAS", then reverse the line and check for "XMAS" again
		matches := regex.FindAllString(lineCopy, -1)
		count += len(matches)
		// Reverse the line
		reversedLine := ""
		for i := len(lineCopy) - 1; i >= 0; i-- {
			reversedLine += string(lineCopy[i])
		}
		matches = regex.FindAllString(reversedLine, -1)
		count += len(matches)
	}
	return count
}

func (ws *WordSearch) FindAllXMASHorizontal() {
	count := ws.FindAllXMASInLines(ws.wordLines)
	ws.xmasCount += count
}

func (ws *WordSearch) FindAllXMASVertical() {
	width := len(ws.wordLines)
	verticalLines := make([]string, width)
	for _, line := range ws.wordLines {
		for i := 0; i < width; i++ {
			verticalLines[i] += string(line[i])
		}
	}

	count := ws.FindAllXMASInLines(verticalLines)
	ws.xmasCount += count
}

func (ws *WordSearch) FindAllXMASDiagonalTopLeftToBottomRight() {
	width := len(ws.wordLines[0]) - ws.widerBy
	if ws.widerBy < 0 {
		width = len(ws.wordLines) + ws.widerBy
	}

	diagonalLines := []string{}

	for i := 0; i < width; i++ {
		thisLine := ""
		for j := 0; j <= i; j++ {
			thisLine += string(ws.wordLines[j][width-i+j-1])
		}
		diagonalLines = append(diagonalLines, thisLine)
	}

	if ws.widerBy > 0 {
		for i := 0; i < ws.widerBy; i++ {
			thisLine := ""
			for j := 0; j < width; j++ {
				thisLine += string(ws.wordLines[j][ws.widerBy-i+j-1])
			}
			diagonalLines = append(diagonalLines, thisLine)
		}
	} else if ws.widerBy < 0 {
		absWiderBy := int(math.Abs(float64(ws.widerBy)))
		for i := 0; i < absWiderBy; i++ {
			thisLine := ""
			for j := 0; j < width; j++ {
				thisLine += string(ws.wordLines[absWiderBy+i+j-1][j])
			}
			diagonalLines = append(diagonalLines, thisLine)
		}
	}

	for i := 0; i < width-1; i++ {
		thisLine := ""
		for j := 0; j <= i; j++ {
			thisLine += string(ws.wordLines[len(ws.wordLines)-i+j-1][j])
		}
		diagonalLines = append(diagonalLines, thisLine)
	}

	count := ws.FindAllXMASInLines(diagonalLines)
	ws.xmasCount += count
}

func (ws *WordSearch) FindAllXMASDiagonalTopRightToBottomLeft() {
	width := len(ws.wordLines[0]) - ws.widerBy
	if ws.widerBy < 0 {
		width = len(ws.wordLines) + ws.widerBy
	}

	diagonalLines := []string{}

	for i := 0; i < width; i++ {
		thisLine := ""
		for j := 0; j <= i; j++ {
			thisLine += string(ws.wordLines[j][i-j])
		}
		diagonalLines = append(diagonalLines, thisLine)
	}

	if ws.widerBy > 0 {
		for i := 0; i < ws.widerBy; i++ {
			thisLine := ""
			for j := 0; j < width; j++ {
				thisLine += string(ws.wordLines[j][width+i-j])
			}
			diagonalLines = append(diagonalLines, thisLine)
		}
	} else if ws.widerBy < 0 {
		absWiderBy := int(math.Abs(float64(ws.widerBy)))
		for i := 0; i < absWiderBy; i++ {
			thisLine := ""
			for j := 0; j < width; j++ {
				thisLine += string(ws.wordLines[absWiderBy+i+j][j])
			}
			diagonalLines = append(diagonalLines, thisLine)
		}
	}

	for i := 0; i < width-1; i++ {
		thisLine := ""
		for j := 0; j <= i; j++ {
			thisLine += string(ws.wordLines[width-i+j-1][width-j-1])
		}
		diagonalLines = append(diagonalLines, thisLine)
	}

	count := ws.FindAllXMASInLines(diagonalLines)
	ws.xmasCount += count
}

func main_old() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wordSearch := WordSearch{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wordSearch.wordLines = append(wordSearch.wordLines, line)
	}

	// Lets find out if the search is wider or taller, and by how much
	widerBy := len(wordSearch.wordLines[0]) - len(wordSearch.wordLines)
	wordSearch.widerBy = widerBy

	wordSearch.FindAllXMASHorizontal()
	wordSearch.FindAllXMASVertical()
	wordSearch.FindAllXMASDiagonalTopLeftToBottomRight()
	wordSearch.FindAllXMASDiagonalTopRightToBottomLeft()

	fmt.Println(wordSearch.xmasCount)
}
