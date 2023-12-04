package aoc2023

import (
	"fmt"
	"strconv"
	"unicode"
)

type Day03 struct {
}

// Part1 is day 3 part 1
func (d Day03) Part1(engine [][]rune) int {
	sum := 0

	currentNumber := 0
	isAdjacent := false
	for i, line := range engine {
		for j, val := range line {
			if unicode.IsDigit(val) {
				currentNumber = currentNumber*10 + int(val-'0')
				if !isAdjacent {
					isAdjacent = d.checkForAdjacentSymbol(&engine, len(line), len(engine), i, j)
				}
			} else {
				if isAdjacent {
					sum += currentNumber
					currentNumber = 0
					isAdjacent = false
				} else {
					currentNumber = 0
				}
			}
		}
		if isAdjacent {
			sum += currentNumber
			currentNumber = 0
			isAdjacent = false
		} else {
			currentNumber = 0
		}
	}
	if isAdjacent {
		sum += currentNumber
		currentNumber = 0
		isAdjacent = false
	}

	return sum
}

// Part2 is day 3 part 2
func (d Day03) Part2(engine [][]rune) int {
	sum := 0

	for i, line := range engine {
		for j, val := range line {
			if val != '*' {
				continue
			}
			numbers := d.getAdjacentNumbers(&engine, len(line), len(engine), i, j)
			if len(numbers) == 2 {
				sum += numbers[0] * numbers[1]
			}
		}
	}

	return sum
}

type engineTuple struct {
	i, j int
}

func (d Day03) checkForAdjacentSymbol(engine *[][]rune, width, height, i, j int) bool {
	checks := []engineTuple{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j - 1},
		{i, j},
		{i, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	}

	var found bool
	for _, check := range checks {
		if check.i >= 0 && check.i < height && check.j >= 0 && check.j < width {
			val := (*engine)[check.i][check.j]
			if !unicode.IsDigit(val) && val != '.' {
				found = true
				break
			}
		}
	}

	return found
}

func (d Day03) getAdjacentNumbers(engine *[][]rune, width, height, i, j int) []int {
	checks := map[engineTuple]bool{
		{i - 1, j - 1}: false,
		{i - 1, j}:     false,
		{i - 1, j + 1}: false,
		{i, j - 1}:     false,
		{i, j}:         false,
		{i, j + 1}:     false,
		{i + 1, j - 1}: false,
		{i + 1, j}:     false,
		{i + 1, j + 1}: false,
	}

	var numbers []int

	for check, checked := range checks {
		if checked || !d.isInRange(width, height, check.i, check.j) {
			continue
		}

		val := (*engine)[check.i][check.j]
		if unicode.IsDigit(val) {
			currentNumber := string(val)
			for x := check.j - 1; x >= 0; x-- {
				if !d.isInRange(width, height, x, j) {
					break
				}
				otherVal := (*engine)[check.i][x]
				tuple := engineTuple{check.i, x}
				if _, found := checks[tuple]; found {
					checks[tuple] = true
				}
				if unicode.IsDigit(otherVal) {
					currentNumber = fmt.Sprintf("%c%s", otherVal, currentNumber)
				} else {
					break
				}
			}
			for x := check.j + 1; x < width; x++ {
				if !d.isInRange(width, height, x, j) {
					break
				}
				otherVal := (*engine)[check.i][x]
				tuple := engineTuple{check.i, x}
				if _, found := checks[tuple]; found {
					checks[tuple] = true
				}
				if unicode.IsDigit(otherVal) {
					currentNumber = fmt.Sprintf("%s%c", currentNumber, otherVal)
				} else {
					break
				}
			}
			num, _ := strconv.Atoi(currentNumber)
			numbers = append(numbers, num)
			currentNumber = ""
		}
		checks[check] = true
	}

	return numbers
}

func (d Day03) isInRange(width, height, i, j int) bool {
	return i >= 0 && i < height && j >= 0 && j < width
}
