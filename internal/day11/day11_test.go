package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {

	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	expected := "374"
	got := (&Day{}).Part1(input)

	require.Equal(t, expected, got)
}
