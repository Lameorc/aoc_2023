package day1

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	in := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	expected := "142"
	d := Day1{}
	result := d.Part1(in)
	if result != expected {
		t.Fatalf("got=%s; want=%s", result, expected)
	}

}

func TestPart2(t *testing.T) {
	in := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	expected := "281"
	d := Day1{}
	result := d.Part2(in)
	if result != expected {
		t.Fatalf("got=%s; want=%s", result, expected)
	}

}

func TestPart2Extra(t *testing.T) {
	in := []string{
		"eighthree",
		"sevenine",
	}

	expected := fmt.Sprintf("%d", 83+79)
	d := Day1{}
	result := d.Part2(in)
	if result != expected {
		t.Fatalf("got=%s; want=%s", result, expected)
	}

}
