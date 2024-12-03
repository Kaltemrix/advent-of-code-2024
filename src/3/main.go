package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type MyRegexes struct {
	MulRegex     *regexp.Regexp
	NumbersRegex *regexp.Regexp
}

func NewMyRegexes() *MyRegexes {
	mulRegex, err := regexp.Compile(`(mul\(\d{1,3}\,\d{1,3}\))`)
	if err != nil {
		panic(err)
	}
	numbersRegex, err := regexp.Compile(`(\d{1,3})`)
	if err != nil {
		panic(err)
	}

	return &MyRegexes{MulRegex: mulRegex, NumbersRegex: numbersRegex}
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Scan all the lines in the file into a string
	scanner := bufio.NewScanner(file)

	goodString := ""
	badString := ""

	shouldAppend := true
	for scanner.Scan() {
		line := scanner.Text()

		for _, char := range line {
			if shouldAppend {
				goodString += string(char)
				if len(goodString) >= 7 && goodString[len(goodString)-7:] == "don't()" {
					shouldAppend = false
					goodString = goodString[:len(goodString)-7]
					// goodString += "|"
				}
			} else {
				badString += string(char)
				// Check if badString ends with "do()"
				if len(badString) >= 4 && badString[len(badString)-4:] == "do()" {
					shouldAppend = true
					badString = badString[:len(badString)-4]
				}
			}
		}
	}

	myRegexes := NewMyRegexes()

	operations := [][]int{}

	operationStrings := myRegexes.MulRegex.FindAllString(goodString, -1)
	if operationStrings == nil {
		fmt.Println("No operations found")
		return
	}

	for _, operationString := range operationStrings {
		numbers := myRegexes.NumbersRegex.FindAllString(operationString, -1)
		if numbers == nil {
			fmt.Println("No numbers found")
			return
		}

		operation := []int{}
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				panic(err)
			}
			operation = append(operation, num)
		}

		operations = append(operations, operation)
	}

	sum := 0
	for _, operation := range operations {
		sum += operation[0] * operation[1]
	}

	fmt.Println(sum)
}
