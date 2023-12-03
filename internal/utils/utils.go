package utils

import (
	"log"
	"regexp"
	"strconv"
)

func AtoiOrFail(s, opIdent string) int {
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
