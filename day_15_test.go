package aoc2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Day15(t *testing.T) {
	d := Day15{}

	t.Run("Part 1 Example", func(t *testing.T) {
		lines, err := getInputExample(15, readCommas)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 1320, value)
	})

	t.Run("Part 1", func(t *testing.T) {
		lines, err := getInput(15, readCommas)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 515210, value)
	})

	t.Run("Part 2 Example", func(t *testing.T) {
		lines, err := getInputExample(15, readCommas)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 145, value)
	})

	t.Run("Part 2", func(t *testing.T) {
		lines, err := getInput(15, readCommas)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 246762, value)
	})
}
