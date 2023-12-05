package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"}
	expected := 142
	result, err := Part1(input)
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen"}
	expected := 281
	result, err := Part2(input)
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
