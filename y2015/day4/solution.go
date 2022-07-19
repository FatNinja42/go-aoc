package main

import (
	"aoc/cmd"
	"aoc/utils"
	"crypto/md5"
	"fmt"
	"strings"
)

type day4 struct{}

func main() {
	doDay()
}

func doDay() {
	day := day4{}
	stringInput := utils.ReadFileAsString("/y2015/day4/input.txt")

	input := cmd.Input{
		SI: stringInput,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day4) Part1(input cmd.Input) interface{} {
	flag := true
	i := 0
	for flag {
		i++
		res := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input.SI, i))))
		if strings.HasPrefix(res, "00000") {
			flag = false
		}
	}
	return i
}

func (d day4) Part2(input cmd.Input) interface{} {
	flag := true
	i := 0
	for flag {
		i++
		res := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input.SI, i))))
		if strings.HasPrefix(res, "000000") {
			flag = false
		}
	}
	return i
}
