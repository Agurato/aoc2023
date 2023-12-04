package aoc2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Day03(t *testing.T) {
	d := Day03{}

	t.Run("Part 1 Example", func(t *testing.T) {
		runes, err := getInputExample(3, readRunes)
		require.NoError(t, err)

		value := d.Part1(runes)
		assert.Equal(t, 4361, value)
	})

	t.Run("Part 1", func(t *testing.T) {
		runes, err := getInput(3, readRunes)
		require.NoError(t, err)

		value := d.Part1(runes)
		assert.Equal(t, 544664, value)
	})

	t.Run("Part 2 Example", func(t *testing.T) {
		runes, err := getInputExample(3, readRunes)
		require.NoError(t, err)

		value := d.Part2(runes)
		assert.Equal(t, 467835, value)
	})

	t.Run("Part 2", func(t *testing.T) {
		runes, err := getInput(3, readRunes)
		require.NoError(t, err)

		value := d.Part2(runes)
		assert.Equal(t, 84495585, value)
	})
}
