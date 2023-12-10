package day10

import (
	"fmt"

	"github.com/lameorc/aoc_2023/internal/solution"
	"github.com/lameorc/aoc_2023/internal/utils"
)

type Day struct{}

// Part1 implements solution.Solution.
func (*Day) Part1(input []string) string {
	m := utils.NewStringMap(input, ".")
	start := findStartCoordinates(m)
	furthestDistance := furthestPoint(m, start)

	return fmt.Sprint(furthestDistance)
}

// Part2 implements solution.Solution.
func (*Day) Part2(input []string) string {
	panic("unimplemented")
}

var _ solution.Solution = (*Day)(nil)

type node struct {
	val  string
	x, y int
}

func findStartCoordinates(m *utils.StringMap) node {
	for rowIdx := 0; rowIdx <= m.MaxY(); rowIdx++ {
		for colIdx := 0; colIdx <= m.MaxX(); colIdx++ {
			if m.ValAt(colIdx, rowIdx) == "S" {
				return newNode(colIdx, rowIdx, "S")
			}
		}
	}

	panic("failed to find start!")
}

type fifo struct {
	nodes []node
}

func (l *fifo) pop() (node, bool) {
	if len(l.nodes) == 0 {
		return node{}, false
	}
	popped := l.nodes[0]
	l.nodes = l.nodes[1:]
	return popped, true
}

func (l *fifo) push(n node) {
	l.nodes = append(l.nodes, n)
}

func furthestPoint(m *utils.StringMap, start node) int {
	distances := map[node]int{
		start: 0,
	}
	furthest := 0

	toScan := fifo{
		nodes: []node{start},
	}

	for current, ok := toScan.pop(); ok; current, ok = toScan.pop() {
		neighbors := neighborsFor(m, current)
		for _, neighbor := range neighbors {
			if _, seen := distances[neighbor]; seen {
				// already seen
				continue
			}
			toScan.push(neighbor)
			distances[neighbor] = distances[current] + 1
		}
	}

	// just iter once again, could have been done in the previous cycle
	for _, distance := range distances {
		if distance > furthest {
			furthest = distance
		}
	}

	return furthest
}

func newNode(x, y int, val string) node {
	return node{
		val: val,
		x:   x,
		y:   y,
	}
}

type offset struct {
	x, y int
}

var (
	mapping map[string][]offset = map[string][]offset{
		".": {},
		"|": {{0, -1}, {0, +1}},
		"-": {{1, 0}, {-1, 0}},
		"L": {{0, -1}, {1, 0}},
		"J": {{0, -1}, {-1, 0}},
		"7": {{0, 1}, {-1, 0}},
		"F": {{0, 1}, {1, 0}},
		"S": {{1, 1}, {0, 1}, {1, 0}, {-1, -1}, {-1, 0}, {0, -1}, {1, -1}, {-1, 1}},
	}
)

func neighborsFor(m *utils.StringMap, n node) []node {
	offsets := mapping[n.val]
	if offsets == nil {
		panic(fmt.Sprintf("invalid node val: %s", n.val))
	}

	neighbors := make([]node, 0)
	for _, o := range offsets {
		nx := n.x + o.x
		ny := n.y + o.y
		nVal := m.ValAt(nx, ny)
		if nVal == "." {
			continue // not a pipe, skip
		}
		newN := newNode(nx, ny, nVal)

		// TODO: handle case where the start pipe is not connected and skip
		if n.val == "S" {
			validNeighbors := mapping[newN.val]
			if validNeighbors == nil {
				panic(fmt.Sprintf("invalid node val: %s", newN.val))
			}
			canConnect := false
			for _, o := range validNeighbors {
				if nx+o.x == n.x && ny+o.y == n.y {
					canConnect = true
					break
				}
			}
			if !canConnect {
				continue
			}
		}
		neighbors = append(
			neighbors,
			newN,
		)
	}
	return neighbors
}
