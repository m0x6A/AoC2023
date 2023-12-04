package main

import (
	"AoC2023/day1"
	"AoC2023/day3"
	"AoC2023/filereader"
	"fmt"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)

	fmt.Println("Adventure of Code 2023")
	fmt.Println("-----------------------")
	fmt.Print("Please enter the day you want to run: ")
	// day, _ := reader.ReadString('\n')
	day := "1"
	switch day {
	case "1":
		fmt.Println("-----------------------")
		fmt.Println("Day 1")
		day1Data, err := filereader.LoadData(day)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		day1Part1Answer, err := day1.Part1(day1Data)

		if err != nil {
			return
		}
		day1part1 := fmt.Sprintf("Answer day1 part: %d", day1Part1Answer)

		fmt.Println(day1part1)
	case "2":
		fmt.Println("Day 2")
	case "3":
		fmt.Println("Day 3")
		day3Data, err := filereader.LoadData("3")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		day3Part1Answer, err := day3.Part1(day3Data)
		if err != nil {
			return
		}
		day3part1 := fmt.Sprintf("Answer day3 part: %d", day3Part1Answer)

		day3Part2Answer, err := day3.Part2(day3Data)

		if err != nil {
			return
		}

		day3part2 := fmt.Sprintf("Answer day3 part: %d", day3Part2Answer)
		fmt.Println(day3part1)
		fmt.Println(day3part2)
	}
}
