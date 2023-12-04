package aoc2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Day04(t *testing.T) {
	d := Day04{}

	t.Run("Part 1 Example", func(t *testing.T) {
		lines, err := getInputExample(4, readLines)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 13, value)
	})

	t.Run("Part 1", func(t *testing.T) {
		lines, err := getInput(4, readLines)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 23847, value)
	})

	t.Run("Part 2 Example", func(t *testing.T) {
		lines, err := getInputExample(4, readLines)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 30, value)
	})

	t.Run("Part 2", func(t *testing.T) {
		lines, err := getInput(4, readLines)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 8570000, value)
	})
}
