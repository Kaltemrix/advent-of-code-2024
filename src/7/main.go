package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

const (
	ADD    = '+'
	MUL    = '*'
	CONCAT = "||"
)

type Equation struct {
	sum        int
	values     []int
	operations [][]string
	iterations int
}

type EquationList struct {
	equations []*Equation
}

func AddEquations(rd io.Reader) (el *EquationList) {
	el = &EquationList{}

	regex := regexp.MustCompile(`\d+`)

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		// first number is sum, rest are values
		nums := regex.FindAllString(scanner.Text(), -1)
		sum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		values := []int{}
		for _, num := range nums[1:] {
			value, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			values = append(values, value)
		}

		allOperations := [][]string{}

		for i := 0; i < int(math.Pow(3, float64(len(values)-1))); i++ {
			ops := []string{}
			for j := 0; j < len(values)-1; j++ {
				switch i / int(math.Pow(3, float64(j))) % 3 {
				case 0:
					ops = append(ops, string(ADD))
				case 1:
					ops = append(ops, string(MUL))
				case 2:
					ops = append(ops, CONCAT)
				}
			}

			allOperations = append(allOperations, ops)
		}

		eq := &Equation{
			sum:    sum,
			values: values,
			// valid:      false,
			operations: allOperations,
			iterations: 0,
		}
		el.equations = append(el.equations, eq)
	}

	return el
}

func (eq *Equation) AttemptSolve(opList []string) int {
	total := eq.values[0]
	for i, value := range eq.values[1:] {
		total = eq.SolvePair(total, value, opList[i])
		if total == -1 {
			return -1
		}
	}
	if total == eq.sum {
		return total
	}
	return -1
}

func (*Equation) SolvePair(a, b int, op string) int {
	switch op {
	case string(ADD):
		return a + b
	case string(MUL):
		return a * b
	case CONCAT:
		val, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
		if err != nil {
			return -1
		}
		return val
	}
	panic("Invalid operation")
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total := 0

	el := AddEquations(file)
	for _, eq := range el.equations {
		for _, ops := range eq.operations {
			if solved := eq.AttemptSolve(ops); solved != -1 {
				total += solved
				break
			}
		}
	}

	println(total)
}
