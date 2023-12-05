package day1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input []string) (int, error) {
	fmt.Println("part1")
	total := 0
	for _, line := range input {
		first, last := findFirstAndLastDigit(line)
		combined, err := strconv.ParseInt(first+last, 10, 64)
		if err != nil {
			return 0, err
		}
		total += int(combined)
	}
	fmt.Println(total)
	return total, nil
}
func findFirstAndLastDigit(input string) (string, string) {
	re := regexp.MustCompile("[0-9]")
	numbers := re.FindAllString(input, -1)

	if len(numbers) == 0 {
		return "0", "0"
	}
	if len(numbers) == 1 {
		return numbers[0], numbers[0]
	}
	return numbers[0], numbers[len(numbers)-1]
}

func Part2(input []string) (int, error) {
	fmt.Println("part2")
	total := 0
	for _, line := range input {
		first, last := findFirstAndLastDigitSpelledOut(line)
		_, err := strconv.Atoi(first)
		if err != nil {
			first, _ = convertToInt(first)
		}

		_, err = strconv.Atoi(last)
		if err != nil {
			last, _ = convertToInt(last)
		}

		combined, err := strconv.ParseInt(first+last, 10, 64)
		if err != nil {
			return 0, err
		}
		total += int(combined)
	}
	return total, nil
}

func findFirstAndLastDigitSpelledOut(input string) (string, string) {
	first := -1
	last := -1
	minpos := len(input) + 1
	maxpos := -1
	numberDict := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
	for key, value := range numberDict {
		position := strings.Index(input, key)
		if position == -1 {
			// substring not found
			continue
		}

		if position < minpos {
			first = value
			minpos = position
		}

		position = strings.LastIndex(input, key)

		if position > maxpos {
			last = value
			maxpos = position
		}
	}
	return strconv.Itoa(first), strconv.Itoa(last)
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func convertToInt(input string) (string, error) {
	switch input {
	case "one":
	case "eno":
		return "1", nil
	case "two":
	case "owt":
		return "2", nil
	case "three":
	case "eerht":
		return "3", nil
	case "four":
	case "rouf":
		return "4", nil
	case "five":
	case "evif":
		return "5", nil
	case "six":
	case "xis":
		return "6", nil
	case "seven":
	case "neves":
		return "7", nil
	case "eight":
	case "thgie":
		return "8", nil
	case "nine":
	case "enin":
		return "9", nil
	}
	return "0", nil
}
