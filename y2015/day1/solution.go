package main

import (
	"aoc/cmd"
	"aoc/utils"
	"fmt"
)

type day1 struct{}

func main() {
	doDay()
}

func doDay() {
	day := day1{}
	stringInput := utils.ReadFileAsString("/y2015/day1/input.txt")

	input := cmd.Input{
		SI: stringInput,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day1) Part1(input cmd.Input) interface{} {
	c := 0
	for _, r := range input.SI {
		if r == '(' {
			c++
		} else if r == ')' {
			c--
		}
	}

	return c
}

func (d day1) Part2(input cmd.Input) interface{} {
	c := 0
	p := 0
	for i, r := range input.SI {
		if r == '(' {
			c++
		} else if r == ')' {
			c--
		}

		if c == -1 {
			p = i + 1
			break
		}
	}

	return p
}
