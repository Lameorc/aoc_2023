package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var solver = &Day{}

func TestPart1(t *testing.T) {
	tcs := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name: "case 1",
			input: []string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			expected: "2",
		},
		{
			name: "case 2",
			input: []string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			expected: "6",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := solver.Part1(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}

	expected := "6"
	got := solver.Part2(input)

	require.Equal(t, expected, got)
}
