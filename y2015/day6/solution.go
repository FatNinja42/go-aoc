package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/cmd"
	"aoc/utils"
)

type day6 struct{}

type point struct {
	x int
	y int
}

func main() {
	doDay()
}

func doDay() {
	day := day6{}
	instr := utils.ReadFileAsStringSlice("/y2015/day6/input_t.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	//fmt.Printf("P2: %d", day.Part2(input))
}

func (d day6) Part1(input cmd.Input) interface{} {
	m := map[point]bool{}

	for _, s := range input.SLI {
		if strings.HasPrefix(s, "turn on") {
			p1, p2 := getCoords(s, "turn on ")
			TurnOn(p1, p2, m)
		} else if strings.HasPrefix(s, "turn off") {
			p1, p2 := getCoords(s, "turn off ")
			TurnOff(p1, p2, m)
		} else if strings.HasPrefix(s, "toggle") {
			p1, p2 := getCoords(s, "toggle ")
			Toggle(p1, p2, m)
		}
		Print(m)
	}

	return Count(m)
}

func (d day6) Part2(input cmd.Input) interface{} {
	return ""
}

func TurnOn(point1, point2 point, m map[point]bool) {
	for i := point1.x; i <= point1.y; i++ {
		for j := point2.x; j <= point2.y; j++ {
			m[point{i, j}] = true
		}
	}
}

func Count(m map[point]bool) int {
	c := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if f, _ := m[point{i, j}]; !f {
				c++
			}
		}
	}

	return c
}

func getCoords(s string, pref string) (point, point) {
	strings.ReplaceAll(s, pref, "")
	split := strings.Split(s, " through ")
	xs := strings.Split(split[0], ",")
	ys := strings.Split(split[1], ",")

	x1, _ := strconv.Atoi(xs[0])
	x2, _ := strconv.Atoi(xs[1])
	y1, _ := strconv.Atoi(ys[0])
	y2, _ := strconv.Atoi(ys[1])

	return point{x1, y1}, point{x2, y2}
}

func TurnOff(point1, point2 point, m map[point]bool) {
	for i := point1.x; i <= point1.y; i++ {
		for j := point2.x; j <= point2.y; j++ {
			m[point{i, j}] = false
		}
	}
}

func Toggle(point1, point2 point, m map[point]bool) {
	for i := point1.x; i <= point1.y; i++ {
		for j := point2.x; j <= point2.y; j++ {
			m[point{i, j}] = !m[point{i, j}]
		}
	}
}

func Print(m map[point]bool) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%t ", m[point{i, j}])
		}
		fmt.Println()
	}
	fmt.Println()
}
