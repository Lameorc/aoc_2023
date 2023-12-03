package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var tester = Day3{}

func TestPart1(t *testing.T) {
	testInput := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expected := "4361"

	require.Equal(t, expected, tester.Part1(testInput))

}

func TestPart1Extensive(t *testing.T) {
	tcs := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name: "above and below",
			input: []string{
				"..100..",
				"...%...",
				"..150..",
			},
			expected: "250",
		},
		{
			name: "cross 1",
			input: []string{
				"100....",
				"...%...",
				"....150",
			},
			expected: "250",
		},
		{
			name: "cross 2",
			input: []string{
				"....150",
				"...%...",
				"100....",
			},
			expected: "250",
		},
		{
			name: "same corners",
			input: []string{
				"150....",
				"...%...",
				"100....",
			},
			expected: "250",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := tester.Part1(tc.input)

			require.Equal(t, tc.expected, got)
		})
	}
}

func TestPart2(t *testing.T) {

	testInput := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expected := "467835"

	require.Equal(t, expected, tester.Part2(testInput))
}
