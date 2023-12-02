package day1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lameorc/aoc_2023/internal/solution"
	"github.com/lameorc/aoc_2023/internal/utils"
)

type Day1 struct{}

var _ solution.Solution = (*Day1)(nil)

const numbers = "0123456789"

var part2Re = regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine)|(1|2|3|4|5|6|7|8|9|0)`)

func (*Day1) Part1(input []string) string {
	result := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		firstIndex := strings.IndexAny(line, numbers)
		lastIndex := strings.LastIndexAny(line, numbers)

		lineVal := fmt.Sprintf("%s%s", string(line[firstIndex]), string(line[lastIndex]))
		interim, err := strconv.Atoi(lineVal)
		if err != nil {
			panic(fmt.Sprintf("failed to parser lineval=%s", lineVal))
		}
		result += interim

	}

	return fmt.Sprintf("%d", result)

}

func numFromString(s string) int {
	val, err := strconv.Atoi(s)
	if err == nil {
		return val
	}

	//wasn't a iteral, do a dumb switch
	switch s {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "zero":
		return 0
	}

	panic(fmt.Sprintf("failed to parse %s as number", s))
}

func findOverlapping(pattern *regexp.Regexp, s string) []string {
	matches := []string{}
	toMatch := s
	for {
		match := pattern.FindStringIndex(toMatch)
		if match == nil {
			break
		}
		matches = append(matches, toMatch[match[0]:match[1]])

		toMatch = toMatch[match[0]+1:]
	}
	return matches
}

func (*Day1) Part2(input []string) string {
	result := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		matches := findOverlapping(part2Re, line)
		first := matches[0]
		last := matches[len(matches)-1]
		lineVal := fmt.Sprintf("%d%d", numFromString(first), numFromString(last))
		interim := utils.AtoiOrFail(lineVal, "lineval")
		result += interim

	}

	return fmt.Sprintf("%d", result)
}
