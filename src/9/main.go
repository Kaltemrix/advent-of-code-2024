package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	diskMap := scanner.Text()

	var blockMap []string
	for i, c := range diskMap {
		lengthOfBlock, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		newBlock := make([]string, lengthOfBlock)
		if i%2 == 0 {
			for j := 0; j < lengthOfBlock; j++ {
				newBlock[j] = strconv.Itoa(i / 2)
			}
		} else {
			for j := 0; j < lengthOfBlock; j++ {
				newBlock[j] = "."
			}
		}
		blockMap = append(blockMap, newBlock...)
	}

	blockMapCompressed := make([]string, len(blockMap))
	copy(blockMapCompressed, blockMap)

	for i := len(blockMapCompressed) - 1; i > 0; i-- {
		if blockMapCompressed[i] == "." {
			continue
		}
		indexOfDot := 0
		for j := 0; j < len(blockMapCompressed); j++ {
			if blockMapCompressed[j] == "." {
				indexOfDot = j
				break
			}
		}
		// If the dot is after i, we're at the end of the blockMap, and don't need to do anything
		if indexOfDot >= i {
			break
		}

		// Replace the dot with the number
		blockMapCompressed[indexOfDot] = blockMapCompressed[i]
		// Replace the number with a dot
		blockMapCompressed[i] = "."
	}

	checksum := 0
	for i, c := range blockMapCompressed {
		if c == "." {
			continue
		}
		block, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		checksum += i * block
	}

	fmt.Println(checksum)

	// fmt.Println(blockMapCompressed[:10])
	// fmt.Println(blockMap[:10])
}
