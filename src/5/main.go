package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	regex := regexp.MustCompile(`\d{2}`)

	rules := [][2]int{}
	updates := [][]int{}

	scanner := bufio.NewScanner(file)
	scanningRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			scanningRules = false
			continue
		}
		if scanningRules {
			matches := regex.FindAllString(line, -1)
			if len(matches) != 2 {
				panic("Invalid rule")
			}
			var rule [2]int
			for i, match := range matches {
				ruleInt, err := strconv.Atoi(match)
				if err != nil {
					panic(err)
				}
				rule[i] = ruleInt
			}
			rules = append(rules, rule)
		} else {
			matches := regex.FindAllString(line, -1)
			update := make([]int, len(matches))
			for i, match := range matches {
				updateInt, err := strconv.Atoi(match)
				if err != nil {
					panic(err)
				}
				update[i] = updateInt
			}
			updates = append(updates, update)
		}
	}

	middleNumberOfAlreadyCorrectSum := 0
	middleNumberOfIncorrectSum := 0
	for _, update := range updates {

		validRules := [][2]int{}
		for _, rule := range rules {
			if slices.Contains(update, rule[0]) && slices.Contains(update, rule[1]) {
				validRules = append(validRules, rule)
			}
		}

		updateIsInvalid := false
		for _, validRule := range validRules {
			indexOfFirstRule := slices.Index(update, validRule[0])
			if indexOfFirstRule == -1 {
				panic("Rule number not found in update")
			}
			if !slices.Contains(update[indexOfFirstRule+1:], validRule[1]) {
				updateIsInvalid = true
				break
			}
		}
		if !updateIsInvalid {
			middleNumber := update[int(math.Round(float64(len(update)/2)))]
			middleNumberOfAlreadyCorrectSum += middleNumber
		} else {
			for i := 0; i < len(validRules); i++ {
				ruleToCheck := validRules[i]
				firstIndex := slices.Index(update, ruleToCheck[0])
				secondIndex := slices.Index(update, ruleToCheck[1])
				diff := firstIndex - secondIndex
				if diff > 0 {
					update[firstIndex], update[secondIndex] = update[secondIndex], update[firstIndex]
					i = -1
				}
			}
			fmt.Println("Updated:", update)

			middleNumber := update[int(math.Round(float64(len(update)/2)))]
			middleNumberOfIncorrectSum += middleNumber
		}
	}

	fmt.Println(middleNumberOfAlreadyCorrectSum)
	fmt.Println(middleNumberOfIncorrectSum)
}
