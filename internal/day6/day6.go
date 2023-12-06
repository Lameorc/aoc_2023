package day6

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/lameorc/aoc_2023/internal/solution"
	"github.com/lameorc/aoc_2023/internal/utils"
)

var (
	valuesRe = regexp.MustCompile(`^(Time|Distance):\W+((\d+\W*)+)$`)
)

type Day struct{}

type race struct {
	timeLimit      int
	distanceRecord int
}

func newRaces(in []string) []race {
	timeM := valuesRe.FindAllStringSubmatch(in[0], -1)
	if timeM == nil {
		panic("failed to match time")
	}
	DistanceM := valuesRe.FindAllStringSubmatch(in[1], -1)
	if DistanceM == nil {
		panic("failed to match distance")
	}
	times := strings.Fields(timeM[0][2])
	dataLen := len(times)

	distances := strings.Fields(DistanceM[0][2])
	if dataLen != len(distances) {
		panic("mismatch between len times and distances")
	}

	races := make([]race, 0, dataLen)
	for i := 0; i < dataLen; i++ {
		races = append(races, race{
			timeLimit:      utils.AtoiOrFail(times[i], "recordTime"),
			distanceRecord: utils.AtoiOrFail(distances[i], "distance"),
		})
	}

	return races
}

func newPart2Race(in []string) race {
	timeM := valuesRe.FindAllStringSubmatch(in[0], -1)
	if timeM == nil {
		panic("failed to match time")
	}
	DistanceM := valuesRe.FindAllStringSubmatch(in[1], -1)
	if DistanceM == nil {
		panic("failed to match distance")
	}
	times := strings.Fields(timeM[0][2])
	dataLen := len(times)

	distances := strings.Fields(DistanceM[0][2])
	if dataLen != len(distances) {
		panic("mismatch between len times and distances")
	}

	timeLimit := utils.AtoiOrFail(strings.Join(times, ""), "time limit")
	distanceToBeat := utils.AtoiOrFail(strings.Join(distances, ""), "distance to beat")

	return race{
		timeLimit:      timeLimit,
		distanceRecord: distanceToBeat,
	}

}

func (r *race) getWinOptions() int {
	options := 0
	toBeat := r.distanceRecord
	minAvgSpeed := float32(r.distanceRecord) / float32(r.timeLimit)
	// ceil, since we can't go fractions of second
	start := int(math.Ceil(float64(minAvgSpeed)))
	for i := start; i < r.timeLimit; i++ {
		timeRemaining := r.timeLimit - i
		distanceDriven := i * timeRemaining
		if distanceDriven > toBeat {
			options++
		}
	}

	return options
}

func (*Day) Part1(input []string) string {
	result := 1
	rs := newRaces(input)
	for _, r := range rs {
		result *= r.getWinOptions()
	}

	return fmt.Sprintf("%d", result)
}

func (*Day) Part2(input []string) string {
	r := newPart2Race(input)
	result := r.getWinOptions()

	return fmt.Sprintf("%d", result)
}

var _ solution.Solution = (*Day)(nil)
