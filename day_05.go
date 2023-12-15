package aoc2023

import (
	"math"
	"strconv"
	"strings"
)

type Day05 struct {
}

// Part1 is day 2 part 1
func (d Day05) Part1(lines []string) int {
	seeds, maps := d.ParseInput(lines)

	minLocation := math.MaxInt

	for _, seed := range seeds {
		val := seed
		for _, seedMapStep := range maps {
			for _, seedMap := range seedMapStep {
				newVal, found := seedMap.GetDestination(val)
				if found {
					val = newVal
					break
				}
			}
		}
		if val < minLocation {
			minLocation = val
		}
	}

	return minLocation
}

// Part2 is day 2 part 2
func (d Day05) Part2(lines []string) int {
	sum := 0

	return sum
}

func (d Day05) ParseInput(lines []string) (seeds []int, maps [7][]SeedMap) {
	for _, seedStr := range strings.Fields(lines[0][strings.Index(lines[0], ":")+2:]) {
		seed, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, seed)
	}

	i := 0
	for _, line := range lines[2:] {
		if len(line) == 0 {
			i++
			continue
		}
		if line[len(line)-1] == ':' {
			continue
		}
		fields := strings.Fields(line)
		destination, _ := strconv.Atoi(fields[0])
		source, _ := strconv.Atoi(fields[1])
		length, _ := strconv.Atoi(fields[2])
		maps[i] = append(maps[i], SeedMap{
			destination: destination,
			source:      source,
			length:      length,
		})
	}

	return seeds, maps
}

type SeedMap struct {
	destination, source, length int
}

func (sm SeedMap) GetDestination(value int) (int, bool) {
	if value >= sm.source && value < sm.source+sm.length {
		return sm.destination + (value - sm.source), true
	}
	return 0, false
}
