package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/cmd"
	"aoc/utils"
)

type day4 struct{}

func newDay() *day4 {
	return &day4{}
}

func main() {
	doDay()
}

func doDay() {
	day := newDay()
	instr := utils.ReadFileAsStringSlice("/y2022/day4/input.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d *day4) Part1(input cmd.Input) interface{} {
	count := 0
	for _, v := range input.SLI {
		a1, b1, a2, b2 := getIntervals(v)

		if (a2 >= a1 && b2 <= b1) || (a1 >= a2 && b1 <= b2) {
			count++
		}
	}

	return count
}

func (d *day4) Part2(input cmd.Input) interface{} {
	count := 0
	for _, v := range input.SLI {
		a1, b1, a2, b2 := getIntervals(v)

		if (a2 >= a1 && a2 <= b1) || (b2 >= a1 && b2 <= b1) || (a1 >= a2 && a1 <= b2) || (b1 >= a2 && b1 <= b2) {
			count++
		}
	}

	return count
}

func getIntervals(s string) (int, int, int, int) {
	intervals := strings.Split(s, ",")
	i1 := strings.Split(intervals[0], "-")
	i2 := strings.Split(intervals[1], "-")

	a1, _ := strconv.Atoi(i1[0])
	b1, _ := strconv.Atoi(i1[1])
	a2, _ := strconv.Atoi(i2[0])
	b2, _ := strconv.Atoi(i2[1])

	return a1, b1, a2, b2
}
