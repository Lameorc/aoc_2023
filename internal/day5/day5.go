package day5

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/lameorc/aoc_2023/internal/solution"
	"github.com/lameorc/aoc_2023/internal/utils"
)

type Day struct{}

type sourceID int
type destinationID int

type mappingType int

const (
	seeds mappingType = iota
	seedToSoil
	soilToFertilizer
	fertilizerToWater
	waterToLight
	lightToTemp
	tempToHumidity
	humidityToLocation
)

type mapping struct {
	sourceStart int
	destStart   int
	rangeSize   int
}

func (m *mapping) offset(val int) (int, bool) {
	if val >= m.sourceStart && val < (m.sourceStart+m.rangeSize) {
		diff := val + m.destStart - int(m.sourceStart)
		return diff, true
	}
	return 0, false
}

type almanac struct {
	seeds []int
	// TODO: this is completely broken memory wise
	mappings map[mappingType][]mapping
}

var (
	seedsRe   = regexp.MustCompile(`^seeds: ((\d+ ?)+)$`)
	mapLineRe = regexp.MustCompile(`^(.+) map:$`)
)

func iterMappingTypes() []mappingType {
	parts := make([]mappingType, 0, humidityToLocation-1)
	for i := seedToSoil; i <= humidityToLocation; i++ {
		parts = append(parts, i)
	}
	return parts
}

func (a *almanac) fillSeeds(in string) {
	m := seedsRe.FindAllStringSubmatch(in, -1)
	if m == nil {
		panic(fmt.Sprintf("failed to parse %s in seeds state", in))
	}
	seedIds := strings.Split(m[0][1], " ")
	a.seeds = make([]int, 0, len(seedIds))
	a.mappings = make(map[mappingType][]mapping)
	for _, ID := range seedIds {
		seedID := utils.AtoiOrFail(ID, "seedID")
		a.seeds = append(a.seeds, seedID)
	}
}

func newAlmanac(input []string) *almanac {
	a := &almanac{}

	var state mappingType = seeds
	for _, line := range input {
		if line == "" {
			continue

		}

		if state == seeds {
			a.fillSeeds(line)
			state = seedToSoil // this is actually state++ but more explicit
		} else { // parsing map
			mapLineMatch := mapLineRe.FindAllStringSubmatch(line, -1)
			if mapLineMatch != nil { // line indicating a state switch
				// the only "special" case which is the first one
				if strings.HasPrefix(mapLineMatch[0][1], "seed-to-soil") {
					continue
				}
				// the value actually doesn't matter since they're always in order
				// so just advance the state
				state++
				continue
			}
			mappingSplit := strings.Split(line, " ")

			destStart := utils.AtoiOrFail(mappingSplit[0], "destStart")
			srcStart := utils.AtoiOrFail(mappingSplit[1], "srcStart")
			rangeSize := utils.AtoiOrFail(mappingSplit[2], "rangeSize")
			a.mappings[state] = append(a.mappings[state], mapping{
				destStart:   destStart,
				sourceStart: srcStart,
				rangeSize:   rangeSize,
			})
		}
	}

	return a
}

func (a *almanac) destForSrc(typ mappingType, src sourceID) destinationID {
	var result destinationID = destinationID(src)
	for _, m := range a.mappings[typ] {
		if offset, matches := m.offset(int(src)); matches {
			return destinationID(offset)
		}
	}

	return result
}

func (a *almanac) seedToLocation(seedID int) int {
	src := seedID
	for _, typ := range iterMappingTypes() {
		dest := a.destForSrc(typ, sourceID(src))
		src = int(dest)
	}
	return src

}

func (a *almanac) lowestLocation() int {
	lowest := math.MaxInt

	for _, seedID := range a.seeds {
		val := a.seedToLocation(seedID)
		if val < lowest {
			lowest = val
		}
	}
	return lowest
}

func (a *almanac) lowestLocationForPart2() int {
	lowest := math.MaxInt
	for i := 0; i < len(a.seeds); i += 2 {
		start := a.seeds[i]
		for j := start; j < start+a.seeds[i+1]; j++ {
			val := a.seedToLocation(j)
			if val < lowest {
				lowest = val
			}
		}
	}
	return lowest

}

func seedsAsRange(start, size int) []int {
	newS := make([]int, 0, size)
	for i := 0; i < size; i++ {
		newS = append(newS, start+i)
	}
	return newS
}

// Part1 implements solution.Solution.
func (*Day) Part1(input []string) string {
	a := newAlmanac(input)

	lowestLoc := a.lowestLocation()
	return fmt.Sprintf("%d", lowestLoc)
}

// Part2 implements solution.Solution.
func (*Day) Part2(input []string) string {
	a := newAlmanac(input)

	lowestLoc := a.lowestLocationForPart2()
	return fmt.Sprintf("%d", lowestLoc)
}

var _ solution.Solution = (*Day)(nil)
