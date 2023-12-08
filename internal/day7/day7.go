package day7

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lameorc/aoc_2023/internal/solution"
	"github.com/lameorc/aoc_2023/internal/utils"
)

type Day struct{}

type card int

const (
	joker card = iota
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

type handType int

const (
	high handType = iota
	pair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	cards []card
	bid   int
}

func parseHands(input []string) []hand {
	hs := make([]hand, 0, len(input))
	for _, l := range input {
		if l == "" {
			continue
		}
		splits := strings.Fields(l)

		cardStr := splits[0]
		cards := make([]card, 0, len(cardStr))
		for _, c := range cardStr {
			cStr := string(c)
			var val card
			switch cStr {
			case "2":
				val = two
			case "3":
				val = three
			case "4":
				val = four
			case "5":
				val = five
			case "6":
				val = six
			case "7":
				val = seven
			case "8":
				val = eight
			case "9":
				val = nine
			case "T":
				val = ten
			case "J":
				val = jack
			case "Q":
				val = queen
			case "K":
				val = king
			case "A":
				val = ace
			}

			cards = append(cards, val)
		}

		bid := utils.AtoiOrFail(splits[1], "bid")
		hs = append(hs, hand{cards: cards, bid: bid})
	}
	return hs
}

func (h *hand) typ() handType {
	counts := make(map[card]int)

	for _, c := range h.cards {
		counts[c]++
	}
	onePair := false
	doublePair := false
	triple := false
	four := false
	five := false

	for face, count := range counts {
		if face == joker {
			continue
		}
		switch count {
		case 2:
			if onePair {
				doublePair = true
			}
			onePair = true
		case 3:
			triple = true
		case 4:
			four = true
		case 5:
			five = true
		}
	}

	numJokers := counts[joker]
	if five {
		return fiveOfAKind
	} else if four {
		if numJokers >= 1 {
			return fiveOfAKind
		}
		return fourOfAKind
	} else if triple && onePair {
		return fullHouse
	} else if triple {
		if numJokers == 1 {
			return fourOfAKind
		} else if numJokers >= 2 {
			return fiveOfAKind
		}
		return threeOfAKind
	} else if doublePair {
		if numJokers == 1 {
			return fullHouse
		} else if numJokers >= 2 {
			return fourOfAKind
		}
		return twoPair
	} else if onePair {
		if numJokers == 1 {
			return threeOfAKind
		} else if numJokers == 2 {
			return fourOfAKind
		} else if numJokers == 3 {
			return fiveOfAKind
		}
		return pair
	} else {
		switch numJokers {
		case 1:
			return pair
		case 2:
			return threeOfAKind
		case 3:
			return fourOfAKind
		case 4:
			return fiveOfAKind
		case 5:
			return fiveOfAKind
		}

		return high
	}
}

type byStrength []hand

// Len implements sort.Interface.
func (b byStrength) Len() int {
	return len(b)
}

// Less implements sort.Interface.
func (b byStrength) Less(i int, j int) bool {
	iTyp := b[i].typ()
	jTyp := b[j].typ()

	if iTyp == jTyp {
		for idx, val := range b[i].cards {
			other := b[j].cards[idx]
			if val == other {
				continue
			}
			return val < other
		}

		return false // completely equal
	}

	return iTyp < jTyp
}

// Swap implements sort.Interface.
func (b byStrength) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func ToDay2(hs []hand) []hand {
	// this is in-place and we're returning it, but whatever
	for _, h := range hs {
		for idx, val := range h.cards {
			if val == jack {
				h.cards[idx] = joker
			}
		}
	}

	return hs
}

// Part1 implements solution.Solution.
func (*Day) Part1(input []string) string {
	winnings := 0
	hands := parseHands(input)
	sort.Sort(byStrength(hands))
	for rank, hand := range hands {
		winnings += (rank + 1) * hand.bid
	}

	return fmt.Sprintf("%d", winnings)
}

// Part2 implements solution.Solution.
func (*Day) Part2(input []string) string {
	winnings := 0
	hands := parseHands(input)
	hands = ToDay2(hands)
	sort.Sort(byStrength(hands))
	for rank, hand := range hands {
		winnings += (rank + 1) * hand.bid
	}

	return fmt.Sprintf("%d", winnings)
}

var _ solution.Solution = (*Day)(nil)
