package main

import (
	"bufio"
	"errors"
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

	sizeOfComparingBlock := 0
	for i := len(blockMapCompressed) - 1; i > 0; i-- {
		if blockMapCompressed[i] == "." {
			continue
		}

		sizeOfComparingBlock++

		thisString := blockMapCompressed[i]
		nextString := blockMapCompressed[i-1]
		if nextString == "." || nextString != thisString {
			sliceToMove := blockMapCompressed[i : i+sizeOfComparingBlock]

			dotSlice, err := FindConsecutiveDotSliceOfSize(blockMapCompressed, sizeOfComparingBlock)
			if err != nil {
				sizeOfComparingBlock = 0
				continue
			}

			if dotSlice[0] >= i {
				sizeOfComparingBlock = 0
				continue
			}

			for j, c := range sliceToMove {
				blockMapCompressed[dotSlice[0]+j] = c
			}
			for j := 0; j < sizeOfComparingBlock; j++ {
				blockMapCompressed[i+j] = "."
			}
			sizeOfComparingBlock = 0
		} else {
			continue
		}
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
}

func FindConsecutiveDotSliceOfSize(blockMapCompressed []string, size int) ([2]int, error) {
	currentSize := 0
	for i := 0; i < len(blockMapCompressed); i++ {
		if blockMapCompressed[i] == "." {
			if currentSize == size {
				return [2]int{i - size, i}, nil
			}
		}
		if blockMapCompressed[i] != "." {
			if currentSize == size {
				return [2]int{i - size, i}, nil
			}
			currentSize = 0
			continue
		}
		currentSize++
	}
	return [2]int{}, errors.New("no consecutive slice of dots found")
}
