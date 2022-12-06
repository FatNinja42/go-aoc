package main

import (
	"fmt"
	"strconv"

	"aoc/cmd"
	"aoc/utils"
)

type day1 struct{}

func main() {
	doDay()
}

func doDay() {
	day := day1{}
	instr := utils.ReadFileAsStringSlice("/y2022/day1/input.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day1) Part1(input cmd.Input) interface{} {
	max := 0
	gmax := 0
	for i := 0; i < len(input.SLI); i++ {
		if input.SLI[i] == "" {
			if max > gmax {
				gmax = max
			}
			max = 0
			continue
		}
		nr, _ := strconv.Atoi(input.SLI[i])
		max += nr
	}

	if max > gmax {
		gmax = max
	}

	return gmax
}

func (d day1) Part2(input cmd.Input) interface{} {
	max := 0
	gmax1 := 0
	gmax2 := 0
	gmax3 := 0
	for i := 0; i < len(input.SLI); i++ {
		if input.SLI[i] == "" {
			if max > gmax3 {
				gmax1 = gmax2
				gmax2 = gmax3
				gmax3 = max
			} else if max > gmax2 {
				gmax1 = gmax2
				gmax2 = max
			} else if max > gmax1 {
				gmax1 = max
			}

			max = 0
			continue
		}
		nr, _ := strconv.Atoi(input.SLI[i])
		max += nr
	}

	if max > gmax3 {
		gmax1 = gmax2
		gmax2 = gmax3
		gmax3 = max
	} else if max > gmax2 {
		gmax1 = gmax2
		gmax2 = max
	} else if max > gmax1 {
		gmax1 = max
	}

	return gmax1 + gmax2 + gmax3
}
