package main

import (
	"fmt"

	"aoc/cmd"
	"aoc/utils"
)

type day6 struct{}

func newDay() *day6 {
	return &day6{}
}

// Template.
func main() {
	doDay()
}

func doDay() {
	day := newDay()
	instr := utils.ReadFileAsString("/y2022/day6/input.txt")

	input := cmd.Input{
		SI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d *day6) Part1(input cmd.Input) interface{} {
	var count map[uint8]int
	var seq = input.SI

	for i := range seq {
		count = map[uint8]int{}
		for j := 0; j < 4; j++ {
			count[seq[i+j]]++
		}

		valid := true
		for _, val := range count {
			if val == 1 {
				continue
			}
			valid = false
		}

		if valid {
			return i + len(count)
		}
	}

	return -1
}

func (d *day6) Part2(input cmd.Input) interface{} {
	var count map[uint8]int
	var seq = input.SI

	for i := range seq {
		count = map[uint8]int{}
		for j := 0; j < 14; j++ {
			count[seq[i+j]]++
		}

		valid := true
		for _, val := range count {
			if val == 1 {
				continue
			}
			valid = false
		}

		if valid {
			return i + len(count)
		}
	}

	return -1
}
