package aoc2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Day05(t *testing.T) {
	d := Day05{}

	t.Run("Part 1 Example", func(t *testing.T) {
		lines, err := getInputExample(5, readLines)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 35, value)
	})

	t.Run("Part 1", func(t *testing.T) {
		lines, err := getInput(5, readLines)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 199602917, value)
	})

	t.Run("Part 2 Example", func(t *testing.T) {
		lines, err := getInputExample(5, readLines)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 0, value)
	})

	t.Run("Part 2", func(t *testing.T) {
		lines, err := getInput(5, readLines)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 0, value)
	})
}
