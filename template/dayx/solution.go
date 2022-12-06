package main

import (
	"fmt"

	"aoc/cmd"
	"aoc/utils"
)

type dayx struct{}

func newDay() *dayx {
	return &dayx{}
}

func main() {
	doDay()
}

func doDay() {
	day := newDay()
	instr := utils.ReadFileAsStringSlice("/y2022/dayx/input.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d *dayx) Part1(input cmd.Input) interface{} {
	return ""
}

func (d *dayx) Part2(input cmd.Input) interface{} {
	return ""
}
