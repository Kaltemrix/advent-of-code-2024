package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Scanning each character of each line
	// If we encounter a M, check if line+1 index+1 is A, and line+2 index+2 is S
	//  Next check if line+0 index+2 is an M or an S
	//  If it's an M, check if line+2 index+0 is an S
	//  If it's an S, check if line+2 index+0 is an M
	// If all checks pass, increment the count
	// If we encounter an S, check if line+1 index+1 is an A, and line+2 index+2 is an M
	//  Next check if line+0 index+2 is an S or an M
	//  If it's an S, check if line+2 index+0 is an M
	//  If it's an M, check if line+2 index+0 is an S
	// If all checks pass, increment the count
	// Also, we need to check if we're not going out of bounds
	// If we do, continue to the next iteration

	count := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if j+2 < len(lines[i]) && i+2 < len(lines) {
				if lines[i][j] == 'M' && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'S' {
					if lines[i][j+2] == 'M' {
						if lines[i+2][j] == 'S' {
							count++
						}
					} else if lines[i][j+2] == 'S' {
						if lines[i+2][j] == 'M' {
							count++
						}
					}
				} else if lines[i][j] == 'S' && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'M' {
					if lines[i][j+2] == 'S' {
						if lines[i+2][j] == 'M' {
							count++
						}
					} else if lines[i][j+2] == 'M' {
						if lines[i+2][j] == 'S' {
							count++
						}
					}
				}
			}
		}
	}

	println(count)
}
