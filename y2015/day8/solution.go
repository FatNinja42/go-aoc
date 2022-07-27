package main

import (
	"fmt"
	"strconv"

	"aoc/cmd"
	"aoc/utils"
)

type day8 struct{}

func main() {
	doDay()
}

func doDay() {
	day := day8{}
	instr := utils.ReadFileAsStringSlice("/y2015/day8/input.txt")
	//instr := utils.ReadFileAsStringSlice("/y2015/day8/input_test.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day8) Part1(input cmd.Input) interface{} {
	total := 0
	for _, v := range input.SLI {
		t := unquote(v)
		total += len(v) - len(t)
	}

	return total
}

func (d day8) Part2(input cmd.Input) interface{} {
	total := 0
	for _, v := range input.SLI {
		t := quote(v)
		total += len(t) - len(v)
	}

	return total
}

func unquote(str string) string {
	s, _ := strconv.Unquote(str)

	return s
}

func quote(str string) string {
	s := strconv.Quote(str)

	return s
}
