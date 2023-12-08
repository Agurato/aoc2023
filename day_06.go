package aoc2023

import (
	"math"
	"strconv"
	"strings"
)

type Day06 struct {
}

// Part1 is day 2 part 1
func (d Day06) Part1(lines []string) int {
	total := 1

	durations := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])

	for i := 1; i < len(durations); i++ {
		duration, _ := strconv.ParseFloat(durations[i], 64)
		distanceToBeat, _ := strconv.ParseFloat(distances[i], 64)
		total *= d.waysToBeat(duration, distanceToBeat)
	}

	return total
}

// Part2 is day 2 part 2
func (d Day06) Part2(lines []string) int {
	duration, _ := strconv.ParseFloat(strings.ReplaceAll(strings.SplitN(lines[0], " ", 2)[1], " ", ""), 64)
	distanceToBeat, _ := strconv.ParseFloat(strings.ReplaceAll(strings.SplitN(lines[1], " ", 2)[1], " ", ""), 64)
	return d.waysToBeat(duration, distanceToBeat)
}

func (d Day06) waysToBeat(duration, distance float64) int {
	delta := math.Sqrt(math.Pow(duration, 2) - 4*distance)
	minDuration := (-duration + delta) / (-2)
	maxDuration := (-duration - delta) / (-2)

	biggerThan := math.Ceil(minDuration)
	if biggerThan == minDuration {
		biggerThan += 1
	}
	lessThan := math.Floor(maxDuration)
	if lessThan == maxDuration {
		lessThan -= 1
	}

	return int(lessThan - biggerThan + 1)
}
