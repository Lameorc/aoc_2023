package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func AtoiOrFail(s, opIdent string) int {
	s = strings.Trim(s, " ")
	out, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse %s from %s", opIdent, s)
	}
	return out

}

func ReMatchOrZero(s string, r *regexp.Regexp, ident string) int {
	matches := r.FindStringSubmatch(s)
	if matches == nil {
		return 0
	}
	return AtoiOrFail(matches[1], ident)
}

func IsNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func Prepend[elemTyp any](s []elemTyp, val elemTyp) []elemTyp {
	return append([]elemTyp{val}, s...)
}

// StringMap assumes that the underlying map is a rectangle
type StringMap struct {
	values     [][]string
	emptyVal   string
	maxX, maxY int
}

func (S *StringMap) inBounds(x, y int) bool {
	return x <= S.maxX && x >= 0 && y <= S.maxY && y >= 0
}

func (S *StringMap) MaxX() int {
	return S.maxX
}

func (S *StringMap) MaxY() int {
	return S.maxY
}

func (S *StringMap) ValAt(x, y int) string {
	if !S.inBounds(x, y) {
		return S.emptyVal
	}
	return S.values[y][x]
}

func NewStringMap(in []string, emptyVal string) *StringMap {
	sm := &StringMap{
		values:   make([][]string, 0, len(in)),
		emptyVal: emptyVal,
	}

	for _, row := range in {
		if row == "" {
			continue
		}
		sm.values = append(sm.values, strings.Split(row, ""))
	}

	sm.maxX = len(sm.values[0]) - 1
	sm.maxY = len(sm.values) - 1
	return sm
}
