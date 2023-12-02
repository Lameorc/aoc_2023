package day2

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/lameorc/aoc_2023/internal/solution"
)

type Day2 struct{}

var _ solution.Solution = (*Day2)(nil)

var (
	redRe   = regexp.MustCompile(`(\d+) red`)
	blueRe  = regexp.MustCompile(`(\d+) blue`)
	greenRe = regexp.MustCompile(`(\d+) green`)
)

type record struct {
	Red   int
	Green int
	Blue  int
}

type game struct {
	Id      int
	Records []record
}

// Checks if the game is possible for the limit as specified by the provided record (the record values set the limit).
func (g *game) isPossibleForLimit(l *record) bool {
	for _, r := range g.Records {
		if l.Blue < r.Blue || l.Green < r.Green || l.Red < r.Red {
			return false
		}
	}

	return true
}

func (g *game) GetFewestCubes() record {
	lowerBound := record{}
	for _, r := range g.Records {
		if r.Blue > lowerBound.Blue {
			lowerBound.Blue = r.Blue
		}
		if r.Green > lowerBound.Green {
			lowerBound.Green = r.Green
		}
		if r.Red > lowerBound.Red {
			lowerBound.Red = r.Red
		}
	}

	return lowerBound
}

func atoiOrFail(s, opIdent string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse %s from %s", opIdent, s)
	}
	return out

}

func reMatchOrZero(s string, r *regexp.Regexp, ident string) int {
	matches := r.FindStringSubmatch(s)
	if matches == nil {
		return 0
	}
	return atoiOrFail(matches[1], ident)
}

func parseGames(input []string) []game {
	games := make([]game, 0, len(input))
	for _, line := range input {
		if line == "" {
			continue
		}
		byColon := strings.Split(line, ":")
		gameId := atoiOrFail(strings.Split(byColon[0], " ")[1], "gameId")
		rawRecords := strings.Split(byColon[1], ";")
		g := game{Id: gameId, Records: make([]record, 0, len(rawRecords))}

		for _, r := range rawRecords {
			reds := reMatchOrZero(r, redRe, "reds")
			greens := reMatchOrZero(r, greenRe, "greens")
			blues := reMatchOrZero(r, blueRe, "blues")

			rec := record{
				Red:   reds,
				Green: greens,
				Blue:  blues,
			}
			g.Records = append(g.Records, rec)
		}

		games = append(games, g)
	}
	return games
}

func (*Day2) Part1(input []string) string {
	recordLimit := &record{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	games := parseGames(input)
	result := 0
	for _, g := range games {
		if g.isPossibleForLimit(recordLimit) {
			result += g.Id
		}
	}

	return fmt.Sprintf("%d", result)
}

func (*Day2) Part2(input []string) string {
	games := parseGames(input)
	result := 0

	for _, g := range games {
		fewest := g.GetFewestCubes()
		setPower := fewest.Blue * fewest.Green * fewest.Red
		result += setPower
	}
	return fmt.Sprintf("%d", result)
}
