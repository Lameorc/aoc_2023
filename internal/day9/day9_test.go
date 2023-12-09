package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var solver = Day{}

func TestPart1(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}

	expected := "114"
	got := solver.Part1(input)

	require.Equal(t, expected, got)

}

func TestPart2(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}

	expected := "2"
	got := solver.Part2(input)

	require.Equal(t, expected, got)

}
