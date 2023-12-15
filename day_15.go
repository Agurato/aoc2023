package aoc2023

import (
	"slices"
)

type Day15 struct {
}

// Part1 is day 15 part 1
func (d Day15) Part1(steps []string) int {
	sum := 0

	for _, step := range steps {
		sum += d.calcHash(step)
	}

	return sum
}

// Part2 is day 15 part 2
func (d Day15) Part2(steps []string) int {
	sum := 0

	var boxes [256][]Lens

	for _, step := range steps {
		if step[len(step)-1] == '-' {
			label := step[:len(step)-1]
			hash := d.calcHash(label)
			boxes[hash] = slices.DeleteFunc(boxes[hash], func(l Lens) bool {
				return l.label == label
			})
		} else {
			label := step[:len(step)-2]
			hash := d.calcHash(label)
			lens := Lens{
				label: label,
				focal: int(step[len(step)-1]) - 48,
			}
			lensIndex := slices.IndexFunc(boxes[hash], func(l Lens) bool {
				return l.label == label
			})
			if lensIndex == -1 {
				boxes[hash] = append(boxes[hash], lens)
			} else {
				boxes[hash][lensIndex] = lens
			}
		}
	}

	for boxIndex, box := range boxes {
		for lensIndex, lens := range box {
			sum += (boxIndex + 1) * (lensIndex + 1) * lens.focal
		}
	}

	return sum
}

func (d Day15) calcHash(s string) int {
	value := 0
	for _, char := range s {
		value += int(char)
		value *= 17
		value %= 256
	}
	return value
}

type Lens struct {
	label string
	focal int
}
