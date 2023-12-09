package aoc2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Day09(t *testing.T) {
	d := Day09{}

	t.Run("Part 1 Example", func(t *testing.T) {
		lines, err := getInputExample(9, readLines)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 114, value)
	})

	t.Run("Part 1", func(t *testing.T) {
		lines, err := getInput(9, readLines)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 1762065988, value)
	})

	t.Run("Part 2 Example", func(t *testing.T) {
		lines, err := getInputExample(9, readLines)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 2, value)
	})

	t.Run("Part 2", func(t *testing.T) {
		lines, err := getInput(9, readLines)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 1066, value)
	})
}
