package main

import (
	"aoc/cmd"
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type day2 struct{}

func main() {
	doDay()
}

func doDay() {
	day := day2{}
	stringSliceInput := utils.ReadFileAsStringSlice("/y2015/day2/input.txt")

	input := cmd.Input{
		SLI: stringSliceInput,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day2) Part1(input cmd.Input) interface{} {
	total := 0

	for _, dim := range input.SLI {
		l, w, h := dimensions(dim)

		sum := (2 * l * w) + (2 * w * h) + (2 * h * l) + minSideArea(l, w, h)

		total += sum
	}

	return total
}

func (d day2) Part2(input cmd.Input) interface{} {
	total := 0

	for _, dim := range input.SLI {
		l, w, h := dimensions(dim)

		total += minSidePerimeter(l, w, h) + l*w*h
	}

	return total
}

func dimensions(dim string) (int, int, int) {
	split := strings.Split(dim, "x")

	l, _ := strconv.Atoi(split[0])
	w, _ := strconv.Atoi(split[1])
	h, _ := strconv.Atoi(split[2])

	return l, w, h
}

func minSideArea(l, w, h int) int {
	return utils.Min(l*w, utils.Min(w*h, l*h))
}

func minSidePerimeter(l, w, h int) int {
	ord := []int{l, w, h}
	sort.Ints(ord)

	return 2*ord[0] + 2*ord[1]
}
