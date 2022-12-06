package main

import (
	"fmt"
	"unicode"

	"aoc/cmd"
	"aoc/utils"
)

type day3 struct{}

func newDay() *day3 {
	return &day3{}
}

/*
 P1: 8243
 P2: 2631
*/
// Template.
func main() {
	doDay()
}

func doDay() {
	day := newDay()
	instr := utils.ReadFileAsStringSlice("/y2022/day3/input.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

var check map[rune]int

func (d *day3) Part1(input cmd.Input) interface{} {
	sum := 0

	for _, v := range input.SLI {
		p1 := v[:len(v)/2]
		p2 := v[len(v)/2:]

		for _, c := range v {

			if _, ok := check[c]; ok {
				continue
			}

			sum += getValue(c)
		}
	}

	return sum
}

func commonChars(parts []string, sum, vals []int) []rune {
	group := 0
	groupNr := len(vals)
	checkDigit := vals[groupNr-1]

	var check map[rune]int
	var runes []rune

	for _, part := range parts {
		if group == 0 {
			for k, v := range check {
				if v == checkDigit {
					runes = append(runes, k)
				}
			}
			check = map[rune]int{}
		}

		for _, c := range part {
			v, ok := check[c]
			if !ok || (group != 0 && v == sum[group-1]) {
				check[c] += vals[group]
			}
		}

		group = (group + 1) % groupNr
	}

}

func getValue(r rune) int {
	if unicode.IsUpper(r) {
		return int(r-'A') + 27
	}

	return int(r-'a') + 1
}

func (d *day3) Part2(input cmd.Input) interface{} {
	group := 0
	sum := []int{1, 4}
	vals := []int{1, 3, 5}

	var check map[rune]int

	s := 0
	for _, line := range input.SLI {
		if group == 0 {
			for k, v := range check {
				if v == 9 {
					if unicode.IsUpper(k) {
						s += int(k-'A') + 27
					} else {
						s += int(k-'a') + 1
					}
				}
			}
			check = map[rune]int{}
		}

		for _, c := range line {
			v, ok := check[c]
			if !ok {
				check[c] += vals[group]
			} else if group != 0 && v == sum[group-1] {
				check[c] += vals[group]
			}
		}

		group = (group + 1) % 3
	}

	for k, v := range check {
		if v == 9 {
			if unicode.IsUpper(k) {
				s += int(k-'A') + 27
			} else {
				s += int(k-'a') + 1
			}
		}
	}

	return s
}
