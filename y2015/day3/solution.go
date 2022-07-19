package main

import (
	"aoc/cmd"
	"aoc/utils"
	"fmt"
	"strings"
)

type day3 struct{}

type point struct {
	x int
	y int
}

func main() {
	doDay()
}

func doDay() {
	day := day3{}
	stringInput := utils.ReadFileAsString("/y2015/day3/input.txt")

	input := cmd.Input{
		SI: stringInput,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day3) Part1(input cmd.Input) interface{} {
	var coords = map[point]bool{}

	x, y := 0, 0
	coords[point{x: x, y: y}] = true

	chars := strings.Split(input.SI, "")
	for _, v := range chars {
		switch v {
		case ">":
			x++
		case "<":
			x--
		case "^":
			y++
		case "v":
			y--
		}
		coords[point{x: x, y: y}] = true
	}

	return len(coords)
}

func (d day3) Part2(input cmd.Input) interface{} {
	var coords = map[point]bool{}
	var ix = []int{0, 0}
	var iy = []int{0, 0}

	coords[point{x: ix[0], y: iy[0]}] = true

	chars := strings.Split(input.SI, "")
	for i, v := range chars {
		switch v {
		case ">":
			ix[i%2]++
		case "<":
			ix[i%2]--
		case "^":
			iy[i%2]++
		case "v":
			iy[i%2]--
		}

		coords[point{x: ix[i%2], y: iy[i%2]}] = true
	}

	return len(coords)
}
