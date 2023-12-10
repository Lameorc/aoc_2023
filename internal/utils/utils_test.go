package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringMap(t *testing.T) {
	input := []string{
		"AAA",
		"BBB",
		"CCC",
	}
	expected := &StringMap{
		values: [][]string{
			{"A", "A", "A"},
			{"B", "B", "B"},
			{"C", "C", "C"},
		},
		emptyVal: ".",
		maxX:     2,
		maxY:     2,
	}

	got := NewStringMap(input, ".")

	require.Equal(t, expected, got)
	require.False(t, got.inBounds(3, 2))
	require.False(t, got.inBounds(2, 3))
	require.Equal(t, got.ValAt(0, 0), "A")
	require.Equal(t, got.ValAt(2, 2), "C")
	require.Equal(t, got.ValAt(3, 2), ".")

}
