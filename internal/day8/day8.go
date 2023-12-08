package day8

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lameorc/aoc_2023/internal/solution"
)

type Day struct{}

const (
	startVal = "AAA"
	endVal   = "ZZZ"
)

type node struct {
	val         string
	left, right string
}

type network struct {
	nodes map[string]*node
}

type instructions struct {
	values []rune
}

type instructionIterator struct {
	instructions
	current int
}

func (i *instructionIterator) next() rune {
	if i.current == len(i.values) {
		i.current = 0
	}
	val := i.values[i.current]
	i.current++

	return val

}

type parsed struct {
	Instructions instructions
	Network      network
}

func newParsed(input []string) *parsed {
	p := parsed{
		Instructions: instructions{values: make([]rune, 0)},
		Network: network{
			nodes: make(map[string]*node),
		},
	}
	for _, r := range input[0] {
		p.Instructions.values = append(p.Instructions.values, r)
	}

	nodeRe := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)`)
	for _, line := range input[2:] {
		if line == "" {
			continue
		}
		match := nodeRe.FindAllStringSubmatch(line, -1)
		if match == nil {
			panic(fmt.Sprintf("failed to match %s", line))
		}

		n := &node{
			val:   match[0][1],
			left:  match[0][2],
			right: match[0][3],
		}
		p.Network.nodes[n.val] = n
	}

	return &p
}

func (n *network) nextNode(current *node, instr rune) *node {
	var nextNodeVal string
	switch instr {
	case 'L':
		nextNodeVal = current.left
	case 'R':
		nextNodeVal = current.right
	}
	return n.nodes[nextNodeVal]
}

// Part1 implements solution.Solution.
func (*Day) Part1(input []string) string {
	p := newParsed(input)
	iter := instructionIterator{
		instructions: p.Instructions,
		current:      0,
	}
	steps := 0
	currentNode := p.Network.nodes[startVal]
	for {
		if currentNode.val == endVal {
			break
		}
		instr := iter.next()
		currentNode = p.Network.nextNode(currentNode, instr)
		steps++
	}
	return fmt.Sprint(steps)
}

type stack struct {
	nodes []*node
}

func (s *stack) push(n *node) {
	s.nodes = append(s.nodes, n)
}

func (s *stack) pop() *node {
	n := s.nodes[0]
	s.nodes = s.nodes[1:]
	return n
}

func (s *stack) endState() bool {
	for _, n := range s.nodes {
		if !strings.HasSuffix(n.val, "Z") {
			return false
		}
	}

	return true
}

// Part2 implements solution.Solution.
func (*Day) Part2(input []string) string {
	p := newParsed(input)
	n := p.Network
	// s := stack{nodes: make([]*node, 0)}
	starts := make([]*node, 0)
	for key, node := range n.nodes {
		if strings.HasSuffix(key, "A") {
			starts = append(starts, node)
		}
	}
	cycleLens := make([]int, len(starts))
	for idx, startNode := range starts {
		current := startNode
		loopStarted := false
		iter := instructionIterator{
			instructions: p.Instructions,
			current:      0,
		}
		var firstNodeOfLoop *node
		for {
			instr := iter.next()
			if strings.HasSuffix(current.val, "Z") {
				if loopStarted {
					potentialNext := n.nextNode(current, instr)
					if potentialNext != firstNodeOfLoop {
						panic(fmt.Sprintf("%s doesnt loop after %s; instead got %s", startNode.val, current.val, potentialNext.val))
					}
					break
				}
				firstNodeOfLoop = n.nextNode(current, instr)
				loopStarted = true
			}
			if loopStarted {
				cycleLens[idx]++
			}
			current = n.nextNode(current, instr)
		}
	}
	steps := LCM(cycleLens[0], cycleLens[1], cycleLens[2:]...)
	return fmt.Sprint(steps)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

var _ solution.Solution = (*Day)(nil)
