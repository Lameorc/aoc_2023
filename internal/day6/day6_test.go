package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var tester = Day{}

func TestPart1(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	expected := "288"

	got := tester.Part1(input)
	require.Equal(t, expected, got)

}

func TestPart2(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	expected := "71503"

	got := tester.Part2(input)
	require.Equal(t, expected, got)

}
