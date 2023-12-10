package day10

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var solver = Day{}

func TestPart1(t *testing.T) {
	tcs := []struct {
		input    []string
		expected string
	}{
		{
			input: []string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			},
			expected: "4",
		},
		{
			input: []string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			},
			expected: "4",
		},
		{
			input: []string{
				"..F7.",
				".FJ|.",
				"SJ.L7",
				"|F--J",
				"LJ...",
			},
			expected: "8",
		},
	}
	for idx, tc := range tcs {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			got := solver.Part1(tc.input)

			require.Equal(t, tc.expected, got)
		})
	}
}
