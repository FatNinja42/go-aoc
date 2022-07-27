package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/cmd"
	"aoc/utils"
)

type operation struct {
	gate string
	op1  int
	op2  int
}

type day7 struct{}

func main() {
	doDay()
}

func doDay() {
	day := day7{}
	instr := utils.ReadFileAsStringSlice("/y2015/day7/input.txt")
	//instr := utils.ReadFileAsStringSlice("/y2015/day7/input_test.txt")

	input := cmd.Input{
		SLI: instr,
	}

	fmt.Printf("P1: %d", day.Part1(input))
	fmt.Println()
	fmt.Printf("P2: %d", day.Part2(input))
}

func (d day7) Part1(input cmd.Input) interface{} {
	circuit, config := buildConfig(input)
	step := true

	for step {
		step = false
		for k, v := range circuit {
			if val, ok := processInstruction(v, config); ok {
				config[k] = val
				step = true
				delete(circuit, k)
			}
		}
	}

	return config["a"]
}

func (d day7) Part2(input cmd.Input) interface{} {
	circuit, config := buildConfig(input)
	circuit["b"] = "3176" //response to part 1
	step := true

	for step {
		step = false
		for k, v := range circuit {
			if val, ok := processInstruction(v, config); ok {
				config[k] = val
				step = true
				delete(circuit, k)
			}
		}
	}

	return config["a"]
}

func processInstruction(v string, config map[string]int) (int, bool) {
	if isNumber(v) {
		return toNumber(v), true
	}

	if val, ok := config[v]; ok && val != -1 {
		return config[v], true
	}

	if !strings.Contains(v, " ") {
		return 0, false
	}

	op, shouldDo := buildOperation(v, config)
	if shouldDo {
		return doOperation(op), true
	}

	return 0, false
}

func buildConfig(input cmd.Input) (map[string]string, map[string]int) {
	var circuit = map[string]string{}
	var values = map[string]int{}

	for _, v := range input.SLI {
		split := strings.Split(v, " -> ")
		circuit[split[1]] = split[0]
		values[split[1]] = -1
	}

	return circuit, values
}

func evaluate(config map[string]int, k string) int {
	if isNumber(k) {
		return toNumber(k)
	}

	return config[k]
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

func toNumber(s string) int {
	v, _ := strconv.Atoi(s)

	return v
}

func doOperation(op operation) int {
	var v = 0
	switch op.gate {
	case "AND":
		v = op.op1 & op.op2
	case "OR":
		v = op.op1 | op.op2
	case "RSHIFT":
		v = op.op1 >> op.op2
	case "LSHIFT":
		v = op.op1 << op.op2
	default:
		v = ^op.op1
	}

	return int(uint16(v))
}

func buildOperation(i string, config map[string]int) (operation, bool) {
	split := strings.Split(i, " ")

	if strings.Contains(i, "NOT") {
		if evaluate(config, split[1]) == -1 {
			return operation{}, false
		}
		return operation{
			gate: split[0],
			op1:  evaluate(config, split[1]),
		}, true
	} else if strings.Contains(i, "AND") || strings.Contains(i, "OR") {
		if evaluate(config, split[0]) == -1 || evaluate(config, split[2]) == -1 {
			return operation{}, false
		}
		return operation{
			gate: split[1],
			op1:  evaluate(config, split[0]),
			op2:  evaluate(config, split[2]),
		}, true
	} else {
		if evaluate(config, split[0]) == -1 {
			return operation{}, false
		}
		return operation{
			gate: split[1],
			op1:  evaluate(config, split[0]),
			op2:  toNumber(split[2]),
		}, true
	}
}
