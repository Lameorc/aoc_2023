package day9

import (
	"fmt"
	"strings"

	"github.com/lameorc/aoc_2023/internal/utils"
)

type Day struct{}

func (Day) Part1(in []string) string {
	hs := ParseHistories(in)
	result := 0
	for _, h := range hs {
		prediction := h.predict()
		result += prediction
	}
	return fmt.Sprint(result)
}

func (Day) Part2(in []string) string {
	hs := ParseHistories(in)
	result := 0
	for _, h := range hs {
		prevExtrapolation := h.extrapolatePrev()
		result += prevExtrapolation
	}
	return fmt.Sprint(result)

}

type history []int

type histories []history

func ParseHistories(in []string) histories {
	hs := make(histories, 0, len(in))
	for _, line := range in {
		if line == "" {
			continue
		}
		rawValues := strings.Fields(line)
		h := make(history, 0, len(rawValues))
		for _, val := range rawValues {
			h = append(h, utils.AtoiOrFail(val, "single reading"))
		}
		hs = append(hs, h)
	}

	return hs
}

func calculateDiffLine(a []int) []int {
	diffLen := len(a) - 1
	diff := make([]int, diffLen)
	for idx := 0; idx < diffLen; idx++ {
		diff[idx] = a[idx+1] - a[idx]
	}

	return diff
}

func allZero(a []int) bool {
	for _, val := range a {
		if val != 0 {
			return false
		}
	}
	return true
}

func (h history) asHistoriesWithDiffs() histories {
	diffs := make(histories, 0)
	// the history is pretty much the first diff
	diffs = append(diffs, h)

	// fill the initial "pyramid" of diffs
	nonZeroDiff := true
	for diffIdx := 0; nonZeroDiff; diffIdx++ {
		newDiff := calculateDiffLine(diffs[diffIdx])
		if allZero(newDiff) {
			nonZeroDiff = false
		}

		diffs = append(diffs, newDiff)
	}

	return diffs
}

func (h history) predict() int {
	diffs := h.asHistoriesWithDiffs()

	// add the trailing zero to last diff
	diffs[len(diffs)-1] = append(diffs[len(diffs)-1], 0)

	// fill in the rest
	for diffIdx := len(diffs) - 2; diffIdx >= 0; diffIdx-- {
		currentLine := diffs[diffIdx]

		valIdx := len(currentLine) - 1

		leftVal := currentLine[valIdx]
		bottomVal := diffs[diffIdx+1][valIdx]

		diffs[diffIdx] = append(diffs[diffIdx], leftVal+bottomVal)
	}

	// finally, check the last value of first diff to get the prediction
	return diffs[0][len(diffs[0])-1]
}

func (h history) extrapolatePrev() int {
	diffs := h.asHistoriesWithDiffs()

	numDiffs := len(diffs)
	// prepend leading zero to last diff
	diffs[numDiffs-1] = utils.Prepend(diffs[numDiffs-1], 0)

	// fill the rest
	// fill in the rest
	for diffIdx := len(diffs) - 2; diffIdx >= 0; diffIdx-- {
		currentLine := diffs[diffIdx]

		rightVal := currentLine[0]
		bottomVal := diffs[diffIdx+1][0]

		diffs[diffIdx] = utils.Prepend(diffs[diffIdx], rightVal-bottomVal)
	}

	return diffs[0][0]
}
