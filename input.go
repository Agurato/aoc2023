package aoc2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput[T any](day int, reader func(string) (T, error)) (T, error) {
	path := fmt.Sprintf("input/day_%02d.txt", day)
	return reader(path)
}

func getInputExample[T any](day int, reader func(string) (T, error)) (T, error) {
	path := fmt.Sprintf("input/day_%02d_example.txt", day)
	return reader(path)
}

func getInputExamplePart[T any](day, part int, reader func(string) (T, error)) (T, error) {
	path := fmt.Sprintf("input/day_%02d_example_%d.txt", day, part)
	return reader(path)
}

// readLines returns slice of the file's lines
func readLines(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.ReplaceAll(string(content), "\r", ""), "\n"), nil
}

// readLines returns slice of the file's lines
func readRunes(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	reader := bufio.NewReader(file)

	var allRunes [][]rune
	var currentRunes []rune
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if r == '\r' {
			continue
		}
		if r == '\n' {
			allRunes = append(allRunes, currentRunes)
			currentRunes = nil
			continue
		}
		currentRunes = append(currentRunes, r)
	}
	allRunes = append(allRunes, currentRunes)

	return allRunes, nil
}
