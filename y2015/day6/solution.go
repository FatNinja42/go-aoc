package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"aoc/cmd"
	"aoc/utils"
)

const gridSize = 1000

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
	instr := utils.ReadFileAsStringSlice("/y2015/day6/input.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day6) Part1(input cmd.Input) interface{} {
	m := map[point]int{}

	for _, s := range input.SLI {
		if strings.HasPrefix(s, "turn on") {
			p1, p2 := getCoords(s, "turn on ")
			turnOn(m, p1, p2)
		} else if strings.HasPrefix(s, "turn off") {
			p1, p2 := getCoords(s, "turn off ")
			turnOff(m, p1, p2)
		} else if strings.HasPrefix(s, "toggle") {
			p1, p2 := getCoords(s, "toggle ")
			toggle(m, p1, p2)
		}
	}

	return count(m)
}

func (d day6) Part2(input cmd.Input) interface{} {
	m := map[point]int{}

	for _, s := range input.SLI {
		if strings.HasPrefix(s, "turn on") {
			p1, p2 := getCoords(s, "turn on ")
			turnOn2(m, p1, p2)
		} else if strings.HasPrefix(s, "turn off") {
			p1, p2 := getCoords(s, "turn off ")
			turnOff2(m, p1, p2)
		} else if strings.HasPrefix(s, "toggle") {
			p1, p2 := getCoords(s, "toggle ")
			toggle2(m, p1, p2)
		}
	}

	return count(m)
}

func getCoords(s string, pref string) (point, point) {
	s = strings.ReplaceAll(s, pref, "")
	split := strings.Split(s, " through ")
	xs := strings.Split(split[0], ",")
	ys := strings.Split(split[1], ",")

	x1, _ := strconv.Atoi(xs[0])
	y1, _ := strconv.Atoi(xs[1])
	x2, _ := strconv.Atoi(ys[0])
	y2, _ := strconv.Atoi(ys[1])

	return point{x1, y1}, point{x2, y2}
}

func turnOn(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			m[point{i, j}] = 1
		}
	}
}

func turnOff(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			m[point{i, j}] = 0
		}
	}
}

func toggle(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			m[point{i, j}] = int(math.Abs(float64(m[point{i, j}] - 1)))
		}
	}
}

func turnOn2(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			m[point{i, j}] += 1
		}
	}
}

func turnOff2(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			if v, _ := m[point{i, j}]; v > 0 {
				m[point{i, j}] -= 1
			}
		}
	}
}

func toggle2(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			m[point{i, j}] += 2
		}
	}
}

func count(m map[point]int) int {
	c := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			v, _ := m[point{i, j}]
			c += v
		}
	}

	return c
}
