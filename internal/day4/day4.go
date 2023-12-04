package day4

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lameorc/aoc_2023/internal/solution"
	"github.com/lameorc/aoc_2023/internal/utils"
)

type Day4 struct{}

var _ solution.Solution = (*Day4)(nil)

var cardRe = regexp.MustCompile(`^Card\W+(\d+):\W*((\d+\W*)+)\|\W*((\d+\W*)+)+$`)

type card struct {
	id int

	winning  map[int]interface{}
	selected map[int]interface{}
}

func (c *card) getPoints() int {
	points := 0
	for i := 0; i < c.numWins(); i++ {
		if points == 0 {
			points = 1
		} else {
			points *= 2
		}
	}
	return points
}

func (c *card) numWins() int {
	wins := 0
	for val := range c.winning {
		if _, ok := c.selected[val]; !ok {
			continue
		}
		wins++
	}
	return wins
}
func parseCardValues(ms string, opid string) map[int]interface{} {
	vals := make(map[int]interface{}, 0)
	valsStr := strings.TrimSpace(ms)
	for _, numStr := range strings.Split(valsStr, " ") {
		numStr = strings.TrimSpace(numStr)
		if numStr == "" {
			continue
		}
		val := utils.AtoiOrFail(numStr, opid)
		vals[val] = nil
	}
	return vals

}

func cardsFromInput(input []string) []card {
	cs := make([]card, 0, len(input))
	for lineIdx, line := range input {
		if line == "" {
			continue
		}

		matches := cardRe.FindAllStringSubmatch(line, -1)
		if len(matches) != 1 {
			panic(fmt.Sprintf("wrong num matches on line %d; should only be one per line", lineIdx))
		}
		m := matches[0]

		id := utils.AtoiOrFail(m[1], "id")

		winning := parseCardValues(m[2], "winning")
		selected := parseCardValues(m[4], "selected")
		cs = append(cs, card{
			id:       id,
			winning:  winning,
			selected: selected,
		})
	}
	return cs
}

func (*Day4) Part1(input []string) string {
	result := 0
	cards := cardsFromInput(input)
	for _, c := range cards {
		result += c.getPoints()
	}
	return fmt.Sprintf("%d", result)
}

func (*Day4) Part2(input []string) string {
	result := 0
	cards := cardsFromInput(input)
	copiesToGet := make(map[int]int, 0)
	for idx, c := range cards {
		result++
		cWins := c.numWins()
		thisExtra := copiesToGet[idx]
		result += thisExtra

		for i := 1; i <= cWins; i++ {
			copiesToGet[idx+i] += 1 + thisExtra // one for base card + one for each extra
		}
	}
	return fmt.Sprintf("%d", result)
}
