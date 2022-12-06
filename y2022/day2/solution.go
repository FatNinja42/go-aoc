package main

import (
	"fmt"
	"strings"

	"aoc/cmd"
	"aoc/utils"
)

type day2 struct{}

func newDay() *day2 {
	return &day2{}
}

func main() {
	doDay()
}

func doDay() {
	day := newDay()
	instr := utils.ReadFileAsStringSlice("/y2022/day2/input.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

var optionScoreP1 = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var roundScoreP1 = map[string]int{
	"AX": 3, "AY": 6, "AZ": 0,
	"BX": 0, "BY": 3, "BZ": 6,
	"CX": 6, "CY": 0, "CZ": 3,
}

func (d *day2) Part1(input cmd.Input) interface{} {
	score := 0
	for _, v := range input.SLI {
		choices := strings.Split(v, " ")
		score += optionScoreP1[choices[1]] + roundScoreP1[fmt.Sprintf("%s%s", choices[0], choices[1])]
	}

	return score
}

var optionScoreP2 = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

/*
		ROCK:     1
		PAPER:    2
	    SCISSORS: 3
*/
var roundScoreP2 = map[string]int{
	"AX": 3, "AY": 1, "AZ": 2,
	"BX": 1, "BY": 2, "BZ": 3,
	"CX": 2, "CY": 3, "CZ": 1,
}

func (d *day2) Part2(input cmd.Input) interface{} {
	score := 0
	for _, v := range input.SLI {
		choices := strings.Split(v, " ")
		score += optionScoreP2[choices[1]] + roundScoreP2[fmt.Sprintf("%s%s", choices[0], choices[1])]
	}

	return score
}
