package main

import (
	"fmt"

	"aoc/cmd"
	"aoc/utils"
)

type dayx struct{}

// Template.
func main() {
	doDay()
}

func doDay() {
	day := dayx{}
	instr := utils.ReadFileAsString("/y2015/dayx/input.txt")

	input := cmd.Input{
		SI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d dayx) Part1(input cmd.Input) interface{} {
	return ""
}

func (d dayx) Part2(input cmd.Input) interface{} {
	return ""
}
