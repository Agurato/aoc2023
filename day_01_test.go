package aoc2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Day01(t *testing.T) {
	d := Day01{}

	t.Run("Part 1 Example", func(t *testing.T) {
		lines, err := getInputExamplePart(1, 1, readLines)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 142, value)
	})

	t.Run("Part 1", func(t *testing.T) {
		lines, err := getInput(1, readLines)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 55816, value)
	})

	t.Run("Part 2 Example", func(t *testing.T) {
		lines, err := getInputExamplePart(1, 2, readLines)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 281, value)
	})

	t.Run("Part 2", func(t *testing.T) {
		lines, err := getInput(1, readLines)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 54980, value)
	})
}
