package aoc2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Day02(t *testing.T) {
	d := Day02{}

	t.Run("Part 1 Example", func(t *testing.T) {
		lines, err := getInputExample(2)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 8, value)
	})

	t.Run("Part 1", func(t *testing.T) {
		lines, err := getInput(2)
		require.NoError(t, err)

		value := d.Part1(lines)
		assert.Equal(t, 2156, value)
	})

	t.Run("Part 2 Example", func(t *testing.T) {
		lines, err := getInputExample(2)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 2286, value)
	})

	t.Run("Part 2", func(t *testing.T) {
		lines, err := getInput(2)
		require.NoError(t, err)

		value := d.Part2(lines)
		assert.Equal(t, 66909, value)
	})
}
