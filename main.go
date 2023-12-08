package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/lameorc/aoc_2023/internal/day1"
	"github.com/lameorc/aoc_2023/internal/day2"
	"github.com/lameorc/aoc_2023/internal/day3"
	"github.com/lameorc/aoc_2023/internal/day4"
	"github.com/lameorc/aoc_2023/internal/day5"
	"github.com/lameorc/aoc_2023/internal/day6"
	"github.com/lameorc/aoc_2023/internal/day7"
	"github.com/lameorc/aoc_2023/internal/day8"
	"github.com/lameorc/aoc_2023/internal/solution"
)

func loadInput(day int) []string {
	filePath := fmt.Sprintf("./inputs/day%d.txt", day)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")

}

func main() {

	dayPtr := flag.Int("day", 1, "The day to solve")
	partPtr := flag.Int("part", 1, "The part to solve")

	flag.Parse()

	var s solution.Solution
	switch *dayPtr {
	case 1:
		s = &day1.Day1{}
	case 2:
		s = &day2.Day2{}
	case 3:
		s = &day3.Day3{}
	case 4:
		s = &day4.Day4{}
	case 5:
		s = &day5.Day{}
	case 6:
		s = &day6.Day{}
	case 7:
		s = &day7.Day{}
	case 8:
		s = &day8.Day{}
	default:
		log.Fatalf("unknown day %d", *dayPtr)
	}

	input := loadInput(*dayPtr)

	var result string
	switch *partPtr {
	case 1:
		result = s.Part1(input)
	case 2:
		result = s.Part2(input)
	case 3:
		log.Fatalf("unknown part %d", *partPtr)
	}

	fmt.Printf("%s\n", result)

}
