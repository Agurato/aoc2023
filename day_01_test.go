package aoc2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Day01Part1_Example(t *testing.T) {
	lines, err := getInputExamplePart(1, 1)
	require.NoError(t, err)

	value := Day01Part1(lines)
	assert.Equal(t, 142, value)
}

func Test_Day01Part1(t *testing.T) {
	lines, err := getInput(1)
	require.NoError(t, err)

	value := Day01Part1(lines)
	assert.Equal(t, 55816, value)
}

func Test_Day01Part2_Example(t *testing.T) {
	lines, err := getInputExamplePart(1, 2)
	require.NoError(t, err)

	value := Day01Part2(lines)
	assert.Equal(t, 281, value)
}

func Test_Day01Part2(t *testing.T) {
	lines, err := getInput(1)
	require.NoError(t, err)

	value := Day01Part2(lines)
	assert.Equal(t, 54980, value)
}
