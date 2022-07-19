package main

import (
	"aoc/cmd"
	"aoc/utils"
	"fmt"
	"strings"
)

type day5 struct{}

func main() {
	doDay()
}

func doDay() {
	day := day5{}
	stringInput := utils.ReadFileAsStringSlice("/y2015/day5/input.txt")

	input := cmd.Input{
		SLI: stringInput,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day5) Part1(input cmd.Input) interface{} {
	c := 0
	for _, s := range input.SLI {
		if IsNice(s) {
			c++
		}
	}

	return c
}

func (d day5) Part2(input cmd.Input) interface{} {
	c := 0
	for _, s := range input.SLI {
		if IsNice2(s) {
			c++
		}
	}

	return c
}

func IsNice(s string) bool {
	return ContainsThreeVowels(s) && HasDuplicateLetter(s) && DoesNotContainStrings(s)
}

func IsNice2(s string) bool {
	return ContainsTwoPairs(s) &&
		HasDuplicateLetterSeparated(s)
}

func HasDuplicateLetterSeparated(s string) bool {
	ss := strings.Split(s, "")
	for i := 0; i < len(ss)-2; i++ {
		if ss[i] == ss[i+2] {
			return true
		}
	}

	return false
}

func ContainsTwoPairs(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}

func DoesNotContainStrings(s string) bool {
	forbidden := []string{"ab", "cd", "pq", "xy"}

	for i := range forbidden {
		if strings.Contains(s, forbidden[i]) {
			return false
		}
	}

	return true
}

func HasDuplicateLetter(s string) bool {
	ss := strings.Split(s, "")
	for i := 0; i < len(ss)-1; i++ {
		if ss[i] == ss[i+1] {
			return true
		}
	}

	return false
}

func ContainsThreeVowels(s string) bool {
	vowels := "aeiou"
	ct := 0
	for _, c := range strings.Split(s, "") {
		if strings.Contains(vowels, c) {
			ct++
		}
	}

	return ct >= 3
}
