package aoc2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Day06(t *testing.T) {
	d := Day06{}

	t.Run("Part 1 Example", func(t *testing.T) {
		lines, err := getInputExample(6, readLines)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 288, value)
	})

	t.Run("Part 1", func(t *testing.T) {
		lines, err := getInput(6, readLines)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 741000, value)
	})

	t.Run("Part 2 Example", func(t *testing.T) {
		lines, err := getInputExample(6, readLines)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 71503, value)
	})

	t.Run("Part 2", func(t *testing.T) {
		lines, err := getInput(6, readLines)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 38220708, value)
	})
}
