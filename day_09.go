package aoc2023

import (
	"strconv"
	"strings"
)

type Day09 struct {
}

// Part1 is day 2 part 1
func (d Day09) Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		var numbers []int
		for _, val := range strings.Fields(line) {
			num, _ := strconv.Atoi(val)
			numbers = append(numbers, num)
		}
		nextNumber := d.getNextNumber(numbers)
		sum += nextNumber
	}

	return sum
}

// Part2 is day 2 part 2
func (d Day09) Part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		var numbers []int
		for _, val := range strings.Fields(line) {
			num, _ := strconv.Atoi(val)
			numbers = append(numbers, num)
		}
		nextNumber := d.getPreviousNumber(numbers)
		sum += nextNumber
	}

	return sum
}

func (d Day09) getNextNumber(numbers []int) int {
	newNumbers := make([]int, len(numbers)-1)
	for i := 0; i < len(newNumbers); i++ {
		newNumbers[i] = numbers[i+1] - numbers[i]
	}

	if d.allZeroes(newNumbers) {
		return numbers[len(numbers)-1]
	}

	return numbers[len(numbers)-1] + d.getNextNumber(newNumbers)
}

func (d Day09) getPreviousNumber(numbers []int) int {
	newNumbers := make([]int, len(numbers)-1)
	for i := 0; i < len(newNumbers); i++ {
		newNumbers[i] = numbers[i+1] - numbers[i]
	}

	if d.allZeroes(newNumbers) {
		return numbers[0]
	}

	return numbers[0] - d.getPreviousNumber(newNumbers)
}

func (d Day09) allZeroes(numbers []int) bool {
	for _, num := range numbers {
		if num != 0 {
			return false
		}
	}
	return true
}
