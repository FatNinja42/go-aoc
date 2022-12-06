package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/cmd"
	"aoc/utils"
)

type node struct {
	from string
	to   string
}

type day9 struct{}

// Template.
func main() {
	doDay()
}

func doDay() {
	day := day9{}
	instr := utils.ReadFileAsStringSlice("/y2015/day9/input_test.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day9) Part1(input cmd.Input) interface{} {
	m := buildGraphMap(input.SLI)
	v := buildVisitedMap(m)

	fmt.Println(v)
	return "doo"
}

func buildGraphMap(input []string) map[node]int {
	var m = map[node]int{}

	for _, v := range input {
		from, to, dist := parse(v)
		m[node{from, to}] = dist
	}

	return m
}

func buildVisitedMap(m map[node]int) map[string]bool {
	var v = map[string]bool{}

	for k := range m {
		v[k.from] = false
		v[k.to] = false
	}

	return v
}

func parse(v string) (string, string, int) {
	split := strings.Split(v, " = ")
	dist, _ := strconv.Atoi(split[1])

	nodes := strings.Split(split[0], " to ")

	return nodes[0], nodes[1], dist
}

func (d day9) Part2(input cmd.Input) interface{} {
	return ""
}
