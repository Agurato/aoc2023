package aoc2023

import (
	"strings"
	"unicode"
)

func Day01Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		runes := []rune(line)

		for _, r := range runes {
			if unicode.IsDigit(r) {
				sum += 10 * int(r-'0')
				break
			}
		}
		for i := len(runes) - 1; i >= 0; i-- {
			r := runes[i]
			if unicode.IsDigit(r) {
				sum += int(r - '0')
				break
			}
		}
	}

	return sum
}

func Day01Part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		runes := []rune(line)

		for i, r := range runes {
			if unicode.IsDigit(r) {
				sum += 10 * int(r-'0')
				break
			}
			val, found := checkSpelledOutDigits(runes[i:])
			if found {
				sum += 10 * val
				break
			}
		}
		for i := len(runes) - 1; i >= 0; i-- {
			r := runes[i]
			if unicode.IsDigit(r) {
				sum += int(r - '0')
				break
			}
			val, found := checkSpelledOutDigits(runes[i:])
			if found {
				sum += val
				break
			}
		}
	}

	return sum
}

func checkSpelledOutDigits(runes []rune) (int, bool) {
	s := string(runes)
	switch {
	case strings.HasPrefix(s, "one"):
		return 1, true
	case strings.HasPrefix(s, "two"):
		return 2, true
	case strings.HasPrefix(s, "three"):
		return 3, true
	case strings.HasPrefix(s, "four"):
		return 4, true
	case strings.HasPrefix(s, "five"):
		return 5, true
	case strings.HasPrefix(s, "six"):
		return 6, true
	case strings.HasPrefix(s, "seven"):
		return 7, true
	case strings.HasPrefix(s, "eight"):
		return 8, true
	case strings.HasPrefix(s, "nine"):
		return 9, true
	}
	return 0, false
}
