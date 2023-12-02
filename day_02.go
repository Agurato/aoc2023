package aoc2023

import (
	"strconv"
	"strings"
)

type Day02 struct {
}

// Part1 is day 2 part 1
func (d Day02) Part1(lines []string) int {
	sum := 0

	maxR, maxG, maxB := 12, 13, 14

	for _, line := range lines {
		colonIndex := strings.IndexRune(line, ':')
		gameNumber, err := strconv.Atoi(line[5:colonIndex])
		if err != nil {
			panic(err)
		}

		lost := false
		grabs := strings.Split(line[colonIndex+1:], ";")
		for _, grab := range grabs {
			r, g, b := d.getCubes(grab)
			if r > maxR || g > maxG || b > maxB {
				lost = true
				break
			}
		}
		if !lost {
			sum += gameNumber
		}
	}

	return sum
}

// Part2 is day 2 part 2
func (d Day02) Part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		var minR, minG, minB int

		colonIndex := strings.IndexRune(line, ':')

		grabs := strings.Split(line[colonIndex+1:], ";")
		for _, grab := range grabs {
			r, g, b := d.getCubes(grab)
			if r > minR {
				minR = r
			}
			if g > minG {
				minG = g
			}
			if b > minB {
				minB = b
			}
		}

		power := minR * minG * minB
		sum += power
	}

	return sum
}

func (d Day02) getCubes(line string) (int, int, int) {
	colorGrabs := strings.Split(line, ", ")
	var r, g, b int

	for _, colorGrab := range colorGrabs {
		grabSplit := strings.Fields(colorGrab)
		number, err := strconv.Atoi(grabSplit[0])
		if err != nil {
			panic(err)
		}
		switch grabSplit[1] {
		case "red":
			r = number
		case "green":
			g = number
		case "blue":
			b = number
		}
	}

	return r, g, b
}
