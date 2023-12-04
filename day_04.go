package aoc2023

import (
	"math"
	"slices"
	"strings"
)

type Day04 struct {
}

// Part1 is day 2 part 1
func (d Day04) Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		colonIndex := strings.IndexRune(line, ':')
		game := strings.Split(line[colonIndex:], " | ")
		winning := strings.Fields(game[0])

		wins := 0
		for _, drawn := range strings.Fields(game[1]) {
			if slices.Contains(winning, drawn) {
				wins++
			}
		}
		if wins != 0 {
			sum += int(math.Pow(2, float64(wins-1)))
		}
	}

	return sum
}

// Part2 is day 2 part 2
func (d Day04) Part2(lines []string) int {
	sum := 0

	winsPerCard := make([]int, len(lines))
	for i, line := range lines {
		colonIndex := strings.IndexRune(line, ':')
		game := strings.Split(line[colonIndex:], " | ")
		winning := strings.Fields(game[0])

		wins := 0
		for _, drawn := range strings.Fields(game[1]) {
			if slices.Contains(winning, drawn) {
				wins++
			}
		}
		winsPerCard[i] = wins
	}

	numberOfCards := make([]int, len(lines))
	for i := range numberOfCards {
		numberOfCards[i] = 1
	}
	for i, wins := range winsPerCard {
		for j := 0; j < wins; j++ {
			numberOfCards[i+j+1] += numberOfCards[i]
		}
		sum += numberOfCards[i]
	}

	return sum
}
