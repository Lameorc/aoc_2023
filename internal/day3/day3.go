package day3

import (
	"fmt"
	"strconv"

	"github.com/lameorc/aoc_2023/internal/solution"
	"github.com/lameorc/aoc_2023/internal/utils"
)

type Day3 struct{}

type schematic struct {
	s    []string
	maxX int
	maxY int
}

func newSchematic(input []string) *schematic {

	return &schematic{
		s:    input,
		maxX: len(input[0]) - 1, // The width is the same all the way
		maxY: len(input) - 1,
	}
}

func (s *schematic) isInBound(x, y int) bool {
	return x >= 0 && x <= s.maxX && y >= 0 && y <= s.maxY
}

func (s *schematic) atPos(x, y int) string {

	if !s.isInBound(x, y) {
		return "."
	}
	return string(s.s[y][x])
}

// get the number that's lies "on" a given position, meaning any of it's digits lie there
func (s *schematic) numAtPos(x, y int) int {
	atPos := s.atPos(x, y)
	if !utils.IsNum(atPos) {
		return 0
	}
	val := atPos

	// all the numbers seem to be 3 digits at most, so just check the neighbors without any recursion
	// though it needs to be by two steps at most
	for i := 1; i <= 2; i++ {
		leftVal := s.atPos(x-i, y)
		if !utils.IsNum(leftVal) {
			break
		}
		val = fmt.Sprintf("%s%s", leftVal, val)
	}
	for i := 1; i <= 2; i++ {
		rightVal := s.atPos(x+i, y)
		if !utils.IsNum(rightVal) {
			break
		}
		val = fmt.Sprintf("%s%s", val, rightVal)
	}

	return utils.AtoiOrFail(val, "numAtPos")
}

func (s *schematic) getNeighboringNums(x, y int) []int {
	// only need to check corners + left, right. Due to nature of numbers,
	// if it's in the space directly above/below, it's going to have one value in corner anyway
	// due to this, if it's the same number, it means it's directly above/below so we need to only
	// count it once
	nums := make([]int, 0)

	top := s.numAtPos(x, y-1)
	if top == 0 { // if top is zero, means corners can be filled
		topLeft := s.numAtPos(x-1, y-1)
		if topLeft != 0 {
			nums = append(nums, topLeft)
		}
		topRight := s.numAtPos(x+1, y-1)
		if topRight != 0 {
			nums = append(nums, topRight)
		}
	} else {
		nums = append(nums, top)
	}

	left := s.numAtPos(x-1, y)
	if left != 0 {
		nums = append(nums, left)
	}

	right := s.numAtPos(x+1, y)
	if right != 0 {
		nums = append(nums, right)
	}
	bot := s.numAtPos(x, y+1)
	if bot == 0 { // same logic as top
		botLeft := s.numAtPos(x-1, y+1)
		if botLeft != 0 {
			nums = append(nums, botLeft)
		}
		botRight := s.numAtPos(x+1, y+1)
		if botRight != 0 {
			nums = append(nums, botRight)
		}

	} else {
		nums = append(nums, bot)
	}

	return nums
}

// Part1 implements solution.Solution.
func (*Day3) Part1(input []string) string {
	result := 0

	s := newSchematic(input)

	// just walk it like a idiot; maybe refactor sometime so we ever only visit one place once?
	for y := 0; y < s.maxY; y++ {
		for x := 0; x < s.maxX; x++ {
			asStr := string(s.atPos(x, y))
			if asStr == "." { // skip empty fields
				continue
			}

			if _, err := strconv.Atoi(asStr); err == nil { // skip numbers
				continue
			}

			// rest is symbols
			neighboringNums := s.getNeighboringNums(x, y)
			for _, num := range neighboringNums {
				result += num
			}
		}
	}
	return fmt.Sprintf("%d", result)
}

// Part2 implements solution.Solution.
func (*Day3) Part2(input []string) string {
	s := newSchematic(input)

	result := 0

	// just walk it like a idiot
	for y := 0; y < s.maxY; y++ {
		for x := 0; x < s.maxX; x++ {
			asStr := string(s.atPos(x, y))
			if asStr != "*" { // not a gear
				continue
			}
			neighboringNums := s.getNeighboringNums(x, y)
			if len(neighboringNums) != 2 { // not a gear
				continue
			}
			result += (neighboringNums[0] * neighboringNums[1])
		}
	}
	return fmt.Sprintf("%d", result)
}

var _ solution.Solution = (*Day3)(nil)
