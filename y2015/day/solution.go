package main

import (
	"aoc/cmd"
	"aoc/utils"
	"fmt"
)

type day struct{}

// Template.
func main() {
	doDay()
}

func doDay() {
	day := day{}
	instr := utils.ReadFileAsString("/y2015/day/input.txt")

	input := cmd.Input{
		SI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day) Part1(input cmd.Input) interface{} {
	return ""
}

func (d day) Part2(input cmd.Input) interface{} {
	return ""
}
