package day11

import (
	"fmt"
	"math"
	"strings"

	"github.com/lameorc/aoc_2023/internal/solution"
)

type Day struct{}

// Part1 implements solution.Solution.
func (*Day) Part1(input []string) string {
	u := parseUniverse(input)
	sumPaths := 0

	for idx, first := range u.galaxies {
		for _, second := range u.galaxies[idx+1:] {
			path := u.shortestPath(first, second)
			sumPaths += len(path)
		}
	}

	return fmt.Sprint(sumPaths)
}

// Part2 implements solution.Solution.
func (*Day) Part2(input []string) string {
	panic("unimplemented")
}

var _ solution.Solution = (*Day)(nil)

type point struct{ x, y int }
type galaxy struct{ point }

type universe struct {
	image    [][]string
	galaxies []galaxy
}

func (u *universe) neighboursFor(p point) []point {
	neighbors := make([]point, 0)
	validOffsets := []struct {
		xOffset, yOffset int
	}{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for _, offset := range validOffsets {
		newCol := p.x + offset.xOffset
		newRow := p.y + offset.yOffset
		newPoint := point{newCol, newRow}
		if u.inBounds(newPoint) {
			neighbors = append(neighbors, newPoint)
		}
	}
	return neighbors
}

func (u *universe) inBounds(p point) bool {
	maxY := len(u.image) - 1
	maxX := len(u.image[0]) - 1
	return p.x >= 0 && p.x <= maxX && p.y >= 0 && p.y <= maxY
}

func (u *universe) shortestPath(start, end galaxy) []point {
	toVisit := map[point]interface{}{start.point: nil}
	sources := make(map[point]point)
	scores := make(map[point]int)
	for col, rowData := range u.image {
		for row := range rowData {
			scores[point{row, col}] = math.MaxInt
		}
	}
	scores[start.point] = 0

	for len(toVisit) > 0 {
		// pop first
		var current point
		minVal := math.MaxInt
		for key := range toVisit {
			if val, ok := scores[key]; ok && val < minVal {
				minVal = val
				current = key
			}
		}
		if current == end.point {
			path := make([]point, 0)

			for ; current != start.point; current = sources[current] {
				path = append(path, current)
			}
			return path
		}
		delete(toVisit, current)

		for _, neighbor := range u.neighboursFor(current) {
			distanceOnThisPath := scores[current] + 1
			if distanceOnThisPath < scores[neighbor] {
				sources[neighbor] = current
				scores[neighbor] = distanceOnThisPath
				if _, ok := toVisit[neighbor]; !ok {
					toVisit[neighbor] = nil
				}
			}
		}

	}

	// no path
	return []point{}
}

func parseUniverse(in []string) universe {
	u := universe{
		image:    make([][]string, 0, len(in)),
		galaxies: make([]galaxy, 0),
	}

	galaxiesPerCol := make(map[int]int)

	for _, line := range in {
		if line == "" {
			continue
		}
		rowVals := strings.Split(line, "")
		u.image = append(u.image, rowVals)

		galaxyOnRow := false
		for x, val := range rowVals {
			if val == "#" {
				galaxyOnRow = true
				galaxiesPerCol[x] += 1
			}
		}

		if !galaxyOnRow {
			// if there's no galaxy on this row, it expands => copy the row and append it
			expanded := make([]string, len(rowVals))
			copy(expanded, rowVals)
			u.image = append(u.image, expanded)
		}
	}

	// start from end so we don't shift next columns
	for colIdx := len(u.image[0]) - 1; colIdx >= 0; colIdx-- {
		numGalaxies := galaxiesPerCol[colIdx]
		if numGalaxies != 0 {
			continue
		}

		for idx, row := range u.image {
			row = append(row[:colIdx+1], row[colIdx:]...)
			row[colIdx] = "."
			u.image[idx] = row
		}
	}

	// go through it expanded to note the correct coords for the galaxies
	for y, row := range u.image {
		for x, val := range row {
			if val == "#" {
				u.galaxies = append(u.galaxies, galaxy{point{x, y}})
			}
		}
	}

	return u
}
