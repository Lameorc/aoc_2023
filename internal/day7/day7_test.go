package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var solver = Day{}

func TestPart1(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	expected := "6440"
	got := solver.Part1(input)

	require.Equal(t, expected, got)
}

func TestPart2(t *testing.T) {
	input := []string{
		// "32T3K 765",
		// "T55J5 684",
		// "KK677 28",
		// "KTJJT 220",
		// "QQQJA 483",
		"2345A 1",
		"Q2KJJ 13",
		"Q2Q2Q 19",
		"T3T3J 17",
		"T3Q33 11",
		"2345J 3",
		"J345A 2",
		"32T3K 5",
		"T55J5 29",
		"KK677 7",
		"KTJJT 34",
		"QQQJA 31",
		"JJJJJ 37",
		"JAAAA 43",
		"AAAAJ 59",
		"AAAAA 61",
		"2AAAA 23",
		"2JJJJ 53",
		"JJJJ2 41",
	}

	// expected := "5905"
	expected := "6839"
	got := solver.Part2(input)

	require.Equal(t, expected, got)
}
