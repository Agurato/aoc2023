package aoc2023

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getInput(day int) ([]string, error) {
	path := fmt.Sprintf("input/day_%02d.txt", day)
	return readLines(path)
}

func getInputExample(day int) ([]string, error) {
	path := fmt.Sprintf("input/day_%02d_example.txt", day)
	return readLines(path)
}

func getInputExamplePart(day, part int) ([]string, error) {
	path := fmt.Sprintf("input/day_%02d_example_%d.txt", day, part)
	return readLines(path)
}

// readLines returns slice of the file's lines
func readLines(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.ReplaceAll(string(content), "\r", ""), "\n"), nil
}
