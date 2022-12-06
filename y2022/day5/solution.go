package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/cmd"
	"aoc/utils"
)

type day5 struct{}

func newDay() *day5 {
	return &day5{}
}

func main() {
	doDay()
}

func doDay() {
	day := newDay()
	instr := utils.ReadFileAsStringSlice("/y2022/day5/input.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %s", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %s", day.Part2(input))
}

func (d *day5) Part1(input cmd.Input) interface{} {
	pConfig := getConfig()
	for _, v := range input.SLI {
		from, to, count := getCols(v)

		for i := 0; i < count; i++ {
			pConfig[to] = append(pConfig[to], pConfig[from][len(pConfig[from])-1])
			pConfig[from] = pConfig[from][:len(pConfig[from])-1]
		}
	}

	output := ""
	for i := 1; i <= len(pConfig); i++ {
		output = fmt.Sprintf("%s%s", output, pConfig[i][len(pConfig[i])-1])
	}

	return output
}
func (d *day5) Part2(input cmd.Input) interface{} {
	pConfig := getConfig()
	for _, v := range input.SLI {
		from, to, count := getCols(v)

		for i := len(pConfig[from]) - count; i < len(pConfig[from]); i++ {
			pConfig[to] = append(pConfig[to], pConfig[from][i])
		}
		for i := 0; i < count; i++ {
			pConfig[from] = pConfig[from][:len(pConfig[from])-1]
		}
	}

	output := ""
	for i := 1; i <= len(pConfig); i++ {
		output = fmt.Sprintf("%s%s", output, pConfig[i][len(pConfig[i])-1])
	}

	return output
}

func getCols(v string) (int, int, int) {
	parts := strings.Split(v, " ")

	count, _ := strconv.Atoi(parts[1])
	from, _ := strconv.Atoi(parts[3])
	to, _ := strconv.Atoi(parts[5])

	return from, to, count
}

func getConfig() map[int][]string {
	//return map[int][]string{
	//	1: {"Z", "N"},
	//	2: {"M", "C", "D"},
	//	3: {"P"},
	//}
	return map[int][]string{
		1: {"D", "H", "N", "Q", "T", "W", "V", "B"},
		2: {"D", "W", "B"},
		3: {"T", "S", "Q", "W", "J", "C"},
		4: {"F", "J", "R", "N", "Z", "T", "P"},
		5: {"G", "P", "V", "J", "M", "S", "T"},
		6: {"B", "W", "F", "T", "N"},
		7: {"B", "L", "D", "Q", "F", "H", "V", "N"},
		8: {"H", "P", "F", "R"},
		9: {"Z", "S", "M", "B", "L", "N", "P", "H"},
	}
}
